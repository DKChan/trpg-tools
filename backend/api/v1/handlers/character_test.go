package handlers

import (
	"os"
	"path/filepath"
	"testing"

	"trpg-sync/backend/domain/character"
	"trpg-sync/backend/infrastructure/storage"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCharacterStorage_SaveAndLoad(t *testing.T) {
	charStorage := storage.NewCharacterStorage()

	// 创建测试人物卡
	char := &character.CharacterCard{
		ID:           1,
		RoomID:       100,
		Name:         "Test Character",
		Race:         "Human",
		Class:        "Fighter",
		Level:        1,
		Strength:     16,
		Dexterity:    14,
		Constitution: 15,
	}

	// 保存人物卡
	err := charStorage.SaveCharacter(char)
	require.NoError(t, err)

	// 验证文件存在
	charPath := charStorage.GetCharacterFilePath(100, 1)
	_, err = os.Stat(charPath)
	require.NoError(t, err)
	require.True(t, !os.IsNotExist(err))

	// 加载人物卡
	loadedChar, err := charStorage.LoadCharacter(100, 1)
	require.NoError(t, err)

	// 验证数据
	assert.Equal(t, char.ID, loadedChar.ID)
	assert.Equal(t, char.Name, loadedChar.Name)
	assert.Equal(t, char.Race, loadedChar.Race)
	assert.Equal(t, char.Class, loadedChar.Class)
}

func TestCharacterStorage_GetRoomCharacters(t *testing.T) {
	charStorage := storage.NewCharacterStorage()

	// 创建测试房间目录
	roomID := uint(100)
	charDir := charStorage.GetRoomCharactersPath(roomID)
	err := os.MkdirAll(charDir, 0755)
	require.NoError(t, err)

	// 创建多个测试人物卡
	chars := []*character.CharacterCard{
		{ID: 1, RoomID: roomID, Name: "Char 1", Race: "Human", Class: "Fighter"},
		{ID: 2, RoomID: roomID, Name: "Char 2", Race: "Elf", Class: "Wizard"},
		{ID: 3, RoomID: roomID, Name: "Char 3", Race: "Dwarf", Class: "Cleric"},
	}

	for _, char := range chars {
		err := charStorage.SaveCharacter(char)
		require.NoError(t, err)
	}

	// 获取房间所有人物卡
	loadedChars, err := charStorage.GetRoomCharacters(roomID)
	require.NoError(t, err)

	// 验证数量
	assert.Len(t, loadedChars, 3)

	// 验证数据
	for i, expectedChar := range chars {
		assert.Equal(t, expectedChar.ID, loadedChars[i].ID)
		assert.Equal(t, expectedChar.Name, loadedChars[i].Name)
	}
}

func TestCharacterStorage_DeleteCharacter(t *testing.T) {
	charStorage := storage.NewCharacterStorage()

	// 创建测试人物卡
	char := &character.CharacterCard{
		ID:      1,
		RoomID:  100,
		Name:    "Test Character",
		Race:    "Human",
		Class:   "Fighter",
	}

	err := charStorage.SaveCharacter(char)
	require.NoError(t, err)

	// 验证文件存在
	charPath := charStorage.GetCharacterFilePath(100, 1)
	_, err = os.Stat(charPath)
	require.NoError(t, err)

	// 删除人物卡
	err = charStorage.DeleteCharacter(100, 1)
	require.NoError(t, err)

	// 验证文件已被删除
	_, err = os.Stat(charPath)
	assert.True(t, os.IsNotExist(err))
}

func TestCharacterStorage_GenerateNextID(t *testing.T) {
	charStorage := storage.NewCharacterStorage()

	// 空目录，应该返回 1
	id1, err := charStorage.GenerateNextID(100)
	require.NoError(t, err)
	assert.Equal(t, uint(1), id1)

	// 创建文件
	char := &character.CharacterCard{ID: 1, RoomID: 100, Name: "Char"}
	charStorage.SaveCharacter(char)

	// 下一个 ID 应该是 2
	id2, err := charStorage.GenerateNextID(100)
	require.NoError(t, err)
	assert.Equal(t, uint(2), id2)
}
