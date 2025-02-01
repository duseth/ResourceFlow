package v1

import (
	"net/http"
	"time"

	"github.com/duseth/ResourceFlow/internal/service"
	"github.com/gin-gonic/gin"
)

type ServerHandler struct {
	monitoringService *service.MonitoringService
}

func NewServerHandler(ms *service.MonitoringService) *ServerHandler {
	return &ServerHandler{
		monitoringService: ms,
	}
}

func (h *ServerHandler) List(c *gin.Context) {
	servers, err := h.monitoringService.GetServers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, servers)
}

func (h *ServerHandler) GetByID(c *gin.Context) {
	serverID := c.Param("id")
	server, err := h.monitoringService.GetServer(c.Request.Context(), serverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, server)
}

func (h *ServerHandler) GetMetrics(c *gin.Context) {
	serverID := c.Param("id")

	// Получаем временной диапазон из query параметров
	from := time.Now().Add(-24 * time.Hour)
	to := time.Now()

	if fromStr := c.Query("from"); fromStr != "" {
		if t, err := time.Parse(time.RFC3339, fromStr); err == nil {
			from = t
		}
	}
	if toStr := c.Query("to"); toStr != "" {
		if t, err := time.Parse(time.RFC3339, toStr); err == nil {
			to = t
		}
	}

	metrics, err := h.monitoringService.GetServerMetrics(c.Request.Context(), serverID, from, to)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, metrics)
}

func (h *ServerHandler) GetAlerts(c *gin.Context) {
	serverID := c.Param("id")
	alerts, err := h.monitoringService.GetServerAlerts(c.Request.Context(), serverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, alerts)
}
