#!/bin/bash
# deploy-cloudrun.sh - Bash deployment script for Google Cloud Run

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

# Configuration
SERVICE_NAME="ivr-api"
REGION="us-central1"

echo -e "${CYAN}"
echo "========================================"
echo "  Q&I IVR API - Cloud Run Deployment"
echo "========================================"
echo -e "${NC}"

# Get project ID
read -p "Enter your GCP Project ID: " PROJECT_ID
IMAGE_NAME="gcr.io/$PROJECT_ID/$SERVICE_NAME"

# Verify gcloud is installed
if ! command -v gcloud &> /dev/null; then
    echo -e "${RED}❌ Google Cloud SDK not found. Please install from: https://cloud.google.com/sdk/docs/install${NC}"
    exit 1
fi

echo -e "${GREEN}✅ Google Cloud SDK found${NC}"

# Set project
echo -e "\n${YELLOW}📋 Setting project: $PROJECT_ID${NC}"
gcloud config set project $PROJECT_ID

# Enable required APIs
echo -e "\n${YELLOW}🔧 Enabling required APIs...${NC}"
gcloud services enable run.googleapis.com
gcloud services enable containerregistry.googleapis.com

# Build Docker image
echo -e "\n${YELLOW}📦 Building Docker image...${NC}"
docker build -t $IMAGE_NAME .

echo -e "${GREEN}✅ Docker build successful!${NC}"

# Configure Docker for GCR
echo -e "\n${YELLOW}🔐 Configuring Docker authentication...${NC}"
gcloud auth configure-docker

# Push to Container Registry
echo -e "\n${YELLOW}📤 Pushing image to Google Container Registry...${NC}"
docker push $IMAGE_NAME

echo -e "${GREEN}✅ Image pushed successfully!${NC}"

# Get Twilio configuration
echo -e "\n${CYAN}⚙️  Enter your Twilio configuration:${NC}"
read -p "Twilio Account SID: " TWILIO_SID
read -sp "Twilio Auth Token: " TWILIO_TOKEN
echo
read -p "Twilio Phone Number (e.g., +1234567890): " TWILIO_PHONE
read -p "Q&I Team Phone (press Enter for default: +917905252436): " QI_PHONE
QI_PHONE=${QI_PHONE:-"+917905252436"}

# Deploy to Cloud Run
echo -e "\n${YELLOW}🚀 Deploying to Cloud Run...${NC}"

gcloud run deploy $SERVICE_NAME \
    --image=$IMAGE_NAME \
    --platform=managed \
    --region=$REGION \
    --allow-unauthenticated \
    --port=8080 \
    --memory=512Mi \
    --cpu=1 \
    --min-instances=0 \
    --max-instances=10 \
    --timeout=300 \
    --set-env-vars="GIN_MODE=release,TWILIO_ACCOUNT_SID=$TWILIO_SID,TWILIO_AUTH_TOKEN=$TWILIO_TOKEN,TWILIO_PHONE_NUMBER=$TWILIO_PHONE,QI_TEAM_PHONE=$QI_PHONE"

# Get the service URL
echo -e "\n${YELLOW}🔍 Getting service URL...${NC}"
SERVICE_URL=$(gcloud run services describe $SERVICE_NAME --region=$REGION --format='value(status.url)')

# Update with SERVER_BASE_URL
echo -e "\n${YELLOW}🔄 Updating SERVER_BASE_URL environment variable...${NC}"
gcloud run services update $SERVICE_NAME \
    --region=$REGION \
    --update-env-vars="SERVER_BASE_URL=$SERVICE_URL"

echo -e "\n${GREEN}"
echo "========================================"
echo "  ✅ Deployment Successful!"
echo "========================================"
echo -e "${NC}"

echo -e "${CYAN}📍 Service URL: $SERVICE_URL${NC}"
echo -e "\n${YELLOW}🧪 Test your API:${NC}"
echo -e "  Health:     $SERVICE_URL/health"
echo -e "  IVR Config: $SERVICE_URL/api/v1/config/ivr"
echo -e "  TwiML:      $SERVICE_URL/api/v1/twiml/welcome"

echo -e "\n${YELLOW}💡 Next Steps:${NC}"
echo -e "  1. Test health endpoint: curl $SERVICE_URL/health"
echo -e "  2. Update Twilio webhook URLs if needed"
echo -e "  3. Test call initiation via Postman or curl"
echo -e "  4. Monitor logs: gcloud run services logs read $SERVICE_NAME --region=$REGION"

echo -e "\n${GREEN}🎉 Your IVR API is now live on Cloud Run!${NC}\n"
