# Twilio Integration Guide for Q&I IVR API

## Overview

This guide will help you integrate Twilio with the Q&I IVR Calling API.

## Prerequisites

1. A Twilio account ([Sign up here](https://www.twilio.com/try-twilio))
2. A Twilio phone number with voice capability
3. Your server publicly accessible (for Twilio to send webhooks)

## Step 1: Get Twilio Credentials

1. Log in to your [Twilio Console](https://console.twilio.com/)
2. Find your **Account SID** and **Auth Token** on the dashboard
3. Purchase a phone number with voice capability (Phone Numbers → Buy a Number)

## Step 2: Configure Environment Variables

Update your `.env` file with Twilio credentials:

```env
# Server Configuration
PORT=8080
GIN_MODE=release
SERVER_BASE_URL=https://yourdomain.com

# Twilio Configuration
TWILIO_ACCOUNT_SID=ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
TWILIO_AUTH_TOKEN=your_auth_token_here
TWILIO_PHONE_NUMBER=+1234567890

# Q&I Configuration
QI_TEAM_PHONE=+917905252436
```

**Important:**

- `SERVER_BASE_URL` must be publicly accessible (use ngrok for testing)
- `TWILIO_PHONE_NUMBER` must be in E.164 format (+[country][number])

## Step 3: Make Your Server Publicly Accessible

### For Development (using ngrok):

1. Install ngrok: https://ngrok.com/download
2. Start your server:
   ```bash
   go run cmd/server/main.go
   ```
3. In another terminal, run ngrok:
   ```bash
   ngrok http 8080
   ```
4. Copy the HTTPS URL (e.g., `https://abc123.ngrok.io`)
5. Update `SERVER_BASE_URL` in your `.env` file:
   ```env
   SERVER_BASE_URL=https://abc123.ngrok.io
   ```
6. Restart your server

### For Production:

Deploy to a cloud service with HTTPS:

- AWS EC2 + Load Balancer
- Google Cloud Run
- Heroku
- DigitalOcean App Platform

## Step 4: Test the Integration

### 1. Health Check

```bash
curl https://yourdomain.com/health
```

### 2. Initiate a Call

```bash
curl -X POST https://yourdomain.com/api/v1/calls/initiate \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "+919876543210",
    "callback_url": "https://yourdomain.com/api/v1/callbacks/ivr"
  }'
```

Expected response:

```json
{
  "call_id": "CAxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "phone_number": "+919876543210",
  "status": "queued",
  "message": "Call initiated successfully via Twilio"
}
```

## Step 5: Verify IVR Flow

When the call is answered, the recipient will hear:

1. **Welcome Message:**

   > "Welcome to Q&I! We are transforming education with smart digital tools..."

2. **Menu Options:**

   - Press 1: Forward to Q&I team (+917905252436)
   - Press 2: Hear detailed information about Q&I
   - Press 3: Repeat the message

3. **End Message:**
   > "Thank you for contacting Q&I. We look forward to helping your school achieve success. Goodbye!"

## API Endpoints

### Initiate Call

```
POST /api/v1/calls/initiate
```

### TwiML Endpoints (used by Twilio)

```
GET/POST /api/v1/twiml/welcome
POST /api/v1/twiml/handle-input
```

### Callbacks

```
POST /api/v1/callbacks/ivr
```

## TwiML Flow Diagram

```
[Call Initiated]
       ↓
[/twiml/welcome] → Play intro message + menu
       ↓
[User presses digit]
       ↓
[/twiml/handle-input]
       ↓
   ┌───┴────┬─────────┐
   ↓        ↓         ↓
[Digit 1] [Digit 2] [Digit 3]
   ↓        ↓         ↓
[Forward] [Info]  [Repeat]
```

## Troubleshooting

### Issue: "Twilio credentials not configured"

**Solution:** Verify `TWILIO_ACCOUNT_SID` and `TWILIO_AUTH_TOKEN` are set in `.env`

### Issue: "Twilio phone number not configured"

**Solution:** Set `TWILIO_PHONE_NUMBER` in `.env` with a valid Twilio number

### Issue: Call doesn't connect

**Solutions:**

- Verify phone number is in E.164 format (+[country][number])
- Check Twilio phone number has voice capability
- Ensure you have sufficient Twilio balance
- Verify the destination phone number is valid

### Issue: IVR menu doesn't play

**Solutions:**

- Ensure `SERVER_BASE_URL` is publicly accessible
- Check Twilio debugger: https://console.twilio.com/debugger
- Verify the `/twiml/welcome` endpoint is accessible

### Issue: Digit presses not working

**Solutions:**

- Check `/twiml/handle-input` endpoint is accessible
- Review Twilio debugger for webhook errors
- Ensure `SERVER_BASE_URL` is correct

## Monitoring Calls

### View Call Logs in Twilio Console

1. Go to [Twilio Console](https://console.twilio.com/)
2. Navigate to **Monitor** → **Logs** → **Calls**
3. Click on a call SID to see details

### Check Webhook Debugger

1. Go to **Monitor** → **Debugger**
2. Look for any failed webhook requests
3. Check error messages and fix issues

## Cost Considerations

- **Outbound Calls:** ~$0.013 - $0.04 per minute (varies by country)
- **Phone Number:** ~$1 - $2 per month
- **Inbound Minutes:** Usually free or very low cost

Check [Twilio Pricing](https://www.twilio.com/voice/pricing) for exact rates.

## Advanced Features

### Add Call Recording

Update the call initiation to include recording:

```go
data.Set("Record", "true")
data.Set("RecordingStatusCallback", callbackURL + "/recording")
```

### Add Voicemail Detection

```go
data.Set("MachineDetection", "Enable")
```

### Custom Voice and Language

Modify the TwiML generation in `twilio_service.go`:

```go
<Say voice="Polly.Aditi" language="en-IN">...</Say>
```

## Security Best Practices

1. **Validate Twilio Requests:** Verify requests come from Twilio
2. **Use HTTPS:** Always use HTTPS for webhooks
3. **Secure Credentials:** Never commit `.env` file
4. **Rate Limiting:** Implement rate limiting for API endpoints
5. **IP Whitelisting:** Restrict access to Twilio's IP ranges

## Example: Validating Twilio Webhooks

```go
import (
    "crypto/hmac"
    "crypto/sha1"
    "encoding/base64"
)

func validateTwilioRequest(authToken, url string, params map[string]string, signature string) bool {
    // Implementation of Twilio signature validation
    // See: https://www.twilio.com/docs/usage/security#validating-requests
}
```

## Support

- **Twilio Documentation:** https://www.twilio.com/docs/voice
- **Twilio Support:** https://support.twilio.com/
- **Q&I API Issues:** Open an issue on GitHub

## Next Steps

1. Test with multiple phone numbers
2. Add call recording and transcription
3. Implement call analytics
4. Add multi-language support
5. Set up call queuing for high volume

---

**Last Updated:** December 9, 2025  
**Version:** 1.0.0
