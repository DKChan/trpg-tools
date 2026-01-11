# TRPG-Sync 后端测试指南

## 测试框架

- **Go Testing**: Go 标准测试框架
- **Testify**: 断言和模拟库
- **GORM**: 数据库 ORM
- **SQLite**: 测试用内存数据库

## 测试结构

```
backend/
├── testutil/
│   ├── setup.go          # 数据库设置
│   ├── http.go           # HTTP 测试辅助
│   ├── jwt.go            # JWT token 生成
│   ├── mock.go           # 模拟数据
│   └── README.md         # 使用文档
├── api/
│   ├── v1/
│   │   ├── handlers/
│   │   │   ├── auth_test.go
│   │   │   ├── user_test.go
│   │   │   ├── room_test.go
│   │   │   └── character_test.go
│   └── middleware/
│       └── auth_test.go
└── infrastructure/
    └── database/
        └── database_test.go
```

## 运行测试

### 运行所有测试
```bash
cd backend
make test
```

### 运行测试并生成覆盖率报告
```bash
make test-cover
```

### 生成 HTML 覆盖率报告
```bash
make test-cover-html
```

### 运行特定包的测试
```bash
make test-package pkg=./api/v1/handlers
```

### 检查覆盖率是否达到阈值
```bash
make test-cover-check threshold=70
```

## 测试覆盖范围

### 已实现的测试

#### 高优先级模块
- ✅ `api/v1/handlers/auth.go` - 注册、登录
- ✅ `api/middleware/auth.go` - JWT 认证中间件
- ✅ `api/v1/handlers/room.go` - 房间创建、查询
- ✅ `api/v1/handlers/character.go` - 人物卡创建、查询

#### 中优先级模块
- ✅ `api/v1/handlers/user.go` - 用户资料管理
- ✅ `infrastructure/database/database.go` - 数据库连接

### 测试覆盖率目标

- 核心业务逻辑: ≥ 90%
- 中间件: ≥ 80%
- 整体覆盖率: ≥ 70%

## 测试最佳实践

1. **使用 table-driven tests** 减少重复代码
2. **使用内存 SQLite** 确保测试隔离
3. **使用 `t.Cleanup()`** 确保资源正确释放
4. **Mock 外部依赖** 避免测试受外部系统影响
5. **测试边界条件和错误场景** 提高测试覆盖率

## 添加新测试

1. 在对应目录创建 `*_test.go` 文件
2. 使用 `testutil.SetupTestDB()` 设置测试数据库
3. 编写测试用例覆盖正常和异常场景
4. 运行 `make test` 验证测试通过
5. 运行 `make test-cover` 检查覆盖率

## 已知问题

1. **UpdatePassword 未验证旧密码**: `api/v1/handlers/user.go:91-119`
2. **JWT_SECRET 默认值不安全**: `api/middleware/auth.go:36-39`
3. **部分功能未实现**: JoinRoom, LeaveRoom, DeleteRoom, UpdateCharacter, DeleteCharacter
