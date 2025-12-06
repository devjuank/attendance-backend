package repositories

import (
	"github.com/juank/attendance-backend/internal/domain/models"
)

type QRCodeRepository interface {
	// Create creates a new QR code
	Create(qr *models.QRCode) error

	// GetActive returns the currently active QR code for an event
	GetActive(eventID uint) (*models.QRCode, error)

	// GetByToken finds a QR code by its token
	GetByToken(token string) (*models.QRCode, error)

	// DeactivateAllForEvent deactivates all QR codes for a specific event
	DeactivateAllForEvent(eventID uint) error

	// DeleteExpired deletes expired QR codes
	DeleteExpired() error
}
