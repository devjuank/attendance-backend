package services

import (
	"time"

	"github.com/juank/attendance-backend/internal/domain/models"
)

type CheckInRequest struct {
	UserID   uint   `json:"user_id" validate:"required"`
	Location string `json:"location"`
	Notes    string `json:"notes"`
}

type CheckOutRequest struct {
	UserID uint   `json:"user_id" validate:"required"`
	Notes  string `json:"notes"`
}

type AttendanceService interface {
	CheckIn(req *CheckInRequest) (*models.Attendance, error)
	CheckOut(req *CheckOutRequest) (*models.Attendance, error)
	GetByID(id uint) (*models.Attendance, error)
	GetUserAttendance(userID uint, page, limit int) ([]models.Attendance, int64, error)
	GetTodayAttendance(userID uint) (*models.Attendance, error)
	GetByDateRange(userID uint, startDate, endDate time.Time) ([]models.Attendance, error)
}
