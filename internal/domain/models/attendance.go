package models

import (
	"time"

	"gorm.io/gorm"
)

type AttendanceStatus string

const (
	StatusPresent AttendanceStatus = "present"
	StatusAbsent  AttendanceStatus = "absent"
	StatusLate    AttendanceStatus = "late"
	StatusOnLeave AttendanceStatus = "on_leave"
)

type Attendance struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `gorm:"not null;index" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user,omitempty"`
	EventID   uint           `gorm:"not null;index" json:"event_id"`
	Event     Event          `gorm:"foreignKey:EventID" json:"event,omitempty"`
	CheckIn   time.Time      `gorm:"not null" json:"check_in"`
	Status    string         `gorm:"type:varchar(20);not null" json:"status"` // present, late, absent
	Notes     string         `gorm:"type:text" json:"notes"`
	Location  string         `gorm:"type:varchar(255)" json:"location"`
	QRToken   string         `gorm:"type:varchar(255);index" json:"qr_token"` // QR code used for marking
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
