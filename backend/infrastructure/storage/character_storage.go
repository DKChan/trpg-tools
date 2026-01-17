package storage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"trpg-sync/backend/domain/character"
)

const DataDir = "data"
const RoomsDir = "rooms"

type CharacterStorage struct {
	basePath string
}

func NewCharacterStorage() *CharacterStorage {
	return &CharacterStorage{
		basePath: DataDir,
	}
}

// GetRoomCharactersPath 获取房间的人物卡目录
func (s *CharacterStorage) GetRoomCharactersPath(roomID uint) string {
	return filepath.Join(s.basePath, RoomsDir, strconv.FormatUint(uint64(roomID), 10), "characters")
}

// GetCharacterFilePath 获取人物卡文件路径
func (s *CharacterStorage) GetCharacterFilePath(roomID uint, characterID uint) string {
	return filepath.Join(s.GetRoomCharactersPath(roomID), strconv.FormatUint(uint64(characterID), 10)+".json")
}

// SaveCharacter 保存人物卡到文件
func (s *CharacterStorage) SaveCharacter(char *character.CharacterCard) error {
	// 确保目录存在
	charPath := s.GetCharacterFilePath(char.RoomID, char.ID)
	charDir := filepath.Dir(charPath)

	if err := os.MkdirAll(charDir, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	// 序列化为 JSON
	data, err := json.MarshalIndent(char, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal character: %w", err)
	}

	// 写入文件
	if err := os.WriteFile(charPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write character file: %w", err)
	}

	return nil
}

// LoadCharacter 从文件加载人物卡
func (s *CharacterStorage) LoadCharacter(roomID uint, characterID uint) (*character.CharacterCard, error) {
	charPath := s.GetCharacterFilePath(roomID, characterID)

	data, err := os.ReadFile(charPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read character file: %w", err)
	}

	var char character.CharacterCard
	if err := json.Unmarshal(data, &char); err != nil {
		return nil, fmt.Errorf("failed to unmarshal character: %w", err)
	}

	return &char, nil
}

// DeleteCharacter 删除人物卡文件
func (s *CharacterStorage) DeleteCharacter(roomID uint, characterID uint) error {
	charPath := s.GetCharacterFilePath(roomID, characterID)

	if err := os.Remove(charPath); err != nil {
		return fmt.Errorf("failed to delete character file: %w", err)
	}

	return nil
}

// GetRoomCharacters 获取房间的所有人物卡
func (s *CharacterStorage) GetRoomCharacters(roomID uint) ([]*character.CharacterCard, error) {
	charDir := s.GetRoomCharactersPath(roomID)

	// 读取目录内容
	entries, err := os.ReadDir(charDir)
	if err != nil {
		if os.IsNotExist(err) {
			// 目录不存在，返回空列表
			return []*character.CharacterCard{}, nil
		}
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	var characters []*character.CharacterCard

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		// 解析文件名获取 ID
		fileName := entry.Name()
		charID, err := strconv.ParseUint(fileName[:len(fileName)-5], 10, 64)
		if err != nil {
			// 文件名格式不正确，跳过
			continue
		}

		// 加载人物卡
		char, err := s.LoadCharacter(roomID, uint(charID))
		if err != nil {
			return nil, fmt.Errorf("failed to load character %s: %w", fileName, err)
		}

		characters = append(characters, char)
	}

	return characters, nil
}

// GenerateNextID 生成下一个人物卡 ID
func (s *CharacterStorage) GenerateNextID(roomID uint) (uint, error) {
	charDir := s.GetRoomCharactersPath(roomID)

	// 确保目录存在
	if err := os.MkdirAll(charDir, 0755); err != nil {
		return 0, fmt.Errorf("failed to create directory: %w", err)
	}

	// 读取目录内容
	entries, err := os.ReadDir(charDir)
	if err != nil {
		return 0, fmt.Errorf("failed to read directory: %w", err)
	}

	maxID := uint(0)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		// 解析文件名获取 ID
		fileName := entry.Name()
		charID, err := strconv.ParseUint(fileName[:len(fileName)-5], 10, 64)
		if err == nil && charID > maxID {
			maxID = uint(charID)
		}
	}

	return maxID + 1, nil
}

// ExportCharacterToFile 导出人物卡到指定路径
func (s *CharacterStorage) ExportCharacterToFile(char *character.CharacterCard, exportPath string) error {
	data, err := json.MarshalIndent(char, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal character: %w", err)
	}

	if err := os.WriteFile(exportPath, data, 0644); err != nil {
		return fmt.Errorf("failed to write export file: %w", err)
	}

	return nil
}

// ImportCharacterFromFile 从文件导入人物卡
func (s *CharacterStorage) ImportCharacterFromFile(filePath string, roomID uint) (*character.CharacterCard, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read import file: %w", err)
	}

	var char character.CharacterCard
	if err := json.Unmarshal(data, &char); err != nil {
		return nil, fmt.Errorf("failed to unmarshal character: %w", err)
	}

	char.RoomID = roomID
	return &char, nil
}

// CopyCharacter 从源文件复制到目标房间
func (s *CharacterStorage) CopyCharacter(srcRoomID uint, srcCharID uint, dstRoomID uint) (*character.CharacterCard, error) {
	char, err := s.LoadCharacter(srcRoomID, srcCharID)
	if err != nil {
		return nil, err
	}

	// 生成新 ID
	newID, err := s.GenerateNextID(dstRoomID)
	if err != nil {
		return nil, err
	}

	char.ID = newID
	char.RoomID = dstRoomID

	// 保存到目标房间
	if err := s.SaveCharacter(char); err != nil {
		return nil, err
	}

	return char, nil
}

// BackupRoomCharacters 备份房间所有人物卡为 ZIP
func (s *CharacterStorage) BackupRoomCharacters(roomID uint, backupPath string) error {
	// 简化实现：复制整个房间目录
	charDir := s.GetRoomCharactersPath(roomID)

	// 读取所有文件
	entries, err := os.ReadDir(charDir)
	if err != nil {
		return fmt.Errorf("failed to read directory: %w", err)
	}

	// 创建备份目录
	if err := os.MkdirAll(backupPath, 0755); err != nil {
		return fmt.Errorf("failed to create backup directory: %w", err)
	}

	// 复制所有文件
	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		srcPath := filepath.Join(charDir, entry.Name())
		dstPath := filepath.Join(backupPath, entry.Name())

		// 复制文件
		if err := copyFile(srcPath, dstPath); err != nil {
			return fmt.Errorf("failed to copy file %s: %w", entry.Name(), err)
		}
	}

	return nil
}

// copyFile 复制文件的辅助函数
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	if _, err := io.Copy(destinationFile, sourceFile); err != nil {
		return err
	}

	return nil
}
