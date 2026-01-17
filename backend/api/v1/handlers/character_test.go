package handlers

import (
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"trpg-sync/backend/domain/character"
	"trpg-sync/backend/domain/room"
	"trpg-sync/backend/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCharacterHandler_CreateCharacter(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// 自动迁移所有相关表
	db.AutoMigrate(&character.CharacterCard{}, &room.Room{}, &room.RoomMember{})

	handler := NewCharacterHandler(db)

	// 创建测试房间和成员
	testRoom := room.Room{
		Name:       "Test Room",
		InviteCode: "TEST123",
		DMID:       1,
		MaxPlayers: 10,
		IsPublic:   true,
	}
	db.Create(&testRoom)

	// 创建测试成员
	testMember := room.RoomMember{
		RoomID:   testRoom.ID,
		UserID:   1,
		Role:     "dm",
		JoinedAt: time.Now(),
	}
	db.Create(&testMember)

	tests := []struct {
		name           string
		requestBody    string
		userID         uint
		expectedStatus int
	}{
		{
			name: "成功创建人物卡",
			requestBody: `{
				"name": "Test Character",
				"race": "Human",
				"class": "Fighter",
				"level": 1,
				"background": "Soldier",
				"alignment": "Lawful Good",
				"strength": 16,
				"dexterity": 14,
				"constitution": 15,
				"intelligence": 10,
				"wisdom": 12,
				"charisma": 8
			}`,
			userID:         1,
			expectedStatus: 200,
		},
		{
			name: "人物卡名称为空",
			requestBody: `{
				"race": "Human",
				"class": "Fighter"
			}`,
			userID:         1,
			expectedStatus: 400,
		},
		{
			name: "成功创建基础人物卡（仅必填字段）",
			requestBody: `{
				"name": "Basic Character"
			}`,
			userID:         1,
			expectedStatus: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/rooms/1/characters", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(rec)
			c.Request = req
			c.Params = gin.Params{{Key: "roomId", Value: "1"}}
			c.Set("user_id", tt.userID)

			handler.CreateCharacter(c)

			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}

func TestCharacterHandler_GetCharacters(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&character.CharacterCard{})

	handler := NewCharacterHandler(db)
	router := testutil.SetupTestRouter()
	router.GET("/rooms/:roomId/characters", handler.GetCharacters)

	// 创建测试人物卡
	db.Create(&character.CharacterCard{
		Name:   "Character 1",
		RoomID: 1,
		UserID: 1,
	})
	db.Create(&character.CharacterCard{
		Name:   "Character 2",
		RoomID: 1,
		UserID: 2,
	})
	db.Create(&character.CharacterCard{
		Name:   "Character 3",
		RoomID: 2,
		UserID: 1,
	})

	req := httptest.NewRequest("GET", "/rooms/1/characters", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "Character 1")
	assert.Contains(t, rec.Body.String(), "Character 2")
	assert.NotContains(t, rec.Body.String(), "Character 3")
}

func TestCharacterHandler_GetCharacter(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&character.CharacterCard{})

	handler := NewCharacterHandler(db)
	router := testutil.SetupTestRouter()
	router.GET("/rooms/:roomId/characters/:id", handler.GetCharacter)

	// 创建测试人物卡
	testCharacter := character.CharacterCard{
		Name:   "Test Character",
		RoomID: 1,
		UserID: 1,
	}
	db.Create(&testCharacter)

	tests := []struct {
		name           string
		characterID    string
		expectedStatus int
	}{
		{
			name:           "成功获取人物卡",
			characterID:    "1",
			expectedStatus: 200,
		},
		{
			name:           "人物卡不存在",
			characterID:    "999",
			expectedStatus: 404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/rooms/1/characters/"+tt.characterID, nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			if tt.expectedStatus == 200 {
				assert.Contains(t, rec.Body.String(), "Test Character")
			}
		})
	}
}

func TestCharacterHandler_UpdateCharacter(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// 自动迁移所有相关表
	db.AutoMigrate(&character.CharacterCard{}, &room.Room{}, &room.RoomMember{})

	handler := NewCharacterHandler(db)

	// 创建测试房间和成员
	testRoom := room.Room{
		Name:       "Test Room",
		InviteCode: "TEST123",
		DMID:       1,
		MaxPlayers: 10,
		IsPublic:   true,
	}
	db.Create(&testRoom)

	testMember := room.RoomMember{
		RoomID:   testRoom.ID,
		UserID:   1,
		Role:     "dm",
		JoinedAt: time.Now(),
	}
	db.Create(&testMember)

	// 创建测试人物卡
	testCharacter := character.CharacterCard{
		UserID: 1,
		RoomID: testRoom.ID,
		Name:   "Test Character",
		Level:  1,
	}
	db.Create(&testCharacter)

	reqBody := `{
		"name": "Updated Character",
		"race": "Elf",
		"class": "Wizard",
		"level": 2
	}`
	req := httptest.NewRequest("PUT", "/rooms/1/characters/1", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rec)
	c.Request = req
	c.Params = gin.Params{{Key: "roomId", Value: "1"}, {Key: "id", Value: "1"}}
	c.Set("user_id", uint(1))

	handler.UpdateCharacter(c)

	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "Updated Character")
}

func TestCharacterHandler_DeleteCharacter(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// 自动迁移所有相关表
	db.AutoMigrate(&character.CharacterCard{}, &room.Room{}, &room.RoomMember{})

	handler := NewCharacterHandler(db)

	// 创建测试房间和成员
	testRoom := room.Room{
		Name:       "Test Room",
		InviteCode: "TEST123",
		DMID:       1,
		MaxPlayers: 10,
		IsPublic:   true,
	}
	db.Create(&testRoom)

	testMember := room.RoomMember{
		RoomID:   testRoom.ID,
		UserID:   1,
		Role:     "dm",
		JoinedAt: time.Now(),
	}
	db.Create(&testMember)

	// 创建测试人物卡
	testCharacter := character.CharacterCard{
		UserID: 1,
		RoomID: testRoom.ID,
		Name:   "Test Character",
		Level:  1,
	}
	db.Create(&testCharacter)

	req := httptest.NewRequest("DELETE", "/rooms/1/characters/1", nil)
	rec := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rec)
	c.Request = req
	c.Params = gin.Params{{Key: "roomId", Value: "1"}, {Key: "id", Value: "1"}}
	c.Set("user_id", uint(1))

	handler.DeleteCharacter(c)

	assert.Equal(t, 200, rec.Code)
}
