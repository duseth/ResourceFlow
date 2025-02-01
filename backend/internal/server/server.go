package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/duseth/ResourceFlow/internal/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Server представляет собой HTTP сервер приложения
type Server struct {
	cfg    *config.Config
	logger *zap.Logger
	router *gin.Engine
	srv    *http.Server
}

// New создает новый экземпляр сервера
func New(cfg *config.Config, logger *zap.Logger) (*Server, error) {
	if cfg.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.New()
	router.Use(gin.Recovery())

	server := &Server{
		cfg:    cfg,
		logger: logger,
		router: router,
	}

	server.setupRoutes()

	server.srv = &http.Server{
		Addr:    fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port),
		Handler: router,
	}

	return server, nil
}

// Run запускает HTTP сервер
func (s *Server) Run() error {
	s.logger.Info("starting server",
		zap.String("addr", s.srv.Addr),
		zap.String("env", s.cfg.App.Env),
	)

	if err := s.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}

// Shutdown gracefully останавливает сервер
func (s *Server) Shutdown(ctx context.Context) error {
	return s.srv.Shutdown(ctx)
}

// setupRoutes настраивает маршруты API
func (s *Server) setupRoutes() {
	api := s.router.Group("/api/v1")
	{
		// Серверы
		servers := api.Group("/servers")
		{
			servers.GET("", s.handleGetServers)
			servers.POST("", s.handleCreateServer)
			servers.GET("/:id", s.handleGetServer)
			servers.PUT("/:id", s.handleUpdateServer)
			servers.DELETE("/:id", s.handleDeleteServer)
		}

		// Метрики
		metrics := api.Group("/metrics")
		{
			metrics.GET("/live/:server_id", s.handleGetLiveMetrics)
			metrics.GET("/history/:server_id", s.handleGetMetricsHistory)
			metrics.GET("/aggregate/:server_id", s.handleGetAggregatedMetrics)
		}

		// Алерты
		alerts := api.Group("/alerts")
		{
			alerts.GET("", s.handleGetAlerts)
			alerts.GET("/:server_id", s.handleGetServerAlerts)
			alerts.POST("/rules", s.handleCreateAlertRule)
			alerts.GET("/rules", s.handleGetAlertRules)
			alerts.PUT("/rules/:id", s.handleUpdateAlertRule)
			alerts.DELETE("/rules/:id", s.handleDeleteAlertRule)
		}

		// Оптимизации
		optimizations := api.Group("/optimizations")
		{
			optimizations.GET("/recommendations/:server_id", s.handleGetOptimizationRecommendations)
			optimizations.POST("/apply/:id", s.handleApplyOptimization)
		}
	}
}

// Временные заглушки для хендлеров
func (s *Server) handleGetServers(c *gin.Context)                     {}
func (s *Server) handleCreateServer(c *gin.Context)                   {}
func (s *Server) handleGetServer(c *gin.Context)                      {}
func (s *Server) handleUpdateServer(c *gin.Context)                   {}
func (s *Server) handleDeleteServer(c *gin.Context)                   {}
func (s *Server) handleGetLiveMetrics(c *gin.Context)                 {}
func (s *Server) handleGetMetricsHistory(c *gin.Context)              {}
func (s *Server) handleGetAggregatedMetrics(c *gin.Context)           {}
func (s *Server) handleGetAlerts(c *gin.Context)                      {}
func (s *Server) handleGetServerAlerts(c *gin.Context)                {}
func (s *Server) handleCreateAlertRule(c *gin.Context)                {}
func (s *Server) handleGetAlertRules(c *gin.Context)                  {}
func (s *Server) handleUpdateAlertRule(c *gin.Context)                {}
func (s *Server) handleDeleteAlertRule(c *gin.Context)                {}
func (s *Server) handleGetOptimizationRecommendations(c *gin.Context) {}
func (s *Server) handleApplyOptimization(c *gin.Context)              {}
