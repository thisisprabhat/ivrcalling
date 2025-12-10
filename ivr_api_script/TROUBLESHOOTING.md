# 🔍 Troubleshooting: https://ivr.waygosquad.com/api/v1/twiml/welcome

## 🎯 Most Likely Issue

**Your Go API is probably not running on the server.**

Caddy is working and proxying requests to `localhost:8080`, but there's nothing listening on that port.

## ✅ Quick Fix

Run this on your server:

```bash
cd /home/prabhat_qanditoday1/ivrcalling/ivr_api_script
sudo bash quick-fix.sh
```

This will:

1. Build the Go binary (if needed)
2. Install systemd service
3. Start the IVR API
4. Restart Caddy
5. Test all endpoints

## 📋 Manual Steps

If you prefer to do it manually:

### 1. Build the Go Binary

```bash
cd /home/prabhat_qanditoday1/ivrcalling/ivr_api_script
go build -o bin/ivr-api cmd/server/main.go
chmod +x bin/ivr-api
```

### 2. Test It Locally

```bash
./bin/ivr-api
```

In another terminal:

```bash
curl http://localhost:8080/health
curl http://localhost:8080/api/v1/twiml/welcome
```

Press Ctrl+C to stop.

### 3. Install as Service (Recommended)

```bash
# Copy service file
sudo cp ivr-api.service /etc/systemd/system/

# Reload systemd
sudo systemctl daemon-reload

# Enable and start
sudo systemctl enable ivr-api
sudo systemctl start ivr-api

# Check status
sudo systemctl status ivr-api
```

### 4. Restart Caddy

```bash
sudo systemctl restart caddy
sudo systemctl status caddy
```

### 5. Test Everything

```bash
# Test local endpoint
curl http://localhost:8080/api/v1/twiml/welcome

# Test through Caddy
curl https://ivr.waygosquad.com/api/v1/twiml/welcome
```

## 🔍 Diagnostic Commands

### Check if Go API is running

```bash
# Check port 8080
sudo netstat -tlnp | grep 8080

# Or using ss
sudo ss -tlnp | grep 8080

# Check service status
sudo systemctl status ivr-api
```

### Check Logs

```bash
# IVR API logs
sudo journalctl -u ivr-api -f

# Caddy logs
sudo journalctl -u caddy -f

# Recent errors
sudo journalctl -u ivr-api -n 50
sudo journalctl -u caddy -n 50
```

### Run Diagnostics Script

```bash
cd /home/prabhat_qanditoday1/ivrcalling/ivr_api_script
bash diagnose.sh
```

## 🐛 Common Issues

### Issue 1: Binary Not Found

**Symptom**: `ivr-api.service` fails with "No such file"

**Fix**:

```bash
cd /home/prabhat_qanditoday1/ivrcalling/ivr_api_script
go build -o bin/ivr-api cmd/server/main.go
```

### Issue 2: Port Already in Use

**Symptom**: "bind: address already in use"

**Fix**:

```bash
# Find what's using port 8080
sudo lsof -i :8080

# Kill the process
sudo kill -9 <PID>

# Or change PORT in .env file
```

### Issue 3: Permission Denied

**Symptom**: Can't read .env file or bind to port

**Fix**:

```bash
# Fix file permissions
chmod +x /home/prabhat_qanditoday1/ivrcalling/ivr_api_script/bin/ivr-api
chmod 644 /home/prabhat_qanditoday1/ivrcalling/ivr_api_script/.env
```

### Issue 4: Caddy Returns 502 Bad Gateway

**Symptom**: Caddy is running but returns 502

**Fix**: Go API is not running

```bash
sudo systemctl start ivr-api
sudo systemctl status ivr-api
```

### Issue 5: Environment Variables Not Loaded

**Symptom**: "Twilio credentials not configured"

**Fix**: Ensure .env file is in the correct location

```bash
cd /home/prabhat_qanditoday1/ivrcalling/ivr_api_script
ls -la .env

# Should show something like:
# -rw-r--r-- 1 prabhat_qanditoday1 prabhat_qanditoday1 XXX bytes .env
```

## 📁 Files Created

1. **`diagnose.sh`** - Diagnostic script to check all components
2. **`ivr-api.service`** - Systemd service file for auto-start
3. **`setup-server.sh`** - Complete setup from scratch
4. **`quick-fix.sh`** - Quick restart and fix script

## 🎯 Expected Output

When working correctly:

```bash
$ curl http://localhost:8080/api/v1/twiml/welcome
<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice">Welcome to Q and I Educational Platform...</Say>
    <Gather numDigits="1" action="https://ivr.waygosquad.com/api/v1/twiml/handle-input" method="POST" timeout="10">
        <Say voice="alice">Press 1 to speak with our team. Press 2 to hear about our services. Press 3 to hear this message again.</Say>
    </Gather>
    <Say voice="alice">We did not receive any input. Goodbye!</Say>
</Response>
```

## 🚀 Next Steps After Fix

Once the endpoint is working:

1. **Update Twilio Webhook**:

   - Go to Twilio Console
   - Phone Numbers → Your Number
   - Voice & Fax → A Call Comes In
   - Set: `https://ivr.waygosquad.com/api/v1/twiml/welcome`

2. **Test a Call**:

   ```bash
   curl -X POST https://ivr.waygosquad.com/api/v1/calls/initiate \
     -H "Content-Type: application/json" \
     -d '{"phone_number": "+919876543210"}'
   ```

3. **Monitor Logs**:
   ```bash
   sudo journalctl -u ivr-api -f
   ```

## 💡 Pro Tips

- **Auto-start on boot**: The systemd service will start automatically
- **View real-time logs**: `journalctl -u ivr-api -f`
- **Restart after changes**: `sudo systemctl restart ivr-api`
- **Check health**: `curl http://localhost:8080/health`

---

**Need help?** Run the diagnostic script first:

```bash
bash diagnose.sh
```
