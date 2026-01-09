package character

import (
	"time"
)

type CharacterCard struct {
	ID           uint      `gorm:"primaryKey"`
	UserID       uint      `gorm:"not null"`
	RoomID       uint      `gorm:"not null"`
	Name         string    `gorm:"not null"`
	Race         string
	Class        string
	Level        int       `gorm:"default:1"`
	Background   string
	Alignment    string
	Strength     int       `gorm:"default:10"`
	Dexterity    int       `gorm:"default:10"`
	Constitution int       `gorm:"default:10"`
	Intelligence int       `gorm:"default:10"`
	Wisdom       int       `gorm:"default:10"`
	Charisma     int       `gorm:"default:10"`
	AC           int       `gorm:"default:10"`
	HP           int       `gorm:"default:10"`
	MaxHP        int       `gorm:"default:10"`
	Speed        int       `gorm:"default:30"`
	Proficiency  int       `gorm:"default:2"`
	Skills       string    `gorm:"type:jsonb"`
	Saves        string    `gorm:"type:jsonb"`
	Equipment    string    `gorm:"type:jsonb"`
	Spells       string    `gorm:"type:jsonb"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (CharacterCard) TableName() string {
	return "character_cards"
}
