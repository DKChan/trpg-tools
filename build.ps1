Write-Host "========================================"
Write-Host "TRPG-Tools Build Script"
Write-Host "========================================"
Write-Host ""

# Check Node.js
$nodeVersion = node --version 2>$null
if (-not $nodeVersion) {
    Write-Host "[Error] Node.js not found" -ForegroundColor Red
    exit 1
}
Write-Host "OK Node.js: $nodeVersion" -ForegroundColor Green

# Check Go
$goVersion = go version 2>$null
if (-not $goVersion) {
    Write-Host "[Error] Go not found" -ForegroundColor Red
    exit 1
}
Write-Host "OK Go installed" -ForegroundColor Green

Write-Host ""
Write-Host "[Step 1/2] Building frontend..."
Set-Location frontend
npm install
if ($LASTEXITCODE -ne 0) {
    Write-Host "[Error] Frontend install failed" -ForegroundColor Red
    Set-Location ..
    exit 1
}

npm run build
if ($LASTEXITCODE -ne 0) {
    Write-Host "[Error] Frontend build failed" -ForegroundColor Red
    Set-Location ..
    exit 1
}
Set-Location ..

Write-Host ""
Write-Host "[Step 2/2] Building backend..."
Set-Location backend
go mod tidy
go build -o ../trpg-tools.exe main.go
if ($LASTEXITCODE -ne 0) {
    Write-Host "[Error] Backend build failed" -ForegroundColor Red
    Set-Location ..
    exit 1
}
Set-Location ..

Write-Host ""
Write-Host "========================================"
Write-Host "Build Success!"
Write-Host "========================================"
Write-Host ""
Write-Host "Executable: trpg-tools.exe (in project root)"
Write-Host "Run: .\trpg-tools.exe"
Write-Host "Access: http://localhost:8080"
Write-Host ""
