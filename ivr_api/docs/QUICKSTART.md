# IVR Calling System - Quick Start Guide

## Setup in 5 Minutes

### 1. Prerequisites
- Go 1.21+ installed
- Twilio account (free trial works)
- ngrok installed (for local testing)

### 2. Get Twilio Credentials

1. Sign up at [twilio.com](https://www.twilio.com/try-twilio)
2. Go to Console Dashboard
3. Note your:
   - Account SID
   - Auth Token
4. Get a phone number:
   - Phone Numbers → Buy a Number
   - Choose a voice-capable number

### 3. Quick Setup

```bash
# Navigate to project
cd /Users/prabhatkumar/Projects/golang/ivrcalling

# Copy environment template
cp .env.example .env

# Edit .env with your credentials
# (Use your favorite editor)
nano .env
```

Update these values in `.env`:
```env
TWILIO_ACCOUNT_SID=AC...your_sid
TWILIO_AUTH_TOKEN=your_token
TWILIO_PHONE_NUMBER=+1234567890
```

### 4. Install Dependencies

```bash
go mod download
```

### 5. Start ngrok (Development)

```bash
# In a separate terminal
ngrok http 8080
```

Copy the HTTPS URL (e.g., `https://abc123.ngrok.io`) and update `.env`:
```env
WEBHOOK_BASE_URL=https://abc123.ngrok.io
```

### 6. Run the Application

```bash
go run main.go
```

You should see:
```
Starting IVR Calling System on port 8080
```

### 7. Test It!

#### Create a Campaign
```bash
curl -X POST http://localhost:8080/api/campaigns \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Test Campaign",
    "language": "en",
    "is_active": true
  }'
```

#### Make a Test Call
```bash
curl -X POST http://localhost:8080/api/calls/bulk \
  -H "Content-Type: application/json" \
  -d '{
    "campaign_id": 1,
    "contacts": [
      {
        "phone_number": "+1YOUR_PHONE_NUMBER",
        "name": "Test User"
      }
    ]
  }'
```

**Replace `+1YOUR_PHONE_NUMBER` with your actual phone number in E.164 format!**

### 8. What Happens Next?

1. You'll receive a call from your Twilio number
2. Answer the call
3. You'll hear a welcome message
4. Follow the IVR menu:
   - Press **1** for product information
   - Press **2** for special offers
   - Press **3** to opt out
   - Press **9** to repeat menu

### 9. Check Call Status

```bash
# Get call details (use the call_id from the previous response)
curl http://localhost:8080/api/calls/1
```

## Common Issues

### "No .env file found"
- Make sure you created `.env` from `.env.example`
- Check file is in the project root directory

### "Failed to create call"
- Verify Twilio credentials are correct
- Check phone number is in E.164 format (+country code + number)
- Ensure Twilio account has sufficient credits

### "Call not connecting"
- Verify ngrok is running
- Check `WEBHOOK_BASE_URL` in `.env` matches ngrok URL
- Ensure ngrok URL is HTTPS

### "Invalid phone number"
- Use E.164 format: +[country][number]
- Examples: +11234567890 (US), +442012345678 (UK)

## Next Steps

1. **Read Full Documentation**: Check `README.md` for all features
2. **API Reference**: See `API_DOCUMENTATION.md` for all endpoints
3. **Customize Messages**: Edit `services/language_service.go`
4. **Add Languages**: Add more language support
5. **Production Deploy**: Set up proper hosting and domains

## Language Examples

Try different languages:

**Spanish Call:**
```bash
curl -X POST http://localhost:8080/api/calls/bulk \
  -H "Content-Type: application/json" \
  -d '{
    "campaign_id": 1,
    "language": "es",
    "contacts": [{"phone_number": "+34...", "name": "María"}]
  }'
```

**Hindi Call:**
```bash
curl -X POST http://localhost:8080/api/calls/bulk \
  -H "Content-Type: application/json" \
  -d '{
    "campaign_id": 1,
    "language": "hi",
    "contacts": [{"phone_number": "+91...", "name": "राज"}]
  }'
```

## Tips

- **Test Numbers**: Use your own phone number for testing
- **Twilio Console**: Check call logs at console.twilio.com
- **Database**: View calls in `ivr_calls.db` using SQLite browser
- **Logs**: Watch the console for real-time events

## Need Help?

1. Check the main `README.md`
2. Review `API_DOCUMENTATION.md`
3. Visit [Twilio Docs](https://www.twilio.com/docs)
4. Check Twilio Console for error logs

---

**You're all set! Start making IVR calls! 📞**
