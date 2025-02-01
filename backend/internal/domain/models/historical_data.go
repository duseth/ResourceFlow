package models

import (
	"time"
)

// HistoricalData представляет агрегированные метрики за период
type HistoricalData struct {
	ID         string    `json:"id" db:"id"`                   // Уникальный идентификатор записи
	ServerID   string    `json:"server_id" db:"server_id"`     // ID сервера
	MetricType string    `json:"metric_type" db:"metric_type"` // Тип метрики
	Period     string    `json:"period" db:"period"`           // Период агрегации
	MinValue   float64   `json:"min_value" db:"min_value"`     // Минимальное значение за период
	MaxValue   float64   `json:"max_value" db:"max_value"`     // Максимальное значение за период
	AvgValue   float64   `json:"avg_value" db:"avg_value"`     // Среднее значение за период
	Timestamp  time.Time `json:"timestamp" db:"timestamp"`     // Время начала периода
}

// HistoricalDataPeriod определяет периоды агрегации
const (
	HistoricalDataPeriodHour  = "hour"  // Почасовая агрегация
	HistoricalDataPeriodDay   = "day"   // Ежедневная агрегация
	HistoricalDataPeriodWeek  = "week"  // Еженедельная агрегация
	HistoricalDataPeriodMonth = "month" // Ежемесячная агрегация
)
