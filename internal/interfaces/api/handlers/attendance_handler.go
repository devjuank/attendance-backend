package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/juank/attendance-backend/internal/domain/services"
)

type AttendanceHandler struct {
	attendanceService services.AttendanceService
	qrService         services.QRService
}

func NewAttendanceHandler(attendanceService services.AttendanceService, qrService services.QRService) *AttendanceHandler {
	return &AttendanceHandler{
		attendanceService: attendanceService,
		qrService:         qrService,
	}
}

// MarkAttendance marks attendance using QR token
// @Summary Mark attendance with QR code
// @Tags Attendance
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param request body services.MarkAttendanceRequest true "Mark Attendance Request"
// @Success 201 {object} models.Attendance
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /attendance/mark [post]
func (h *AttendanceHandler) MarkAttendance(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	var req struct {
		QRToken  string `json:"qr_token" binding:"required"`
		Location string `json:"location"`
		Notes    string `json:"notes"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate QR token
	_, err := h.qrService.ValidateToken(req.QRToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Mark attendance
	markReq := &services.MarkAttendanceRequest{
		UserID:   userID.(uint),
		QRToken:  req.QRToken,
		Location: req.Location,
		Notes:    req.Notes,
	}

	attendance, err := h.attendanceService.MarkAttendance(markReq)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, attendance)
}

// GetToday gets today's attendance for current user
// @Summary Get today's attendance
// @Tags Attendance
// @Security BearerAuth
// @Success 200 {object} models.Attendance
// @Failure 404 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /attendance/today [get]
func (h *AttendanceHandler) GetToday(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	attendance, err := h.attendanceService.GetTodayAttendance(userID.(uint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendance)
}

// GetMyHistory gets attendance history for current user
// @Summary Get my attendance history
// @Tags Attendance
// @Security BearerAuth
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]string
// @Router /attendance/history [get]
func (h *AttendanceHandler) GetMyHistory(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	attendances, total, err := h.attendanceService.GetUserAttendance(userID.(uint), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  attendances,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

// GetByDateRange gets attendance records for a date range
// @Summary Get attendance by date range
// @Tags Attendance
// @Security BearerAuth
// @Param start_date query string true "Start date (YYYY-MM-DD)"
// @Param end_date query string true "End date (YYYY-MM-DD)"
// @Success 200 {array} models.Attendance
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /attendance/range [get]
func (h *AttendanceHandler) GetByDateRange(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not authenticated"})
		return
	}

	startDateStr := c.Query("start_date")
	endDateStr := c.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start_date format, use YYYY-MM-DD"})
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end_date format, use YYYY-MM-DD"})
		return
	}

	attendances, err := h.attendanceService.GetByDateRange(userID.(uint), startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, attendances)
}
