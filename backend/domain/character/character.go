package character

type CharacterCard struct {
	ID           uint      `json:"id"`
	RoomID       uint      `json:"room_id"`
	Name         string    `json:"name"`
	Race         string    `json:"race"`
	Class        string    `json:"class"`
	Level        int       `json:"level"`
	Background   string    `json:"background"`
	Alignment    string    `json:"alignment"`
	Strength     int       `json:"strength"`
	Dexterity    int       `json:"dexterity"`
	Constitution int       `json:"constitution"`
	Intelligence int       `json:"intelligence"`
	Wisdom       int       `json:"wisdom"`
	Charisma     int       `json:"charisma"`
	AC           int       `json:"ac"`
	HP           int       `json:"hp"`
	MaxHP        int       `json:"max_hp"`
	Speed        int       `json:"speed"`
	Proficiency  int       `json:"proficiency"`
	Skills       string    `json:"skills"`
	Saves        string    `json:"saves"`
	Equipment    string    `json:"equipment"`
	Spells       string    `json:"spells"`
}
