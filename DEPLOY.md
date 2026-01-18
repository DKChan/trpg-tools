# 一键部署指南

## 快速开始

### 1. 准备环境

确保已安装以下工具：
- **Node.js** 18+ (用于构建前端)
- **Go** 1.24+ (用于构建和运行后端)

### 2. 一键构建

**Windows:**
```powershell
.\build.ps1
```

**Linux/Mac:**
```bash
./build.sh
```

### 3. 运行应用

```bash
cd backend

# Windows
trpg-tools.exe

# Linux/Mac
./trpg-tools
```

### 4. 访问应用

打开浏览器访问: http://localhost:8080

---

## 手动构建步骤

如果需要更细粒度的控制：

### 构建前端

```bash
cd frontend
npm install
npm run build
cd ..
```

前端构建产物会自动输出到 `backend/dist/` 目录。

### 构建后端

```bash
cd backend
go mod tidy
go build -o trpg-tools main.go
```

### 运行

```bash
# Windows
.\trpg-tools.exe

# Linux/Mac
./trpg-tools
```

---

## 部署说明

### 架构说明

项目采用 **Go Embed** 技术，将前端静态资源嵌入到二进制文件中：

- ✅ 单一可执行文件
- ✅ 前后端集成
- ✅ 无需额外配置
- ✅ 端口统一 (8080)

### 运行时文件

程序首次运行时会自动创建：

```
backend/
├── trpg-tools.exe          # 可执行文件
├── sqlite.db               # SQLite数据库（自动创建）
└── data/                   # 数据目录（自动创建）
    └── rooms/
        └── {room_id}/
            └── characters/
                └── {character_id}.json
```

### 配置端口

默认使用 8080 端口，可通过环境变量修改：

```bash
# Windows
set SERVER_PORT=3000
trpg-tools.exe

# Linux/Mac
export SERVER_PORT=3000
./trpg-tools
```

### 跨平台部署

构建不同平台的可执行文件：

```bash
# Windows
cd backend
go build -o trpg-tools-windows.exe main.go

# Linux
GOOS=linux GOARCH=amd64 go build -o trpg-tools-linux main.go

# macOS
GOOS=darwin GOARCH=amd64 go build -o trpg-tools-mac main.go
```

---

## 常见问题

### Q: 构建失败 "pattern dist: no matching files"

**A:** 需要先构建前端，执行：
```bash
cd frontend
npm run build
cd ..
```

### Q: 如何备份数据？

**A:** 备份以下文件/目录：
- `sqlite.db` - 数据库文件
- `data/` - 人物卡数据目录

### Q: 可以修改端口吗？

**A:** 可以，通过环境变量 `SERVER_PORT` 修改

### Q: 如何更新应用？

**A:** 
1. 拉取最新代码
2. 运行 `build.ps1` 或 `build.sh`
3. 停止旧进程，运行新版本

---

## 开发模式

如需开发调试：

### 启动后端（端口8080）
```bash
cd backend
go run main.go
```

### 启动前端（端口5173）
```bash
cd frontend
npm run dev
```

前端会自动代理API请求到后端。

---

## 技术细节

### Go Embed

使用 `embed.FS` 将前端静态资源嵌入到Go二进制文件：

```go
//go:embed all:dist
var frontendFS embed.FS
```

优势：
- 无需外部依赖
- 单文件部署
- 启动速度快

### 路由处理

- `/api/*` - 后端API路由
- `/assets/*` - 前端静态资源
- `/*` - 前端应用路由（React Router）

所有前端路由请求都会返回 `index.html`，由前端路由处理。
