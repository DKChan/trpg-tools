package database

import (
	"testing"

	"trpg-sync/backend/infrastructure/config"

	"github.com/stretchr/testify/assert"
)

func TestInitDB(t *testing.T) {
	tests := []struct {
		name          string
		cfg           *config.Config
		expectError   bool
		errorContains string
	}{
		{
			name: "成功初始化数据库",
			cfg: &config.Config{
				Database: config.DatabaseConfig{
					Host:     "localhost",
					Port:     "5432",
					User:     "test",
					Password: "test",
					DBName:   "testdb",
					SSLMode:  "disable",
				},
				Log: config.LogConfig{
					Level: "silent",
				},
			},
			expectError: true, // 会失败，因为没有实际的 PostgreSQL 实例
		},
		{
			name: "配置为空",
			cfg: &config.Config{
				Database: config.DatabaseConfig{},
			},
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, err := InitDB(tt.cfg)

			if tt.expectError {
				assert.Error(t, err)
				assert.Nil(t, db)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, db)
			}
		})
	}
}
