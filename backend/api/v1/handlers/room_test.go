package handlers

import (
	"net/http/httptest"
	"testing"

	"trpg-sync/backend/domain/room"
	"trpg-sync/backend/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoomHandler_CreateRoom(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&room.Room{})

	handler := NewRoomHandler(db)
	router := testutil.SetupTestRouter()
	router.POST("/rooms", handler.CreateRoom)

	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
	}{
		{
			name: "成功创建房间",
			requestBody: `{
				"name": "Test Room",
				"description": "A test room",
				"rule_system": "DND5e"
			}`,
			expectedStatus: 200,
		},
		{
			name: "房间名称为空",
			requestBody: `{
				"description": "A test room",
				"rule_system": "DND5e"
			}`,
			expectedStatus: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/rooms", nil)
			req.Body = nil
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(rec)
			c.Request = req
			c.Request.Body = nil
			c.Request.Header.Set("Content-Type", "application/json")

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
	room1 := &room.Room{
		Name:        "Room 1",
		Description: "Description 1",
		RuleSystem:  "DND5e",
	}
	db.Create(room1)

	room2 := &room.Room{
		Name:        "Room 2",
		Description: "Description 2",
		RuleSystem:  "DND5e",
	}
	db.Create(room2)

	req := httptest.NewRequest("GET", "/rooms", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "Room 1")
	assert.Contains(t, rec.Body.String(), "Room 2")
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
		Name:        "Test Room",
		Description: "Test description",
		RuleSystem:  "DND5e",
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

func TestRoomHandler_DeleteRoom(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&room.Room{})

	handler := NewRoomHandler(db)
	router := testutil.SetupTestRouter()
	router.DELETE("/rooms/:id", handler.DeleteRoom)

	// 创建测试房间
	testRoom := room.Room{
		Name:        "Test Room",
		Description: "Test description",
		RuleSystem:  "DND5e",
	}
	db.Create(&testRoom)

	req := httptest.NewRequest("DELETE", "/rooms/1", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)

	// 验证房间已被删除
	var count int64
	db.Model(&room.Room{}).Count(&count)
	assert.Equal(t, int64(0), count)
}
