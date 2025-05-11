package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/duseth/ResourceFlow/internal/api"
	"github.com/duseth/ResourceFlow/internal/config"
	"github.com/duseth/ResourceFlow/internal/repository/postgres"
	"github.com/duseth/ResourceFlow/internal/service"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

// @title ResourceFlow API
// @version 1.0
// @description API сервиса мониторинга и оптимизации серверных ресурсов
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// Инициализация логгера
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer logger.Sync()

	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("failed to load config", zap.Error(err))
	}

	// Подключение к базе данных
	db, err := sql.Open("postgres", cfg.Database.DSN())
	if err != nil {
		logger.Fatal("failed to connect to database", zap.Error(err))
	}
	defer db.Close()

	// Инициализация репозиториев
	serverRepo := postgres.NewServerRepository(db)
	metricRepo := postgres.NewMetricRepository(db)
	alertRepo := postgres.NewAlertRepository(db)
	optimizationRepo := postgres.NewOptimizationRepository(db)

	// Инициализация сервисов
	monitoringService := service.NewMonitoringService(serverRepo, metricRepo, alertRepo, &cfg.SSH)
	alertService := service.NewAlertService(alertRepo)
	optimizationService := service.NewOptimizationService(optimizationRepo, metricRepo)

	// Инициализация роутера
	router := api.NewRouter(monitoringService, alertService, optimizationService)
	handler := router.Setup()

	// Создание HTTP-сервера
	srv := &http.Server{
		Addr:    cfg.Server.Address(),
		Handler: handler,
	}

	// Запуск сервера в горутине
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("failed to start server", zap.Error(err))
		}
	}()

	// Ожидание сигнала для graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server...")

	// Установка таймаута для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal("server forced to shutdown", zap.Error(err))
	}

	logger.Info("server stopped")
}
