# TRPG-Sync 后端测试运行脚本 (PowerShell)

$ErrorActionPreference = "Stop"

Write-Host "========================================="
Write-Host "TRPG-Sync 后端测试套件"
Write-Host "========================================="
Write-Host ""

# 1. 检查 Go
Write-Host "1. 检查 Go..."
try {
    $goVersion = go version
    Write-Host "   ✓ Go 已安装" -ForegroundColor Green
    Write-Host "   $goVersion"
} catch {
    Write-Host "   ✗ Go 未安装，请先安装 Go 1.24+" -ForegroundColor Red
    exit 1
}

# 2. 检查 go.mod
Write-Host ""
Write-Host "2. 检查 go.mod..."
if (Test-Path "go.mod") {
    Write-Host "   ✓ go.mod 存在" -ForegroundColor Green
    Write-Host "   主要依赖:"
    Select-String -Path "go.mod" -Pattern "^require" | Select-Object -First 5 | ForEach-Object {
        Write-Host "     $($_.Line.Trim())"
    }
} else {
    Write-Host "   ✗ go.mod 不存在" -ForegroundColor Red
    exit 1
}

# 3. 检查 testify 依赖
Write-Host ""
Write-Host "3. 检查 testify 依赖..."
if (Select-String -Path "go.mod" -Pattern "github.com/stretchr/testify" -Quiet) {
    Write-Host "   ✓ testify 依赖已添加" -ForegroundColor Green
} else {
    Write-Host "   ⚠ testify 依赖可能未正确安装" -ForegroundColor Yellow
    Write-Host "   尝试安装 testify..."
    go get github.com/stretchr/testify@latest
}

# 4. 运行测试
Write-Host ""
Write-Host "4. 运行测试套件..."

$testType = $args[0]

if ($testType) {
    switch ($testType.ToLower()) {
        "auth" {
            Write-Host "   运行认证测试..."
            go test -v ./api/v1/handlers/... -run TestAuthHandler
        }
        "room" {
            Write-Host "   运行房间测试..."
            go test -v ./api/v1/handlers/... -run TestRoomHandler
        }
        "character" {
            Write-Host "   运行人物卡测试..."
            go test -v ./api/v1/handlers/... -run TestCharacterHandler
        }
        "user" {
            Write-Host "   运行用户测试..."
            go test -v ./api/v1/handlers/... -run TestUserHandler
        }
        "middleware" {
            Write-Host "   运行中间件测试..."
            go test -v ./api/middleware/... -run TestAuthMiddleware
        }
        "all" {
            Write-Host "   运行所有测试..."
            go test -v ./...
        }
        default {
            Write-Host "   ✗ 未知的测试类型: $testType" -ForegroundColor Red
            Write-Host "   可用选项: auth, room, character, user, middleware, all"
            exit 1
        }
    }
} else {
    Write-Host "   运行所有测试（使用 ./...）..."
    go test -v ./...
}

if ($LASTEXITCODE -eq 0) {
    Write-Host "   ✓ 所有测试通过！" -ForegroundColor Green
} else {
    Write-Host "   ✗ 测试失败，请检查错误信息" -ForegroundColor Red
    exit 1
}

# 5. 生成覆盖率报告（可选）
Write-Host ""
$genCover = $args[1] -eq "--cover" -or $testType -eq "all"

if ($genCover) {
    Write-Host "5. 生成测试覆盖率报告..."
    go test -coverprofile=coverage.out -covermode=atomic ./...

    if (Test-Path "coverage.out") {
        Write-Host "   ✓ 覆盖率报告已生成: coverage.out" -ForegroundColor Green
        Write-Host ""
        Write-Host "   覆盖率统计:"
        $coverageOutput = go tool cover -func=coverage.out | Select-String "total"
        Write-Host "     $coverageOutput"

        # 提取总覆盖率
        $totalCoverage = ($coverageOutput -split '\s+')[2] -replace '%',''

        Write-Host ""
        Write-Host "   总覆盖率: $totalCoverage%"

        # 检查是否达到 70%
        if ([double]$totalCoverage -lt 70) {
            Write-Host "   ⚠ 覆盖率 $totalCoverage% 低于目标 70%" -ForegroundColor Yellow
            Write-Host "   建议添加更多测试用例"
        } else {
            Write-Host "   ✓ 覆盖率 $totalCoverage% 达到目标 70%" -ForegroundColor Green
        }

        Write-Host ""
        Write-Host "   生成 HTML 覆盖率报告: coverage.html"
        go tool cover -html=coverage.out -o coverage.html
        Write-Host "   ✓ HTML 报告已生成: coverage.html" -ForegroundColor Green
    } else {
        Write-Host "   ✗ 覆盖率报告生成失败" -ForegroundColor Red
    }
} else {
    Write-Host "5. 跳过覆盖率报告生成（使用 --cover 参数启用）"
}

Write-Host ""
Write-Host "========================================="
Write-Host "测试完成！"
Write-Host "========================================="
Write-Host ""
Write-Host "快速命令参考:"
Write-Host "  .\test.ps1 auth          - 运行认证测试"
Write-Host "  .\test.ps1 room          - 运行房间测试"
Write-Host "  .\test.ps1 character     - 运行人物卡测试"
Write-Host "  .\test.ps1 user          - 运行用户测试"
Write-Host "  .\test.ps1 middleware     - 运行中间件测试"
Write-Host "  .\test.ps1 all           - 运行所有测试（包括覆盖率）"
Write-Host "  .\test.ps1 all --cover   - 运行测试并生成覆盖率报告"
Write-Host ""
