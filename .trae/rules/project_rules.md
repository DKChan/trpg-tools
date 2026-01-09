# TRPG人物卡共享项目规则

## 1. 项目概述
本项目是一个TRPG游戏人物卡共享平台，支持DM（地下城主）和玩家在线管理和共享人物卡信息。

## 2. 后端规则（Go）

### 2.1 技术栈
- **Web框架**：Gin（高性能、轻量级）
- **ORM**：GORM（强大的数据库操作能力）
- **认证**：JWT（JSON Web Token）
- **数据库**：PostgreSQL（支持复杂数据结构）
- **日志**：zap（结构化日志）
- **配置管理**：viper

### 2.2 目录结构
```
backend/
├── api/             # 接口层：HTTP路由和控制器
│   ├── v1/          # API版本控制
│   └── middleware/  # 中间件
├── application/     # 应用层：协调领域对象完成业务逻辑
│   ├── commands/    # 命令处理
│   └── queries/     # 查询处理
├── domain/          # 领域层：核心业务逻辑
│   └── domain_name/ # 具体业务领域名称
├── infrastructure/  # 基础设施层：技术支持
│   ├── database/    # 数据库访问
│   ├── cache/       # 缓存
│   ├── config/      # 配置管理
│   └── utils/       # 工具函数
├── docs/            # API文档
├── main.go          # 应用入口
├── go.mod           # 依赖管理
└── .env.example     # 环境变量示例
```

### 2.3 代码规范
- 遵循Go官方代码规范（使用`gofmt`和`golint`）
- 函数注释清晰，包含参数和返回值说明
- 错误处理统一，使用自定义错误类型
- 日志采用结构化日志
- 变量命名使用驼峰命名法
- 常量使用全大写加下划线

### 2.4 API设计
- RESTful风格API
- 版本控制（如`/v1/`）
- 统一的响应格式
```json
{
  "code": 200,
  "message": "success",
  "data": {}
}
```
- 完善的错误码体系

### 2.5 数据库设计
- 采用PostgreSQL数据库
- 表名使用复数形式，下划线分隔
- 字段名使用下划线分隔
- 每个表必须包含`id`、`created_at`、`updated_at`字段
- 外键约束明确

## 3. 前端规则

### 3.1 技术栈
- **框架**：React 18
- **构建工具**：Vite
- **状态管理**：Zustand（轻量级、易用）
- **路由**：React Router v6
- **UI库**：Ant Design（组件丰富、美观）
- **API请求**：Axios
- **样式**：TailwindCSS（快速开发、响应式）
- **类型检查**：TypeScript 5

### 3.2 目录结构
```
frontend/
├── src/
│   ├── components/  # 可复用组件
│   ├── pages/       # 页面组件
│   ├── store/       # 状态管理
│   ├── services/    # API调用
│   ├── hooks/       # 自定义Hooks
│   ├── utils/       # 工具函数
│   ├── types/       # TypeScript类型定义
│   ├── styles/      # 全局样式
│   ├── App.tsx      # 应用入口组件
│   └── main.tsx     # 应用入口文件
├── public/          # 静态资源
├── index.html       # HTML模板
├── vite.config.ts   # Vite配置
├── tsconfig.json    # TypeScript配置
├── package.json     # 依赖管理
└── tailwind.config.js # TailwindCSS配置
```

### 3.3 代码规范
- TypeScript严格模式
- 使用ESLint + Prettier进行代码格式化
- 组件命名采用PascalCase
- 函数组件优先
- 清晰的props和state类型定义
- 变量命名使用驼峰命名法
- 常量使用全大写加下划线

### 3.4 组件设计
- 组件化开发，高内聚低耦合
- 容器组件与展示组件分离
- 组件props使用接口定义
- 合理使用React Hooks
- 避免不必要的重渲染

### 3.5 状态管理
- 使用Zustand进行状态管理
- 状态按功能模块划分
- 避免全局状态过多
- 异步操作处理清晰

### 3.6 API调用
- 统一的API请求封装
- 错误处理统一
- 请求状态管理
- 超时处理

## 4. 开发流程

### 4.1 分支管理
- `main`：主分支，用于生产环境
- `develop`：开发分支，用于集成测试
- `feature/xxx`：功能分支，用于开发新功能
- `bugfix/xxx`：bug修复分支

### 4.2 代码提交规范
- 提交信息格式：`type(scope): description`
  - type：feat（新功能）、fix（修复bug）、docs（文档）、style（代码格式）、refactor（重构）、test（测试）、chore（构建过程或辅助工具的变动）
  - scope：影响的模块
  - description：简短的描述

### 4.3 测试规范
- 后端：使用`testing`包进行单元测试和集成测试
- 前端：使用Jest + React Testing Library进行单元测试和组件测试
- 代码覆盖率要求：核心功能≥80%

## 5. 部署规范

### 5.1 环境配置
- 开发环境：使用`.env.development`
- 测试环境：使用`.env.test`
- 生产环境：使用`.env.production`

### 5.2 容器化
- 使用Docker进行容器化部署
- 提供Dockerfile和docker-compose.yml

### 5.3 CI/CD
- 使用GitHub Actions进行持续集成和持续部署
- 自动运行测试和构建
- 自动部署到测试环境和生产环境

## 6. 安全规范

### 6.1 后端安全
- 使用HTTPS协议
- 密码加密存储（bcrypt）
- JWT token设置合理的过期时间
- 防止SQL注入
- 防止XSS攻击
- 防止CSRF攻击
- 合理的权限控制

### 6.2 前端安全
- 防止XSS攻击
- 防止CSRF攻击
- 敏感信息加密传输
- 输入验证
- 避免在客户端存储敏感信息

## 7. 性能优化

### 7.1 后端优化
- 使用连接池
- 合理的索引设计
- 缓存策略（Redis）
- 异步处理
- 分页查询

### 7.2 前端优化
- 代码分割
- 懒加载
- 图片优化
- 减少HTTP请求
- 合理使用缓存
- 避免不必要的重渲染

## 8. 文档规范

### 8.1 API文档
- 使用Swagger/OpenAPI生成API文档
- 包含API描述、参数、返回值、错误码等
- 定期更新

### 8.2 代码文档
- 核心函数和类必须有注释
- 复杂逻辑必须有注释
- README.md包含项目说明、安装和运行步骤

## 9. 版本管理

- 使用语义化版本控制（Semantic Versioning）
- 格式：MAJOR.MINOR.PATCH
- MAJOR：不兼容的API更改
- MINOR：向下兼容的功能性新增
- PATCH：向下兼容的问题修正

# 变更记录
- v1.0.0：初始版本
