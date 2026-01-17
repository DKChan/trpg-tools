package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"
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
	MaxPlayers  int    `json:"max_players" binding:"min=1,max=100"`
	IsPublic    bool   `json:"is_public"`
}

// generateInviteCode 生成8位随机邀请码
func generateInviteCode() string {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		// 如果随机数生成失败，使用时间戳作为备选
		return strings.ToUpper(fmt.Sprintf("%08X", time.Now().Unix())[:8])
	}
	return strings.ToUpper(hex.EncodeToString(bytes)[:8])
}

// GenerateInviteCode 导出的邀请码生成函数，用于测试
func GenerateInviteCode() string {
	return generateInviteCode()
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
		InviteCode:  generateInviteCode(),
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
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    nil,
		})
		return
	}

	roomIDStr := c.Param("id")

	// 验证房间存在
	var targetRoom room.Room
	if err := h.db.First(&targetRoom, roomIDStr).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Room not found",
			"data":    nil,
		})
		return
	}

	// 验证房间是否已满
	var memberCount int64
	h.db.Model(&room.RoomMember{}).Where("room_id = ?", targetRoom.ID).Count(&memberCount)
	if memberCount >= int64(targetRoom.MaxPlayers) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Room is full",
			"data":    nil,
		})
		return
	}

	// 验证用户是否已在房间中
	var existingMember room.RoomMember
	err := h.db.Where("room_id = ? AND user_id = ?", targetRoom.ID, userID.(uint)).First(&existingMember).Error
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Already in room",
			"data":    nil,
		})
		return
	}

	// 创建房间成员记录
	newMember := room.RoomMember{
		RoomID:   targetRoom.ID,
		UserID:   userID.(uint),
		Role:     "player",
		JoinedAt: time.Now(),
	}

	if err := h.db.Create(&newMember).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to join room",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Joined room successfully",
		"data":    newMember,
	})
}

func (h *RoomHandler) LeaveRoom(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    nil,
		})
		return
	}

	roomIDStr := c.Param("id")

	// 验证房间存在
	var targetRoom room.Room
	if err := h.db.First(&targetRoom, roomIDStr).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Room not found",
			"data":    nil,
		})
		return
	}

	// DM不能离开房间（需要删除房间或转让DM权限）
	if targetRoom.DMID == userID.(uint) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "DM cannot leave room without deleting or transferring",
			"data":    nil,
		})
		return
	}

	// 验证用户是否在房间中
	var member room.RoomMember
	if err := h.db.Where("room_id = ? AND user_id = ?", targetRoom.ID, userID.(uint)).First(&member).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Not in room",
			"data":    nil,
		})
		return
	}

	// 删除房间成员记录
	if err := h.db.Delete(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to leave room",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Left room successfully",
		"data":    nil,
	})
}

func (h *RoomHandler) DeleteRoom(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    nil,
		})
		return
	}

	roomIDStr := c.Param("id")

	// 验证房间存在
	var targetRoom room.Room
	if err := h.db.First(&targetRoom, roomIDStr).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Room not found",
			"data":    nil,
		})
		return
	}

	// 验证用户是DM
	if targetRoom.DMID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "Only DM can delete room",
			"data":    nil,
		})
		return
	}

	// 删除房间（GORM会级联删除关联的RoomMember）
	if err := h.db.Delete(&targetRoom).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to delete room",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Room deleted successfully",
		"data":    nil,
	})
}

func (h *RoomHandler) KickMember(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    nil,
		})
		return
	}

	roomIDStr := c.Param("id")
	targetUserIDStr := c.Param("userId")

	// 验证房间存在
	var targetRoom room.Room
	if err := h.db.First(&targetRoom, roomIDStr).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Room not found",
			"data":    nil,
		})
		return
	}

	// 验证用户是DM
	if targetRoom.DMID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "Only DM can kick members",
			"data":    nil,
		})
		return
	}

	// 不能踢出DM（DM不能自己）
	if targetRoom.DMID == userID.(uint) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Cannot kick DM",
			"data":    nil,
		})
		return
	}

	// 验证目标用户在房间中
	var member room.RoomMember
	if err := h.db.Where("room_id = ? AND user_id = ?", targetRoom.ID, targetUserIDStr).First(&member).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Member not found in room",
			"data":    nil,
		})
		return
	}

	// 删除房间成员
	if err := h.db.Delete(&member).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to kick member",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Member kicked successfully",
		"data":    nil,
	})
}

func (h *RoomHandler) TransferDM(c *gin.Context) {
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "Unauthorized",
			"data":    nil,
		})
		return
	}

	roomIDStr := c.Param("id")

	// 验证房间存在
	var targetRoom room.Room
	if err := h.db.First(&targetRoom, roomIDStr).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Room not found",
			"data":    nil,
		})
		return
	}

	// 验证用户是DM
	if targetRoom.DMID != userID.(uint) {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    403,
			"message": "Only DM can transfer ownership",
			"data":    nil,
		})
		return
	}

	// 获取要转让给的用户ID（从请求体获取）
	var req struct {
		UserID uint `json:"user_id" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid request",
			"data":    nil,
		})
		return
	}

	// 验证目标用户在房间中
	var targetMember room.RoomMember
	if err := h.db.Where("room_id = ? AND user_id = ?", targetRoom.ID, req.UserID).First(&targetMember).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Target user not found in room",
			"data":    nil,
		})
		return
	}

	// 更新房间的DMID
	targetRoom.DMID = req.UserID
	if err := h.db.Save(&targetRoom).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to transfer DM",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "DM transferred successfully",
		"data":    targetRoom,
	})
}

func (h *RoomHandler) GetRoomMembers(c *gin.Context) {
	roomIDStr := c.Param("id")

	// 验证房间存在
	var targetRoom room.Room
	if err := h.db.First(&targetRoom, roomIDStr).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "Room not found",
			"data":    nil,
		})
		return
	}

	// 获取房间成员列表
	var members []room.RoomMember
	if err := h.db.Preload("User").Where("room_id = ?", targetRoom.ID).Find(&members).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to get room members",
			"data":    nil,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Success",
		"data":    members,
	})
}



