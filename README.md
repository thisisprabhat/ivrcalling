# IVR Calling System for Marketing

A comprehensive Go-based Interactive Voice Response (IVR) system for marketing campaigns with Twilio integration and multilanguage support.

## Features

✨ **Core Features:**
- 🌍 **Multilanguage Support** - English, Spanish, French, German, and Hindi
- 📞 **Bulk Call Initiation** - Send calls to multiple contacts simultaneously
- 🎯 **Campaign Management** - Create and manage marketing campaigns
- 📊 **Call Tracking** - Real-time call status and detailed logging
- 🔄 **IVR Menu System** - Interactive voice menus with digit input
- 💾 **MongoDB Database** - Persistent storage for campaigns and call logs
- 🔗 **Twilio Integration** - Enterprise-grade telephony service
- 📝 **Comprehensive API** - RESTful endpoints for all operations

## Prerequisites

- Go 1.21 or higher
- MongoDB 4.4 or higher (local or MongoDB Atlas)
- Twilio account with:
  - Account SID
  - Auth Token
  - Active phone number
- Public webhook URL (for production) or ngrok (for development)

## Installation

1. **Clone the repository:**
```bash
cd /Users/prabhatkumar/Projects/golang/ivrcalling
```

2. **Install dependencies:**
```bash
go mod download
```

3. **Configure environment variables:**
```bash
cp .env.example .env
```

Edit `.env` with your Twilio credentials:
```env
PORT=8080
ENV=development

TWILIO_ACCOUNT_SID=your_account_sid_here
TWILIO_AUTH_TOKEN=your_auth_token_here
TWILIO_PHONE_NUMBER=+1234567890

MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=ivr_calling_system
DEFAULT_LANGUAGE=en
WEBHOOK_BASE_URL=https://your-domain.com
```

4. **Build and run:**
```bash
go build -o ivr-system
./ivr-system
```

Or run directly:
```bash
go run main.go
```

## Project Structure

```
ivrcalling/
├── main.go                 # Application entry point
├── config/
│   └── config.go          # Configuration management
├── models/
│   └── models.go          # Data models
├── database/
│   └── database.go        # Database initialization
├── services/
│   ├── twilio_service.go  # Twilio API integration
│   ├── language_service.go # Multilanguage support
│   └── twiml_service.go   # TwiML generation
├── handlers/
│   ├── campaign_handler.go # Campaign endpoints
│   ├── call_handler.go    # Call management
│   └── webhook_handler.go # Twilio webhooks
└── routes/
    └── routes.go          # API route definitions
```

## API Documentation

### Base URL
```
http://localhost:8080/api
```

### Endpoints

#### 1. Health Check
```http
GET /api/health
```

**Response:**
```json
{
  "status": "healthy",
  "service": "IVR Calling System"
}
```

#### 2. Get Supported Languages
```http
GET /api/languages
```

**Response:**
```json
{
  "languages": ["en", "es", "fr", "de", "hi"]
}
```

### Campaign Management

#### Create Campaign
```http
POST /api/campaigns
Content-Type: application/json

{
  "name": "Summer Sale 2025",
  "description": "Promotional campaign for summer products",
  "language": "en",
  "is_active": true
}
```

**Response:**
```json
{
  "id": 1,
  "name": "Summer Sale 2025",
  "description": "Promotional campaign for summer products",
  "language": "en",
  "is_active": true,
  "created_at": "2025-11-30T10:00:00Z",
  "updated_at": "2025-11-30T10:00:00Z"
}
```

#### List All Campaigns
```http
GET /api/campaigns
```

#### Get Campaign Details
```http
GET /api/campaigns/{id}
```

#### Update Campaign
```http
PUT /api/campaigns/{id}
Content-Type: application/json

{
  "name": "Updated Campaign Name",
  "is_active": false
}
```

#### Delete Campaign
```http
DELETE /api/campaigns/{id}
```

### Call Management

#### Initiate Bulk Calls (Main Endpoint)
```http
POST /api/calls/bulk
Content-Type: application/json

{
  "campaign_id": 1,
  "language": "en",
  "contacts": [
    {
      "phone_number": "+1234567890",
      "name": "John Doe"
    },
    {
      "phone_number": "+0987654321",
      "name": "Jane Smith"
    }
  ]
}
```

**Parameters:**
- `campaign_id` (required): ID of the campaign
- `language` (optional): Language code (defaults to campaign language)
- `contacts` (required): Array of contact objects
  - `phone_number` (required): E.164 format phone number
  - `name` (optional): Customer name

**Response:**
```json
{
  "message": "Bulk calls initiated",
  "success_count": 2,
  "fail_count": 0,
  "call_ids": [1, 2]
}
```

#### Get Call Status
```http
GET /api/calls/{id}
```

**Response:**
```json
{
  "id": 1,
  "campaign_id": 1,
  "phone_number": "+1234567890",
  "customer_name": "John Doe",
  "status": "completed",
  "twilio_call_sid": "CA1234567890abcdef",
  "language": "en",
  "duration": 45,
  "created_at": "2025-11-30T10:00:00Z",
  "updated_at": "2025-11-30T10:01:00Z",
  "call_logs": [
    {
      "id": 1,
      "call_id": 1,
      "event": "initiated",
      "details": "Call initiated to +1234567890",
      "created_at": "2025-11-30T10:00:00Z"
    }
  ]
}
```

#### Get Campaign Calls with Statistics
```http
GET /api/campaigns/{id}/calls
```

**Response:**
```json
{
  "calls": [...],
  "stats": {
    "total": 10,
    "pending": 2,
    "initiated": 3,
    "completed": 4,
    "failed": 1
  }
}
```

## IVR Menu Flow

When a call is initiated, the recipient experiences:

1. **Welcome Message**: Personalized greeting with customer name
2. **Main Menu**:
   - Press 1: Product information
   - Press 2: Special offers
   - Press 3: Opt out from calls
   - Press 9: Repeat menu

3. **Sub-menus**:
   - Product info: Detailed product description
   - Offers: Current promotional details
   - Opt-out: Confirmation flow

4. **Actions**:
   - Press 0: Return to main menu
   - Press 9: Repeat current menu

## Call Statuses

- `pending`: Call created, not yet initiated
- `initiated`: Call sent to Twilio
- `in-progress`: Call is active
- `completed`: Call finished successfully
- `failed`: Call failed or was not answered

## Multilanguage Support

The system supports 5 languages with complete IVR scripts:

| Code | Language | Voice Locale |
|------|----------|--------------|
| `en` | English  | en-US        |
| `es` | Spanish  | es-ES        |
| `fr` | French   | fr-FR        |
| `de` | German   | de-DE        |
| `hi` | Hindi    | hi-IN        |

All IVR messages are automatically translated and spoken in the selected language.

## Development Setup

### Using ngrok for local development:

1. **Install ngrok:**
```bash
brew install ngrok
# or download from https://ngrok.com
```

2. **Start ngrok:**
```bash
ngrok http 8080
```

3. **Update `.env`:**
```env
WEBHOOK_BASE_URL=https://your-ngrok-url.ngrok.io
```

4. **Run the application:**
```bash
go run main.go
```

## Database Schema (MongoDB)

### Campaigns Collection
- `_id`: ObjectId
- `name`: Campaign name
- `description`: Campaign description
- `language`: Default language
- `is_active`: Active status
- `created_at`, `updated_at`: Timestamps

### Calls Collection
- `_id`: ObjectId
- `campaign_id`: Reference to campaigns collection
- `phone_number`: Recipient phone number
- `customer_name`: Customer name
- `status`: Call status
- `twilio_call_sid`: Twilio identifier
- `language`: Call language
- `duration`: Call duration in seconds
- `error_message`: Error details (if failed)
- `created_at`, `updated_at`: Timestamps

### Call Logs Collection
- `_id`: ObjectId
- `call_id`: Reference to calls collection
- `event`: Event type
- `details`: Event details
- `user_input`: User DTMF input
- `created_at`: Timestamp

## Testing the API

### Example: Create a campaign and make calls

1. **Create a campaign:**
```bash
curl -X POST http://localhost:8080/api/campaigns \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Black Friday Sale",
    "description": "Promotional calls for Black Friday",
    "language": "en",
    "is_active": true
  }'
```

2. **Initiate bulk calls:**
```bash
curl -X POST http://localhost:8080/api/calls/bulk \
  -H "Content-Type: application/json" \
  -d '{
    "campaign_id": 1,
    "language": "en",
    "contacts": [
      {
        "phone_number": "+1234567890",
        "name": "Alice Johnson"
      },
      {
        "phone_number": "+0987654321",
        "name": "Bob Williams"
      }
    ]
  }'
```

3. **Check call status:**
```bash
curl http://localhost:8080/api/calls/1
```

## Error Handling

The API returns standard HTTP status codes:

- `200 OK`: Success
- `201 Created`: Resource created
- `400 Bad Request`: Invalid request
- `404 Not Found`: Resource not found
- `500 Internal Server Error`: Server error

Error response format:
```json
{
  "error": "Error message description"
}
```

## Security Considerations

1. **Validate Twilio Webhooks**: Implement Twilio signature validation
2. **Rate Limiting**: Add rate limiting to prevent abuse
3. **Authentication**: Implement API authentication (JWT, API keys)
4. **HTTPS**: Always use HTTPS in production
5. **Phone Number Validation**: Validate E.164 format
6. **Opt-out List**: Maintain and respect opt-out list

## Customization

### Adding New Languages

Edit `services/language_service.go`:

```go
"it": { // Italian
    Welcome:     "Ciao %s, benvenuto...",
    MainMenu:    "Premi 1 per...",
    // ... other strings
}
```

### Customizing IVR Menu

Edit `services/twiml_service.go` to modify menu options and flow.

### Changing Voice

Modify the `voice` parameter in TwiML generation:
- Options: `alice`, `man`, `woman`, `Polly.*`

## Production Deployment

1. **Set environment to production:**
```env
ENV=production
```

2. **Use production MongoDB:**
```env
MONGODB_URI=mongodb+srv://user:password@cluster.mongodb.net/
MONGODB_DATABASE=ivr_calling_system
```

3. **Configure reverse proxy** (nginx/caddy)

4. **Set up monitoring** and logging

5. **Implement backup strategy** for database

## Troubleshooting

### Calls not initiating:
- Verify Twilio credentials
- Check phone number format (E.164)
- Ensure webhook URL is accessible
- Check Twilio account balance

### Webhooks not working:
- Verify `WEBHOOK_BASE_URL` is publicly accessible
- Check ngrok is running (development)
- Review Twilio webhook logs

### Database errors:
- Ensure MongoDB is running (start with `mongod` or check MongoDB Atlas connection)
- Verify `MONGODB_URI` in `.env` is correct
- Check network connectivity to MongoDB server

## Contributing

1. Fork the repository
2. Create feature branch
3. Commit changes
4. Push to branch
5. Create Pull Request

## License

MIT License - feel free to use this project for commercial purposes.

## Support

For issues and questions:
- Create an issue in the repository
- Check Twilio documentation: https://www.twilio.com/docs
- Review Go Gin documentation: https://gin-gonic.com/docs/

## Changelog

### v1.0.0 (2025-11-30)
- Initial release
- Multilanguage support (5 languages)
- Bulk call initiation
- Campaign management
- Complete IVR menu system
- Call tracking and logging

---

**Built with ❤️ using Go, Gin, GORM, and Twilio**
