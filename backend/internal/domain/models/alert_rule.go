package models

// AlertRule представляет правило для генерации алертов
type AlertRule struct {
	ID         string  `json:"id" db:"id"`                   // Уникальный идентификатор правила
	ServerID   string  `json:"server_id" db:"server_id"`     // ID сервера
	MetricType string  `json:"metric_type" db:"metric_type"` // Тип метрики
	Condition  string  `json:"condition" db:"condition"`     // Условие (>, <, =, etc.)
	Threshold  float64 `json:"threshold" db:"threshold"`     // Пороговое значение
	Duration   int     `json:"duration" db:"duration"`       // Длительность в секундах
}

// AlertRuleCondition определяет возможные условия для правил
const (
	AlertRuleConditionGreaterThan = ">"
	AlertRuleConditionLessThan    = "<"
	AlertRuleConditionEqual       = "="
)
