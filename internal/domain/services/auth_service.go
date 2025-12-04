package services

import "github.com/juank/attendance-backend/internal/domain/models"

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=6"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type AuthService interface {
	Register(req *RegisterRequest) (*models.User, error)
	Login(req *LoginRequest) (*TokenResponse, error)
	RefreshToken(token string) (*TokenResponse, error)
	Logout(token string) error
}
