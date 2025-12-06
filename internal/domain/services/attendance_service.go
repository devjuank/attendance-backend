package services

import (
	"time"

	"github.com/juank/attendance-backend/internal/domain/models"
)

type MarkAttendanceRequest struct {
	UserID   uint   `json:"user_id" validate:"required"`
	QRToken  string `json:"qr_token" validate:"required"`
	Location string `json:"location"`
	Notes    string `json:"notes"`
}

type AttendanceService interface {
	MarkAttendance(req *MarkAttendanceRequest) (*models.Attendance, error)
	GetByID(id uint) (*models.Attendance, error)
	GetUserAttendance(userID uint, page, limit int) ([]models.Attendance, int64, error)
	GetTodayAttendance(userID uint) (*models.Attendance, error)
	GetByDateRange(userID uint, startDate, endDate time.Time) ([]models.Attendance, error)
	GetEventAttendance(eventID uint) ([]models.Attendance, error)
	MarkManualAttendance(eventID, userID uint, notes string) (*models.Attendance, error)
}
