package api

import (
	v1 "github.com/duseth/ResourceFlow/internal/api/v1"
	"github.com/duseth/ResourceFlow/internal/service"
	"github.com/gin-gonic/gin"
)

type Router struct {
	monitoringService   *service.MonitoringService
	alertService        *service.AlertService
	optimizationService *service.OptimizationService
}

func NewRouter(ms *service.MonitoringService, as *service.AlertService, os *service.OptimizationService) *Router {
	return &Router{
		monitoringService:   ms,
		alertService:        as,
		optimizationService: os,
	}
}

func (r *Router) Setup() *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	// API v1
	v1Group := router.Group("/api/v1")
	{
		// Инициализация хендлеров
		serverHandler := v1.NewServerHandler(r.monitoringService)
		alertHandler := v1.NewAlertHandler(r.alertService)
		optimizationHandler := v1.NewOptimizationHandler(r.optimizationService)

		// Серверы
		servers := v1Group.Group("/servers")
		{
			servers.GET("", serverHandler.List)
			servers.GET("/:id", serverHandler.GetByID)
			servers.GET("/:id/metrics", serverHandler.GetMetrics)
			servers.GET("/:id/alerts", serverHandler.GetAlerts)
		}

		// Алерты
		alerts := v1Group.Group("/alerts")
		{
			alerts.GET("", alertHandler.List)
			alerts.POST("", alertHandler.Create)
			alerts.PUT("/:id/resolve", alertHandler.Resolve)
			alerts.PUT("/:id/acknowledge", alertHandler.Acknowledge)
		}

		// Оптимизации
		optimizations := v1Group.Group("/optimizations")
		{
			optimizations.GET("", optimizationHandler.List)
			optimizations.POST("", optimizationHandler.CreateRecommendation)
			optimizations.GET("/server/:id", optimizationHandler.GetServerOptimizations)
			optimizations.POST("/:id/apply", optimizationHandler.Apply)
			optimizations.PUT("/:id", optimizationHandler.UpdateRecommendation)
			optimizations.DELETE("/:id", optimizationHandler.DeleteRecommendation)
		}
	}

	return router
}
