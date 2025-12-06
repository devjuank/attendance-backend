package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/juank/attendance-backend/internal/domain/services"
)

type QRHandler struct {
	qrService services.QRService
}

func NewQRHandler(qrService services.QRService) *QRHandler {
	return &QRHandler{
		qrService: qrService,
	}
}

// GetActive returns the currently active QR code (auto-generates if none exists)
// @Summary Get active QR code
// @Tags QR
// @Security BearerAuth
// @Success 200 {object} models.QRCode
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /qr/active [get]
func (h *QRHandler) GetActive(c *gin.Context) {
	eventIDStr := c.Query("event_id")
	if eventIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "event_id is required"})
		return
	}

	eventID, err := strconv.ParseUint(eventIDStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid event_id"})
		return
	}

	qr, err := h.qrService.GetOrCreateActive(uint(eventID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, qr)
}

// Generate creates a new QR code (invalidates previous)
// @Summary Generate new QR code
// @Tags QR
// @Security BearerAuth
// @Success 201 {object} models.QRCode
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /qr/generate [post]
func (h *QRHandler) Generate(c *gin.Context) {
	var req struct {
		EventID uint `json:"event_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	qr, err := h.qrService.GenerateNew(req.EventID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, qr)
}

// Deactivate deactivates the current QR code for an event
// @Summary Deactivate QR code
// @Tags QR
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /qr/deactivate [post]
func (h *QRHandler) Deactivate(c *gin.Context) {
	var req struct {
		EventID uint `json:"event_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.qrService.DeactivateActiveForEvent(req.EventID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "QR code deactivated",
		"event_id": req.EventID,
	})
}
