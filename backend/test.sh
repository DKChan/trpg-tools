#!/bin/bash

# TRPG-Sync 后端测试运行脚本

set -e  # 遇到错误立即退出

echo "========================================="
echo "TRPG-Sync 后端测试套件"
echo "========================================="
echo ""

# 颜色定义
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# 函数：打印成功消息
print_success() {
    echo -e "${GREEN}✓ $1${NC}"
}

# 函数：打印警告消息
print_warning() {
    echo -e "${YELLOW}⚠ $1${NC}"
}

# 函数：打印错误消息
print_error() {
    echo -e "${RED}✗ $1${NC}"
}

# 1. 检查依赖
echo "1. 检查依赖..."
if command -v go &> /dev/null; then
    print_success "Go 已安装"
    go version
else
    print_error "Go 未安装，请先安装 Go 1.24+"
    exit 1
fi

if [ -f "go.mod" ]; then
    print_success "go.mod 存在"
    echo "   依赖列表:"
    grep "^require" go.mod | head -10
else
    print_error "go.mod 不存在"
    exit 1
fi

# 2. 检查 testify 依赖
echo ""
echo "2. 检查 testify 依赖..."
if grep -q "github.com/stretchr/testify" go.mod go.sum 2>/dev/null; then
    print_success "testify 依赖已添加"
else
    print_warning "testify 依赖可能未正确安装"
    echo "   尝试安装 testify..."
    go get github.com/stretchr/testify@latest
fi

# 3. 运行测试
echo ""
echo "3. 运行测试套件..."

# 选择测试类型
if [ -n "$1" ]; then
    case "$1" in
        auth)
            echo "   运行认证测试..."
            go test -v ./api/v1/handlers/... -run TestAuthHandler
            ;;
        room)
            echo "   运行房间测试..."
            go test -v ./api/v1/handlers/... -run TestRoomHandler
            ;;
        character)
            echo "   运行人物卡测试..."
            go test -v ./api/v1/handlers/... -run TestCharacterHandler
            ;;
        user)
            echo "   运行用户测试..."
            go test -v ./api/v1/handlers/... -run TestUserHandler
            ;;
        middleware)
            echo "   运行中间件测试..."
            go test -v ./api/middleware/... -run TestAuthMiddleware
            ;;
        all)
            echo "   运行所有测试..."
            go test -v ./...
            ;;
        *)
            print_error "未知的测试类型: $1"
            echo "   可用选项: auth, room, character, user, middleware, all"
            exit 1
            ;;
    esac
else
    echo "   运行所有测试（使用 ./...）..."
    go test -v ./...
fi

if [ $? -eq 0 ]; then
    print_success "所有测试通过！"
else
    print_error "测试失败，请检查错误信息"
    exit 1
fi

# 4. 生成覆盖率报告（可选）
echo ""
if [ "$2" == "--cover" ] || [ "$1" == "all" ]; then
    echo "4. 生成测试覆盖率报告..."
    go test -coverprofile=coverage.out -covermode=atomic ./...

    if [ -f "coverage.out" ]; then
        print_success "覆盖率报告已生成: coverage.out"
        echo ""
        echo "   覆盖率统计:"
        go tool cover -func=coverage.out | grep total

        # 检查是否达到 70% 阈值
        total_coverage=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | sed 's/%//')
        echo ""
        echo "   总覆盖率: ${total_coverage}%"

        if (( $(echo "$total_coverage < 70" | bc -l) )); then
            print_warning "覆盖率 ${total_coverage}% 低于目标 70%"
            echo "   建议添加更多测试用例"
        else
            print_success "覆盖率 ${total_coverage}% 达到目标 70%"
        fi

        echo ""
        echo "   生成 HTML 覆盖率报告: coverage.html"
        go tool cover -html=coverage.out -o coverage.html
        print_success "HTML 报告已生成: coverage.html"
    else
        print_error "覆盖率报告生成失败"
    fi
else
    echo "4. 跳过覆盖率报告生成（使用 --cover 参数启用）"
fi

echo ""
echo "========================================="
echo "测试完成！"
echo "========================================="
echo ""
echo "快速命令参考:"
echo "  ./test.sh auth          - 运行认证测试"
echo "  ./test.sh room          - 运行房间测试"
echo "  ./test.sh character     - 运行人物卡测试"
echo "  ./test.sh user          - 运行用户测试"
echo "  ./test.sh middleware     - 运行中间件测试"
echo "  ./test.sh all           - 运行所有测试（包括覆盖率）"
echo "  ./test.sh all --cover   - 运行测试并生成覆盖率报告"
echo ""
