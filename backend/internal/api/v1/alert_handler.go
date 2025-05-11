package v1

import (
	"net/http"

	"github.com/duseth/ResourceFlow/internal/domain/models"
	"github.com/duseth/ResourceFlow/internal/service"
	"github.com/gin-gonic/gin"
)

type AlertHandler struct {
	alertService *service.AlertService
}

func NewAlertHandler(as *service.AlertService) *AlertHandler {
	return &AlertHandler{
		alertService: as,
	}
}

func (h *AlertHandler) List(c *gin.Context) {
	alerts, err := h.alertService.GetActiveAlerts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, alerts)
}

func (h *AlertHandler) Create(c *gin.Context) {
	var alert models.Alert
	if err := c.ShouldBindJSON(&alert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.alertService.CreateAlert(c.Request.Context(), &alert); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, alert)
}

func (h *AlertHandler) Resolve(c *gin.Context) {
	alertID := c.Param("id")
	if err := h.alertService.ResolveAlert(c.Request.Context(), alertID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (h *AlertHandler) Acknowledge(c *gin.Context) {
	alertID := c.Param("id")
	if err := h.alertService.AcknowledgeAlert(c.Request.Context(), alertID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

// ... методы AlertHandler ...
