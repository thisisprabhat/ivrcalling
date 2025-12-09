#!/bin/bash

# Q&I IVR API Test Script
# Usage: ./examples/test_api.sh

BASE_URL="http://localhost:8080"

echo "============================================"
echo "Q&I IVR API Test Suite"
echo "============================================"
echo ""

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Test 1: Health Check
echo "Test 1: Health Check"
echo "--------------------"
response=$(curl -s -w "\n%{http_code}" "$BASE_URL/health")
http_code=$(echo "$response" | tail -n 1)
body=$(echo "$response" | sed '$d')

if [ "$http_code" -eq 200 ]; then
    echo -e "${GREEN}✓ PASSED${NC}"
    echo "Response: $body"
else
    echo -e "${RED}✗ FAILED${NC}"
    echo "HTTP Code: $http_code"
fi
echo ""

# Test 2: Get IVR Configuration
echo "Test 2: Get IVR Configuration"
echo "------------------------------"
response=$(curl -s -w "\n%{http_code}" "$BASE_URL/api/v1/config/ivr")
http_code=$(echo "$response" | tail -n 1)
body=$(echo "$response" | sed '$d')

if [ "$http_code" -eq 200 ]; then
    echo -e "${GREEN}✓ PASSED${NC}"
    echo "Response: $body" | head -n 5
    echo "..."
else
    echo -e "${RED}✗ FAILED${NC}"
    echo "HTTP Code: $http_code"
fi
echo ""

# Test 3: Initiate Call (Valid)
echo "Test 3: Initiate Call (Valid Phone Number)"
echo "------------------------------------------"
response=$(curl -s -w "\n%{http_code}" -X POST "$BASE_URL/api/v1/calls/initiate" \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+919876543210", "callback_url": "https://example.com/callback"}')
http_code=$(echo "$response" | tail -n 1)
body=$(echo "$response" | sed '$d')

if [ "$http_code" -eq 200 ]; then
    echo -e "${GREEN}✓ PASSED${NC}"
    echo "Response: $body"
else
    echo -e "${RED}✗ FAILED${NC}"
    echo "HTTP Code: $http_code"
    echo "Response: $body"
fi
echo ""

# Test 4: Initiate Call (Invalid - No Phone)
echo "Test 4: Initiate Call (Invalid - Missing Phone)"
echo "-----------------------------------------------"
response=$(curl -s -w "\n%{http_code}" -X POST "$BASE_URL/api/v1/calls/initiate" \
  -H "Content-Type: application/json" \
  -d '{}')
http_code=$(echo "$response" | tail -n 1)
body=$(echo "$response" | sed '$d')

if [ "$http_code" -eq 400 ]; then
    echo -e "${GREEN}✓ PASSED${NC} (Error handling works)"
    echo "Response: $body"
else
    echo -e "${RED}✗ FAILED${NC}"
    echo "HTTP Code: $http_code"
fi
echo ""

# Test 5: IVR Callback
echo "Test 5: IVR Callback"
echo "-------------------"
response=$(curl -s -w "\n%{http_code}" -X POST "$BASE_URL/api/v1/callbacks/ivr" \
  -H "Content-Type: application/json" \
  -d '{"call_id": "test_123", "event": "digit_pressed", "digit_input": "1", "timestamp": "2025-12-09T10:30:00Z"}')
http_code=$(echo "$response" | tail -n 1)
body=$(echo "$response" | sed '$d')

if [ "$http_code" -eq 200 ]; then
    echo -e "${GREEN}✓ PASSED${NC}"
    echo "Response: $body"
else
    echo -e "${RED}✗ FAILED${NC}"
    echo "HTTP Code: $http_code"
fi
echo ""

echo "============================================"
echo "Test Suite Completed"
echo "============================================"
