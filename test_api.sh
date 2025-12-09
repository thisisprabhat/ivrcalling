#!/bin/bash

# Simple API Test Script
# Run this to verify the backend is working correctly

echo "================================"
echo "Testing IVR Backend API"
echo "================================"
echo ""

# Test 1: Health Check
echo "1. Testing Health Endpoint..."
curl -s http://localhost:8080/api/health | jq
echo ""

# Test 2: Get Languages
echo "2. Testing Languages Endpoint..."
curl -s http://localhost:8080/api/languages | jq
echo ""

# Test 3: List Campaigns
echo "3. Testing List Campaigns..."
curl -s http://localhost:8080/api/campaigns | jq
echo ""

# Test 4: Create Test Campaign
echo "4. Creating Test Campaign..."
curl -s -X POST http://localhost:8080/api/campaigns \
  -H "Content-Type: application/json" \
  -d '{
    "name": "API Test Campaign",
    "description": "Testing via API",
    "language": "en",
    "intro_text": "Thank you for calling our test line.",
    "actions": [
      {
        "action_type": "information",
        "action_input": "1",
        "message": "This is a test message"
      },
      {
        "action_type": "forward",
        "action_input": "2",
        "forward_phone": "+1234567890"
      }
    ],
    "is_active": true
  }' | jq

echo ""
echo "================================"
echo "Test Complete!"
echo "Check the backend terminal for detailed logs"
echo "================================"
