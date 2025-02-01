package models

import "time"

// Alert представляет собой оповещение о проблеме
type Alert struct {
	ID         string     `json:"id" db:"id"`                   // Уникальный идентификатор алерта
	ServerID   string     `json:"server_id" db:"server_id"`     // ID сервера
	RuleID     string     `json:"rule_id" db:"rule_id"`         // ID правила, вызвавшего алерт
	Status     string     `json:"status" db:"status"`           // Статус алерта
	Message    string     `json:"message" db:"message"`         // Сообщение алерта
	CreatedAt  time.Time  `json:"created_at" db:"created_at"`   // Время создания
	ResolvedAt *time.Time `json:"resolved_at" db:"resolved_at"` // Время разрешения
}

// AlertType определяет типы алертов
const (
	AlertTypeWarning  = "warning"
	AlertTypeCritical = "critical"
	AlertTypeError    = "error"
)

// AlertStatus определяет возможные статусы алерта
const (
	AlertStatusActive       = "active"       // Активный алерт
	AlertStatusResolved     = "resolved"     // Разрешенный алерт
	AlertStatusAcknowledged = "acknowledged" // Подтвержденный алерт
)
