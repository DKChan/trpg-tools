package room

import (
	"time"
)

type Room struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name" gorm:"not null"`
	Description string    `json:"description"`
	RuleSystem  string    `json:"rule_system" gorm:"not null;default:'DND5e'"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Room) TableName() string {
	return "rooms"
}
