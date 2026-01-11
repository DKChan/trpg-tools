# TRPG-Sync 测试运行指南

本文档说明如何运行 TRPG-Sync 项目的单元测试。

## 前置条件

### 后端测试
- Go 1.24 或更高版本
- testify 依赖（已在 go.mod 中配置）
- SQLite（用于测试）

### 前端测试
- Node.js 18 或更高版本
- npm 或 yarn

## 后端测试

### Windows (PowerShell)

运行所有测试：
```powershell
cd backend
.\test.ps1 all
```

运行特定测试：
```powershell
# 认证测试
.\test.ps1 auth

# 房间测试
.\test.ps1 room

# 人物卡测试
.\test.ps1 character

# 用户测试
.\test.ps1 user

# 中间件测试
.\test.ps1 middleware
```

生成覆盖率报告：
```powershell
.\test.ps1 all --cover
```

### Linux/Mac (Bash)

```bash
cd backend

# 运行所有测试
make test

# 运行特定测试
make test-cover

# 生成 HTML 覆盖率报告
make test-cover-html
```

或使用测试脚本：
```bash
# 运行所有测试
chmod +x test.sh
./test.sh all

# 运行特定测试
./test.sh auth
./test.sh room
./test.sh character
./test.sh user
./test.sh middleware

# 生成覆盖率报告
./test.sh all --cover
```

### 使用 Go 标准命令

```bash
cd backend

# 运行所有测试
go test -v ./...

# 运行特定包的测试
go test -v ./api/v1/handlers

# 运行特定测试函数
go test -v ./api/v1/handlers -run TestAuthHandler_Register

# 运行测试并生成覆盖率
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
go tool cover -html=coverage.out -o coverage.html
```

## 前端测试

### 安装依赖
```bash
cd frontend
npm install
```

### 运行测试
```bash
# 运行所有测试
npm test

# 运行测试并监听文件变化
npm test:ui

# 运行测试一次
npm test:run

# 生成测试覆盖率报告
npm test:coverage
```

### 查看测试覆盖率报告
测试覆盖率报告将生成在：
- 文本报告：`frontend/coverage/coverage-summary.txt`
- HTML 报告：`frontend/coverage/index.html`

## 测试覆盖率目标

- **后端整体覆盖率**: ≥ 70%
- **前端整体覆盖率**: ≥ 70%
- **核心业务逻辑**: ≥ 90%

## 已修复的测试文件

### 后端测试文件
- ✅ `api/v1/handlers/auth_test.go` - 认证处理
- ✅ `api/v1/handlers/room_test.go` - 房间处理
- ✅ `api/v1/handlers/character_test.go` - 人物卡处理
- ✅ `api/v1/handlers/user_test.go` - 用户处理
- ✅ `api/middleware/auth_test.go` - 认证中间件
- ✅ `infrastructure/database/database_test.go` - 数据库连接

### 前端测试文件
- ✅ `store/authStore.test.ts` - 状态管理
- ✅ `services/api.test.ts` - API 拦截器
- ✅ `services/index.test.ts` - API 服务
- ✅ `pages/Login.test.tsx` - 登录页面
- ✅ `pages/Register.test.tsx` - 注册页面
- ✅ `pages/Home.test.tsx` - 房间列表页面

## 故障排查

### 后端测试失败

1. **testify 依赖问题**
   ```bash
   go get github.com/stretchr/testify@latest
   ```

2. **数据库连接问题**
   - 测试使用内存 SQLite，不需要实际数据库
   - 确保测试函数正确设置数据库

3. **gin.Context 创建错误**
   - 使用 `gin.CreateTestContext(rec)` 而不是 `c.Writer = rec`
   - 已在所有测试文件中修复

### 前端测试失败

1. **依赖未安装**
   ```bash
   npm install
   ```

2. **TypeScript 类型错误**
   - 检查 mock 类型是否正确
   - 确保 tsconfig.json 配置正确

3. **组件渲染问题**
   - 确保正确 mock 外部依赖
   - 检查 React Testing Library 版本

## CI/CD 集成

GitHub Actions 工作流已配置在 `.github/workflows/ci.yml`，包括：
- 后端测试任务
- 前端测试任务
- 代码检查任务
- 构建任务

每次 push 或 pull request 时自动运行。

## 相关文档

- 后端测试指南: `backend/TESTING.md`
- 前端测试指南: `frontend/TESTING.md`
- 测试修复总结: `backend/TEST_FIXES.md`
- 测试实施总结: `TESTING_SUMMARY.md`
