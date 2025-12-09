#!/bin/bash

# Q&I IVR API - Twilio Test Script
# Usage: ./examples/test_twilio.sh +919876543210

if [ $# -eq 0 ]; then
    echo "Usage: ./examples/test_twilio.sh <phone_number>"
    echo "Example: ./examples/test_twilio.sh +919876543210"
    exit 1
fi

PHONE_NUMBER=$1
BASE_URL="http://localhost:8080"

# Colors for output
GREEN='\033[0;32m'
BLUE='\033[0;34m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo "============================================"
echo "Q&I IVR API - Twilio Integration Test"
echo "============================================"
echo ""

# Test 1: Health Check
echo -e "${BLUE}Test 1: Health Check${NC}"
echo "--------------------"
response=$(curl -s "$BASE_URL/health")
echo "Response: $response"
echo ""

# Test 2: Get IVR Configuration
echo -e "${BLUE}Test 2: Get IVR Configuration${NC}"
echo "------------------------------"
response=$(curl -s "$BASE_URL/api/v1/config/ivr")
echo "$response" | python3 -m json.tool 2>/dev/null || echo "$response"
echo ""

# Test 3: Initiate Call via Twilio
echo -e "${BLUE}Test 3: Initiate Call to $PHONE_NUMBER${NC}"
echo "--------------------------------------------"
response=$(curl -s -X POST "$BASE_URL/api/v1/calls/initiate" \
  -H "Content-Type: application/json" \
  -d "{\"phone_number\": \"$PHONE_NUMBER\", \"callback_url\": \"$BASE_URL/api/v1/callbacks/ivr\"}")

echo "$response" | python3 -m json.tool 2>/dev/null || echo "$response"
echo ""

# Extract Call SID if successful
call_id=$(echo "$response" | grep -o '"call_id":"[^"]*' | cut -d'"' -f4)

if [ -n "$call_id" ]; then
    echo -e "${GREEN}✓ Call initiated successfully!${NC}"
    echo "Call SID: $call_id"
    echo ""
    echo -e "${YELLOW}Next steps:${NC}"
    echo "1. The recipient will receive a call shortly"
    echo "2. They will hear the Q&I welcome message"
    echo "3. They can press:"
    echo "   - 1: Connect to Q&I team"
    echo "   - 2: Hear more about Q&I"
    echo "   - 3: Repeat the message"
else
    echo -e "${YELLOW}⚠ Check the response above for errors${NC}"
fi

echo ""
echo "============================================"
echo "Test Completed"
echo "============================================"
