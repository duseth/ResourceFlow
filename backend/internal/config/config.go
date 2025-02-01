package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

// Config содержит конфигурацию приложения
type Config struct {
	App      AppConfig      `yaml:"app"`
	HTTP     HTTPConfig     `yaml:"http"`
	Redis    RedisConfig    `yaml:"redis"`
	JWT      JWTConfig      `yaml:"jwt"`
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
	SSH      SSHConfig      `yaml:"ssh"`
}

// AppConfig содержит общие настройки приложения
type AppConfig struct {
	Name    string `yaml:"name"`
	Version string `yaml:"version"`
	Debug   bool   `yaml:"debug"`
	Env     string `yaml:"env"`
}

// HTTPConfig содержит настройки HTTP сервера
type HTTPConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

// RedisConfig содержит настройки Redis
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// JWTConfig содержит настройки JWT
type JWTConfig struct {
	Secret     string `yaml:"secret"`
	Expiration string `yaml:"expiration"`
}

// ServerConfig содержит настройки сервера
type ServerConfig struct {
	Port int    `yaml:"port"`
	Host string `yaml:"host"`
}

// Address возвращает адрес сервера в формате host:port
func (c *ServerConfig) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

// DatabaseConfig содержит настройки базы данных
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

// DSN возвращает строку подключения к базе данных
func (c *DatabaseConfig) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode)
}

// SSHConfig содержит настройки SSH
type SSHConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Timeout  int    `yaml:"timeout"` // в секундах
}

// Load загружает конфигурацию из переменных окружения
func Load() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	viper.SetDefault("app.name", "ResourceFlow")
	viper.SetDefault("app.version", "1.0.0")
	viper.SetDefault("app.debug", true)
	viper.SetDefault("app.env", "development")

	viper.SetDefault("http.host", "0.0.0.0")
	viper.SetDefault("http.port", "8080")

	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", 5432)
	viper.SetDefault("database.sslmode", "disable")

	viper.SetDefault("redis.host", "localhost")
	viper.SetDefault("redis.port", "6379")
	viper.SetDefault("redis.db", 0)

	viper.SetDefault("ssh.user", "monitoring")
	viper.SetDefault("ssh.timeout", 5)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %v", err)
		}
		// Если файл не найден, используем значения по умолчанию
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("RF")

	// Заменяем точки на подчеркивания в именах переменных окружения
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
