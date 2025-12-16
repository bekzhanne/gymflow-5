package user

import (
	"net/http"

	"gymflow/internal/config"
	"gymflow/internal/token"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	cfg     *config.Config
	service Service
}

func NewHandler(cfg *config.Config, service Service) *Handler {
	return &Handler{cfg: cfg, service: service}
}

// POST /api/v1/auth/register
func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.service.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tkn, err := token.GenerateToken(h.cfg, u.ID, u.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot generate token"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user":  ToUserResponse(u),
		"token": tkn,
	})
}

// POST /api/v1/auth/login
func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := h.service.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	tkn, err := token.GenerateToken(h.cfg, u.ID, u.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot generate token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user":  ToUserResponse(u),
		"token": tkn,
	})
}

// GET /api/v1/users (admin)
func (h *Handler) ListUsers(c *gin.Context) {
	users, err := h.service.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to list users"})
		return
	}
	resp := make([]*UserResponse, 0, len(users))
	for i := range users {
		resp = append(resp, ToUserResponse(&users[i]))
	}
	c.JSON(http.StatusOK, resp)
}

