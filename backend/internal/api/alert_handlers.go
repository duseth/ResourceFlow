package api

import (
	"net/http"

	"github.com/duseth/ResourceFlow/internal/domain/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) GetAlerts(c *gin.Context) {
	alerts, err := h.alertService.GetActiveAlerts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, alerts)
}

func (h *Handler) CreateAlert(c *gin.Context) {
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

func (h *Handler) ResolveAlert(c *gin.Context) {
	alertID := c.Param("id")
	if err := h.alertService.ResolveAlert(c.Request.Context(), alertID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) AcknowledgeAlert(c *gin.Context) {
	alertID := c.Param("id")
	if err := h.alertService.AcknowledgeAlert(c.Request.Context(), alertID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) GetServerAlerts(c *gin.Context) {
	serverID := c.Param("id")
	alerts, err := h.alertService.GetServerAlerts(c.Request.Context(), serverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, alerts)
}
