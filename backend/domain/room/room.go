package room

import (
	"time"
)

type Room struct {
	ID          uint      `gorm:"primaryKey"`
	Name        string    `gorm:"not null"`
	Description string
	RuleSystem  string    `gorm:"not null;default:'DND5e'"`
	Password    string
	InviteCode  string    `gorm:"uniqueIndex;not null"`
	DMID        uint      `gorm:"not null"`
	MaxPlayers  int       `gorm:"not null;default:10"`
	IsPublic    bool      `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (Room) TableName() string {
	return "rooms"
}

type RoomMember struct {
	ID       uint      `gorm:"primaryKey"`
	RoomID   uint      `gorm:"not null"`
	UserID   uint      `gorm:"not null"`
	Role     string    `gorm:"not null;default:'player'"`
	JoinedAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (RoomMember) TableName() string {
	return "room_members"
}
