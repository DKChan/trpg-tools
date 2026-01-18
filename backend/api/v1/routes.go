package v1

import (
	"trpg-sync/backend/api/v1/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	api := r.Group("/api/v1")

	// 房间路由
	roomHandler := handlers.NewRoomHandler(db)
	api.POST("/rooms", roomHandler.CreateRoom)
	api.GET("/rooms", roomHandler.GetRooms)
	api.GET("/rooms/:id", roomHandler.GetRoom)
	api.DELETE("/rooms/:id", roomHandler.DeleteRoom)

	// 人物卡路由 - 使用独立路径避免Gin路由冲突
	characterHandler := handlers.NewCharacterHandler()
	api.POST("/characters/:roomId", characterHandler.CreateCharacter)
	api.GET("/characters/:roomId", characterHandler.GetCharacters)
	api.GET("/characters/:roomId/:charId", characterHandler.GetCharacter)
	api.PUT("/characters/:roomId/:charId", characterHandler.UpdateCharacter)
	api.DELETE("/characters/:roomId/:charId", characterHandler.DeleteCharacter)
}
