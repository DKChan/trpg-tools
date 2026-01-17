package config

import (
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type Config struct {
	Server ServerConfig
	Database DatabaseConfig
	CORS    CORSConfig
	Log     LogConfig
}

type ServerConfig struct {
	Port string
	Mode string
}

type DatabaseConfig struct {
	Path string
}

type CORSConfig struct {
	AllowedOrigins []string
}

type LogConfig struct {
	Level  string
	Format string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	cfg := &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "8080"),
			Mode: getEnv("SERVER_MODE", "development"),
		},
		Database: DatabaseConfig{
			Path: getEnv("DB_PATH", "./sqlite.db"),
		},
		CORS: CORSConfig{
			AllowedOrigins: []string{
				getEnv("CORS_ALLOWED_ORIGINS", "http://localhost:5173,http://localhost:3000"),
			},
		},
		Log: LogConfig{
			Level:  getEnv("LOG_LEVEL", "debug"),
			Format: getEnv("LOG_FORMAT", "json"),
		},
	}

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvInt(key string, defaultValue int) int {
	if value := os.Getenv(key); value != "" {
		if intVal, err := strconv.Atoi(value); err == nil {
			return intVal
		}
	}
	return defaultValue
}
