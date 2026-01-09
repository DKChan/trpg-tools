# TRPG人物卡共享项目初始化计划

## 任务一：生成项目规则

### 后端规则（Go）
1. **目录结构**：采用清晰的分层架构
   - `api/`：HTTP路由和控制器
   - `services/`：业务逻辑层
   - `models/`：数据模型和数据库交互
   - `middleware/`：中间件（认证、日志等）
   - `config/`：配置管理
   - `utils/`：工具函数
   - `docs/`：API文档

2. **框架选择**：
   - Web框架：Gin（高性能、轻量级）
   - ORM：GORM（强大的数据库操作能力）
   - 认证：JWT
   - 数据库：PostgreSQL（支持复杂数据结构）

3. **代码规范**：
   - 遵循Go官方代码规范（gofmt/golint）
   - 函数注释清晰，包含参数和返回值说明
   - 错误处理统一，使用自定义错误类型
   - 日志采用结构化日志（zap）

4. **API设计**：
   - RESTful风格API
   - 版本控制（如/v1/）
   - 统一的响应格式
   - 完善的错误码体系

### 前端规则
1. **技术栈**：
   - 框架：React 18
   - 构建工具：Vite
   - 状态管理：Zustand（轻量级、易用）
   - 路由：React Router v6
   - UI库：Ant Design（组件丰富、美观）
   - API请求：Axios
   - 样式：TailwindCSS（快速开发、响应式）

2. **目录结构**：
   - `src/`：
     - `components/`：可复用组件
     - `pages/`：页面组件
     - `store/`：状态管理
     - `services/`：API调用
     - `hooks/`：自定义Hooks
     - `utils/`：工具函数
     - `types/`：TypeScript类型定义
     - `styles/`：全局样式

3. **代码规范**：
   - TypeScript严格模式
   - ESLint + Prettier代码格式化
   - 组件命名采用PascalCase
   - 函数组件优先
   - 清晰的props和state类型定义

4. **开发流程**：
   - 组件化开发，高内聚低耦合
   - 响应式设计，适配不同设备
   - 良好的用户体验，加载状态和错误处理
   - 单元测试和集成测试

## 任务二：创建需求文档

### 需求文档内容
1. **项目概述**：TRPG游戏人物卡共享平台，支持DM和玩家在线管理人物卡

2. **核心功能**：
   - 房间管理：创建、加入、退出战役房间
   - 角色管理：
     - 玩家：创建DND5e人物卡，更新属性
     - DM：查看所有玩家人物卡信息
   - 实时同步：游戏过程中人物属性实时更新
   - 权限管理：区分DM和玩家权限

3. **数据模型**：
   - 用户模型
   - 房间模型
   - 人物卡模型（DND5e规则）
   - 游戏会话模型

4. **技术实现**：
   - 后端：Go + Gin + GORM + PostgreSQL
   - 前端：React + TypeScript + Ant Design + TailwindCSS
   - 实时通信：WebSocket

5. **后续扩展**：
   - 支持更多TRPG规则（如COC、PF等）
   - 人物卡模板系统
   - 战斗管理功能
   - 骰子系统集成

## 执行计划

1. 创建项目规则文件：`.trae/rules/project_rules.md`
2. 创建需求文档：`requirements.md`
3. 初始化Go后端项目结构
4. 初始化前端React项目结构
5. 实现核心功能模块

现在我将开始创建项目规则和需求文档。