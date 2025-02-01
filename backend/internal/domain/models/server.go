package models

import "time"

// Server представляет собой сервер для мониторинга
type Server struct {
	ID          string    `json:"id" db:"id"`                       // Уникальный идентификатор сервера
	Name        string    `json:"name" db:"name"`                   // Имя сервера
	Host        string    `json:"host" db:"host"`                   // Хост или IP-адрес сервера
	Port        int       `json:"port" db:"port"`                   // SSH порт для подключения
	Status      string    `json:"status" db:"status"`               // Текущий статус сервера
	Tags        []string  `json:"tags" db:"tags"`                   // Теги для группировки серверов
	CreatedAt   time.Time `json:"created_at" db:"created_at"`       // Время создания записи
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`       // Время последнего обновления
	LastCheckAt time.Time `json:"last_check_at" db:"last_check_at"` // Время последней проверки
}

// ServerStatus определяет возможные статусы сервера
const (
	ServerStatusActive   = "active"   // Сервер активен и доступен
	ServerStatusInactive = "inactive" // Сервер неактивен
	ServerStatusError    = "error"    // Ошибка при подключении к серверу
)
