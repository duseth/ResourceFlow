package models

import "time"

// Metric представляет собой метрику сервера
type Metric struct {
	ID        string    `json:"id" db:"id"`
	ServerID  string    `json:"server_id" db:"server_id"`
	Type      string    `json:"type" db:"type"`
	Value     float64   `json:"value" db:"value"`
	Timestamp time.Time `json:"timestamp" db:"timestamp"`
}

// ResourceUsage представляет собой использование ресурсов сервера
type ResourceUsage struct {
	CPU struct {
		Usage       float64 `json:"usage"`       // Использование CPU в процентах
		Temperature float64 `json:"temperature"` // Температура CPU в градусах Цельсия
		Processes   int     `json:"processes"`   // Количество активных процессов
	} `json:"cpu"`

	Memory struct {
		Total     int64   `json:"total"`      // Общий объем памяти в байтах
		Used      int64   `json:"used"`       // Использованная память в байтах
		Free      int64   `json:"free"`       // Свободная память в байтах
		SwapUsage float64 `json:"swap_usage"` // Использование swap в процентах
	} `json:"memory"`

	Disk struct {
		Total       int64   `json:"total"`       // Общий объем диска в байтах
		Used        int64   `json:"used"`        // Использованное пространство в байтах
		Free        int64   `json:"free"`        // Свободное пространство в байтах
		IORead      float64 `json:"io_read"`     // Скорость чтения в байтах/сек
		IOWrite     float64 `json:"io_write"`    // Скорость записи в байтах/сек
		Utilization float64 `json:"utilization"` // Утилизация диска в процентах
	} `json:"disk"`

	Network struct {
		BytesReceived   int64   `json:"bytes_received"`   // Принято байт
		BytesSent       int64   `json:"bytes_sent"`       // Отправлено байт
		PacketsReceived int64   `json:"packets_received"` // Принято пакетов
		PacketsSent     int64   `json:"packets_sent"`     // Отправлено пакетов
		Bandwidth       float64 `json:"bandwidth"`        // Пропускная способность в байтах/сек
	} `json:"network"`
}

// MetricType определяет типы метрик
const (
	MetricTypeCPU     = "cpu"     // Использование CPU в процентах
	MetricTypeMemory  = "memory"  // Использование памяти в процентах
	MetricTypeDisk    = "disk"    // Использование диска в процентах
	MetricTypeNetwork = "network" // Сетевая активность в байтах/сек
)

// MetricThresholds определяет пороговые значения для метрик
const (
	CPUWarningThreshold     = 80.0 // Предупреждение при использовании CPU > 80%
	CPUCriticalThreshold    = 90.0 // Критическое при использовании CPU > 90%
	MemoryWarningThreshold  = 80.0 // Предупреждение при использовании памяти > 80%
	MemoryCriticalThreshold = 90.0 // Критическое при использовании памяти > 90%
	DiskWarningThreshold    = 85.0 // Предупреждение при использовании диска > 85%
	DiskCriticalThreshold   = 95.0 // Критическое при использовании диска > 95%
)

// MetricPeriods определяет периоды агрегации метрик
const (
	MetricPeriodHour  = "hour"
	MetricPeriodDay   = "day"
	MetricPeriodWeek  = "week"
	MetricPeriodMonth = "month"
)
