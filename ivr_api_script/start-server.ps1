# Q&I IVR API Server Startup Script
Write-Host "Starting Q&I IVR API Server..." -ForegroundColor Cyan
Write-Host ""

# Check if .env file exists
if (-not (Test-Path ".env")) {
    Write-Host "WARNING: .env file not found!" -ForegroundColor Yellow
    Write-Host "Please copy .env.example to .env and configure your Twilio credentials." -ForegroundColor Yellow
    Write-Host ""
    Read-Host "Press Enter to exit"
    exit 1
}

# Start the server
Write-Host "Server starting on port 8080..." -ForegroundColor Green
Write-Host "Press Ctrl+C to stop the server" -ForegroundColor Yellow
Write-Host ""
Write-Host "Twilio webhook URL: http://localhost:8080/api/v1/twiml/welcome" -ForegroundColor Cyan
Write-Host ""

.\bin\ivr-api.exe
