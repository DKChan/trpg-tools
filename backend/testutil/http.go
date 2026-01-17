package testutil

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

// SetupTestRouter 设置测试路由
func SetupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	return gin.New()
}

// MakeJSONRequest 创建 JSON 格式的测试请求
func MakeJSONRequest(method, url string, body interface{}) (*http.Request, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req := httptest.NewRequest(method, url, bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

// ParseResponse 解析响应体
func ParseResponse(rec *httptest.ResponseRecorder, v interface{}) error {
	return json.Unmarshal(rec.Body.Bytes(), v)
}

// GetResponseBody 获取响应体字符串
func GetResponseBody(rec *httptest.ResponseRecorder) string {
	return rec.Body.String()
}

// GetResponseCode 获取响应状态码
func GetResponseCode(rec *httptest.ResponseRecorder) int {
	return rec.Code
}
