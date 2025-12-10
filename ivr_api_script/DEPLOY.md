# 🚀 Google Cloud Run Deployment - Quick Guide

## Prerequisites

1. **Google Cloud SDK** - [Download here](https://cloud.google.com/sdk/docs/install)
2. **Docker Desktop** - [Download here](https://www.docker.com/products/docker-desktop)
3. **GCP Project** with billing enabled

## 🎯 One-Command Deployment

```powershell
.\deploy-cloudrun.ps1
```

The script will:

- ✅ Build Docker image
- ✅ Push to Google Container Registry
- ✅ Deploy to Cloud Run
- ✅ Configure environment variables
- ✅ Get your HTTPS URL

## 📋 Step-by-Step Manual Deployment

### Step 1: Authenticate

```powershell
gcloud auth login
gcloud config set project YOUR_PROJECT_ID
```

### Step 2: Enable APIs

```powershell
gcloud services enable run.googleapis.com
gcloud services enable containerregistry.googleapis.com
```

### Step 3: Test Locally (Optional)

```powershell
.\test-docker-local.ps1
```

### Step 4: Deploy

```powershell
.\deploy-cloudrun.ps1 -ProjectId YOUR_PROJECT_ID
```

## ✅ What You'll Get

After deployment:

- 🌐 HTTPS URL: `https://ivr-api-xxxxx-uc.a.run.app`
- 📊 Automatic scaling (0 to 10 instances)
- 💰 Pay only when used
- 🔒 Built-in SSL/TLS

## 🧪 Testing After Deployment

```powershell
# Get your service URL
$URL = gcloud run services describe ivr-api --region=us-central1 --format='value(status.url)'

# Test health
Invoke-WebRequest "$URL/health"

# Test TwiML
Invoke-WebRequest "$URL/api/v1/twiml/welcome"

# Make a test call
$body = @{ phone_number = "+919876543210" } | ConvertTo-Json
Invoke-RestMethod -Uri "$URL/api/v1/calls/initiate" -Method Post -Body $body -ContentType "application/json"
```

## 📊 View Logs

```powershell
gcloud run services logs read ivr-api --region=us-central1 --limit=50
```

## 🔄 Update Environment Variables

```powershell
gcloud run services update ivr-api `
    --region=us-central1 `
    --update-env-vars="TWILIO_ACCOUNT_SID=new_value"
```

## 💰 Cost Estimate

- First 2 million requests/month: **FREE**
- Additional requests: ~$0.40/million
- Typical cost: **$0.50-2.00/month** for 1000 calls

## 🛠️ Troubleshooting

### Build Fails

```powershell
# Check Docker is running
docker ps

# Rebuild without cache
docker build --no-cache -t test .
```

### Push Fails

```powershell
gcloud auth configure-docker
```

### Deployment Fails

```powershell
# Check logs
gcloud run services logs read ivr-api --region=us-central1

# Check service status
gcloud run services describe ivr-api --region=us-central1
```

## 🔐 Security Notes

- ✅ Non-root user in container
- ✅ HTTPS enforced by Cloud Run
- ✅ Minimal Alpine base image
- ⚠️ Use Secret Manager for production (see below)

### Using Secret Manager (Recommended for Production)

```powershell
# Create secret
echo "your_token" | gcloud secrets create twilio-auth-token --data-file=-

# Update service to use secret
gcloud run services update ivr-api `
    --update-secrets=TWILIO_AUTH_TOKEN=twilio-auth-token:latest
```

## 🎉 You're Done!

Your IVR API is now running on Google Cloud Run with:

- ✅ Automatic HTTPS
- ✅ Auto-scaling
- ✅ Global availability
- ✅ Production-ready setup

Access your API at: `https://ivr-api-xxxxx-uc.a.run.app`
