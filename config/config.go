package config

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server     ServerConfig
	Database   DatabaseConfig
	JWT        JWTConfig
	CORS       CORSConfig
	Logging    LoggingConfig
	RateLimit  RateLimitConfig
	Pagination PaginationConfig
}

type ServerConfig struct {
	Port string
	Env  string
	Host string
}

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

type JWTConfig struct {
	Secret            string
	Expiration        time.Duration
	RefreshExpiration time.Duration
}

type CORSConfig struct {
	AllowedOrigins []string
}

type LoggingConfig struct {
	Level string
}

type RateLimitConfig struct {
	Requests int
	Duration string
}

type PaginationConfig struct {
	DefaultPageSize int
	MaxPageSize     int
}

// LoadConfig carga la configuración desde variables de entorno y archivos
func LoadConfig() (*Config, error) {
	// Configurar Viper para leer variables de entorno
	viper.AutomaticEnv()

	// Configurar archivo de configuración (opcional)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.AddConfigPath(".")

	// Leer archivo de configuración (si existe)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return nil, fmt.Errorf("error reading config file: %w", err)
		}
		// No hay problema si no existe el archivo, usaremos solo env vars
		log.Println("No config file found, using environment variables")
	}

	// Establecer valores por defecto
	setDefaults()

	// Mapear configuración a struct
	config := &Config{
		Server: ServerConfig{
			Port: viper.GetString("PORT"),
			Env:  viper.GetString("ENV"),
			Host: viper.GetString("HOST"),
		},
		Database: DatabaseConfig{
			Host:     viper.GetString("DB_HOST"),
			Port:     viper.GetString("DB_PORT"),
			User:     viper.GetString("DB_USER"),
			Password: viper.GetString("DB_PASSWORD"),
			DBName:   viper.GetString("DB_NAME"),
			SSLMode:  viper.GetString("DB_SSLMODE"),
		},
		JWT: JWTConfig{
			Secret:            viper.GetString("JWT_SECRET"),
			Expiration:        viper.GetDuration("JWT_EXPIRATION"),
			RefreshExpiration: viper.GetDuration("JWT_REFRESH_EXPIRATION"),
		},
		CORS: CORSConfig{
			AllowedOrigins: parseAllowedOrigins(),
		},
		Logging: LoggingConfig{
			Level: viper.GetString("LOG_LEVEL"),
		},
		RateLimit: RateLimitConfig{
			Requests: viper.GetInt("RATE_LIMIT_REQUESTS"),
			Duration: viper.GetString("RATE_LIMIT_DURATION"),
		},
		Pagination: PaginationConfig{
			DefaultPageSize: viper.GetInt("DEFAULT_PAGE_SIZE"),
			MaxPageSize:     viper.GetInt("MAX_PAGE_SIZE"),
		},
	}

	// Validar configuración crítica
	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("invalid configuration: %w", err)
	}

	return config, nil
}

// setDefaults establece valores por defecto
func setDefaults() {
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("ENV", "development")
	viper.SetDefault("HOST", "localhost")

	viper.SetDefault("DB_HOST", "localhost")
	viper.SetDefault("DB_PORT", "5432")
	viper.SetDefault("DB_SSLMODE", "disable")

	viper.SetDefault("JWT_EXPIRATION", "24h")
	viper.SetDefault("JWT_REFRESH_EXPIRATION", "168h")

	viper.SetDefault("LOG_LEVEL", "info")

	viper.SetDefault("RATE_LIMIT_REQUESTS", 100)
	viper.SetDefault("RATE_LIMIT_DURATION", "1m")

	viper.SetDefault("DEFAULT_PAGE_SIZE", 20)
	viper.SetDefault("MAX_PAGE_SIZE", 100)
}

// parseAllowedOrigins parsea ALLOWED_ORIGINS desde variable de entorno
// Maneja formato: "http://localhost:3000,http://localhost:5173"
func parseAllowedOrigins() []string {
	originsStr := viper.GetString("ALLOWED_ORIGINS")
	if originsStr == "" {
		return []string{}
	}

	// Split por coma y limpiar espacios
	origins := []string{}
	for _, origin := range strings.Split(originsStr, ",") {
		trimmed := strings.TrimSpace(origin)
		if trimmed != "" {
			origins = append(origins, trimmed)
		}
	}

	return origins
}

// validateConfig valida que la configuración tenga los valores críticos
func validateConfig(config *Config) error {
	if config.Database.Host == "" {
		return fmt.Errorf("DB_HOST is required")
	}
	if config.Database.User == "" {
		return fmt.Errorf("DB_USER is required")
	}
	if config.Database.DBName == "" {
		return fmt.Errorf("DB_NAME is required")
	}
	if config.JWT.Secret == "" {
		return fmt.Errorf("JWT_SECRET is required")
	}
	return nil
}

// GetDSN retorna el Data Source Name para PostgreSQL
func (c *DatabaseConfig) GetDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host, c.Port, c.User, c.Password, c.DBName, c.SSLMode,
	)
}
