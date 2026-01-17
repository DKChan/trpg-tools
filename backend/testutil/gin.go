package testutil

import (
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/gin-gonic/gin"
)

// CreateTestContext 创建测试用的 gin.Context
func CreateTestContext(recorder *httptest.ResponseRecorder, method, path string, body string) *gin.Context {
	c, _ := gin.CreateTestContext(recorder)

	// 设置请求
	req := httptest.NewRequest(method, path, nil)
	if body != "" {
		req = httptest.NewRequest(method, path, strings.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
	}
	c.Request = req

	return c
}

// CreateTestContextWithRequest 创建测试用的 gin.Context 并设置自定义请求
func CreateTestContextWithRequest(recorder *httptest.ResponseRecorder, req *http.Request) *gin.Context {
	c, _ := gin.CreateTestContext(recorder)
	c.Request = req
	return c
}
