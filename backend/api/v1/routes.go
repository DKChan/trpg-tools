package v1

import (
	"trpg-sync/backend/api/middleware"
	"trpg-sync/backend/api/v1/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		authHandler := handlers.NewAuthHandler(db)
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		userHandler := handlers.NewUserHandler(db)
		user := api.Group("/user")
		user.Use(middleware.AuthMiddleware())
		{
			user.GET("/profile", userHandler.GetProfile)
			user.PUT("/profile", userHandler.UpdateProfile)
			user.PUT("/password", userHandler.UpdatePassword)
		}

		roomHandler := handlers.NewRoomHandler(db)
		room := api.Group("/rooms")
		room.Use(middleware.AuthMiddleware())
		{
			room.POST("", roomHandler.CreateRoom)
			room.GET("", roomHandler.GetRooms)
			room.GET("/:id", roomHandler.GetRoom)
			room.POST("/:id/join", roomHandler.JoinRoom)
			room.POST("/:id/leave", roomHandler.LeaveRoom)
			room.PUT("/:id/members/:userId/kick", roomHandler.KickMember)
			room.PUT("/:id/transfer-dm", roomHandler.TransferDM)
			room.DELETE("/:id", roomHandler.DeleteRoom)
		}

characterHandler := handlers.NewCharacterHandler(db)
	character := api.Group("/rooms/:roomId/characters")
	character.Use(middleware.AuthMiddleware())
	{
		character.POST("", characterHandler.CreateCharacter)
		character.GET("", characterHandler.GetCharacters)
		character.GET("/:id", characterHandler.GetCharacter)
		character.PUT("/:id", characterHandler.UpdateCharacter)
		character.DELETE("/:id", characterHandler.DeleteCharacter)
	}

	room.GET("/:id/members", roomHandler.GetRoomMembers)
}
}
