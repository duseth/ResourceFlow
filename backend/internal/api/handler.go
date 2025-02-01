package api

import (
	"github.com/duseth/ResourceFlow/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	monitoringService   *service.MonitoringService
	alertService        *service.AlertService
	optimizationService *service.OptimizationService
}

func NewHandler(ms *service.MonitoringService, as *service.AlertService, os *service.OptimizationService) *Handler {
	return &Handler{
		monitoringService:   ms,
		alertService:        as,
		optimizationService: os,
	}
}

func (h *Handler) Router() *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	api := router.Group("/api/v1")
	{
		servers := api.Group("/servers")
		{
			servers.GET("", h.ListServers)
			servers.POST("", h.CreateServer)
			servers.GET("/:id", h.GetServer)
			servers.PUT("/:id", h.UpdateServer)
			servers.DELETE("/:id", h.DeleteServer)
			servers.GET("/:id/metrics", h.GetServerMetrics)
			servers.GET("/:id/alerts", h.GetServerAlerts)
		}

		alerts := api.Group("/alerts")
		{
			alerts.GET("", h.GetAlerts)
			alerts.POST("", h.CreateAlert)
			alerts.PUT("/:id/resolve", h.ResolveAlert)
			alerts.PUT("/:id/acknowledge", h.AcknowledgeAlert)
		}

		rules := api.Group("/rules")
		{
			rules.GET("", h.GetRules)
			rules.POST("", h.CreateRule)
			rules.PUT("/:id", h.UpdateRule)
			rules.DELETE("/:id", h.DeleteRule)
		}

		optimizations := api.Group("/optimizations")
		{
			optimizations.GET("", h.GetOptimizations)
			optimizations.POST("/:id/apply", h.ApplyOptimization)
		}
	}

	return router
}
