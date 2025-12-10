# Quick Start Instructions

## Running the Server

### Option 1: Using PowerShell Script (Recommended)

```powershell
.\start-server.ps1
```

### Option 2: Using Batch File

```cmd
start-server.bat
```

### Option 3: Direct Command

```powershell
go run cmd/server/main.go
```

Or if already built:

```powershell
.\bin\ivr-api.exe
```

### Option 4: In Background (PowerShell)

```powershell
Start-Process -FilePath ".\bin\ivr-api.exe" -NoNewWindow
```

## First Time Setup

1. **Copy environment file:**

   ```powershell
   Copy-Item .env.example .env
   ```

2. **Edit .env with your Twilio credentials:**

   - Get credentials from https://console.twilio.com/
   - Update TWILIO_ACCOUNT_SID
   - Update TWILIO_AUTH_TOKEN
   - Update TWILIO_PHONE_NUMBER

3. **Build the application:**

   ```powershell
   go build -o bin/ivr-api.exe cmd/server/main.go
   ```

4. **Run the server:**
   ```powershell
   .\start-server.ps1
   ```

## Testing the Server

```powershell
# Test health endpoint
Invoke-WebRequest -Uri http://localhost:8080/health -UseBasicParsing

# Test IVR config
Invoke-WebRequest -Uri http://localhost:8080/api/v1/config/ivr -UseBasicParsing

# Test TwiML generation
Invoke-WebRequest -Uri http://localhost:8080/api/v1/twiml/welcome -UseBasicParsing
```

## Making a Test Call

```powershell
$body = @{
    phone_number = "+919876543210"
} | ConvertTo-Json

Invoke-RestMethod -Uri http://localhost:8080/api/v1/calls/initiate `
  -Method Post `
  -Body $body `
  -ContentType "application/json"
```

## Using Python Example

```powershell
python examples/twilio_example.py +919876543210
```

## Stopping the Server

Press `Ctrl+C` in the terminal where the server is running.

## Common Issues

### "Press any key to execute code"

This happens when running Go files directly. Solution:

1. Build first: `go build -o bin/ivr-api.exe cmd/server/main.go`
2. Then run: `.\bin\ivr-api.exe`

### Port 8080 already in use

```powershell
# Find and kill process on port 8080
Get-NetTCPConnection -LocalPort 8080 | Select-Object -ExpandProperty OwningProcess | ForEach-Object { Stop-Process -Id $_ -Force }
```

### .env file not found

```powershell
Copy-Item .env.example .env
# Then edit .env with your credentials
```

## Next Steps

1. ✅ Server is running on http://localhost:8080
2. ✅ Test endpoints work
3. ⬜ Set up ngrok for public URL (see docs/TWILIO_SETUP.md)
4. ⬜ Configure Twilio credentials in .env
5. ⬜ Make your first test call!
