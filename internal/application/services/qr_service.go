package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/juank/attendance-backend/internal/domain/models"
	"github.com/juank/attendance-backend/internal/domain/repositories"
	domainServices "github.com/juank/attendance-backend/internal/domain/services"
)

const QRExpirationMinutes = 10

type QRServiceImpl struct {
	qrRepo repositories.QRCodeRepository
}

func NewQRService(qrRepo repositories.QRCodeRepository) domainServices.QRService {
	return &QRServiceImpl{
		qrRepo: qrRepo,
	}
}

func (s *QRServiceImpl) GetOrCreateActive() (*models.QRCode, error) {
	// Try to get active QR code
	qr, err := s.qrRepo.GetActive()
	if err == nil && qr != nil {
		// Valid active QR found
		return qr, nil
	}

	// No active QR or error, create new one
	return s.GenerateNew()
}

func (s *QRServiceImpl) GenerateNew() (*models.QRCode, error) {
	// Deactivate all existing QR codes
	err := s.qrRepo.DeactivateAll()
	if err != nil {
		return nil, err
	}

	// Generate new QR code
	token := uuid.New().String()
	expiresAt := time.Now().Add(QRExpirationMinutes * time.Minute)

	qr := &models.QRCode{
		Token:     token,
		ExpiresAt: expiresAt,
		IsActive:  true,
	}

	err = s.qrRepo.Create(qr)
	if err != nil {
		return nil, err
	}

	// Clean up expired QR codes (async cleanup)
	go s.qrRepo.DeleteExpired()

	return qr, nil
}

func (s *QRServiceImpl) ValidateToken(token string) (*models.QRCode, error) {
	qr, err := s.qrRepo.GetByToken(token)
	if err != nil {
		return nil, errors.New("invalid QR code")
	}

	if !qr.IsValid() {
		return nil, errors.New("QR code expired or inactive")
	}

	return qr, nil
}
