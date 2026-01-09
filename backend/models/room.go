package models

import (
	"time"

	"gorm.io/gorm"
)

// Room 房间模型
type Room struct {
	ID          string         `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name        string         `gorm:"type:varchar(100);not null" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	RuleType    string         `gorm:"type:varchar(20);not null;default:'DND5e'" json:"rule_type"`
	Password    string         `gorm:"type:varchar(50)" json:"-"`
	CreatorID   string         `gorm:"type:uuid;not null" json:"creator_id"`
	Creator     User           `gorm:"foreignKey:CreatorID" json:"creator"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Status      string         `gorm:"type:varchar(20);not null;default:'active'" json:"status"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

// RoomMember 房间成员模型
type RoomMember struct {
	ID        string         `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	RoomID    string         `gorm:"type:uuid;not null" json:"room_id"`
	Room      Room           `gorm:"foreignKey:RoomID" json:"room"`
	UserID    string         `gorm:"type:uuid;not null" json:"user_id"`
	User      User           `gorm:"foreignKey:UserID" json:"user"`
	Role      string         `gorm:"type:varchar(20);not null;default:'player'" json:"role"`
	JoinedAt  time.Time      `json:"joined_at"`
	Status    string         `gorm:"type:varchar(20);not null;default:'active'" json:"status"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
