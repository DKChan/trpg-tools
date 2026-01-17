package handlers

import (
	"net/http"
	"strconv"
	"trpg-sync/backend/domain/character"
	"trpg-sync/backend/domain/room"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CharacterHandler struct {
	db *gorm.DB
}

func NewCharacterHandler(db *gorm.DB) *CharacterHandler {
	return &CharacterHandler{db: db}
}

type CreateCharacterRequest struct {
	Name         string `json:"name" binding:"required"`
	Race         string `json:"race"`
	Class        string `json:"class"`
	Level        int    `json:"level"`
	Background   string `json:"background"`
	Alignment    string `json:"alignment"`
	Strength     int    `json:"strength"`
	Dexterity    int    `json:"dexterity"`
	Constitution int    `json:"constitution"`
	Intelligence int    `json:"intelligence"`
	Wisdom       int    `json:"wisdom"`
	Charisma     int    `json:"charisma"`
}

func (h *CharacterHandler) CreateCharacter(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    nil,
		})
		return
	}

	roomIDStr := c.Param("roomId")
	roomID, err := strconv.ParseUint(roomIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid room ID",
			"data":    nil,
		})
		return
	}

	// 验证房间是否存在
	var roomExists bool
	if err := h.db.Model(&room.Room{}).Select("count(*) > 0").Where("id = ?", roomID).Find(&roomExists).Error; err != nil || !roomExists {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Room not found",
			"data":    nil,
		})
		return
	}

	// 验证用户是否在房间中
	var memberExists bool
	if err := h.db.Model(&room.RoomMember{}).Select("count(*) > 0").Where("room_id = ? AND user_id = ?", roomID, userID).Find(&memberExists).Error; err != nil || !memberExists {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "You are not a member of this room",
			"data":    nil,
		})
		return
	}

	var req CreateCharacterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	newCharacter := character.CharacterCard{
		UserID:       userID.(uint),
		RoomID:       uint(roomID),
		Name:         req.Name,
		Race:         req.Race,
		Class:        req.Class,
		Level:        req.Level,
		Background:   req.Background,
		Alignment:    req.Alignment,
		Strength:     req.Strength,
		Dexterity:    req.Dexterity,
		Constitution: req.Constitution,
		Intelligence: req.Intelligence,
		Wisdom:       req.Wisdom,
		Charisma:     req.Charisma,
	}

	if err := h.db.Create(&newCharacter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to create character",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Character created successfully",
		"data":    newCharacter,
	})
}

func (h *CharacterHandler) GetCharacters(c *gin.Context) {
	roomID := c.Param("roomId")

	var characters []character.CharacterCard
	if err := h.db.Where("room_id = ?", roomID).Find(&characters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get characters",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success",
		"data":    characters,
	})
}

func (h *CharacterHandler) GetCharacter(c *gin.Context) {
	characterID := c.Param("id")

	var character character.CharacterCard
	if err := h.db.First(&character, characterID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Character not found",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success",
		"data":    character,
	})
}

func (h *CharacterHandler) UpdateCharacter(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    nil,
		})
		return
	}

	characterIDStr := c.Param("id")

	// 获取要更新的人物卡
	var targetCharacter character.CharacterCard
	if err := h.db.First(&targetCharacter, characterIDStr).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Character not found",
			"data":    nil,
		})
		return
	}

	// 验证所有权
	if targetCharacter.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "Only character owner can update",
			"data":    nil,
		})
		return
	}

	// 解析请求体
	var req CreateCharacterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// 更新字段
	targetCharacter.Name = req.Name
	targetCharacter.Race = req.Race
	targetCharacter.Class = req.Class
	targetCharacter.Level = req.Level
	targetCharacter.Background = req.Background
	targetCharacter.Alignment = req.Alignment
	targetCharacter.Strength = req.Strength
	targetCharacter.Dexterity = req.Dexterity
	targetCharacter.Constitution = req.Constitution
	targetCharacter.Intelligence = req.Intelligence
	targetCharacter.Wisdom = req.Wisdom
	targetCharacter.Charisma = req.Charisma

	if err := h.db.Save(&targetCharacter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to update character",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Character updated successfully",
		"data":    targetCharacter,
	})
}

func (h *CharacterHandler) DeleteCharacter(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    nil,
		})
		return
	}

	characterIDStr := c.Param("id")

	// 获取要删除的人物卡
	var targetCharacter character.CharacterCard
	if err := h.db.First(&targetCharacter, characterIDStr).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Character not found",
			"data":    nil,
		})
		return
	}

	// 验证所有权
	if targetCharacter.UserID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "Only character owner can delete",
			"data":    nil,
		})
		return
	}

	// 删除人物卡
	if err := h.db.Delete(&targetCharacter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete character",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Character deleted successfully",
		"data":    nil,
	})
}
