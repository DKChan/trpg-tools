#!/bin/bash

echo "========================================"
echo "TRPG-Tools 一键构建脚本"
echo "========================================"
echo ""

# 检查Node.js
if ! command -v node &> /dev/null; then
    echo "[错误] 未找到 Node.js，请先安装 Node.js"
    exit 1
fi

# 检查Go
if ! command -v go &> /dev/null; then
    echo "[错误] 未找到 Go，请先安装 Go"
    exit 1
fi

echo "[1/2] 构建前端..."
cd frontend
npm install
if [ $? -ne 0 ]; then
    echo "[错误] 前端依赖安装失败"
    exit 1
fi

npm run build
if [ $? -ne 0 ]; then
    echo "[错误] 前端构建失败"
    exit 1
fi
cd ..

echo "[2/2] 构建后端..."
cd backend
go mod tidy
go build -o ../trpg-tools main.go
if [ $? -ne 0 ]; then
    echo "[错误] 后端构建失败"
    exit 1
fi
cd ..

echo ""
echo "========================================"
echo "[构建成功！]"
echo "========================================"
echo ""
echo "可执行文件: trpg-tools (项目根目录)"
echo "运行命令: ./trpg-tools"
echo ""
echo "服务启动后访问: http://localhost:8080"
echo ""
