package persistence

import (
	"github.com/juank/attendance-backend/internal/domain/models"
	"github.com/juank/attendance-backend/internal/domain/repositories"
	"gorm.io/gorm"
)

type RefreshTokenRepositoryImpl struct {
	db *gorm.DB
}

func NewRefreshTokenRepository(db *gorm.DB) repositories.RefreshTokenRepository {
	return &RefreshTokenRepositoryImpl{db: db}
}

func (r *RefreshTokenRepositoryImpl) Create(token *models.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r *RefreshTokenRepositoryImpl) GetByToken(token string) (*models.RefreshToken, error) {
	var refreshToken models.RefreshToken
	if err := r.db.Preload("User").Where("token = ?", token).First(&refreshToken).Error; err != nil {
		return nil, err
	}
	return &refreshToken, nil
}

func (r *RefreshTokenRepositoryImpl) Revoke(id uint) error {
	return r.db.Model(&models.RefreshToken{}).Where("id = ?", id).Update("revoked", true).Error
}

func (r *RefreshTokenRepositoryImpl) RevokeByUserID(userID uint) error {
	return r.db.Model(&models.RefreshToken{}).Where("user_id = ?", userID).Update("revoked", true).Error
}
