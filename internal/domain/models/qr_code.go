package models

import (
	"time"

	"gorm.io/gorm"
)

type QRCode struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	Token     string         `gorm:"uniqueIndex;not null" json:"qr_token"`
	ExpiresAt time.Time      `gorm:"not null" json:"expires_at"`
	IsActive  bool           `gorm:"default:true;index" json:"is_active"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// TableName specifies the table name for QRCode model
func (QRCode) TableName() string {
	return "qr_codes"
}

// IsExpired checks if the QR code has expired
func (q *QRCode) IsExpired() bool {
	return time.Now().After(q.ExpiresAt)
}

// IsValid checks if the QR code is valid (active and not expired)
func (q *QRCode) IsValid() bool {
	return q.IsActive && !q.IsExpired()
}
