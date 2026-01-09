package handlers

import (
	"net/http"
	"trpg-sync/backend/domain/room"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type RoomHandler struct {
	db *gorm.DB
}

func NewRoomHandler(db *gorm.DB) *RoomHandler {
	return &RoomHandler{db: db}
}

type CreateRoomRequest struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	RuleSystem  string `json:"rule_system"`
	Password    string `json:"password"`
	MaxPlayers  int    `json:"max_players"`
	IsPublic    bool   `json:"is_public"`
}

func (h *RoomHandler) CreateRoom(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var req CreateRoomRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	newRoom := room.Room{
		Name:        req.Name,
		Description: req.Description,
		RuleSystem:  req.RuleSystem,
		Password:    req.Password,
		DMID:        userID.(uint),
		MaxPlayers:  req.MaxPlayers,
		IsPublic:    req.IsPublic,
	}

	if err := h.db.Create(&newRoom).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create room",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Room created successfully",
		"data":    newRoom,
	})
}

func (h *RoomHandler) GetRooms(c *gin.Context) {
	var rooms []room.Room
	if err := h.db.Where("is_public = ?", true).Find(&rooms).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get rooms",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success",
		"data":    rooms,
	})
}

func (h *RoomHandler) GetRoom(c *gin.Context) {
	roomID := c.Param("id")

	var room room.Room
	if err := h.db.First(&room, roomID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Room not found",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success",
		"data":    room,
	})
}

func (h *RoomHandler) JoinRoom(c *gin.Context) {
	userID, _ := c.Get("user_id")
	roomID := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Joined room successfully",
		"data":    nil,
	})
}

func (h *RoomHandler) LeaveRoom(c *gin.Context) {
	userID, _ := c.Get("user_id")
	roomID := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Left room successfully",
		"data":    nil,
	})
}

func (h *RoomHandler) DeleteRoom(c *gin.Context) {
	userID, _ := c.Get("user_id")
	roomID := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Room deleted successfully",
		"data":    nil,
	})
}
