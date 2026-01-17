package handlers

import (
	"net/http"
	"strconv"
	"trpg-sync/backend/domain/character"
	"trpg-sync/backend/infrastructure/storage"

	"github.com/gin-gonic/gin"
)

type CharacterHandler struct {
	storage *storage.CharacterStorage
}

func NewCharacterHandler() *CharacterHandler {
	return &CharacterHandler{
		storage: storage.NewCharacterStorage(),
	}
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
	AC           int    `json:"ac"`
	HP           int    `json:"hp"`
	MaxHP        int    `json:"max_hp"`
	Speed        int    `json:"speed"`
	Proficiency  int    `json:"proficiency"`
	Skills       string `json:"skills"`
	Saves        string `json:"saves"`
	Equipment    string `json:"equipment"`
	Spells       string `json:"spells"`
}

func (h *CharacterHandler) CreateCharacter(c *gin.Context) {
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

	var req CreateCharacterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": err.Error(),
			"data":    nil,
		})
		return
	}

	// 生成新 ID
	charID, err := h.storage.GenerateNextID(uint(roomID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to generate character ID",
			"data":    nil,
		})
		return
	}

	// 创建人物卡
	newCharacter := &character.CharacterCard{
		ID:           charID,
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
		AC:           req.AC,
		HP:           req.HP,
		MaxHP:        req.MaxHP,
		Speed:        req.Speed,
		Proficiency:  req.Proficiency,
		Skills:       req.Skills,
		Saves:        req.Saves,
		Equipment:    req.Equipment,
		Spells:       req.Spells,
	}

	if err := h.storage.SaveCharacter(newCharacter); err != nil {
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

	characters, err := h.storage.GetRoomCharacters(uint(roomID))
	if err != nil {
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

	characterIDStr := c.Param("id")
	characterID, err := strconv.ParseUint(characterIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid character ID",
			"data":    nil,
		})
		return
	}

	char, err := h.storage.LoadCharacter(uint(roomID), uint(characterID))
	if err != nil {
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
		"data":    char,
	})
}

func (h *CharacterHandler) UpdateCharacter(c *gin.Context) {
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

	characterIDStr := c.Param("id")
	characterID, err := strconv.ParseUint(characterIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid character ID",
			"data":    nil,
		})
		return
	}

	// 加载现有人物卡
	targetCharacter, err := h.storage.LoadCharacter(uint(roomID), uint(characterID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Character not found",
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
	targetCharacter.AC = req.AC
	targetCharacter.HP = req.HP
	targetCharacter.MaxHP = req.MaxHP
	targetCharacter.Speed = req.Speed
	targetCharacter.Proficiency = req.Proficiency
	targetCharacter.Skills = req.Skills
	targetCharacter.Saves = req.Saves
	targetCharacter.Equipment = req.Equipment
	targetCharacter.Spells = req.Spells

	if err := h.storage.SaveCharacter(targetCharacter); err != nil {
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

	characterIDStr := c.Param("id")
	characterID, err := strconv.ParseUint(characterIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid character ID",
			"data":    nil,
		})
		return
	}

	if err := h.storage.DeleteCharacter(uint(roomID), uint(characterID)); err != nil {
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
