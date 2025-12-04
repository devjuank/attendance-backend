package services

import (
	"errors"
	"time"

	"github.com/juank/attendance-backend/internal/domain/models"
	"github.com/juank/attendance-backend/internal/domain/repositories"
	"github.com/juank/attendance-backend/internal/domain/services"
)

type AttendanceServiceImpl struct {
	attendanceRepo repositories.AttendanceRepository
}

func NewAttendanceService(attendanceRepo repositories.AttendanceRepository) services.AttendanceService {
	return &AttendanceServiceImpl{
		attendanceRepo: attendanceRepo,
	}
}

func (s *AttendanceServiceImpl) CheckIn(req *services.CheckInRequest) (*models.Attendance, error) {
	// Verificar si ya hizo check-in hoy y no ha hecho check-out
	lastAttendance, err := s.attendanceRepo.GetLastAttendance(req.UserID)
	if err == nil && lastAttendance != nil {
		// Si el último registro es de hoy y no tiene check-out, error
		today := time.Now().Truncate(24 * time.Hour)
		lastDate := lastAttendance.CheckIn.Truncate(24 * time.Hour)

		if today.Equal(lastDate) && lastAttendance.CheckOut == nil {
			return nil, errors.New("user already checked in")
		}
	}

	// Determinar estado (ejemplo simple: tarde si es después de las 9:00 AM)
	now := time.Now()
	status := models.StatusPresent

	// Lógica de negocio simple para determinar si llega tarde (ej: después de las 9:15)
	limitTime := time.Date(now.Year(), now.Month(), now.Day(), 9, 15, 0, 0, now.Location())
	if now.After(limitTime) {
		status = models.StatusLate
	}

	attendance := &models.Attendance{
		UserID:   req.UserID,
		CheckIn:  now,
		Status:   status,
		Location: req.Location,
		Notes:    req.Notes,
	}

	if err := s.attendanceRepo.Create(attendance); err != nil {
		return nil, err
	}

	return attendance, nil
}

func (s *AttendanceServiceImpl) CheckOut(req *services.CheckOutRequest) (*models.Attendance, error) {
	// Buscar el último registro sin check-out
	attendance, err := s.attendanceRepo.GetLastAttendance(req.UserID)
	if err != nil {
		return nil, errors.New("no active check-in found")
	}

	if attendance.CheckOut != nil {
		return nil, errors.New("user already checked out")
	}

	now := time.Now()
	attendance.CheckOut = &now

	if req.Notes != "" {
		if attendance.Notes != "" {
			attendance.Notes += "; " + req.Notes
		} else {
			attendance.Notes = req.Notes
		}
	}

	if err := s.attendanceRepo.Update(attendance); err != nil {
		return nil, err
	}

	return attendance, nil
}

func (s *AttendanceServiceImpl) GetByID(id uint) (*models.Attendance, error) {
	return s.attendanceRepo.GetByID(id)
}

func (s *AttendanceServiceImpl) GetUserAttendance(userID uint, page, limit int) ([]models.Attendance, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	return s.attendanceRepo.GetByUserID(userID, page, limit)
}

func (s *AttendanceServiceImpl) GetTodayAttendance(userID uint) (*models.Attendance, error) {
	lastAttendance, err := s.attendanceRepo.GetLastAttendance(userID)
	if err != nil {
		return nil, err
	}

	today := time.Now().Truncate(24 * time.Hour)
	lastDate := lastAttendance.CheckIn.Truncate(24 * time.Hour)

	if !today.Equal(lastDate) {
		return nil, errors.New("no attendance record for today")
	}

	return lastAttendance, nil
}

func (s *AttendanceServiceImpl) GetByDateRange(userID uint, startDate, endDate time.Time) ([]models.Attendance, error) {
	return s.attendanceRepo.GetByDateRange(userID, startDate, endDate)
}
