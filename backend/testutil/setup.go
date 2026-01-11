package testutil

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SetupTestDB 创建测试用的内存 SQLite 数据库
func SetupTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("Failed to setup test database: %v", err)
	}

	// 获取底层 SQL DB 以便在测试后关闭
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to get database instance: %v", err)
	}
	t.Cleanup(func() {
		sqlDB.Close()
	})

	return db
}

// CleanupTestDB 清理测试数据
func CleanupTestDB(db *gorm.DB, tables []interface{}) error {
	for _, table := range tables {
		if err := db.Exec("DELETE FROM " + db.Statement.Table).Error; err != nil {
			return err
		}
	}
	return nil
}
