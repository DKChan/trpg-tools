package room

import (
	"time"
)

type Room struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string
	RuleSystem  string    `gorm:"not null;default:'DND5e'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Room) TableName() string {
	return "rooms"
}
