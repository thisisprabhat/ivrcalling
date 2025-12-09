# Q&I IVR Calling API

A robust Go-based REST API for initiating and managing IVR (Interactive Voice Response) calls with Q&I educational platform information.

## Features

- 🚀 Initiate outbound IVR calls programmatically
- 📞 Interactive voice menu with 3 options for callers
- 🔄 Handle callbacks from IVR provider
- 📋 RESTful API design
- 🔒 Environment-based configuration
- 📚 Comprehensive API documentation
- ✅ Production-ready architecture

## Quick Start

### Prerequisites

- Go 1.21 or higher
- An IVR service provider account (Twilio, Exotel, etc.)

### Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd ivr_api_script
```

2. Install dependencies:

```bash
go mod download
```

3. Configure environment variables:

```bash
cp .env.example .env
```

Edit `.env` with your Twilio credentials:

```env
PORT=8080
GIN_MODE=release
SERVER_BASE_URL=http://localhost:8080

# Get these from https://console.twilio.com/
TWILIO_ACCOUNT_SID=ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
TWILIO_AUTH_TOKEN=your_auth_token_here
TWILIO_PHONE_NUMBER=+1234567890

QI_TEAM_PHONE=+917905252436
```

4. Run the server:

```bash
go run cmd/server/main.go
```

The API will be available at `http://localhost:8080`

## Twilio Setup

For detailed Twilio integration instructions, see [Twilio Setup Guide](docs/TWILIO_SETUP.md).

**Quick Setup:**

1. Create a Twilio account at https://www.twilio.com/try-twilio
2. Get your Account SID and Auth Token from the console
3. Purchase a phone number with voice capability
4. Update `.env` with your credentials
5. For local testing, use ngrok to expose your server publicly

## API Endpoints

### 1. Initiate Call

**POST** `/api/v1/calls/initiate`

Initiates an IVR call to a phone number using Twilio.

**Request Body:**

```json
{
  "phone_number": "+919876543210",
  "callback_url": "https://yourapp.com/callback"
}
```

**Response:**

```json
{
  "call_id": "CAxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  "phone_number": "+919876543210",
  "status": "queued",
  "message": "Call initiated successfully via Twilio"
}
  "message": "Call initiated successfully"
}
```

### 2. Handle IVR Callback

**POST** `/api/v1/callbacks/ivr`

Receives callbacks from the IVR provider.

**Request Body:**

```json
{
  "call_id": "call_1733739600000000000",
  "event": "digit_pressed",
  "digit_input": "1",
  "timestamp": "2025-12-09T10:30:00Z"
}
```

### 3. Get IVR Configuration

**GET** `/api/v1/config/ivr`

Returns the current IVR flow configuration.

**Response:**

```json
{
  "intro_text": "Welcome to Q&I! We are transforming education...",
  "actions": [
    {
      "key": "1",
      "message": "To talk to Q&I team, press 1",
      "action": "forward",
      "forward_to": "+917905252436"
    },
    {
      "key": "2",
      "message": "To know more about Q&I, press 2",
      "action": "inform",
      "description": "Q&I is an AI-powered educational platform..."
    },
    {
      "key": "3",
      "message": "To hear this message again, press 3",
      "action": "repeat"
    }
  ],
  "end_message": "Thank you for contacting Q&I..."
}
```

### 4. Health Check

**GET** `/health`

Returns API health status.

**Response:**

```json
{
  "status": "healthy",
  "version": "1.0.0"
}
```

## IVR Flow

When a call is initiated, the caller hears:

1. **Intro Message**: Welcome message about Q&I platform
2. **Menu Options**:
   - Press 1: Forward to Q&I team (+917905252436)
   - Press 2: Hear detailed information about Q&I
   - Press 3: Repeat the message
3. **End Message**: Thank you message before call ends

## Project Structure

```
ivr_api_script/
├── cmd/
│   └── server/
│       └── main.go              # Application entry point
├── internal/
│   ├── api/
│   │   └── routes.go            # API route definitions
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── handlers/
│   │   └── call_handler.go      # HTTP request handlers
│   ├── models/
│   │   ├── call.go              # Call-related models
│   │   └── ivr_config.go        # IVR configuration
│   └── service/
│       ├── twilio_service.go    # Twilio integration & business logic
│       └── ivr_service.go       # Generic IVR service (deprecated)
├── docs/
│   ├── API.md                   # API documentation
│   ├── DEVELOPER_GUIDE.md       # Developer guide
│   └── TWILIO_SETUP.md          # Twilio integration guide
├── .env.example                 # Example environment variables
├── .gitignore
├── go.mod
└── readme.md
```

## Development

### Building

```bash
go build -o bin/ivr-api cmd/server/main.go
```

### Running

```bash
./bin/ivr-api
```

### Testing

```bash
go test ./...
```

## Configuration

All configuration is done through environment variables:

| Variable              | Description                        | Default               |
| --------------------- | ---------------------------------- | --------------------- |
| `PORT`                | Server port                        | 8080                  |
| `GIN_MODE`            | Gin framework mode (debug/release) | debug                 |
| `SERVER_BASE_URL`     | Public URL of your server          | http://localhost:8080 |
| `TWILIO_ACCOUNT_SID`  | Twilio Account SID                 | -                     |
| `TWILIO_AUTH_TOKEN`   | Twilio Auth Token                  | -                     |
| `TWILIO_PHONE_NUMBER` | Twilio phone number (E.164 format) | -                     |
| `QI_TEAM_PHONE`       | Q&I team phone number              | +917905252436         |

## Twilio Integration

This API is built specifically for Twilio. See the [Twilio Setup Guide](docs/TWILIO_SETUP.md) for detailed instructions.

### Quick Setup

1. **Get Twilio Credentials:**

   - Sign up at https://www.twilio.com/try-twilio
   - Get Account SID and Auth Token from console
   - Purchase a phone number with voice capability

2. **Configure Environment:**

   ```env
   TWILIO_ACCOUNT_SID=ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
   TWILIO_AUTH_TOKEN=your_auth_token_here
   TWILIO_PHONE_NUMBER=+1234567890
   SERVER_BASE_URL=https://yourdomain.com
   ```

3. **For Local Testing (ngrok):**

   ```bash
   # Terminal 1: Start server
   go run cmd/server/main.go

   # Terminal 2: Expose with ngrok
   ngrok http 8080

   # Update SERVER_BASE_URL with ngrok URL
   ```

### Twilio-Specific Features

- **TwiML Generation:** Automatic generation of voice responses
- **Call Forwarding:** Direct forwarding to Q&I team
- **DTMF Detection:** Capture user digit inputs (1, 2, 3)
- **Status Callbacks:** Real-time call status updates

## Security

- Always use HTTPS in production
- Keep your `.env` file secure and never commit it
- Validate all incoming phone numbers
- Implement rate limiting for production use
- Use API keys for authentication

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT License - see LICENSE file for details

## Support

For support, email support@qandi.com or open an issue in the repository.

## Roadmap

- [ ] Add database support for call logging
- [ ] Implement authentication middleware
- [ ] Add rate limiting
- [ ] Support for multiple languages
- [ ] Call recording capability
- [ ] Analytics dashboard
- [ ] Webhook retry mechanism
- [ ] SMS fallback option

---

Built with ❤️ for Q&I Educational Platform
