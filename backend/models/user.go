package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID           string         `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Username     string         `gorm:"type:varchar(50);unique;not null" json:"username"`
	Email        string         `gorm:"type:varchar(100);unique;not null" json:"email"`
	Phone        string         `gorm:"type:varchar(20)" json:"phone"`
	PasswordHash string         `gorm:"type:varchar(255);not null" json:"-"`
	Avatar       string         `gorm:"type:varchar(255)" json:"avatar"`
	Nickname     string         `gorm:"type:varchar(50)" json:"nickname"`
	Bio          string         `gorm:"type:text" json:"bio"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	LastLoginAt  time.Time      `json:"last_login_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}
