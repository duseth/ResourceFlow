package service

import (
	"context"
	"fmt"
	"time"

	"github.com/duseth/ResourceFlow/internal/domain/models"
	"github.com/duseth/ResourceFlow/internal/domain/repository"
	"github.com/google/uuid"
)

// OptimizationService предоставляет методы для анализа и оптимизации серверных ресурсов
type OptimizationService struct {
	optimizationRepo repository.OptimizationRepository
	metricRepo       repository.MetricRepository
}

// NewOptimizationService создает новый экземпляр OptimizationService
func NewOptimizationService(or repository.OptimizationRepository, mr repository.MetricRepository) *OptimizationService {
	return &OptimizationService{
		optimizationRepo: or,
		metricRepo:       mr,
	}
}

// AnalyzeServerMetrics анализирует метрики сервера и создает рекомендации по оптимизации
func (s *OptimizationService) AnalyzeServerMetrics(ctx context.Context, serverID string) error {
	// Получаем метрики за последние 24 часа
	metrics, err := s.metricRepo.GetRange(ctx, serverID, time.Now().Add(-24*time.Hour), time.Now())
	if err != nil {
		return fmt.Errorf("failed to get metrics: %v", err)
	}

	// Анализируем метрики и создаем рекомендации
	recommendations := s.analyzeMetrics(metrics)

	// Сохраняем рекомендации
	for _, rec := range recommendations {
		rec.ID = uuid.New().String()
		rec.ServerID = serverID
		rec.CreatedAt = time.Now()
		rec.Status = models.OptimizationStatusPending

		if err := s.optimizationRepo.CreateRecommendation(ctx, rec); err != nil {
			return fmt.Errorf("failed to create recommendation: %v", err)
		}
	}

	return nil
}

// GetPendingRecommendations возвращает список ожидающих применения рекомендаций
func (s *OptimizationService) GetPendingRecommendations(ctx context.Context) ([]*models.OptimizationRecommendation, error) {
	return s.optimizationRepo.GetPending(ctx)
}

// ApplyRecommendation применяет рекомендацию по оптимизации
func (s *OptimizationService) ApplyRecommendation(ctx context.Context, recID string) error {
	// Здесь должна быть реальная логика применения оптимизации
	// Например, изменение конфигурации сервера или масштабирование ресурсов

	// Обновляем статус рекомендации
	return s.optimizationRepo.UpdateStatus(ctx, recID, models.OptimizationStatusApplied)
}

// analyzeMetrics анализирует метрики и создает рекомендации по оптимизации
func (s *OptimizationService) analyzeMetrics(metrics []*models.Metric) []*models.OptimizationRecommendation {
	var recommendations []*models.OptimizationRecommendation

	// Анализ CPU метрик
	cpuRec := s.analyzeCPUMetrics(metrics)
	if cpuRec != nil {
		recommendations = append(recommendations, cpuRec)
	}

	// Анализ Memory метрик
	memRec := s.analyzeMemoryMetrics(metrics)
	if memRec != nil {
		recommendations = append(recommendations, memRec)
	}

	return recommendations
}

// analyzeCPUMetrics анализирует CPU метрики и создает рекомендации
func (s *OptimizationService) analyzeCPUMetrics(metrics []*models.Metric) *models.OptimizationRecommendation {
	var cpuMetrics []*models.Metric
	for _, m := range metrics {
		if m.Type == models.MetricTypeCPU {
			cpuMetrics = append(cpuMetrics, m)
		}
	}

	if len(cpuMetrics) == 0 {
		return nil
	}

	// Вычисляем среднее использование CPU
	var totalUsage float64
	for _, m := range cpuMetrics {
		totalUsage += m.Value
	}
	avgUsage := totalUsage / float64(len(cpuMetrics))

	// Если среднее использование CPU высокое, рекомендуем масштабирование
	if avgUsage > 80 {
		return &models.OptimizationRecommendation{
			Type:        models.OptimizationTypeScaling,
			Description: fmt.Sprintf("High CPU usage (%.2f%%). Consider scaling up CPU resources.", avgUsage),
			Impact:      "Improved performance and response times",
		}
	}

	return nil
}

// analyzeMemoryMetrics анализирует Memory метрики и создает рекомендации
func (s *OptimizationService) analyzeMemoryMetrics(metrics []*models.Metric) *models.OptimizationRecommendation {
	var memMetrics []*models.Metric
	for _, m := range metrics {
		if m.Type == models.MetricTypeMemory {
			memMetrics = append(memMetrics, m)
		}
	}

	if len(memMetrics) == 0 {
		return nil
	}

	// Вычисляем среднее использование памяти
	var totalUsage float64
	for _, m := range memMetrics {
		totalUsage += m.Value
	}
	avgUsage := totalUsage / float64(len(memMetrics))

	// Если среднее использование памяти высокое, рекомендуем увеличение
	if avgUsage > 85 {
		return &models.OptimizationRecommendation{
			Type:        models.OptimizationTypePerformance,
			Description: fmt.Sprintf("High memory usage (%.2f%%). Consider increasing memory allocation.", avgUsage),
			Impact:      "Reduced swapping and improved application performance",
		}
	}

	return nil
}

// GetServerRecommendations возвращает список рекомендаций для сервера
func (s *OptimizationService) GetServerRecommendations(ctx context.Context, serverID string) ([]*models.OptimizationRecommendation, error) {
	return s.optimizationRepo.GetByServer(ctx, serverID)
}

// CreateRecommendation создает новую рекомендацию
func (s *OptimizationService) CreateRecommendation(ctx context.Context, rec *models.OptimizationRecommendation) error {
	rec.ID = uuid.New().String()
	rec.CreatedAt = time.Now()
	rec.Status = models.OptimizationStatusPending
	return s.optimizationRepo.CreateRecommendation(ctx, rec)
}

// DeleteRecommendation удаляет рекомендацию
func (s *OptimizationService) DeleteRecommendation(ctx context.Context, id string) error {
	// Проверяем существование рекомендации перед удалением
	recs, err := s.optimizationRepo.GetByServer(ctx, id)
	if err != nil {
		return err
	}
	if len(recs) == 0 {
		return repository.ErrNotFound
	}

	return s.optimizationRepo.Delete(ctx, id)
}

// UpdateStatus обновляет статус рекомендации по оптимизации
func (s *OptimizationService) UpdateStatus(ctx context.Context, id string, status string) error {
	// Проверяем существование рекомендации
	recs, err := s.optimizationRepo.GetByServer(ctx, id)
	if err != nil {
		return err
	}
	if len(recs) == 0 {
		return repository.ErrNotFound
	}

	// Проверяем валидность статуса
	validStatuses := map[string]bool{
		models.OptimizationStatusPending:    true,
		models.OptimizationStatusApplied:    true,
		models.OptimizationStatusRejected:   true,
		models.OptimizationStatusInProgress: true,
		models.OptimizationStatusFailed:     true,
	}
	if !validStatuses[status] {
		return fmt.Errorf("invalid status: %s", status)
	}

	// Обновляем статус
	return s.optimizationRepo.UpdateStatus(ctx, id, status)
}
