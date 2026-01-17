# TRPG-Tools 个人人物卡管理工具

一个专为TRPG（桌面角色扮演游戏）玩家设计的个人人物卡管理工具，支持创建、编辑和管理多个规则系统下的人物卡数据。

## 项目概述

TRPG-Tools是一个轻量级的Web应用，专为个人使用设计：
- 🏠 创建和管理多个游戏房间（用于分类不同战役）
- 🎭 创建和编辑多规则系统的人物卡（目前支持D&D 5e）
- 📊 便捷管理人物卡属性、技能、装备等信息
- 💾 本地数据存储，无需网络同步

## 技术栈

### 后端
- **语言**: Go 1.24+
- **Web框架**: Gin
- **ORM**: GORM
- **数据库**: SQLite（单机使用，无需外部数据库）
- **测试**: Testify + go test

### 前端
- **框架**: React 18
- **语言**: TypeScript 5
- **构建工具**: Vite
- **状态管理**: Zustand
- **UI组件**: Ant Design + TailwindCSS
- **测试**: Vitest

## 快速开始

### 前置要求

- Go 1.24 或更高版本
- Node.js 18+ 和 npm

### 后端设置

1. **克隆仓库**
```bash
git clone https://github.com/DKChan/trpg-tools.git
cd trpg-tools/backend
```

2. **安装依赖**
```bash
# 使用 go mod 安装Go依赖
go mod tidy
go mod download
```

3. **启动后端服务**
```bash
# 开发模式
go run main.go

# 或使用Makefile
make run
```

后端服务将在 `http://localhost:8080` 启动，自动创建SQLite数据库文件

### 前端设置

1. **进入前端目录**
```bash
cd trpg-sync/frontend
```

2. **安装依赖**
```bash
npm install
```

3. **配置API地址** (可选)
```bash
# 前端默认使用代理到后端
# 如需修改，编辑 vite.config.ts
```

4. **启动开发服务器**
```bash
npm run dev
```

前端应用将在 `http://localhost:5173` 启动

### 生产环境部署

#### 后端部署

1. **构建可执行文件**
```bash
cd backend
go build -o trpg-tools main.go
```

2. **运行服务**
```bash
./trpg-tools
# 程序会自动在当前目录创建sqlite.db文件
```

#### 前端部署

1. **构建生产版本**
```bash
cd frontend
npm run build
```

2. **部署到Web服务器**
```bash
# 构建产物在 dist/ 目录
# 可以部署到Nginx、Apache或静态托管服务
```

## 依赖管理

### 更新后端依赖

```bash
cd backend

# 更新所有依赖到最新版本
go get -u ./...

# 清理未使用的依赖
go mod tidy

# 验证依赖
go mod verify
```

### 更新前端依赖

```bash
cd frontend

# 检查过期依赖
npm outdated

# 更新依赖
npm update

# 更新到最新版本（谨慎使用）
npx npm-check-updates -u
npm install
```

## 测试

### 运行后端测试

```bash
cd backend

# 运行所有测试
go test ./...

# 运行带覆盖率的测试
go test -cover ./...

# 生成HTML覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html

# 运行特定包的测试
go test ./api/v1/handlers -v

# 使用测试脚本
# Windows
.\test.ps1 all --cover
# Linux/Mac
./test.sh all --cover
```

### 运行前端测试

```bash
cd frontend

# 运行所有测试
npm test

# 运行测试并生成覆盖率
npm run test:coverage

# 运行测试一次（不监听）
npm run test:run
```

## 开发指南

### 项目结构

```
trpg-sync/
├── backend/                 # 后端代码
│   ├── api/                # API层
│   │   ├── middleware/     # 中间件
│   │   └── v1/            # API v1版本
│   │       ├── handlers/  # 请求处理器
│   │       └── routes.go  # 路由配置
│   ├── domain/            # 领域层
│   │   ├── character/     # 人物卡领域
│   │   ├── room/         # 房间领域
│   │   └── user/         # 用户领域
│   ├── application/      # 应用层（预留）
│   ├── infrastructure/   # 基础设施层
│   │   ├── config/      # 配置管理
│   │   └── database/    # 数据库连接
│   ├── testutil/       # 测试工具
│   └── main.go         # 应用入口
├── frontend/            # 前端代码
│   ├── src/
│   │   ├── components/  # 组件
│   │   ├── pages/      # 页面
│   │   ├── services/   # API服务
│   │   └── store/      # 状态管理
│   └── public/         # 静态资源
├── docs/               # 文档
└── .github/workflows/ # CI/CD配置
```

### 添加新功能

#### 后端

1. **定义领域模型** - 在 `backend/domain/` 中添加实体
2. **创建处理器** - 在 `backend/api/v1/handlers/` 中实现业务逻辑
3. **注册路由** - 在 `backend/api/v1/routes.go` 中添加路由
4. **编写测试** - 在 `backend/api/v1/handlers/*_test.go` 中编写单元测试

#### 前端

1. **创建页面组件** - 在 `frontend/src/pages/` 中添加页面
2. **添加路由** - 在 `frontend/src/App.tsx` 中配置路由
3. **创建API服务** - 在 `frontend/src/services/index.ts` 中添加接口
4. **管理状态** - 在 `frontend/src/store/` 中添加Zustand store

### 代码规范

#### Go代码规范
- 使用 `gofmt` 格式化代码
- 遵循Go命名约定（PascalCase导出，camelCase私有）
- 错误处理必须显式，不使用panic
- 单元测试覆盖率 > 70%

#### TypeScript代码规范
- 使用TypeScript严格模式
- 组件使用函数式组件和Hooks
- 类型定义使用接口（interface）
- 使用ESLint进行代码检查

### 数据库迁移

项目使用GORM的AutoMigrate自动管理表结构。如需手动迁移：

```go
// 在main.go中添加
db.AutoMigrate(
    &user.User{},
    &room.Room{},
    &room.RoomMember{},
    &character.CharacterCard{},
)
```

## 环境变量

### 后端环境变量 (.env)

```bash
# 服务器配置（可选）
SERVER_PORT=8080

# 数据库配置（使用SQLite，无需配置）
# DB_PATH=./sqlite.db  # 可自定义数据库文件路径

# 日志级别
LOG_LEVEL=info
```

### 前端环境变量 (.env.local)

```bash
# API基础URL
VITE_API_BASE_URL=http://localhost:8080/api/v1
```

## 常见问题

### 数据库连接失败
- 检查PostgreSQL服务是否运行
- 验证数据库配置是否正确
- 确保数据库已创建

### JWT认证失败
- 检查JWT_SECRET是否配置
- 确保token未过期
- 验证Authorization头格式：Bearer <token>

### 前端无法连接后端
- 检查后端服务是否运行
- 确认CORS配置
- 检查代理配置（vite.config.ts）

### 测试失败
- 确保所有依赖已安装
- 检查测试数据库配置
- 运行 `go mod tidy` 清理依赖

## 贡献指南

1. Fork 仓库
2. 创建特性分支 (`git checkout -b feature/amazing-feature`)
3. 提交更改 (`git commit -m 'feat: 添加 amazing feature'`)
4. 推送到分支 (`git push origin feature/amazing-feature`)
5. 创建 Pull Request

## 许可证

MIT License

## 联系方式

- 项目地址: https://github.com/DKChan/trpg-tools
- 问题反馈: https://github.com/DKChan/trpg-tools/issues

## 更新日志

### v0.2.0 (2026-01-18)
- ✅ 调整为个人工具定位
- ✅ 移除用户系统和认证
- ✅ 简化房间功能（仅分类用途）
- ✅ 保留多规则系统支持
- ✅ 本地SQLite数据存储

### v0.1.0 (2026-01-17)
- ✅ 基础架构搭建
- ✅ D&D 5e人物卡管理
- ✅ 房间分类功能
- ✅ 单元测试覆盖率达到57.1%
- ✅ CI/CD工作流
