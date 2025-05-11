package repository

import (
	"context"
	"time"

	"github.com/duseth/ResourceFlow/internal/domain/models"
)

// ServerFilter содержит параметры фильтрации серверов
type ServerFilter struct {
	Status string
	Tags   []string
	Search string
}

// ServerRepository определяет методы для работы с серверами
type ServerRepository interface {
	Create(ctx context.Context, server *models.Server) error
	Update(ctx context.Context, server *models.Server) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*models.Server, error)
	List(ctx context.Context, filter ServerFilter) ([]*models.Server, error)
}

// MetricRepository определяет методы для работы с метриками
type MetricRepository interface {
	Store(ctx context.Context, metric *models.Metric) error
	GetLatest(ctx context.Context, serverID string, metricType string) (*models.Metric, error)
	GetRange(ctx context.Context, serverID string, from, to time.Time) ([]*models.Metric, error)
	GetAggregated(ctx context.Context, serverID string, period string) ([]*models.HistoricalData, error)
}

// AlertRepository определяет методы для работы с алертами
type AlertRepository interface {
	Create(ctx context.Context, alert *models.Alert) error
	Update(ctx context.Context, alert *models.Alert) error
	GetByID(ctx context.Context, id string) (*models.Alert, error)
	GetActive(ctx context.Context) ([]*models.Alert, error)
	GetByServer(ctx context.Context, serverID string) ([]*models.Alert, error)
	GetRules(ctx context.Context, serverID string) ([]*models.AlertRule, error)
	CreateRule(ctx context.Context, rule *models.AlertRule) error
	UpdateRule(ctx context.Context, rule *models.AlertRule) error
	DeleteRule(ctx context.Context, id string) error
}

// OptimizationRepository определяет методы для работы с оптимизациями
type OptimizationRepository interface {
	CreateRecommendation(ctx context.Context, rec *models.OptimizationRecommendation) error
	UpdateStatus(ctx context.Context, id string, status string) error
	GetPending(ctx context.Context) ([]*models.OptimizationRecommendation, error)
	GetByServer(ctx context.Context, serverID string) ([]*models.OptimizationRecommendation, error)
	Delete(ctx context.Context, id string) error
}
