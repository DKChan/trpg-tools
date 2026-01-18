# TRPG-Tools 个人人物卡管理工具

一个专为TRPG（桌面角色扮演游戏）玩家设计的个人人物卡管理工具，支持创建、编辑和管理多个规则系统下的人物卡数据。

## 项目概述

TRPG-Tools是一个轻量级的Web应用，专为个人使用设计：
|- 🏠 创建和管理多个游戏房间（用于分类不同战役）
|- 🎭 创建和编辑多规则系统的人物卡（目前支持D&D 5e）
|- 📊 便捷管理人物卡属性、技能、装备等信息
|- 💾 混合存储方案（SQLite + JSON 文件）
|- 📂 人物卡以 JSON 文件存储，易于备份和版本控制

## 技术栈

### 后端
|- **语言**: Go 1.24+
|- **Web框架**: Gin
|- **ORM**: GORM（仅用于 Room）
|- **数据库**: SQLite（存储房间信息）
|- **文件存储**: JSON 文件（存储人物卡）
|- **测试**: Testify + go test

### 前端
|- **框架**: React 18
|- **语言**: TypeScript 5
|- **构建工具**: Vite
|- **状态管理**: Zustand
|- **UI组件**: Ant Design + TailwindCSS
|- **测试**: Vitest

## 数据存储方案

### 混合存储架构

项目采用**混合存储方案**，结合关系型数据库和文件存储的优势：

```
data/
  rooms/
    {room_id}/
      characters/
        {character_id}.json    # 每个人物卡一个 JSON 文件
```

**SQLite 数据库**（`sqlite.db`）：
- 存储房间信息（ID、名称、描述、规则系统）
- 支持高效查询和索引
- 零配置，自动创建

**JSON 文件存储**：
- 存储人物卡数据
- 易于查看和编辑（文本格式）
- 支持版本控制（git）
- 易于备份和迁移

**优势**：
- ✅ 查询效率高（房间列表）
- ✅ 数据可读性强（人物卡为 JSON）
- ✅ 备份简单（复制 data 目录）
- ✅ 版本控制友好（可用 git 管理人物卡）
- ✅ 灵活性高（易于导出和导入）

## 快速开始

### 前置要求

- Go 1.24 或更高版本
- Node.js 18+ 和 npm

### 一键部署（推荐）

**Windows:**
```powershell
.\build.ps1
cd backend
.\trpg-tools.exe
```

**Linux/Mac:**
```bash
./build.sh
cd backend
./trpg-tools
```

应用将在 `http://localhost:8080` 启动，所有功能都已集成！

### 开发模式

如需开发调试，可以分别启动前后端：

**后端:**
```bash
cd backend
go run main.go
# 服务在 http://localhost:8080
```

**前端:**
```bash
cd frontend
npm run dev
# 服务在 http://localhost:5173
```

详细的部署说明请查看 [DEPLOY.md](DEPLOY.md)

## 生产环境部署

### 一键构建（推荐）

使用一键构建脚本，自动完成前端和后端的构建：

**Windows:**
```powershell
.\build.ps1
```

**Linux/Mac:**
```bash
./build.sh
```

构建完成后，只需运行一个可执行文件：

```bash
cd backend
# Windows
trpg-tools.exe
# Linux/Mac
./trpg-tools
```

服务启动后访问: http://localhost:8080

### 手动部署

如果需要手动构建：

1. **构建前端**
```bash
cd frontend
npm install
npm run build
# 前端产物自动嵌入到backend/dist目录
```

2. **构建后端**
```bash
cd backend
go mod tidy
go build -o trpg-tools main.go
```

3. **运行服务**
```bash
./trpg-tools
```

### 部署架构

项目使用 **Go Embed** 技术，将前端静态资源嵌入到Go二进制文件中：

- ✅ **单一可执行文件**: 只需部署一个二进制文件
- ✅ **一键启动**: 无需额外配置Web服务器
- ✅ **包含前端**: 所有静态资源已打包
- ✅ **端口统一**: 前后端共享同一个端口(默认8080)
- ✅ **自动路由**: API路由 `/api/*` 和前端路由自动分离

### 运行时生成的文件

程序运行时会自动创建：

- `sqlite.db`: SQLite数据库文件（存储房间信息）
- `data/`: 数据目录（存储人物卡JSON文件）

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
│   │   └── room/         # 房间领域
│   ├── infrastructure/   # 基础设施层
│   │   ├── config/      # 配置管理
│   │   ├── database/    # 数据库连接
│   │   └── storage/     # 文件存储
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

### 数据存储规范

- **房间数据**：存储在 SQLite `rooms` 表
- **人物卡数据**：存储为 JSON 文件，路径 `data/rooms/{room_id}/characters/{character_id}.json`
- **文件操作**：使用 `storage.CharacterStorage` 统一管理
- **备份**：可以备份整个 `data/` 目录

## 环境变量

### 后端环境变量 (.env)

```bash
# 服务器配置（可选）
SERVER_PORT=8080

# 数据库配置（使用SQLite，无需配置）
DB_PATH=./sqlite.db  # 可自定义数据库文件路径

# 日志级别
LOG_LEVEL=info
```

### 前端环境变量 (.env.local)

```bash
# API基础URL
VITE_API_BASE_URL=http://localhost:8080/api/v1
```

## 常见问题

### 数据连接失败
- 确保data目录有读写权限
- 验证SQLite数据库文件是否可写

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

### v0.4.0 (2026-01-18)
- ✅ 实现一键部署（Go Embed 技术）
- ✅ 前端静态资源嵌入到Go二进制文件
- ✅ 单一可执行文件，无需额外Web服务器
- ✅ 添加一键构建脚本（Windows/Linux/Mac）
- ✅ 统一前后端端口，简化部署

### v0.3.0 (2026-01-17)
- ✅ 采用混合存储方案（SQLite + JSON 文件）
- ✅ 房间使用 SQLite 存储
- ✅ 人物卡使用 JSON 文件存储
- ✅ 添加文件导出和备份功能
- ✅ 优化数据备份和版本控制

### v0.2.0 (2026-01-18)
- ✅ 调整为个人工具定位
- ✅ 移除用户系统和认证
- ✅ 简化房间功能（仅分类用途）
- ✅ 保留多规则系统支持
- ✅ 本地SQLite数据存储
- ✅ 添加房间搜索功能
- ✅ 添加人物卡删除功能

### v0.1.0 (2026-01-17)
- ✅ 基础架构搭建
- ✅ D&D 5e人物卡管理
- ✅ 房间分类功能
- ✅ 单元测试覆盖率达到57.1%
- ✅ CI/CD工作流
