package handlers

import (
	"net/http"
	"trpg-sync/backend/domain/character"

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
	userID, _ := c.Get("user_id")
	roomID := c.Param("roomId")

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
		RoomID:       0,
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
	characterID := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Character updated successfully",
		"data":    nil,
	})
}

func (h *CharacterHandler) DeleteCharacter(c *gin.Context) {
	characterID := c.Param("id")

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Character deleted successfully",
		"data":    nil,
	})
}
