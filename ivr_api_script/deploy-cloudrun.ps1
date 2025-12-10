# deploy-cloudrun.ps1 - PowerShell deployment script for Google Cloud Run

param(
    [string]$ProjectId = "",
    [string]$ServiceName = "ivr-api",
    [string]$Region = "us-central1"
)

# Colors for output
function Write-ColorOutput {
    param([string]$Message, [string]$Color = "White")
    Write-Host $Message -ForegroundColor $Color
}

Write-ColorOutput "`n========================================" "Cyan"
Write-ColorOutput "  Q&I IVR API - Cloud Run Deployment" "Cyan"
Write-ColorOutput "========================================`n" "Cyan"

# Get project ID if not provided
if ([string]::IsNullOrEmpty($ProjectId)) {
    $ProjectId = Read-Host "Enter your GCP Project ID"
}

$ImageName = "gcr.io/$ProjectId/$ServiceName"

# Verify gcloud is installed
try {
    $null = gcloud --version
    Write-ColorOutput "✅ Google Cloud SDK found" "Green"
}
catch {
    Write-ColorOutput "❌ Google Cloud SDK not found. Please install from: https://cloud.google.com/sdk/docs/install" "Red"
    exit 1
}

# Set project
Write-ColorOutput "`n📋 Setting project: $ProjectId" "Yellow"
gcloud config set project $ProjectId

# Enable required APIs
Write-ColorOutput "`n🔧 Enabling required APIs..." "Yellow"
gcloud services enable run.googleapis.com
gcloud services enable containerregistry.googleapis.com

# Build Docker image
Write-ColorOutput "`n📦 Building Docker image..." "Yellow"
docker build -t $ImageName .

if ($LASTEXITCODE -ne 0) {
    Write-ColorOutput "❌ Docker build failed!" "Red"
    exit 1
}

Write-ColorOutput "✅ Docker build successful!" "Green"

# Configure Docker for GCR
Write-ColorOutput "`n🔐 Configuring Docker authentication..." "Yellow"
gcloud auth configure-docker

# Push to Container Registry
Write-ColorOutput "`n📤 Pushing image to Google Container Registry..." "Yellow"
docker push $ImageName

if ($LASTEXITCODE -ne 0) {
    Write-ColorOutput "❌ Docker push failed!" "Red"
    exit 1
}

Write-ColorOutput "✅ Image pushed successfully!" "Green"

# Get Twilio configuration
Write-ColorOutput "`n⚙️  Enter your Twilio configuration:" "Cyan"
$TwilioSid = Read-Host "Twilio Account SID"
$TwilioToken = Read-Host "Twilio Auth Token" -AsSecureString
$TwilioTokenPlain = [Runtime.InteropServices.Marshal]::PtrToStringAuto([Runtime.InteropServices.Marshal]::SecureStringToBSTR($TwilioToken))
$TwilioPhone = Read-Host "Twilio Phone Number (e.g., +1234567890)"
$QiPhone = Read-Host "Q&I Team Phone (press Enter for default: +917905252436)"

if ([string]::IsNullOrEmpty($QiPhone)) {
    $QiPhone = "+917905252436"
}

# Deploy to Cloud Run
Write-ColorOutput "`n🚀 Deploying to Cloud Run..." "Yellow"

gcloud run deploy $ServiceName `
    --image=$ImageName `
    --platform=managed `
    --region=$Region `
    --allow-unauthenticated `
    --port=8080 `
    --memory=512Mi `
    --cpu=1 `
    --min-instances=0 `
    --max-instances=10 `
    --timeout=300 `
    --set-env-vars="GIN_MODE=release,TWILIO_ACCOUNT_SID=$TwilioSid,TWILIO_AUTH_TOKEN=$TwilioTokenPlain,TWILIO_PHONE_NUMBER=$TwilioPhone,QI_TEAM_PHONE=$QiPhone"

if ($LASTEXITCODE -ne 0) {
    Write-ColorOutput "❌ Cloud Run deployment failed!" "Red"
    exit 1
}

# Get the service URL
Write-ColorOutput "`n🔍 Getting service URL..." "Yellow"
$ServiceUrl = gcloud run services describe $ServiceName --region=$Region --format='value(status.url)'

# Update with SERVER_BASE_URL
Write-ColorOutput "`n🔄 Updating SERVER_BASE_URL environment variable..." "Yellow"
gcloud run services update $ServiceName `
    --region=$Region `
    --update-env-vars="SERVER_BASE_URL=$ServiceUrl"

Write-ColorOutput "`n========================================" "Green"
Write-ColorOutput "  ✅ Deployment Successful!" "Green"
Write-ColorOutput "========================================`n" "Green"

Write-ColorOutput "📍 Service URL: $ServiceUrl" "Cyan"
Write-ColorOutput "`n🧪 Test your API:" "Yellow"
Write-ColorOutput "  Health:     $ServiceUrl/health" "White"
Write-ColorOutput "  IVR Config: $ServiceUrl/api/v1/config/ivr" "White"
Write-ColorOutput "  TwiML:      $ServiceUrl/api/v1/twiml/welcome" "White"

Write-ColorOutput "`n💡 Next Steps:" "Yellow"
Write-ColorOutput "  1. Test health endpoint in browser: $ServiceUrl/health" "White"
Write-ColorOutput "  2. Update Twilio webhook URLs if needed" "White"
Write-ColorOutput "  3. Test call initiation via Postman or curl" "White"
Write-ColorOutput "  4. Monitor logs: gcloud run services logs read $ServiceName --region=$Region" "White"

Write-ColorOutput "`n🎉 Your IVR API is now live on Cloud Run!`n" "Green"
