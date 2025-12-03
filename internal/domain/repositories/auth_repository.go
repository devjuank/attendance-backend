package repositories

import "github.com/juank/attendance-backend/internal/domain/models"

type RefreshTokenRepository interface {
	Create(token *models.RefreshToken) error
	GetByToken(token string) (*models.RefreshToken, error)
	Revoke(id uint) error
	RevokeByUserID(userID uint) error
}
