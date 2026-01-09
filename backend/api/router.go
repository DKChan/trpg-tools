package api

import (
	"github.com/gin-gonic/gin"
	"trpg-sync/backend/api/middleware"
	"trpg-sync/backend/api/v1"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	// 创建Gin引擎
	router := gin.Default()

	// 添加中间件
	router.Use(middleware.CORS())
	router.Use(middleware.Logger())

	// API路由组
	api := router.Group("/api/v1")
	{
		// 健康检查
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})

		// 认证路由
		auth := api.Group("/auth")
		{
			auth.POST("/register", v1.Register)
			auth.POST("/login", v1.Login)
			auth.POST("/logout", v1.Logout)
			auth.POST("/reset-password", v1.ResetPassword)
		}

		// 用户路由
		users := api.Group("/users")
		users.Use(middleware.Auth())
		{
			users.GET("/me", v1.GetCurrentUser)
			users.PUT("/me", v1.UpdateCurrentUser)
		}

		// 房间路由
		rooms := api.Group("/rooms")
		rooms.Use(middleware.Auth())
		{
			rooms.POST("", v1.CreateRoom)
			rooms.GET("", v1.GetRooms)
			rooms.GET("/:id", v1.GetRoom)
			rooms.PUT("/:id", v1.UpdateRoom)
			rooms.DELETE("/:id", v1.DeleteRoom)
			rooms.POST("/:id/join", v1.JoinRoom)
			rooms.POST("/:id/leave", v1.LeaveRoom)
			rooms.POST("/:id/kick", v1.KickPlayer)
		}

		// 人物卡路由
		characterSheets := api.Group("/character-sheets")
		characterSheets.Use(middleware.Auth())
		{
			characterSheets.POST("", v1.CreateCharacterSheet)
			characterSheets.GET("", v1.GetCharacterSheets)
			characterSheets.GET("/:id", v1.GetCharacterSheet)
			characterSheets.PUT("/:id", v1.UpdateCharacterSheet)
			characterSheets.DELETE("/:id", v1.DeleteCharacterSheet)
			characterSheets.GET("/:id/versions", v1.GetCharacterSheetVersions)
		}
	}

	return router
}
