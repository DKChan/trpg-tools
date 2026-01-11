# 测试辅助工具包

本包提供测试辅助函数和工具，用于后端单元测试和集成测试。

## 目录结构

```
testutil/
├── setup.go    # 数据库设置和清理
├── http.go      # HTTP 测试辅助函数
├── jwt.go       # JWT token 生成
└── mock.go      # 模拟数据生成
```

## 使用示例

### 1. 设置测试数据库

```go
func TestMyHandler(t *testing.T) {
    db := testutil.SetupTestDB(t)
    defer db.Close()

    // 自动迁移表结构
    db.AutoMigrate(&domain.User{})
}
```

### 2. HTTP 测试

```go
func TestAuthHandler_Register(t *testing.T) {
    router := testutil.SetupTestRouter()
    db := testutil.SetupTestDB(t)

    handler := handlers.NewAuthHandler(db)
    router.POST("/auth/register", handler.Register)

    req := testutil.MakeJSONRequest("POST", "/auth/register", map[string]interface{}{
        "email":    "test@example.com",
        "password": "password123",
        "nickname": "Test User",
    })

    rec := httptest.NewRecorder()
    router.ServeHTTP(rec, req)

    assert.Equal(t, 200, testutil.GetResponseCode(rec))
}
```

### 3. 生成测试 Token

```go
func TestProtectedEndpoint(t *testing.T) {
    token, _ := testutil.GenerateTestToken(1, "test@example.com", "test-secret")

    req := testutil.MakeJSONRequest("GET", "/user/profile", nil)
    req.Header.Set("Authorization", "Bearer "+token)

    // ... 测试逻辑
}
```

### 4. 使用模拟数据

```go
func TestRoomHandler(t *testing.T) {
    mockRoom := testutil.NewMockRoom()

    // 创建测试数据
    db.Create(&mockRoom)

    // 测试逻辑
}
```

## 测试命令

```bash
# 运行所有测试
make test

# 运行测试并生成覆盖率报告
make test-cover

# 生成 HTML 覆盖率报告
make test-cover-html

# 检查覆盖率是否达到阈值（70%）
make test-cover-check threshold=70

# 运行特定包的测试
make test-package pkg=./api/v1/handlers
```

## 注意事项

1. 测试数据库使用内存 SQLite，每个测试独立运行
2. 使用 `t.Cleanup()` 确保资源正确释放
3. Mock 数据仅用于测试，不应在生产代码中使用
4. JWT secret 在测试环境中应使用专用密钥
