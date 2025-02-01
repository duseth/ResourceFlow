package api

import (
	"net/http"
	"time"

	"github.com/duseth/ResourceFlow/internal/domain/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) ListServers(c *gin.Context) {
	servers, err := h.monitoringService.GetServers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, servers)
}

func (h *Handler) CreateServer(c *gin.Context) {
	var server models.Server
	if err := c.ShouldBindJSON(&server); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.monitoringService.CreateServer(c.Request.Context(), &server); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, server)
}

// ... другие методы для серверов

func (h *Handler) GetServerMetrics(c *gin.Context) {
	serverID := c.Param("id")
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

func (h *Handler) GetServer(c *gin.Context) {
	serverID := c.Param("id")
	server, err := h.monitoringService.GetServer(c.Request.Context(), serverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, server)
}

func (h *Handler) UpdateServer(c *gin.Context) {
	var server models.Server
	if err := c.ShouldBindJSON(&server); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	server.ID = c.Param("id")
	if err := h.monitoringService.UpdateServer(c.Request.Context(), &server); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, server)
}

func (h *Handler) DeleteServer(c *gin.Context) {
	serverID := c.Param("id")
	if err := h.monitoringService.DeleteServer(c.Request.Context(), serverID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
