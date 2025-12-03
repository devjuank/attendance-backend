package repositories

import (
	"time"

	"github.com/juank/attendance-backend/internal/domain/models"
)

type AttendanceRepository interface {
	Create(attendance *models.Attendance) error
	GetByID(id uint) (*models.Attendance, error)
	GetByUserID(userID uint, page, limit int) ([]models.Attendance, int64, error)
	GetByDateRange(userID uint, startDate, endDate time.Time) ([]models.Attendance, error)
	Update(attendance *models.Attendance) error
	GetLastAttendance(userID uint) (*models.Attendance, error)
}
