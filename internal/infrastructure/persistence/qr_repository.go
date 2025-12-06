package persistence

import (
	"time"

	"github.com/juank/attendance-backend/internal/domain/models"
	"github.com/juank/attendance-backend/internal/domain/repositories"
	"gorm.io/gorm"
)

type QRCodeRepositoryImpl struct {
	db *gorm.DB
}

func NewQRCodeRepository(db *gorm.DB) repositories.QRCodeRepository {
	return &QRCodeRepositoryImpl{db: db}
}

func (r *QRCodeRepositoryImpl) Create(qr *models.QRCode) error {
	return r.db.Create(qr).Error
}

func (r *QRCodeRepositoryImpl) GetActive(eventID uint) (*models.QRCode, error) {
	var qr models.QRCode
	err := r.db.Where("event_id = ? AND is_active = ? AND expires_at > ?", eventID, true, time.Now()).
		Order("created_at DESC").
		First(&qr).Error

	if err != nil {
		return nil, err
	}

	return &qr, nil
}

func (r *QRCodeRepositoryImpl) GetByToken(token string) (*models.QRCode, error) {
	var qr models.QRCode
	err := r.db.Preload("Event").Where("token = ?", token).First(&qr).Error
	if err != nil {
		return nil, err
	}
	return &qr, nil
}

func (r *QRCodeRepositoryImpl) DeactivateAllForEvent(eventID uint) error {
	return r.db.Model(&models.QRCode{}).
		Where("event_id = ? AND is_active = ?", eventID, true).
		Update("is_active", false).Error
}

func (r *QRCodeRepositoryImpl) DeleteExpired() error {
	return r.db.Where("expires_at < ?", time.Now()).
		Delete(&models.QRCode{}).Error
}
