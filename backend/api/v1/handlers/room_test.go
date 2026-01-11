package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"trpg-sync/backend/domain/room"
	"trpg-sync/backend/domain/user"
	"trpg-sync/backend/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

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
	db.Create(&room.Room{
		Name:       "Public Room 1",
		IsPublic:   true,
		MaxPlayers: 5,
	})
	db.Create(&room.Room{
		Name:       "Public Room 2",
		IsPublic:   true,
		MaxPlayers: 3,
	})
	db.Create(&room.Room{
		Name:       "Private Room",
		IsPublic:   false,
		MaxPlayers: 4,
	})

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
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&room.Room{})

	handler := NewRoomHandler(db)

	// 创建测试房间
	testRoom := room.Room{
		Name:       "Test Room",
		IsPublic:   true,
		MaxPlayers: 5,
	}
	db.Create(&testRoom)

	req := httptest.NewRequest("POST", "/rooms/1/join", nil)
	rec := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rec)
	c.Request = req
	c.Set("user_id", uint(1))
	c.Set("room_id", uint(1))

	handler.JoinRoom(c)

	// 当前实现只返回 200，未实际加入房间
	assert.Equal(t, 200, rec.Code)
}

func TestRoomHandler_LeaveRoom(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&room.Room{})

	handler := NewRoomHandler(db)

	req := httptest.NewRequest("POST", "/rooms/1/leave", nil)
	rec := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rec)
	c.Request = req
	c.Set("user_id", uint(1))
	c.Set("room_id", uint(1))

	handler.LeaveRoom(c)

	// 当前实现只返回 200，未实际退出房间
	assert.Equal(t, 200, rec.Code)
}

func TestRoomHandler_DeleteRoom(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&room.Room{})

	handler := NewRoomHandler(db)

	req := httptest.NewRequest("DELETE", "/rooms/1", nil)
	rec := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rec)
	c.Request = req
	c.Set("user_id", uint(1))
	c.Set("room_id", uint(1))

	handler.DeleteRoom(c)

	// 当前实现只返回 200，未实际删除房间
	assert.Equal(t, 200, rec.Code)
}
