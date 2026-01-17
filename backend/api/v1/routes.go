package v1

import (
	"trpg-sync/backend/api/v1/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api/v1")
	{
		roomHandler := handlers.NewRoomHandler(db)
		room := api.Group("/rooms")
		{
			room.POST("", roomHandler.CreateRoom)
			room.GET("", roomHandler.GetRooms)
			room.GET("/:id", roomHandler.GetRoom)
			room.DELETE("/:id", roomHandler.DeleteRoom)
		}

		characterHandler := handlers.NewCharacterHandler(db)
		character := api.Group("/rooms/:roomId/characters")
		{
			character.POST("", characterHandler.CreateCharacter)
			character.GET("", characterHandler.GetCharacters)
			character.GET("/:id", characterHandler.GetCharacter)
			character.PUT("/:id", characterHandler.UpdateCharacter)
			character.DELETE("/:id", characterHandler.DeleteCharacter)
		}
	}
}
