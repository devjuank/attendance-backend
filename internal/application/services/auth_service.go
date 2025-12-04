package services

import (
	"errors"
	"fmt"
	"time"

	"github.com/juank/attendance-backend/config"
	"github.com/juank/attendance-backend/internal/domain/models"
	"github.com/juank/attendance-backend/internal/domain/repositories"
	"github.com/juank/attendance-backend/internal/domain/services"
	"github.com/juank/attendance-backend/pkg/utils"
)

type AuthServiceImpl struct {
	userRepo         repositories.UserRepository
	refreshTokenRepo repositories.RefreshTokenRepository
	cfg              *config.Config
}

func NewAuthService(
	userRepo repositories.UserRepository,
	refreshTokenRepo repositories.RefreshTokenRepository,
	cfg *config.Config,
) services.AuthService {
	return &AuthServiceImpl{
		userRepo:         userRepo,
		refreshTokenRepo: refreshTokenRepo,
		cfg:              cfg,
	}
}

func (s *AuthServiceImpl) Register(req *services.RegisterRequest) (*models.User, error) {
	// Verificar si el email ya existe
	existingUser, _ := s.userRepo.GetByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Role:      models.RoleEmployee, // Por defecto
		IsActive:  true,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *AuthServiceImpl) Login(req *services.LoginRequest) (*services.TokenResponse, error) {
	user, err := s.userRepo.GetByEmail(req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(req.Password, user.Password) {
		return nil, errors.New("invalid credentials")
	}

	if !user.IsActive {
		return nil, errors.New("user is inactive")
	}

	return s.generateTokens(user)
}

func (s *AuthServiceImpl) RefreshToken(tokenString string) (*services.TokenResponse, error) {
	// Validar refresh token
	_, err := utils.ValidateToken(tokenString, s.cfg.JWT.Secret)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	// Verificar si existe en BD y no está revocado
	storedToken, err := s.refreshTokenRepo.GetByToken(tokenString)
	if err != nil {
		return nil, errors.New("refresh token not found")
	}

	if storedToken.Revoked {
		return nil, errors.New("refresh token revoked")
	}

	// Obtener usuario
	user, err := s.userRepo.GetByID(storedToken.UserID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Revocar token anterior (rotación de refresh tokens)
	if err := s.refreshTokenRepo.Revoke(storedToken.ID); err != nil {
		return nil, err
	}

	return s.generateTokens(user)
}

func (s *AuthServiceImpl) Logout(tokenString string) error {
	// Simplemente revocamos el refresh token si se proporciona
	// En una implementación stateless pura, el logout es del lado del cliente (borrar token)
	// Pero aquí invalidamos el refresh token
	storedToken, err := s.refreshTokenRepo.GetByToken(tokenString)
	if err == nil && storedToken != nil {
		return s.refreshTokenRepo.Revoke(storedToken.ID)
	}
	return nil
}

func (s *AuthServiceImpl) generateTokens(user *models.User) (*services.TokenResponse, error) {
	accessToken, refreshToken, err := utils.GenerateTokenPair(user.ID, user.Email, string(user.Role), s.cfg)
	if err != nil {
		return nil, err
	}

	// Guardar refresh token
	rt := &models.RefreshToken{
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(s.cfg.JWT.RefreshExpiration),
		Revoked:   false,
	}

	if err := s.refreshTokenRepo.Create(rt); err != nil {
		return nil, fmt.Errorf("failed to store refresh token: %w", err)
	}

	return &services.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
