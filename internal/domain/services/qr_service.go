package services

import "github.com/juank/attendance-backend/internal/domain/models"

type QRService interface {
	// GetOrCreateActive returns the active QR code for an event or creates a new one
	GetOrCreateActive(eventID uint) (*models.QRCode, error)

	// GenerateNew generates a new QR code for an event and deactivates all previous ones
	GenerateNew(eventID uint) (*models.QRCode, error)

	// ValidateToken validates a QR token and returns true if valid
	ValidateToken(token string) (*models.QRCode, error)

	// DeactivateActiveForEvent deactivates the current active QR code for an event
	DeactivateActiveForEvent(eventID uint) error
}
