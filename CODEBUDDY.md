# CODEBUDDY.md This file provides guidance to CodeBuddy when working with code in this repository.

## 项目概述

TRPG人物卡共享平台 - 支持DM和玩家在线管理和共享人物卡信息，实现实时同步的桌面角色扮演游戏平台。

- **后端**: Go 1.24+ | Gin | GORM | PostgreSQL | JWT
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
│   │   ├── handlers/ # 请求处理器 (AuthHandler, UserHandler, RoomHandler, CharacterHandler)
│   │   └── routes.go # 路由注册
│   ├── middleware/   # 中间件 (Auth, CORS, Logger, Recovery)
│   └── router.go    # 路由配置（备用，当前使用 v1/routes.go）
├── application/      # 应用层 - 业务用例编排（待实现）
│   ├── commands/    # 命令处理（写操作）
│   └── queries/     # 查询处理（读操作）
├── domain/          # 领域层 - 核心业务逻辑和实体
│   ├── aggregate/   # 聚合根（待实现）
│   ├── character/   # 人物卡领域 (CharacterCard)
│   ├── entity/      # 基础实体（待实现）
│   ├── repository/  # 仓储接口定义（待实现）
│   ├── room/        # 房间领域 (Room, RoomMember)
│   ├── service/     # 领域服务（待实现）
│   ├── user/        # 用户领域 (User)
│   └── vo/          # 值对象（待实现）
├── infrastructure/  # 基础设施层 - 技术实现
│   ├── cache/       # 缓存（待实现）
│   ├── config/      # 配置管理
│   ├── database/    # 数据库连接 (InitDB)
│   └── utils/       # 工具函数（待实现）
├── models/          # 数据模型（备用，实际使用 domain/ 中的结构体）
└── main.go          # 应用入口
```

**核心设计原则**：
- **依赖方向**: infrastructure → application → domain（domain层不依赖任何外层）
- **实体定义**: 领域实体在 `domain/*/` 包中，使用 GORM 标签
- **API层**: handlers 通过 GORM DB 直接操作实体（当前实现），后续应引入 repository 模式
- **中间件链**: CORS → Logger → Recovery → Auth（按需）
- **JWT认证**: 通过 `AuthMiddleware` 验证 Token，解析用户信息存入 context

### 前端架构

```
frontend/
├── src/
│   ├── components/  # 可复用组件 (Layout等)
│   ├── pages/       # 页面组件 (Login, Register, Home, Room, RoomDetail, CharacterCard)
│   ├── store/       # Zustand状态管理 (authStore)
│   ├── services/    # API封装 (api.ts统一实例, index.ts具体接口)
│   ├── hooks/       # 自定义Hooks（待实现）
│   ├── utils/       # 工具函数（待实现）
│   ├── types/       # TypeScript类型 (User, Room, CharacterCard, ApiResponse)
│   ├── styles/      # 全局样式 (TailwindCSS)
│   ├── App.tsx      # 应用根组件
│   └── main.tsx     # 应用入口
├── index.html       # HTML模板
├── vite.config.ts   # Vite配置
└── package.json
```

**核心设计原则**：
- **状态管理**: Zustand + persist中间件，状态按功能模块划分
- **API客户端**: Axios实例统一拦截器处理 JWT 和 401 错误
- **路由守卫**: 在组件中通过 `useAuthStore` 检查登录状态
- **组件设计**: 函数组件优先，props使用接口定义，避免不必要的重渲染
- **样式策略**: Ant Design组件库 + TailwindCSS工具类

## 数据模型

### 核心实体 (domain/)

- **User**: 用户信息 (ID, Email, Password, Nickname, Avatar)
- **Room**: 房间信息 (ID, Name, Description, RuleSystem, Password, InviteCode, DMID, MaxPlayers, IsPublic)
- **RoomMember**: 房间成员关系 (ID, RoomID, UserID, Role: 'dm'|'player')
- **CharacterCard**: 人物卡 (ID, UserID, RoomID, 属性/技能/装备等DND5e字段)

### 表命名规范

- 表名：复数形式，下划线分隔（users, rooms, room_members, character_cards）
- 字段：下划线分隔，必需字段 `gorm:"not null"`
- 索引：uniqueIndex, 普通索引在需要性能的字段上

## API接口规范

### RESTful风格

- 版本前缀：`/api/v1/`
- 认证路由：`POST /api/v1/auth/login`, `POST /api/v1/auth/register`
- 用户路由：`GET/PUT /api/v1/user/profile`, `PUT /api/v1/user/password`
- 房间路由：`GET/POST /api/v1/rooms`, `GET/PUT/DELETE /api/v1/rooms/:id`, `POST /api/v1/rooms/:id/join`
- 人物卡路由：`GET/POST /api/v1/rooms/:roomId/characters`, `GET/PUT/DELETE /api/v1/rooms/:roomId/characters/:id`

### 响应格式

```json
{
  "code": 200,
  "message": "success",
  "data": {}
}
```

### 认证机制

- JWT Token 存储在 Zustand persist store
- Axios请求拦截器自动添加 `Authorization: Bearer {token}`
- 401响应时自动登出并跳转登录页

## 关键开发注意事项

### 后端开发

1. **新增API端点**: 在 `api/v1/handlers/` 创建handler，在 `api/v1/routes.go` 注册路由
2. **数据库迁移**: 如修改实体结构，AutoMigrate会自动更新表结构（开发环境）
3. **认证保护**: 在路由组上添加 `room.Use(AuthMiddleware())`
4. **错误处理**: 统一使用gin.JSON返回错误，避免直接暴露内部错误
5. **Context传递**: 从gin.Context获取当前用户ID: `userID := c.GetUint("user_id")`

### 前端开发

1. **新增页面**: 在 `pages/` 创建组件，在 `App.tsx` 添加路由
2. **API调用**: 在 `services/index.ts` 定义接口函数，使用 `api.ts` 的axios实例
3. **状态管理**: 在 `store/` 创建Zustand store，使用persist中间件持久化
4. **类型定义**: 在 `types/index.ts` 添加接口，保持前后端类型一致
5. **认证检查**: 使用 `useAuthStore.getState().token` 判断登录状态

## 环境配置

### 后端

复制 `backend/.env.example` 为 `backend/.env`，配置数据库连接和JWT密钥。

### 前端

在 `frontend/` 创建 `.env` 文件（如需要自定义API地址），默认使用 `/api/v1` 代理。
