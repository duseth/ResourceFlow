package api

import (
	v1 "github.com/duseth/ResourceFlow/internal/api/v1"
	"github.com/duseth/ResourceFlow/internal/service"
	"github.com/gin-gonic/gin"
)

func NewRouter(ms *service.MonitoringService) *gin.Engine {
	router := gin.Default()

	// Middleware
	router.Use(gin.Recovery())

	// API v1
	v1Group := router.Group("/api/v1")
	{
		serverHandler := v1.NewServerHandler(ms)

		servers := v1Group.Group("/servers")
		{
			servers.GET("", serverHandler.List)
			servers.GET("/:id", serverHandler.GetByID)
			servers.GET("/:id/metrics", serverHandler.GetMetrics)
			servers.GET("/:id/alerts", serverHandler.GetAlerts)
		}
	}

	return router
}
