package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetOptimizations(c *gin.Context) {
	recommendations, err := h.optimizationService.GetPendingRecommendations(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, recommendations)
}

func (h *Handler) ApplyOptimization(c *gin.Context) {
	recID := c.Param("id")
	if err := h.optimizationService.ApplyRecommendation(c.Request.Context(), recID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
