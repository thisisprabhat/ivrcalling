#!/bin/bash
# Quick fix script - restarts everything in the right order

echo "🔧 Quick Fix - Restarting IVR API stack..."
echo ""

# Check if running as root
if [ "$EUID" -ne 0 ]; then 
    echo "Please run with sudo: sudo bash quick-fix.sh"
    exit 1
fi

PROJECT_DIR="/home/prabhat_qanditoday1/ivrcalling/ivr_api_script"

# Check if binary exists
if [ ! -f "$PROJECT_DIR/bin/ivr-api" ]; then
    echo "⚠️  Binary not found. Building..."
    cd "$PROJECT_DIR"
    su - prabhat_qanditoday1 -c "cd $PROJECT_DIR && go build -o bin/ivr-api cmd/server/main.go"
    chmod +x "$PROJECT_DIR/bin/ivr-api"
fi

# Stop services
echo "1️⃣ Stopping services..."
systemctl stop ivr-api 2>/dev/null || true
systemctl stop caddy 2>/dev/null || true

# Install/update systemd service
echo ""
echo "2️⃣ Installing systemd service..."
cp "$PROJECT_DIR/ivr-api.service" /etc/systemd/system/ivr-api.service
systemctl daemon-reload

# Start IVR API
echo ""
echo "3️⃣ Starting IVR API..."
systemctl start ivr-api
sleep 2

if systemctl is-active --quiet ivr-api; then
    echo "✅ IVR API is running"
    curl -s http://localhost:8080/health
else
    echo "❌ IVR API failed to start"
    echo "Checking logs..."
    journalctl -u ivr-api -n 20 --no-pager
    exit 1
fi

# Start Caddy
echo ""
echo "4️⃣ Starting Caddy..."
systemctl start caddy
sleep 2

if systemctl is-active --quiet caddy; then
    echo "✅ Caddy is running"
else
    echo "❌ Caddy failed to start"
    journalctl -u caddy -n 20 --no-pager
    exit 1
fi

# Test endpoints
echo ""
echo "5️⃣ Testing endpoints..."
echo "Local: http://localhost:8080/api/v1/twiml/welcome"
curl -s http://localhost:8080/api/v1/twiml/welcome | head -n 3

echo ""
echo "Domain: https://ivr.waygosquad.com/api/v1/twiml/welcome"
curl -s https://ivr.waygosquad.com/api/v1/twiml/welcome | head -n 3

echo ""
echo "✅ All services restarted!"
echo ""
echo "Check status:"
echo "  systemctl status ivr-api"
echo "  systemctl status caddy"
