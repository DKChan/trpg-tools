package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"trpg-sync/backend/api/middleware"
	"trpg-sync/backend/api/v1"
	"trpg-sync/backend/domain/room"
	"trpg-sync/backend/infrastructure/config"
	"trpg-sync/backend/infrastructure/database"

	"github.com/gin-gonic/gin"
)

//go:embed all:dist
var frontendFS embed.FS

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

	// 注册API路由
	v1.SetupRoutes(r, db)

	// 获取前端静态文件系统
	distFS, err := fs.Sub(frontendFS, "dist")
	if err != nil {
		log.Fatalf("Failed to create sub filesystem: %v", err)
	}

	// 静态文件服务
	fileServer := http.FileServer(http.FS(distFS))
	r.GET("/assets/*filepath", func(c *gin.Context) {
		fileServer.ServeHTTP(c.Writer, c.Request)
	})

	// 前端路由处理
	r.NoRoute(func(c *gin.Context) {
		c.Request.URL.Path = "/"
		fileServer.ServeHTTP(c.Writer, c.Request)
	})

	log.Printf("Server starting on port %s", cfg.Server.Port)
	log.Printf("Frontend: http://localhost:%s", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
