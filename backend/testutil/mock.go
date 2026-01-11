package testutil

// MockUser 模拟用户数据
type MockUser struct {
	ID       uint
	Email    string
	Password string
	Nickname string
	Avatar   string
}

// NewMockUser 创建模拟用户
func NewMockUser() MockUser {
	return MockUser{
		ID:       1,
		Email:    "test@example.com",
		Password: "hashedpassword",
		Nickname: "Test User",
		Avatar:   "https://example.com/avatar.png",
	}
}

// MockRoom 模拟房间数据
type MockRoom struct {
	ID          uint
	Name        string
	Description string
	RuleSystem  string
	DMID        uint
	MaxPlayers  int
	IsPublic    bool
}

// NewMockRoom 创建模拟房间
func NewMockRoom() MockRoom {
	return MockRoom{
		ID:          1,
		Name:        "Test Room",
		Description: "A test room for testing",
		RuleSystem:  "DND5e",
		DMID:        1,
		MaxPlayers:  5,
		IsPublic:    true,
	}
}

// MockCharacter 模拟人物卡数据
type MockCharacter struct {
	ID         uint
	Name       string
	Race       string
	Class      string
	Level      int
	Background string
	Alignment  string
	Strength   int
	Dexterity  int
}

// NewMockCharacter 创建模拟人物卡
func NewMockCharacter() MockCharacter {
	return MockCharacter{
		ID:         1,
		Name:       "Test Character",
		Race:       "Human",
		Class:      "Fighter",
		Level:      1,
		Background: "Soldier",
		Alignment:  "Lawful Good",
		Strength:   16,
		Dexterity:  14,
	}
}
