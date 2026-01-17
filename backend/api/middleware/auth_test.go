package middleware

import (
	"net/http/httptest"
	"os"
	"testing"

	"trpg-sync/backend/testutil"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	// 设置测试 JWT secret
	os.Setenv("JWT_SECRET", "test-secret-key-for-testing")
	defer os.Unsetenv("JWT_SECRET")

	// 设置测试路由
	gin.SetMode(gin.TestMode)
	router := gin.New()

	// 添加认证中间件
	router.Use(AuthMiddleware())

	// 添加测试端点
	router.GET("/protected", func(c *gin.Context) {
		userID, exists := c.Get("user_id")
		email, emailExists := c.Get("email")

		c.JSON(200, gin.H{
			"user_id": userID,
			"email":   email,
			"exists":  exists && emailExists,
		})
	})

	tests := []struct {
		name           string
		authHeader     string
		expectedStatus int
		checkBody      func(t *testing.T, body string)
	}{
		{
			name:           "缺少 Authorization 头",
			authHeader:     "",
			expectedStatus: 401,
		},
		{
			name:           "错误的 Authorization 头格式（缺少 Bearer）",
			authHeader:     "token-value",
			expectedStatus: 401,
		},
		{
			name:           "错误的 Authorization 头格式（不是 Bearer）",
			authHeader:     "Basic token-value",
			expectedStatus: 401,
		},
		{
			name:           "无效的 token",
			authHeader:     "Bearer invalid-token",
			expectedStatus: 401,
		},
		{
			name:           "有效的 token",
			authHeader:     "Bearer valid-token",
			expectedStatus: 401, // 因为 "valid-token" 不是有效的 JWT
		},
		{
			name:           "有效的 JWT token",
			authHeader:     "Bearer " + generateValidToken(t),
			expectedStatus: 200,
			checkBody: func(t *testing.T, body string) {
				assert.Contains(t, body, `"user_id":1`)
				assert.Contains(t, body, `"email":"test@example.com"`)
				assert.Contains(t, body, `"exists":true`)
			},
		},
		{
			name:           "过期的 token",
			authHeader:     "Bearer " + generateExpiredToken(t),
			expectedStatus: 401,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/protected", nil)
			if tt.authHeader != "" {
				req.Header.Set("Authorization", tt.authHeader)
			}
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			assert.Equal(t, tt.expectedStatus, rec.Code)

			if tt.checkBody != nil {
				tt.checkBody(t, rec.Body.String())
			}
		})
	}
}

func TestAuthMiddleware_MissingHeader(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")

	router := gin.New()
	router.Use(AuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	assert.Equal(t, 401, rec.Code)
	assert.Contains(t, rec.Body.String(), "Authorization header is required")
}

func TestAuthMiddleware_InvalidFormat(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")

	router := gin.New()
	router.Use(AuthMiddleware())
	router.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "success"})
	})

	tests := []string{
		"token-value",           // 缺少 Bearer
		"Basic token-value",     // 不是 Bearer
		"Bearer",                 // 缺少 token
		"Bearer token1 token2",   // 多余的部分
	}

	for _, authHeader := range tests {
		t.Run(authHeader, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/test", nil)
			req.Header.Set("Authorization", authHeader)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			assert.Equal(t, 401, rec.Code)
			assert.Contains(t, rec.Body.String(), "Invalid authorization header format")
		})
	}
}

// 生成有效的测试 token
func generateValidToken(t *testing.T) string {
	token, err := testutil.GenerateTestToken(1, "test@example.com", "test-secret-key-for-testing")
	assert.NoError(t, err)
	return token
}

// 生成过期的测试 token
func generateExpiredToken(t *testing.T) string {
	token, err := testutil.GenerateExpiredToken(1, "test@example.com", "test-secret-key-for-testing")
	assert.NoError(t, err)
	return token
}
