package services

import "github.com/juank/attendance-backend/internal/domain/models"

type QRService interface {
	// GetOrCreateActive returns the active QR code or creates a new one if none exists/expired
	GetOrCreateActive() (*models.QRCode, error)

	// GenerateNew generates a new QR code and deactivates all previous ones
	GenerateNew() (*models.QRCode, error)

	// ValidateToken validates a QR token and returns true if valid
	ValidateToken(token string) (*models.QRCode, error)
}
