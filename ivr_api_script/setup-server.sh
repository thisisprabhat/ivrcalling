#!/bin/bash
# Setup script for IVR API on Ubuntu server

set -e

echo "🚀 Setting up IVR API Service..."
echo ""

# Check if running as root
if [ "$EUID" -ne 0 ]; then 
    echo "❌ Please run as root (use sudo)"
    exit 1
fi

PROJECT_DIR="/home/prabhat_qanditoday1/ivrcalling/ivr_api_script"
USER="prabhat_qanditoday1"

# Step 1: Build the Go binary
echo "1️⃣ Building Go binary..."
cd "$PROJECT_DIR"
if ! command -v go &> /dev/null; then
    echo "❌ Go is not installed. Installing Go 1.21..."
    wget https://go.dev/dl/go1.21.5.linux-amd64.tar.gz
    rm -rf /usr/local/go
    tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
    export PATH=$PATH:/usr/local/go/bin
    echo 'export PATH=$PATH:/usr/local/go/bin' >> /home/$USER/.bashrc
fi

# Build as the user
su - $USER -c "cd $PROJECT_DIR && go build -o bin/ivr-api cmd/server/main.go"
chmod +x "$PROJECT_DIR/bin/ivr-api"
echo "✅ Binary built successfully"

# Step 2: Install systemd service
echo ""
echo "2️⃣ Installing systemd service..."
cp "$PROJECT_DIR/ivr-api.service" /etc/systemd/system/ivr-api.service
systemctl daemon-reload
echo "✅ Service installed"

# Step 3: Enable and start the service
echo ""
echo "3️⃣ Starting IVR API service..."
systemctl enable ivr-api
systemctl restart ivr-api
sleep 2

if systemctl is-active --quiet ivr-api; then
    echo "✅ IVR API service is running"
else
    echo "❌ IVR API service failed to start"
    echo "Check logs with: journalctl -u ivr-api -n 50"
    exit 1
fi

# Step 4: Test the service
echo ""
echo "4️⃣ Testing the service..."
sleep 1
if curl -s http://localhost:8080/health > /dev/null 2>&1; then
    echo "✅ Health endpoint responding"
    curl -s http://localhost:8080/health
else
    echo "❌ Health endpoint not responding"
    exit 1
fi

# Step 5: Restart Caddy
echo ""
echo "5️⃣ Restarting Caddy..."
systemctl restart caddy
sleep 2

if systemctl is-active --quiet caddy; then
    echo "✅ Caddy is running"
else
    echo "❌ Caddy failed to start"
    echo "Check logs with: journalctl -u caddy -n 50"
fi

# Step 6: Test the full stack
echo ""
echo "6️⃣ Testing full stack..."
echo "Testing: https://ivr.waygosquad.com/health"
if curl -s https://ivr.waygosquad.com/health > /dev/null 2>&1; then
    echo "✅ Domain is accessible"
else
    echo "⚠️  Domain might not be accessible yet (DNS or firewall)"
fi

echo ""
echo "Testing: https://ivr.waygosquad.com/api/v1/twiml/welcome"
response=$(curl -s https://ivr.waygosquad.com/api/v1/twiml/welcome)
if [[ $response == *"<?xml"* ]]; then
    echo "✅ TwiML endpoint working!"
else
    echo "⚠️  TwiML endpoint might have issues"
    echo "Response: $response"
fi

echo ""
echo "✅ Setup complete!"
echo ""
echo "📋 Useful commands:"
echo "  • Check IVR API status:   systemctl status ivr-api"
echo "  • Check IVR API logs:     journalctl -u ivr-api -f"
echo "  • Check Caddy status:     systemctl status caddy"
echo "  • Check Caddy logs:       journalctl -u caddy -f"
echo "  • Restart IVR API:        systemctl restart ivr-api"
echo "  • Restart Caddy:          systemctl restart caddy"
echo ""
echo "🌐 Your API is accessible at: https://ivr.waygosquad.com"
