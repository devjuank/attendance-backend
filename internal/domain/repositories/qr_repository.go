package repositories

import (
	"github.com/juank/attendance-backend/internal/domain/models"
)

type QRCodeRepository interface {
	// Create creates a new QR code
	Create(qr *models.QRCode) error

	// GetActive returns the currently active QR code
	GetActive() (*models.QRCode, error)

	// GetByToken finds a QR code by its token
	GetByToken(token string) (*models.QRCode, error)

	// DeactivateAll deactivates all QR codes
	DeactivateAll() error

	// DeleteExpired deletes expired QR codes
	DeleteExpired() error
}
