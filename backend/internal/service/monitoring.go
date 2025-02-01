package service

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/duseth/ResourceFlow/internal/config"
	"github.com/duseth/ResourceFlow/internal/domain/models"
	"github.com/duseth/ResourceFlow/internal/domain/repository"
	"github.com/duseth/ResourceFlow/internal/ssh"
	"github.com/google/uuid"
)

type MonitoringService struct {
	serverRepo repository.ServerRepository
	metricRepo repository.MetricRepository
	alertRepo  repository.AlertRepository
	sshClient  *ssh.Client
}

func NewMonitoringService(sr repository.ServerRepository, mr repository.MetricRepository, ar repository.AlertRepository, cfg *config.SSHConfig) *MonitoringService {
	sshClient := ssh.NewClient(cfg.User, cfg.Password)
	return &MonitoringService{
		serverRepo: sr,
		metricRepo: mr,
		alertRepo:  ar,
		sshClient:  sshClient,
	}
}

func (s *MonitoringService) CreateServer(ctx context.Context, server *models.Server) error {
	server.ID = uuid.New().String()
	server.CreatedAt = time.Now()
	server.UpdatedAt = time.Now()
	server.LastCheckAt = time.Now()
	server.Status = models.ServerStatusActive

	return s.serverRepo.Create(ctx, server)
}

func (s *MonitoringService) UpdateServer(ctx context.Context, server *models.Server) error {
	server.UpdatedAt = time.Now()
	return s.serverRepo.Update(ctx, server)
}

func (s *MonitoringService) DeleteServer(ctx context.Context, id string) error {
	return s.serverRepo.Delete(ctx, id)
}

func (s *MonitoringService) GetServer(ctx context.Context, id string) (*models.Server, error) {
	return s.serverRepo.GetByID(ctx, id)
}

func (s *MonitoringService) GetServers(ctx context.Context) ([]*models.Server, error) {
	return s.serverRepo.List(ctx, repository.ServerFilter{})
}

func (s *MonitoringService) StoreMetric(ctx context.Context, metric *models.Metric) error {
	metric.ID = uuid.New().String()
	metric.Timestamp = time.Now()
	return s.metricRepo.Store(ctx, metric)
}

func (s *MonitoringService) GetServerMetrics(ctx context.Context, serverID string, from, to time.Time) ([]*models.Metric, error) {
	return s.metricRepo.GetRange(ctx, serverID, from, to)
}

func (s *MonitoringService) GetLatestMetric(ctx context.Context, serverID string, metricType string) (*models.Metric, error) {
	return s.metricRepo.GetLatest(ctx, serverID, metricType)
}

func (s *MonitoringService) GetHistoricalData(ctx context.Context, serverID string, period string) ([]*models.HistoricalData, error) {
	return s.metricRepo.GetAggregated(ctx, serverID, period)
}

func (s *MonitoringService) CollectMetrics(ctx context.Context, serverID string) error {
	server, err := s.serverRepo.GetByID(ctx, serverID)
	if err != nil {
		return err
	}

	// Проверяем доступность сервера перед сбором метрик
	if err := s.checkServerAvailability(server); err != nil {
		server.Status = models.ServerStatusError
		_ = s.serverRepo.Update(ctx, server)
		return err
	}

	// Если сервер был в статусе ошибки, возвращаем его в активный статус
	if server.Status == models.ServerStatusError {
		server.Status = models.ServerStatusActive
		if err := s.serverRepo.Update(ctx, server); err != nil {
			return err
		}
	}

	// Сбор метрик CPU
	cpuMetric := &models.Metric{
		ID:        uuid.New().String(),
		ServerID:  serverID,
		Type:      models.MetricTypeCPU,
		Timestamp: time.Now(),
	}

	// Здесь должна быть реальная логика сбора метрик CPU
	cpuMetric.Value = s.collectCPUMetrics(server)
	if err := s.metricRepo.Store(ctx, cpuMetric); err != nil {
		return err
	}

	// Сбор метрик памяти
	memoryMetric := &models.Metric{
		ID:        uuid.New().String(),
		ServerID:  serverID,
		Type:      models.MetricTypeMemory,
		Timestamp: time.Now(),
	}

	// Здесь должна быть реальная логика сбора метрик памяти
	memoryMetric.Value = s.collectMemoryMetrics(server)
	if err := s.metricRepo.Store(ctx, memoryMetric); err != nil {
		return err
	}

	// Обновляем время последней проверки сервера
	server.LastCheckAt = time.Now()
	return s.serverRepo.Update(ctx, server)
}

func (s *MonitoringService) CheckAlerts(ctx context.Context, serverID string) error {
	// Получаем последние метрики
	cpuMetric, err := s.metricRepo.GetLatest(ctx, serverID, models.MetricTypeCPU)
	if err != nil {
		if err == sql.ErrNoRows {
			// Если метрик нет, это не ошибка
			return nil
		}
		return err
	}

	memoryMetric, err := s.metricRepo.GetLatest(ctx, serverID, models.MetricTypeMemory)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		return err
	}

	// Проверяем CPU метрики
	if cpuMetric.Value > models.CPUCriticalThreshold {
		alert := &models.Alert{
			ID:        uuid.New().String(),
			ServerID:  serverID,
			Status:    models.AlertStatusActive,
			Message:   fmt.Sprintf("Critical CPU usage: %.2f%%", cpuMetric.Value),
			CreatedAt: time.Now(),
		}
		if err := s.alertRepo.Create(ctx, alert); err != nil {
			return err
		}
	}

	// Проверяем Memory метрики
	if memoryMetric.Value > models.MemoryCriticalThreshold {
		alert := &models.Alert{
			ID:        uuid.New().String(),
			ServerID:  serverID,
			Status:    models.AlertStatusActive,
			Message:   fmt.Sprintf("Critical Memory usage: %.2f%%", memoryMetric.Value),
			CreatedAt: time.Now(),
		}
		if err := s.alertRepo.Create(ctx, alert); err != nil {
			return err
		}
	}

	return nil
}

// Обновляем вспомогательные методы для сбора метрик
func (s *MonitoringService) collectCPUMetrics(server *models.Server) float64 {
	// Реализация сбора метрик CPU через SSH
	usage, err := s.getServerCPUUsage(server)
	if err != nil {
		// Логируем ошибку и возвращаем 0
		return 0.0
	}
	return usage
}

func (s *MonitoringService) collectMemoryMetrics(server *models.Server) float64 {
	// Реализация сбора метрик памяти через SSH
	usage, err := s.getServerMemoryUsage(server)
	if err != nil {
		// Логируем ошибку и возвращаем 0
		return 0.0
	}
	return usage
}

func (s *MonitoringService) getServerCPUUsage(server *models.Server) (float64, error) {
	command := `top -bn1 | grep "Cpu(s)" | awk '{print $2 + $4}'`
	output, err := s.sshClient.ExecuteCommand(server.Host, server.Port, command)
	if err != nil {
		return 0.0, fmt.Errorf("failed to get CPU usage: %v", err)
	}

	usage, err := strconv.ParseFloat(strings.TrimSpace(output), 64)
	if err != nil {
		return 0.0, fmt.Errorf("failed to parse CPU usage: %v", err)
	}

	return usage, nil
}

func (s *MonitoringService) getServerMemoryUsage(server *models.Server) (float64, error) {
	command := `free | grep Mem | awk '{print $3/$2 * 100.0}'`
	output, err := s.sshClient.ExecuteCommand(server.Host, server.Port, command)
	if err != nil {
		return 0.0, fmt.Errorf("failed to get memory usage: %v", err)
	}

	usage, err := strconv.ParseFloat(strings.TrimSpace(output), 64)
	if err != nil {
		return 0.0, fmt.Errorf("failed to parse memory usage: %v", err)
	}

	return usage, nil
}

// Добавляем метод для проверки доступности сервера
func (s *MonitoringService) checkServerAvailability(server *models.Server) error {
	// Пробуем выполнить простую команду для проверки доступности
	command := "echo 1"
	_, err := s.sshClient.ExecuteCommand(server.Host, server.Port, command)
	if err != nil {
		return fmt.Errorf("server is not available: %v", err)
	}
	return nil
}

func (s *MonitoringService) GetServerAlerts(ctx context.Context, serverID string) ([]*models.Alert, error) {
	return s.alertRepo.GetByServer(ctx, serverID)
}
