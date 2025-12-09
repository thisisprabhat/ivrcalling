# PowerShell API Test Script
# Run this to verify the backend is working correctly

Write-Host "================================" -ForegroundColor Cyan
Write-Host "Testing IVR Backend API" -ForegroundColor Cyan
Write-Host "================================" -ForegroundColor Cyan
Write-Host ""

# Test 1: Health Check
Write-Host "1. Testing Health Endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/health" -Method Get
    $response | ConvertTo-Json
    Write-Host "✓ Health check passed" -ForegroundColor Green
} catch {
    Write-Host "✗ Health check failed: $_" -ForegroundColor Red
}
Write-Host ""

# Test 2: Get Languages
Write-Host "2. Testing Languages Endpoint..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/languages" -Method Get
    $response | ConvertTo-Json
    Write-Host "✓ Languages endpoint working" -ForegroundColor Green
} catch {
    Write-Host "✗ Languages endpoint failed: $_" -ForegroundColor Red
}
Write-Host ""

# Test 3: List Campaigns
Write-Host "3. Testing List Campaigns..." -ForegroundColor Yellow
try {
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/campaigns" -Method Get
    Write-Host "Found $($response.Count) campaign(s)" -ForegroundColor Cyan
    $response | ConvertTo-Json
    Write-Host "✓ List campaigns working" -ForegroundColor Green
} catch {
    Write-Host "✗ List campaigns failed: $_" -ForegroundColor Red
}
Write-Host ""

# Test 4: Create Test Campaign
Write-Host "4. Creating Test Campaign..." -ForegroundColor Yellow
$campaignData = @{
    name = "API Test Campaign"
    description = "Testing via PowerShell API"
    language = "en"
    intro_text = "Thank you for calling our test line."
    actions = @(
        @{
            action_type = "information"
            action_input = "1"
            message = "This is a test message"
        },
        @{
            action_type = "forward"
            action_input = "2"
            forward_phone = "+1234567890"
        }
    )
    is_active = $true
}

try {
    $json = $campaignData | ConvertTo-Json -Depth 10
    $response = Invoke-RestMethod -Uri "http://localhost:8080/api/campaigns" -Method Post -Body $json -ContentType "application/json"
    Write-Host "Campaign Created:" -ForegroundColor Cyan
    $response | ConvertTo-Json
    Write-Host "✓ Campaign creation successful" -ForegroundColor Green
    Write-Host "Campaign ID: $($response.id)" -ForegroundColor Cyan
} catch {
    Write-Host "✗ Campaign creation failed: $_" -ForegroundColor Red
    Write-Host "Error details: $($_.Exception.Message)" -ForegroundColor Red
}

Write-Host ""
Write-Host "================================" -ForegroundColor Cyan
Write-Host "Test Complete!" -ForegroundColor Cyan
Write-Host "Check the backend terminal for detailed logs" -ForegroundColor Cyan
Write-Host "================================" -ForegroundColor Cyan
