package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"trpg-sync/backend/domain/user"
	"trpg-sync/backend/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestUserHandler_GetProfile(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&user.User{})

	handler := NewUserHandler(db)
	router := testutil.SetupTestRouter()
	router.GET("/user/profile", handler.GetProfile)

	// 创建测试用户
	testUser := user.User{
		Email:    "test@example.com",
		Password: "hashedpassword",
		Nickname: "Test User",
		Avatar:   "https://example.com/avatar.png",
	}
	db.Create(&testUser)

	tests := []struct {
		name           string
		userID         uint
		expectedStatus int
	}{
		{
			name:           "成功获取用户资料",
			userID:         1,
			expectedStatus: 200,
		},
		{
			name:           "用户不存在",
			userID:         999,
			expectedStatus: 404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/user/profile", nil)
			rec := httptest.NewRecorder()

			c := gin.Context{}
			c.Set("user_id", tt.userID)
			c.Request = req
			c.Writer = rec

			handler.GetProfile(&c)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			if tt.expectedStatus == 200 {
				assert.Contains(t, rec.Body.String(), "test@example.com")
				assert.Contains(t, rec.Body.String(), "Test User")
			}
		})
	}
}

func TestUserHandler_UpdateProfile(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&user.User{})

	handler := NewUserHandler(db)

	// 创建测试用户
	testUser := user.User{
		Email:    "test@example.com",
		Password: "hashedpassword",
		Nickname: "Test User",
	}
	db.Create(&testUser)

	tests := []struct {
		name           string
		requestBody    string
		userID         uint
		expectedStatus int
	}{
		{
			name: "成功更新昵称",
			requestBody: `{
				"nickname": "Updated Name"
			}`,
			userID:         1,
			expectedStatus: 200,
		},
		{
			name: "成功更新头像",
			requestBody: `{
				"avatar": "https://example.com/new-avatar.png"
			}`,
			userID:         1,
			expectedStatus: 200,
		},
		{
			name: "成功更新昵称和头像",
			requestBody: `{
				"nickname": "New Name",
				"avatar": "https://example.com/avatar.png"
			}`,
			userID:         1,
			expectedStatus: 200,
		},
		{
			name:           "空请求体（不更新任何字段）",
			requestBody:    `{}`,
			userID:         1,
			expectedStatus: 200,
		},
		{
			name: "用户不存在",
			requestBody: `{
				"nickname": "Test"
			}`,
			userID:         999,
			expectedStatus: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("PUT", "/user/profile", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(rec)
			c.Request = req
			c.Set("user_id", tt.userID)

			handler.UpdateProfile(c)

			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}

func TestUserHandler_UpdatePassword(t *testing.T) {
	db := testutil.SetupTestDB(t)
	defer func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	}()

	db.AutoMigrate(&user.User{})

	handler := NewUserHandler(db)

	// 创建测试用户
	testUser := user.User{
		Email:    "test@example.com",
		Password: "oldhashedpassword",
		Nickname: "Test User",
	}
	db.Create(&testUser)

	tests := []struct {
		name           string
		requestBody    string
		userID         uint
		expectedStatus int
	}{
		{
			name: "成功更新密码（注意：当前实现未验证旧密码）",
			requestBody: `{
				"old_password": "oldpassword",
				"new_password": "newpassword123"
			}`,
			userID:         1,
			expectedStatus: 200,
		},
		{
			name: "新密码少于6位",
			requestBody: `{
				"old_password": "oldpassword",
				"new_password": "123"
			}`,
			userID:         1,
			expectedStatus: 400,
		},
		{
			name: "缺少旧密码",
			requestBody: `{
				"new_password": "newpassword123"
			}`,
			userID:         1,
			expectedStatus: 400,
		},
		{
			name: "用户不存在",
			requestBody: `{
				"old_password": "oldpassword",
				"new_password": "newpassword123"
			}`,
			userID:         999,
			expectedStatus: 404,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("PUT", "/user/password", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			rec := httptest.NewRecorder()

			c, _ := gin.CreateTestContext(rec)
			c.Request = req
			c.Set("user_id", tt.userID)

			handler.UpdatePassword(c)

			assert.Equal(t, tt.expectedStatus, rec.Code)
		})
	}
}
