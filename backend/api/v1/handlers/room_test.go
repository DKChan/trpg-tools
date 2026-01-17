package handlers

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	"trpg-sync/backend/domain/room"
	"trpg-sync/backend/domain/user"
	"trpg-sync/backend/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerateInviteCode(t *testing.T) {
	// 测试生成多个邀请码，验证格式和唯一性
	codes := make(map[string]bool)
	for i := 0; i < 100; i++ {
		code := GenerateInviteCode()
		require.Len(t, code, 8, "邀请码必须是8位")
		assert.Regexp(t, `^[A-Z0-9]{8}$`, code, "邀请码必须是大写字母和数字")

		// 验证唯一性（实际使用中会有冲突处理，这里简单验证）
		if !codes[code] {
			codes[code] = true
		}
	}

	// 至少应该生成95个以上的唯一邀请码（考虑到可能有极低概率的重复）
	assert.GreaterOrEqual(t, len(codes), 95, "应该生成大量唯一的邀请码")
}

func TestRoomHandler_CreateRoom(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&room.Room{}, &user.User{})

	handler := NewRoomHandler(db)
	router := testutil.SetupTestRouter()
	router.POST("/rooms", handler.CreateRoom)

	tests := []struct {
		name           string
		requestBody    string
		userID         uint
		expectedStatus int
	}{
		{
			name: "成功创建公开房间",
			requestBody: `{
				"name": "Test Room",
				"description": "A test room",
				"rule_system": "DND5e",
				"max_players": 5,
				"is_public": true
			}`,
			userID:         1,
			expectedStatus: 200,
		},
		{
			name: "成功创建私有房间",
			requestBody: `{
				"name": "Private Room",
				"description": "A private room",
				"rule_system": "DND5e",
				"password": "secret123",
				"max_players": 3,
				"is_public": false
			}`,
			userID:         1,
			expectedStatus: 200,
		},
		{
			name: "房间名称为空",
			requestBody: `{
				"description": "A test room",
				"rule_system": "DND5e"
			}`,
			userID:         1,
			expectedStatus: 400,
		},
		{
			name: "max_players 为负数",
			requestBody: `{
				"name": "Test Room",
				"max_players": -1
			}`,
			userID:         1,
			expectedStatus: 400,
		},
		{
			name: "成功创建房间并自动生成邀请码",
			requestBody: `{
				"name": "Room with Invite",
				"description": "Test invite code",
				"rule_system": "DND5e",
				"max_players": 10
			}`,
			userID:         1,
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/rooms", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(rec)
			c.Request = req
			c.Set("user_id", tt.userID)

			handler.CreateRoom(c)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			// 验证邀请码生成
			if tt.expectedStatus == 200 && tt.name == "成功创建房间并自动生成邀请码" {
				var response map[string]interface{}
				err := json.Unmarshal(rec.Body.Bytes(), &response)
				assert.NoError(t, err)
				if data, ok := response["data"].(map[string]interface{}); ok {
					if inviteCode, ok := data["invite_code"].(string); ok {
						assert.Equal(t, 8, len(inviteCode))
						assert.Regexp(t, `^[A-Z0-9]{8}$`, inviteCode)
					}
				}
			}
		})
	}
}

func TestRoomHandler_GetRooms(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&room.Room{})

	handler := NewRoomHandler(db)
	router := testutil.SetupTestRouter()
	router.GET("/rooms", handler.GetRooms)

	// 创建测试房间
	room1 := &room.Room{
		Name:       "Public Room 1",
		InviteCode: "invite1",
		IsPublic:   true,
		MaxPlayers: 5,
		DMID:       1,
	}
	db.Create(room1)
	
	room2 := &room.Room{
		Name:       "Public Room 2",
		InviteCode: "invite2",
		IsPublic:   true,
		MaxPlayers: 3,
		DMID:       2,
	}
	db.Create(room2)
	
	room3 := &room.Room{
		Name:       "Private Room",
		InviteCode: "invite3",
		IsPublic:   false,
		MaxPlayers: 4,
		DMID:       1,
	}
	db.Create(room3)
	
	// 验证房间是否正确创建
	var count int64
	db.Model(&room.Room{}).Count(&count)
	assert.Equal(t, int64(3), count)
	
	// 验证Private Room的IsPublic是否为false
	var privateRoom room.Room
	db.Where("invite_code = ?", "invite3").First(&privateRoom)
	assert.Equal(t, false, privateRoom.IsPublic)

	req := httptest.NewRequest("GET", "/rooms", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "Public Room 1")
	assert.Contains(t, rec.Body.String(), "Public Room 2")
	assert.NotContains(t, rec.Body.String(), "Private Room")
}

func TestRoomHandler_GetRoom(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&room.Room{})

	handler := NewRoomHandler(db)
	router := testutil.SetupTestRouter()
	router.GET("/rooms/:id", handler.GetRoom)

	// 创建测试房间
	testRoom := room.Room{
		Name:       "Test Room",
		IsPublic:   true,
		MaxPlayers: 5,
	}
	db.Create(&testRoom)

	tests := []struct {
		name           string
		roomID         string
		expectedStatus int
	}{
		{
			name:           "成功获取房间",
			roomID:         "1",
			expectedStatus: 200,
		},
		{
			name:           "房间不存在",
			roomID:         "999",
			expectedStatus: 404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/rooms/"+tt.roomID, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			if tt.expectedStatus == 200 {
				assert.Contains(t, rec.Body.String(), "Test Room")
			}
		})
	}
}

func TestRoomHandler_JoinRoom(t *testing.T) {
	// 设置JWT secret环境变量
	os.Setenv("JWT_SECRET", "test-secret")
	defer os.Unsetenv("JWT_SECRET")
	
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&room.Room{}, &room.RoomMember{})

	handler := NewRoomHandler(db)
	router := testutil.SetupTestRouter()
	// 创建路由组，使用中间件设置user_id
	joinGroup := router.Group("/rooms/:id")
	joinGroup.Use(func(c *gin.Context) {
		// 从自定义头获取user_id
		userIDStr := c.GetHeader("X-Test-User-ID")
		if userIDStr != "" {
			// 解析user_id并设置到上下文
			var userID uint
			fmt.Sscanf(userIDStr, "%d", &userID)
			c.Set("user_id", userID)
		}
		c.Next()
	})
	joinGroup.POST("/join", handler.JoinRoom)

	// 创建测试房间
	testRoom := room.Room{
		Name:       "Test Room",
		IsPublic:   true,
		MaxPlayers: 2,
		DMID:       1,
	}
	db.Create(&testRoom)

	tests := []struct {
		name           string
		roomID         string
		userID         uint
		expectedStatus int
		description    string
		setup          func() // 为每个测试设置不同的初始状态
	}{
		{
			name:           "成功加入房间",
			roomID:         "1",
			userID:         2,
			expectedStatus: 200,
			description:    "用户成功加入房间",
			setup:          func() {},
		},
		{
			name:           "房间不存在",
			roomID:         "999",
			userID:         2,
			expectedStatus: 404,
			description:    "尝试加入不存在的房间",
			setup:          func() {},
		},
		{
			name:           "房间已满",
			roomID:         "1",
			userID:         3,
			expectedStatus: 400,
			description:    "房间人数已达到上限",
			setup: func() {
				// 添加1个成员使房间满员（MaxPlayers=2，已有DM=1）
				db.Create(&room.RoomMember{
					RoomID:   1,
					UserID:   4,
					Role:     "player",
					JoinedAt: time.Now(),
				})
			},
		},
		{
			name:           "重复加入房间",
			roomID:         "1",
			userID:         2,
			expectedStatus: 400,
			description:    "用户已在房间中",
			setup: func() {
				// 先让用户2加入房间
				db.Create(&room.RoomMember{
					RoomID:   1,
					UserID:   2,
					Role:     "player",
					JoinedAt: time.Now(),
				})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 为每个测试设置初始状态
			tt.setup()
			
			req := httptest.NewRequest("POST", "/rooms/"+tt.roomID+"/join", nil)
			
			// 使用自定义头传递user_id，绕过认证
			req.Header.Set("X-Test-User-ID", fmt.Sprintf("%d", tt.userID))
			
			rec := httptest.NewRecorder()

			// 使用router.ServeHTTP来正确处理路由参数
			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code, tt.description)
		})
	}
}

func TestRoomHandler_LeaveRoom(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&room.Room{}, &room.RoomMember{})

	handler := NewRoomHandler(db)
	router := testutil.SetupTestRouter()
	// 创建路由组，使用中间件设置user_id
	leaveGroup := router.Group("/rooms/:id")
	leaveGroup.Use(func(c *gin.Context) {
		// 从自定义头获取user_id
		userIDStr := c.GetHeader("X-Test-User-ID")
		if userIDStr != "" {
			// 解析user_id并设置到上下文
			var userID uint
			fmt.Sscanf(userIDStr, "%d", &userID)
			c.Set("user_id", userID)
		}
		c.Next()
	})
	leaveGroup.POST("/leave", handler.LeaveRoom)
	
	// 创建测试房间
	testRoom := room.Room{
		Name:       "Test Room",
		IsPublic:   true,
		MaxPlayers: 5,
		DMID:       2, // DM是user 2，不是user 1
	}
	db.Create(&testRoom)
	
	// 添加user 1到房间
	db.Create(&room.RoomMember{
		RoomID:   1,
		UserID:   1,
		Role:     "player",
		JoinedAt: time.Now(),
	})

	req := httptest.NewRequest("POST", "/rooms/1/leave", nil)
	req.Header.Set("X-Test-User-ID", "1") // user 1离开房间
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
}

func TestRoomHandler_DeleteRoom(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&room.Room{}, &room.RoomMember{})

	handler := NewRoomHandler(db)
	router := testutil.SetupTestRouter()
	// 创建路由组，使用中间件设置user_id
	deleteGroup := router.Group("/rooms/:id")
	deleteGroup.Use(func(c *gin.Context) {
		// 从自定义头获取user_id
		userIDStr := c.GetHeader("X-Test-User-ID")
		if userIDStr != "" {
			// 解析user_id并设置到上下文
			var userID uint
			fmt.Sscanf(userIDStr, "%d", &userID)
			c.Set("user_id", userID)
		}
		c.Next()
	})
	deleteGroup.DELETE("", handler.DeleteRoom)
	
	// 创建测试房间，DM是user 1
	testRoom := room.Room{
		Name:       "Test Room",
		IsPublic:   true,
		MaxPlayers: 5,
		DMID:       1,
	}
	db.Create(&testRoom)

	req := httptest.NewRequest("DELETE", "/rooms/1", nil)
	req.Header.Set("X-Test-User-ID", "1") // user 1是DM，可以删除房间
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
}
