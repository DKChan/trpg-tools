package main

import (
	"log"
	"trpg-sync/backend/infrastructure/config"
	"trpg-sync/backend/infrastructure/database"
	"trpg-sync/backend/api/v1"
	"trpg-sync/backend/api/middleware"
	"trpg-sync/backend/domain/room"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 自动同步表结构（只迁移 Room 表）
	if err := db.AutoMigrate(&room.Room{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	r := gin.Default()

	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())

	v1.SetupRoutes(r, db)

	log.Printf("Server starting on port %s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
