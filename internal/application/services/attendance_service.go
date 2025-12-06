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
	qrService      services.QRService
}

func NewAttendanceService(attendanceRepo repositories.AttendanceRepository, qrService services.QRService) services.AttendanceService {
	return &AttendanceServiceImpl{
		attendanceRepo: attendanceRepo,
		qrService:      qrService,
	}
}

func (s *AttendanceServiceImpl) MarkAttendance(req *services.MarkAttendanceRequest) (*models.Attendance, error) {
	// Validate QR token
	qr, err := s.qrService.ValidateToken(req.QRToken)
	if err != nil {
		return nil, err
	}

	// Check if user already marked attendance for this event
	existingAttendance, err := s.attendanceRepo.GetByEventAndUser(qr.EventID, req.UserID)
	if err == nil && existingAttendance != nil {
		return nil, errors.New("user already marked attendance for this event")
	}

	// Determine status based on check-in time
	now := time.Now()
	status := s.calculateStatus(now)

	// Create attendance record
	attendance := &models.Attendance{
		UserID:   req.UserID,
		EventID:  qr.EventID,
		CheckIn:  now,
		Status:   string(status),
		Location: req.Location,
		Notes:    req.Notes,
		QRToken:  req.QRToken,
	}

	err = s.attendanceRepo.Create(attendance)
	if err != nil {
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

// calculateStatus determines the attendance status based on check-in time
func (s *AttendanceServiceImpl) calculateStatus(checkInTime time.Time) models.AttendanceStatus {
	// Business rule: Late if after 9:15 AM
	limitTime := time.Date(
		checkInTime.Year(),
		checkInTime.Month(),
		checkInTime.Day(),
		9, 15, 0, 0,
		checkInTime.Location(),
	)

	if checkInTime.After(limitTime) {
		return models.StatusLate
	}

	return models.StatusPresent
}

func (s *AttendanceServiceImpl) GetEventAttendance(eventID uint) ([]models.Attendance, error) {
	return s.attendanceRepo.GetByEventID(eventID)
}

func (s *AttendanceServiceImpl) MarkManualAttendance(eventID, userID uint, notes string) (*models.Attendance, error) {
	// Check if user already marked attendance for this event
	existingAttendance, err := s.attendanceRepo.GetByEventAndUser(eventID, userID)
	if err == nil && existingAttendance != nil {
		return nil, errors.New("user already marked attendance for this event")
	}

	now := time.Now()
	// Manual entry is typically considered "present" or we could calculate logic.
	// For simplicity, let's calculate status normally.
	status := s.calculateStatus(now)

	attendance := &models.Attendance{
		UserID:   userID,
		EventID:  eventID,
		CheckIn:  now,
		Status:   string(status),
		Notes:    notes,
		Location: "Manual Entry",
	}

	if err := s.attendanceRepo.Create(attendance); err != nil {
		return nil, err
	}

	return attendance, nil
}
