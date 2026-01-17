package handlers

import (
	"net/http/httptest"
	"strings"
	"testing"

	"trpg-sync/backend/domain/character"
	"trpg-sync/backend/domain/room"
	"trpg-sync/backend/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCharacterHandler_CreateCharacter(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&character.CharacterCard{}, &room.Room{})

	handler := NewCharacterHandler(db)

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
		requestBody    string
		expectedStatus int
	}{
		{
			name:   "成功创建人物卡",
			roomID: "1",
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
			expectedStatus: 200,
		},
		{
			name:   "人物卡名称为空",
			roomID: "1",
			requestBody: `{
				"race": "Human",
				"class": "Fighter"
			}`,
			expectedStatus: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/rooms/"+tt.roomID+"/characters", nil)
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(rec)
			c.Request = req
			c.Params = []gin.Param{
				{Key: "roomId", Value: tt.roomID},
			}
			c.Request.Body = nil
			c.Request.Header.Set("Content-Type", "application/json")

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

	db.AutoMigrate(&character.CharacterCard{}, &room.Room{})

	handler := NewCharacterHandler(db)

	// 创建测试房间
	testRoom := room.Room{
		Name:        "Test Room",
		Description: "Test description",
		RuleSystem:  "DND5e",
	}
	db.Create(&testRoom)

	// 创建测试人物卡
	char1 := &character.CharacterCard{
		RoomID: testRoom.ID,
		Name:   "Character 1",
		Race:   "Human",
		Class:  "Fighter",
	}
	db.Create(char1)

	char2 := &character.CharacterCard{
		RoomID: testRoom.ID,
		Name:   "Character 2",
		Race:   "Elf",
		Class:  "Wizard",
	}
	db.Create(char2)

	req := httptest.NewRequest("GET", "/rooms/1/characters", nil)
	rec := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rec)
	c.Request = req
	c.Params = []gin.Param{
		{Key: "roomId", Value: "1"},
	}

	handler.GetCharacters(c)

	assert.Equal(t, 200, rec.Code)
	assert.Contains(t, rec.Body.String(), "Character 1")
	assert.Contains(t, rec.Body.String(), "Character 2")
}

func TestCharacterHandler_GetCharacter(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&character.CharacterCard{}, &room.Room{})

	handler := NewCharacterHandler(db)

	// 创建测试房间和人物卡
	testRoom := room.Room{
		Name:        "Test Room",
		Description: "Test description",
		RuleSystem:  "DND5e",
	}
	db.Create(&testRoom)

	testChar := character.CharacterCard{
		RoomID: testRoom.ID,
		Name:   "Test Character",
		Race:   "Human",
		Class:  "Fighter",
	}
	db.Create(&testChar)

	tests := []struct {
		name           string
		roomID         string
		characterID    string
		expectedStatus int
	}{
		{
			name:           "成功获取人物卡",
			roomID:         "1",
			characterID:    "1",
			expectedStatus: 200,
		},
		{
			name:           "人物卡不存在",
			roomID:         "1",
			characterID:    "999",
			expectedStatus: 404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/rooms/"+tt.roomID+"/characters/"+tt.characterID, nil)
			rec := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(rec)
			c.Request = req
			c.Params = []gin.Param{
				{Key: "roomId", Value: tt.roomID},
				{Key: "id", Value: tt.characterID},
			}

			handler.GetCharacter(c)

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

	db.AutoMigrate(&character.CharacterCard{}, &room.Room{})

	handler := NewCharacterHandler(db)

	// 创建测试房间和人物卡
	testRoom := room.Room{
		Name:        "Test Room",
		Description: "Test description",
		RuleSystem:  "DND5e",
	}
	db.Create(&testRoom)

	testChar := character.CharacterCard{
		RoomID: testRoom.ID,
		Name:   "Test Character",
		Race:   "Human",
		Class:  "Fighter",
	}
	db.Create(&testChar)

	// 更新人物卡
	updateBody := `{
		"name": "Updated Character",
		"race": "Elf"
	}`

	req := httptest.NewRequest("PUT", "/rooms/1/characters/1", strings.NewReader(updateBody))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rec)
	c.Request = req
	c.Params = []gin.Param{
		{Key: "roomId", Value: "1"},
		{Key: "id", Value: "1"},
	}

	handler.UpdateCharacter(c)

	assert.Equal(t, 200, rec.Code)

	// 验证数据已更新
	var updatedChar character.CharacterCard
	db.First(&updatedChar, 1)
	assert.Equal(t, "Updated Character", updatedChar.Name)
	assert.Equal(t, "Elf", updatedChar.Race)
}

func TestCharacterHandler_DeleteCharacter(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&character.CharacterCard{}, &room.Room{})

	handler := NewCharacterHandler(db)

	// 创建测试房间和人物卡
	testRoom := room.Room{
		Name:        "Test Room",
		Description: "Test description",
		RuleSystem:  "DND5e",
	}
	db.Create(&testRoom)

	testChar := character.CharacterCard{
		RoomID: testRoom.ID,
		Name:   "Test Character",
		Race:   "Human",
		Class:  "Fighter",
	}
	db.Create(&testChar)

	req := httptest.NewRequest("DELETE", "/rooms/1/characters/1", nil)
	rec := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(rec)
	c.Request = req
	c.Params = []gin.Param{
		{Key: "roomId", Value: "1"},
		{Key: "id", Value: "1"},
	}

	handler.DeleteCharacter(c)

	assert.Equal(t, 200, rec.Code)

	// 验证人物卡已被删除
	var count int64
	db.Model(&character.CharacterCard{}).Count(&count)
	assert.Equal(t, int64(0), count)
}
