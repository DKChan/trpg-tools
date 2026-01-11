package handlers

import (
	"net/http/httptest"
	"strings"
	"testing"

	"trpg-sync/backend/domain/user"
	"trpg-sync/backend/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthHandler_Register(t *testing.T) {
	// 设置测试数据库
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// 自动迁移表结构
	db.AutoMigrate(&user.User{})

	// 创建 handler
	handler := NewAuthHandler(db)
	router := testutil.SetupTestRouter()
	router.POST("/auth/register", handler.Register)

	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
		expectedBody   map[string]interface{}
	}{
		{
			name: "成功注册新用户",
			requestBody: `{
				"email": "test@example.com",
				"password": "password123",
				"nickname": "Test User"
			}`,
			expectedStatus: 200,
			expectedBody: map[string]interface{}{
				"code":    float64(200),
				"message": "User registered successfully",
			},
		},
		{
			name: "邮箱格式无效",
			requestBody: `{
				"email": "invalid-email",
				"password": "password123",
				"nickname": "Test User"
			}`,
			expectedStatus: 400,
			expectedBody: map[string]interface{}{
				"code": float64(400),
			},
		},
		{
			name: "密码少于6位",
			requestBody: `{
				"email": "test@example.com",
				"password": "123",
				"nickname": "Test User"
			}`,
			expectedStatus: 400,
			expectedBody: map[string]interface{}{
				"code": float64(400),
			},
		},
		{
			name: "昵称为空",
			requestBody: `{
				"email": "test@example.com",
				"password": "password123"
			}`,
			expectedStatus: 400,
			expectedBody: map[string]interface{}{
				"code": float64(400),
			},
		},
		{
			name: "重复邮箱注册",
			requestBody: `{
				"email": "duplicate@example.com",
				"password": "password123",
				"nickname": "Test User"
			}`,
			expectedStatus: 409,
			expectedBody: map[string]interface{}{
				"code":    float64(409),
				"message": "Email already registered",
			},
		},
	}

	// 先注册一个用户，用于测试重复邮箱
	db.Create(&user.User{
		Email:    "duplicate@example.com",
		Password: "hashedpassword",
		Nickname: "Existing User",
	})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/auth/register", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}

func TestAuthHandler_Login(t *testing.T) {
	// 设置测试数据库
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	// 自动迁移表结构
	db.AutoMigrate(&user.User{})

	// 创建 handler
	handler := NewAuthHandler(db)
	router := testutil.SetupTestRouter()
	router.POST("/auth/login", handler.Login)

	// 创建测试用户
	hashedPassword := "$2a$10$N9qo8uLOickgx2ZMRZoMy.MrqJq4v4v4v4v4v4v4v4v4v4v4v4v4" // "password123" bcrypt hash
	testUser := user.User{
		Email:    "test@example.com",
		Password: hashedPassword,
		Nickname: "Test User",
	}
	db.Create(&testUser)

	tests := []struct {
		name           string
		requestBody    string
		expectedStatus int
	}{
		{
			name: "成功注册新用户",
			requestBody: `{
				"email": "test@example.com",
				"password": "password123",
				"nickname": "Test User"
			}`,
			expectedStatus: 200,
		},
	}
		{
			name: "成功登录",
			requestBody: `{
				"email": "test@example.com",
				"password": "password123"
			}`,
			expectedStatus: 200,
		},
		{
			name: "邮箱不存在",
			requestBody: `{
				"email": "nonexistent@example.com",
				"password": "password123"
			}`,
			expectedStatus: 401,
		},
		{
			name: "密码错误",
			requestBody: `{
				"email": "test@example.com",
				"password": "wrongpassword"
			}`,
			expectedStatus: 401,
		},
		{
			name: "邮箱格式无效",
			requestBody: `{
				"email": "invalid-email",
				"password": "password123"
			}`,
			expectedStatus: 400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/auth/login", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}
