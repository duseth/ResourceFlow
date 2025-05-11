package v1

import (
	"net/http"

	"github.com/duseth/ResourceFlow/internal/domain/models"
	"github.com/duseth/ResourceFlow/internal/service"
	"github.com/gin-gonic/gin"
)

type OptimizationHandler struct {
	optimizationService *service.OptimizationService
}

func NewOptimizationHandler(os *service.OptimizationService) *OptimizationHandler {
	return &OptimizationHandler{
		optimizationService: os,
	}
}

func (h *OptimizationHandler) List(c *gin.Context) {
	serverID := c.Query("server_id")
	var recommendations []*models.OptimizationRecommendation
	var err error

	if serverID != "" {
		recommendations, err = h.optimizationService.GetServerRecommendations(c.Request.Context(), serverID)
	} else {
		recommendations, err = h.optimizationService.GetPendingRecommendations(c.Request.Context())
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recommendations)
}

func (h *OptimizationHandler) Apply(c *gin.Context) {
	recID := c.Param("id")
	if recID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "recommendation ID is required"})
		return
	}

	if err := h.optimizationService.ApplyRecommendation(c.Request.Context(), recID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

func (h *OptimizationHandler) GetServerOptimizations(c *gin.Context) {
	serverID := c.Param("id")
	if serverID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "server ID is required"})
		return
	}

	// Сначала анализируем текущие метрики
	if err := h.optimizationService.AnalyzeServerMetrics(c.Request.Context(), serverID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Затем получаем все рекомендации для сервера
	recommendations, err := h.optimizationService.GetServerRecommendations(c.Request.Context(), serverID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, recommendations)
}

func (h *OptimizationHandler) CreateRecommendation(c *gin.Context) {
	var rec models.OptimizationRecommendation
	if err := c.ShouldBindJSON(&rec); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid recommendation data"})
		return
	}

	if rec.ServerID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "server ID is required"})
		return
	}

	if err := h.optimizationService.CreateRecommendation(c.Request.Context(), &rec); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, rec)
}

func (h *OptimizationHandler) UpdateRecommendation(c *gin.Context) {
	recID := c.Param("id")
	if recID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "recommendation ID is required"})
		return
	}

	var update struct {
		Status string `json:"status" binding:"required"`
	}

	if err := c.ShouldBindJSON(&update); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status"})
		return
	}

	if !isValidStatus(update.Status) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status value"})
		return
	}

	if err := h.optimizationService.UpdateStatus(c.Request.Context(), recID, update.Status); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (h *OptimizationHandler) DeleteRecommendation(c *gin.Context) {
	recID := c.Param("id")
	if recID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "recommendation ID is required"})
		return
	}

	if err := h.optimizationService.DeleteRecommendation(c.Request.Context(), recID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

// Вспомогательная функция для проверки статуса
func isValidStatus(status string) bool {
	validStatuses := map[string]bool{
		models.OptimizationStatusPending:    true,
		models.OptimizationStatusApplied:    true,
		models.OptimizationStatusRejected:   true,
		models.OptimizationStatusInProgress: true,
		models.OptimizationStatusFailed:     true,
	}
	return validStatuses[status]
}
