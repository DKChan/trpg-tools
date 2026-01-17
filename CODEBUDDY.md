# CODEBUDDY.md This file provides guidance to CodeBuddy when working with code in this repository.

## 项目概述

TRPG个人人物卡管理工具 - 专为TRPG玩家设计的单机人物卡管理应用，支持多规则系统的人物卡创建和管理。

- **后端**: Go 1.24+ | Gin | GORM | SQLite
- **前端**: React 18 | TypeScript 5 | Vite | Zustand | Ant Design | TailwindCSS

## 常用命令

### 后端开发 (backend/)

```bash
# 启动开发服务器
cd backend
go run main.go

# 依赖管理
go mod tidy
go mod download

# 代码格式化
gofmt -w .

# 代码检查
golangci-lint run
```

### 前端开发 (frontend/)

```bash
# 安装依赖
cd frontend
npm install

# 启动开发服务器 (http://localhost:5173)
npm run dev

# 构建生产版本
npm run build

# 预览生产构建
npm run preview

# 代码检查
npm run lint
```

### 数据库

项目使用 GORM AutoMigrate 进行表结构同步，启动后端时会自动创建/更新表。如需手动控制迁移，可在 `main.go` 中添加迁移代码。

## 架构概览

### 后端架构 (DDD - 领域驱动设计)

```
backend/
├── api/              # 接口层 - HTTP路由和控制器
│   ├── v1/          # API v1版本路由和处理器
│   │   ├── handlers/ # 请求处理器 (RoomHandler, CharacterHandler)
│   │   └── routes.go # 路由注册
│   └── middleware/   # 中间件 (CORS, Logger, Recovery)
├── domain/          # 领域层 - 核心业务逻辑和实体
│   ├── character/   # 人物卡领域 (CharacterCard)
│   └── room/        # 房间领域 (Room)
├── infrastructure/  # 基础设施层 - 技术实现
│   ├── config/      # 配置管理
│   └── database/    # 数据库连接 (InitDB)
└── main.go          # 应用入口
```

**核心设计原则**：
- **依赖方向**: api → domain; infrastructure → domain（domain层不依赖任何外层）
- **实体定义**: 领域实体在 `domain/*/` 包中，使用 GORM 标签
- **API层**: handlers 通过 GORM DB 直接操作实体
- **中间件链**: CORS → Logger → Recovery
- **单机应用**: 使用SQLite数据库，无需认证

### 前端架构

```
frontend/
├── src/
│   ├── components/  # 可复用组件 (Layout等)
│   ├── pages/       # 页面组件 (Home, Room, RoomDetail, CharacterCard)
│   ├── store/       # Zustand状态管理 (roomStore, characterStore)
│   ├── services/    # API封装 (api.ts统一实例, index.ts具体接口)
│   ├── hooks/       # 自定义Hooks（待实现）
│   ├── utils/       # 工具函数（待实现）
│   ├── types/       # TypeScript类型 (Room, CharacterCard, ApiResponse)
│   ├── styles/      # 全局样式 (TailwindCSS)
│   ├── App.tsx      # 应用根组件
│   └── main.tsx     # 应用入口
├── index.html       # HTML模板
├── vite.config.ts   # Vite配置
└── package.json
```

**核心设计原则**：
- **状态管理**: Zustand + persist中间件，本地存储数据
- **API客户端**: Axios实例统一处理请求
- **组件设计**: 函数组件优先，props使用接口定义
- **样式策略**: Ant Design组件库 + TailwindCSS工具类
- **单机应用**: 无认证，直接使用API

## 数据模型

### 核心实体 (domain/)

- **Room**: 房间信息 (ID, Name, Description, RuleSystem)
- **CharacterCard**: 人物卡 (ID, RoomID, 属性/技能/装备等DND5e字段)

### 表命名规范

- 表名：复数形式，下划线分隔（users, rooms, room_members, character_cards）
- 字段：下划线分隔，必需字段 `gorm:"not null"`
- 索引：uniqueIndex, 普通索引在需要性能的字段上

## API接口规范

### RESTful风格

- 版本前缀：`/api/v1/`
- 房间路由：`GET/POST /api/v1/rooms`, `GET/PUT/DELETE /api/v1/rooms/:id`
- 人物卡路由：`GET/POST /api/v1/rooms/:roomId/characters`, `GET/PUT/DELETE /api/v1/rooms/:roomId/characters/:id`

### 响应格式

```json
{
  "code": 200,
  "message": "success",
  "data": {}
}
```

### 数据存储

- 使用SQLite本地数据库
- 数据文件为 `sqlite.db`
- GORM自动管理表结构

## 关键开发注意事项

### 后端开发

1. **新增API端点**: 在 `api/v1/handlers/` 创建handler，在 `api/v1/routes.go` 注册路由
2. **数据库迁移**: 如修改实体结构，AutoMigrate会自动更新表结构
3. **错误处理**: 统一使用gin.JSON返回错误，避免直接暴露内部错误
4. **SQLite使用**: 所有数据存储在本地sqlite.db文件

### 前端开发

1. **新增页面**: 在 `pages/` 创建组件，在 `App.tsx` 添加路由
2. **API调用**: 在 `services/index.ts` 定义接口函数，使用 `api.ts` 的axios实例
3. **状态管理**: 在 `store/` 创建Zustand store，使用persist中间件本地持久化
4. **类型定义**: 在 `types/index.ts` 添加接口，保持前后端类型一致

## 环境配置

### 后端

复制 `backend/.env.example` 为 `backend/.env`，配置数据库连接和JWT密钥。

### 前端

在 `frontend/` 创建 `.env` 文件（如需要自定义API地址），默认使用 `/api/v1` 代理。
