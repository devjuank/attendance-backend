package persistence

import (
	"time"

	"github.com/juank/attendance-backend/internal/domain/models"
	"github.com/juank/attendance-backend/internal/domain/repositories"
	"gorm.io/gorm"
)

type AttendanceRepositoryImpl struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) repositories.AttendanceRepository {
	return &AttendanceRepositoryImpl{db: db}
}

func (r *AttendanceRepositoryImpl) Create(attendance *models.Attendance) error {
	return r.db.Create(attendance).Error
}

func (r *AttendanceRepositoryImpl) GetByID(id uint) (*models.Attendance, error) {
	var attendance models.Attendance
	if err := r.db.Preload("User").First(&attendance, id).Error; err != nil {
		return nil, err
	}
	return &attendance, nil
}

func (r *AttendanceRepositoryImpl) GetByUserID(userID uint, page, limit int) ([]models.Attendance, int64, error) {
	var attendances []models.Attendance
	var total int64

	offset := (page - 1) * limit

	if err := r.db.Model(&models.Attendance{}).Where("user_id = ?", userID).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	if err := r.db.Where("user_id = ?", userID).Order("check_in desc").Offset(offset).Limit(limit).Find(&attendances).Error; err != nil {
		return nil, 0, err
	}

	return attendances, total, nil
}

func (r *AttendanceRepositoryImpl) GetByDateRange(userID uint, startDate, endDate time.Time) ([]models.Attendance, error) {
	var attendances []models.Attendance
	if err := r.db.Where("user_id = ? AND check_in BETWEEN ? AND ?", userID, startDate, endDate).Order("check_in asc").Find(&attendances).Error; err != nil {
		return nil, err
	}
	return attendances, nil
}

func (r *AttendanceRepositoryImpl) Update(attendance *models.Attendance) error {
	return r.db.Save(attendance).Error
}

func (r *AttendanceRepositoryImpl) GetLastAttendance(userID uint) (*models.Attendance, error) {
	var attendance models.Attendance
	if err := r.db.Where("user_id = ?", userID).Order("check_in desc").First(&attendance).Error; err != nil {
		return nil, err
	}
	return &attendance, nil
}
