package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Config 应用配置结构体
type Config struct {
	ServerPort   int    `mapstructure:"SERVER_PORT"`
	DatabaseURL  string `mapstructure:"DATABASE_URL"`
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	RedisURL     string `mapstructure:"REDIS_URL"`
}

// DB 全局数据库连接
var DB *gorm.DB

// LoadConfig 加载配置
func LoadConfig() (*Config, error) {
	// 加载.env文件
	if err := godotenv.Load(); err != nil {
		fmt.Println("Warning: .env file not found")
	}

	// 从环境变量读取配置
	cfg := &Config{
		ServerPort:   getEnvAsInt("SERVER_PORT", 8000),
		DatabaseURL:  getEnv("DATABASE_URL", "postgres://postgres:password@localhost:5432/trpg_sync?sslmode=disable"),
		JWTSecret:    getEnv("JWT_SECRET", "your-secret-key"),
		RedisURL:     getEnv("REDIS_URL", "redis://localhost:6379/0"),
	}

	return cfg, nil
}

// InitDB 初始化数据库连接
func InitDB(cfg *Config) error {
	var err error
	DB, err = gorm.Open(postgres.Open(cfg.DatabaseURL), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	fmt.Println("Database connection established")
	return nil
}

// 从环境变量获取字符串值
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// 从环境变量获取整数值
func getEnvAsInt(key string, defaultValue int) int {
	valueStr := getEnv(key, "")
	if valueStr == "" {
		return defaultValue
	}

	value := 0
	fmt.Sscanf(valueStr, "%d", &value)
	return value
}
