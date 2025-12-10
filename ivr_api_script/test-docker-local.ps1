# test-docker-local.ps1 - Test Docker image locally before deploying

Write-Host "`n========================================" -ForegroundColor Cyan
Write-Host "  Testing Docker Image Locally" -ForegroundColor Cyan
Write-Host "========================================`n" -ForegroundColor Cyan

# Build the image
Write-Host "📦 Building Docker image..." -ForegroundColor Yellow
docker build -t ivr-api:local .

if ($LASTEXITCODE -ne 0) {
    Write-Host "❌ Build failed!" -ForegroundColor Red
    exit 1
}

Write-Host "✅ Build successful!" -ForegroundColor Green

# Stop and remove existing container if running
Write-Host "`n🧹 Cleaning up old containers..." -ForegroundColor Yellow
docker stop ivr-api-test 2>$null
docker rm ivr-api-test 2>$null

# Run the container
Write-Host "`n🚀 Starting container..." -ForegroundColor Yellow

# Check if .env file exists
if (Test-Path .env) {
    Write-Host "✅ Using .env file for configuration" -ForegroundColor Green
    docker run -d `
        --name ivr-api-test `
        -p 8080:8080 `
        --env-file .env `
        ivr-api:local
} else {
    Write-Host "⚠️  .env file not found. Using example values..." -ForegroundColor Yellow
    docker run -d `
        --name ivr-api-test `
        -p 8080:8080 `
        -e PORT=8080 `
        -e GIN_MODE=release `
        -e SERVER_BASE_URL=http://localhost:8080 `
        -e TWILIO_ACCOUNT_SID=ACxxxxxxxx `
        -e TWILIO_AUTH_TOKEN=xxxxxxxx `
        -e TWILIO_PHONE_NUMBER=+1234567890 `
        -e QI_TEAM_PHONE=+917905252436 `
        ivr-api:local
}

# Wait for container to start
Write-Host "`n⏳ Waiting for container to start..." -ForegroundColor Yellow
Start-Sleep -Seconds 3

# Check if container is running
$status = docker ps --filter "name=ivr-api-test" --format "{{.Status}}"

if ($status) {
    Write-Host "✅ Container is running!" -ForegroundColor Green
    
    Write-Host "`n📊 Container logs:" -ForegroundColor Cyan
    Write-Host "----------------------------------------" -ForegroundColor Gray
    docker logs ivr-api-test
    Write-Host "----------------------------------------" -ForegroundColor Gray
    
    Write-Host "`n🧪 Testing endpoints..." -ForegroundColor Yellow
    Start-Sleep -Seconds 2
    
    # Test health endpoint
    try {
        $health = Invoke-WebRequest http://localhost:8080/health -UseBasicParsing
        Write-Host "✅ Health check: $($health.StatusCode) - $($health.Content)" -ForegroundColor Green
    } catch {
        Write-Host "❌ Health check failed: $_" -ForegroundColor Red
    }
    
    # Test IVR config endpoint
    try {
        $config = Invoke-WebRequest http://localhost:8080/api/v1/config/ivr -UseBasicParsing
        Write-Host "✅ IVR Config: $($config.StatusCode)" -ForegroundColor Green
    } catch {
        Write-Host "❌ IVR Config failed: $_" -ForegroundColor Red
    }
    
    # Test TwiML endpoint
    try {
        $twiml = Invoke-WebRequest http://localhost:8080/api/v1/twiml/welcome -UseBasicParsing
        Write-Host "✅ TwiML: $($twiml.StatusCode)" -ForegroundColor Green
    } catch {
        Write-Host "❌ TwiML failed: $_" -ForegroundColor Red
    }
    
    Write-Host "`n========================================" -ForegroundColor Cyan
    Write-Host "  📍 Access Your API" -ForegroundColor Cyan
    Write-Host "========================================" -ForegroundColor Cyan
    Write-Host "  Health:     http://localhost:8080/health" -ForegroundColor White
    Write-Host "  IVR Config: http://localhost:8080/api/v1/config/ivr" -ForegroundColor White
    Write-Host "  TwiML:      http://localhost:8080/api/v1/twiml/welcome" -ForegroundColor White
    
    Write-Host "`n💡 Useful Commands:" -ForegroundColor Yellow
    Write-Host "  View logs:      docker logs -f ivr-api-test" -ForegroundColor White
    Write-Host "  Stop:           docker stop ivr-api-test" -ForegroundColor White
    Write-Host "  Remove:         docker rm ivr-api-test" -ForegroundColor White
    Write-Host "  Restart:        docker restart ivr-api-test" -ForegroundColor White
    Write-Host "  Shell access:   docker exec -it ivr-api-test /bin/sh" -ForegroundColor White
    
    Write-Host "`n🧪 Test Call Initiation:" -ForegroundColor Yellow
    Write-Host '  $body = @{ phone_number = "+919876543210" } | ConvertTo-Json' -ForegroundColor White
    Write-Host '  Invoke-RestMethod -Uri http://localhost:8080/api/v1/calls/initiate -Method Post -Body $body -ContentType "application/json"' -ForegroundColor White
    
    Write-Host "`n✅ Docker container is ready for testing!`n" -ForegroundColor Green
    
} else {
    Write-Host "❌ Container failed to start!" -ForegroundColor Red
    Write-Host "`n📊 Container logs:" -ForegroundColor Yellow
    docker logs ivr-api-test
    exit 1
}
