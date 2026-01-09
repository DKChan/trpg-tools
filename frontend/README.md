# TRPG人物卡共享平台 - 前端

基于 React + TypeScript + Vite + Ant Design + TailwindCSS 构建的 TRPG 人物卡共享平台前端应用。

## 技术栈

- **框架**: React 18
- **构建工具**: Vite
- **状态管理**: Zustand
- **路由**: React Router v6
- **UI库**: Ant Design
- **API请求**: Axios
- **样式**: TailwindCSS
- **类型检查**: TypeScript 5

## 项目结构

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
│   └── main.tsx    # 应用入口文件
├── public/          # 静态资源
├── index.html       # HTML模板
├── vite.config.ts   # Vite配置
├── tsconfig.json    # TypeScript配置
├── package.json     # 依赖管理
└── tailwind.config.js # TailwindCSS配置
```

## 开发

### 安装依赖

```bash
npm install
```

### 启动开发服务器

```bash
npm run dev
```

开发服务器将在 http://localhost:5173 启动

### 构建生产版本

```bash
npm run build
```

### 预览生产版本

```bash
npm run preview
```

### 代码检查

```bash
npm run lint
```

## 环境变量

创建 `.env` 文件配置环境变量：

```
VITE_API_BASE_URL=http://localhost:8080
```

## 功能特性

- 用户注册和登录
- 房间管理（创建、加入、退出）
- 人物卡管理（DND5e规则）
- 实时同步（WebSocket）
- 响应式设计

## 浏览器支持

- Chrome (最新版)
- Firefox (最新版)
- Safari (最新版)
- Edge (最新版)
