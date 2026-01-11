package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"trpg-sync/backend/domain/character"
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

	db.AutoMigrate(&character.CharacterCard{})

	handler := NewCharacterHandler(db)

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
			c.Set("user_id", tt.userID)
			c.Set("room_id", uint(1))

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

	db.AutoMigrate(&character.CharacterCard{})

	handler := NewCharacterHandler(db)

	req := httptest.NewRequest("PUT", "/rooms/1/characters/1", nil)
	rec := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rec)
	c.Request = req
	c.Set("character_id", uint(1))

	handler.UpdateCharacter(c)

	// 当前实现只返回 200，未实际更新人物卡
	assert.Equal(t, 200, rec.Code)
}

func TestCharacterHandler_DeleteCharacter(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&character.CharacterCard{})

	handler := NewCharacterHandler(db)

	req := httptest.NewRequest("DELETE", "/rooms/1/characters/1", nil)
	rec := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rec)
	c.Request = req
	c.Set("character_id", uint(1))

	handler.DeleteCharacter(c)

	// 当前实现只返回 200，未实际删除人物卡
	assert.Equal(t, 200, rec.Code)
}
