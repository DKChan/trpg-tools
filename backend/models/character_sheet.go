package models

import (
	"time"

	"gorm.io/gorm"
)

// CharacterSheet 人物卡模型
type CharacterSheet struct {
	ID                     string         `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	RoomID                 string         `gorm:"type:uuid;not null" json:"room_id"`
	Room                   Room           `gorm:"foreignKey:RoomID" json:"room"`
	UserID                 string         `gorm:"type:uuid;not null" json:"user_id"`
	User                   User           `gorm:"foreignKey:UserID" json:"user"`
	Name                   string         `gorm:"type:varchar(100);not null" json:"name"`
	Race                   string         `gorm:"type:varchar(50)" json:"race"`
	Class                  string         `gorm:"type:varchar(50)" json:"class"`
	Level                  int            `gorm:"not null;default:1" json:"level"`
	Background             string         `gorm:"type:varchar(100)" json:"background"`
	Alignment              string         `gorm:"type:varchar(50)" json:"alignment"`
	Strength               int            `gorm:"not null;default:10" json:"strength"`
	Dexterity              int            `gorm:"not null;default:10" json:"dexterity"`
	Constitution           int            `gorm:"not null;default:10" json:"constitution"`
	Intelligence           int            `gorm:"not null;default:10" json:"intelligence"`
	Wisdom                 int            `gorm:"not null;default:10" json:"wisdom"`
	Charisma               int            `gorm:"not null;default:10" json:"charisma"`
	Skills                 string         `gorm:"type:jsonb" json:"skills"` // JSON格式存储技能数据
	HitPoints              int            `gorm:"not null;default:10" json:"hit_points"`
	MaxHitPoints           int            `gorm:"not null;default:10" json:"max_hit_points"`
	TemporaryHitPoints     int            `gorm:"not null;default:0" json:"temporary_hit_points"`
	Magic                  string         `gorm:"type:jsonb" json:"magic"` // JSON格式存储魔法数据
	Equipment              string         `gorm:"type:jsonb" json:"equipment"` // JSON格式存储装备数据
	Features               string         `gorm:"type:jsonb" json:"features"` // JSON格式存储特性与专长数据
	Backstory              string         `gorm:"type:text" json:"backstory"`
	CreatedAt              time.Time      `json:"created_at"`
	UpdatedAt              time.Time      `json:"updated_at"`
	DeletedAt              gorm.DeletedAt `gorm:"index" json:"-"`
}

// CharacterSheetVersion 人物卡版本模型
type CharacterSheetVersion struct {
	ID                string         `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CharacterSheetID  string         `gorm:"type:uuid;not null" json:"character_sheet_id"`
	CharacterSheet    CharacterSheet `gorm:"foreignKey:CharacterSheetID" json:"character_sheet"`
	UserID            string         `gorm:"type:uuid;not null" json:"user_id"`
	User              User           `gorm:"foreignKey:UserID" json:"user"`
	Data              string         `gorm:"type:jsonb;not null" json:"data"` // JSON格式存储人物卡数据快照
	CreatedAt         time.Time      `json:"created_at"`
	Note              string         `gorm:"type:varchar(255)" json:"note"`
}
