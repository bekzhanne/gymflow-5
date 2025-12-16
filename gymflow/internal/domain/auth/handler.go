package auth

import (
	"gymflow/internal/domain/user"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	userHandler *user.Handler
}

func NewHandler(userHandler *user.Handler) *Handler {
	return &Handler{userHandler: userHandler}
}

func (h *Handler) Register(c *gin.Context) {
	h.userHandler.Register(c)
}

func (h *Handler) Login(c *gin.Context) {
	h.userHandler.Login(c)
}
