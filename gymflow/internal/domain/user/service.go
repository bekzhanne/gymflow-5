package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Service interface {
	Register(req RegisterRequest) (*User, error)
	Login(req LoginRequest) (*User, error)
	GetByID(id uint) (*User, error)
	ListUsers() ([]User, error)
}

type service struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &service{repo: repo}
}

func (s *service) Register(req RegisterRequest) (*User, error) {
	_, err := s.repo.FindByEmail(req.Email)
	if err == nil {
		return nil, errors.New("email already used")
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return nil, err
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	u := &User{
		Name:           req.Name,
		Email:          req.Email,
		PasswordHash:   string(hash),
		Role:           req.Role,
		MembershipTier: req.MembershipTier,
		Active:         true,
	}

	if err := s.repo.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}

func (s *service) Login(req LoginRequest) (*User, error) {
	u, err := s.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid credentials")
	}

	return u, nil
}

func (s *service) GetByID(id uint) (*User, error) {
	return s.repo.FindByID(id)
}

func (s *service) ListUsers() ([]User, error) {
	return s.repo.List()
}

