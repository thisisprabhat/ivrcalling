#!/bin/bash
# Diagnostic script to check IVR API issues

echo "=== IVR API Diagnostics ==="
echo ""

# Check if Go API is running
echo "1. Checking if Go API is running on port 8080..."
if curl -s http://localhost:8080/health > /dev/null 2>&1; then
    echo "✅ Go API is running"
    curl -s http://localhost:8080/health | jq '.' 2>/dev/null || curl -s http://localhost:8080/health
else
    echo "❌ Go API is NOT running on localhost:8080"
    echo "   Start it with: cd /home/prabhat_qanditoday1/ivrcalling/ivr_api_script && ./bin/ivr-api"
fi

echo ""

# Check Caddy status
echo "2. Checking Caddy status..."
systemctl is-active --quiet caddy && echo "✅ Caddy is running" || echo "❌ Caddy is not running"

echo ""

# Test TwiML endpoint locally
echo "3. Testing TwiML endpoint locally (direct to Go API)..."
response=$(curl -s -w "\n%{http_code}" http://localhost:8080/api/v1/twiml/welcome)
http_code=$(echo "$response" | tail -n1)
body=$(echo "$response" | head -n-1)

if [ "$http_code" = "200" ]; then
    echo "✅ Local endpoint works (HTTP $http_code)"
    echo "Response:"
    echo "$body"
else
    echo "❌ Local endpoint failed (HTTP $http_code)"
    echo "$body"
fi

echo ""

# Test TwiML endpoint through Caddy
echo "4. Testing TwiML endpoint through Caddy (via domain)..."
response=$(curl -s -w "\n%{http_code}" https://ivr.waygosquad.com/api/v1/twiml/welcome)
http_code=$(echo "$response" | tail -n1)
body=$(echo "$response" | head -n-1)

if [ "$http_code" = "200" ]; then
    echo "✅ Domain endpoint works (HTTP $http_code)"
    echo "Response:"
    echo "$body"
else
    echo "❌ Domain endpoint failed (HTTP $http_code)"
    echo "$body"
fi

echo ""

# Check logs
echo "5. Recent Caddy errors (if any)..."
journalctl -u caddy -n 20 --no-pager | grep -i error || echo "No recent errors"

echo ""

# Check if Go binary exists
echo "6. Checking Go binary..."
if [ -f "/home/prabhat_qanditoday1/ivrcalling/ivr_api_script/bin/ivr-api" ]; then
    echo "✅ Go binary exists"
else
    echo "❌ Go binary not found. Build it with:"
    echo "   cd /home/prabhat_qanditoday1/ivrcalling/ivr_api_script"
    echo "   go build -o bin/ivr-api cmd/server/main.go"
fi

echo ""

# Check environment variables
echo "7. Checking .env configuration..."
if [ -f "/home/prabhat_qanditoday1/ivrcalling/ivr_api_script/.env" ]; then
    echo "✅ .env file exists"
    echo "SERVER_BASE_URL=$(grep SERVER_BASE_URL /home/prabhat_qanditoday1/ivrcalling/ivr_api_script/.env | grep -v '^#')"
else
    echo "❌ .env file not found"
fi

echo ""
echo "=== End Diagnostics ==="
