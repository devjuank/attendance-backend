package handlers

import (
	"net/http"

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
	qr, err := h.qrService.GetOrCreateActive()
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
	qr, err := h.qrService.GenerateNew()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, qr)
}
