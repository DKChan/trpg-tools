package testutil

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// GenerateTestToken 生成用于测试的 JWT token
func GenerateTestToken(userID uint, email string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

// GenerateExpiredToken 生成已过期的测试 token
func GenerateExpiredToken(userID uint, email string, secret string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(-time.Hour).Unix(), // 已过期
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
