package models

import "time"

// OptimizationRecommendation представляет собой рекомендацию по оптимизации
type OptimizationRecommendation struct {
	ID          string     `json:"id" db:"id"`                           // Уникальный идентификатор рекомендации
	ServerID    string     `json:"server_id" db:"server_id"`             // ID сервера
	Type        string     `json:"type" db:"type"`                       // Тип оптимизации
	Description string     `json:"description" db:"description"`         // Описание рекомендации
	Impact      string     `json:"impact" db:"impact"`                   // Ожидаемый эффект
	Status      string     `json:"status" db:"status"`                   // Статус рекомендации
	CreatedAt   time.Time  `json:"created_at" db:"created_at"`           // Время создания
	AppliedAt   *time.Time `json:"applied_at,omitempty" db:"applied_at"` // Время применения
}

// OptimizationType определяет типы оптимизаций
const (
	OptimizationTypeScaling     = "scaling"     // Масштабирование ресурсов
	OptimizationTypePerformance = "performance" // Оптимизация производительности
	OptimizationTypeCost        = "cost"        // Оптимизация затрат
)

// OptimizationPriority определяет приоритеты оптимизаций
const (
	OptimizationPriorityLow    = "low"
	OptimizationPriorityMedium = "medium"
	OptimizationPriorityHigh   = "high"
)

// OptimizationStatus определяет статусы оптимизаций
const (
	OptimizationStatusPending    = "pending"     // Ожидает применения
	OptimizationStatusApplied    = "applied"     // Применена
	OptimizationStatusRejected   = "rejected"    // Отклонена
	OptimizationStatusInProgress = "in_progress" // В процессе применения
	OptimizationStatusFailed     = "failed"      // Ошибка при применении
)
