# Q&I IVR Calling API - Developer Guide

## Table of Contents

1. [Introduction](#introduction)
2. [Architecture Overview](#architecture-overview)
3. [Getting Started](#getting-started)
4. [API Usage Examples](#api-usage-examples)
5. [Integration with IVR Providers](#integration-with-ivr-providers)
6. [Error Handling](#error-handling)
7. [Best Practices](#best-practices)
8. [Testing](#testing)
9. [Deployment](#deployment)

## Introduction

The Q&I IVR Calling API is a Go-based REST API designed to initiate and manage automated phone calls with an Interactive Voice Response (IVR) system. This guide will help developers integrate, extend, and deploy the API.

### Key Features

- RESTful API design
- Modular and scalable architecture
- Easy integration with major IVR providers (Twilio, Exotel, etc.)
- Comprehensive error handling
- Environment-based configuration

## Architecture Overview

### Project Structure

```
ivr_api_script/
├── cmd/
│   └── server/          # Application entry point
├── internal/
│   ├── api/             # Route definitions
│   ├── config/          # Configuration management
│   ├── handlers/        # HTTP request handlers
│   ├── models/          # Data models and DTOs
│   └── service/         # Business logic layer
├── docs/                # Documentation
├── .env.example         # Environment template
└── go.mod              # Go dependencies
```

### Architecture Layers

1. **Handler Layer** (`internal/handlers/`)

   - Receives HTTP requests
   - Validates input
   - Calls service layer
   - Returns HTTP responses

2. **Service Layer** (`internal/service/`)

   - Contains business logic
   - Interacts with IVR providers
   - Processes callbacks
   - Manages IVR flow

3. **Model Layer** (`internal/models/`)

   - Defines data structures
   - Provides IVR configuration
   - Data transfer objects (DTOs)

4. **API Layer** (`internal/api/`)
   - Route registration
   - Middleware configuration

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Git
- An IVR service provider account
- Basic understanding of REST APIs

### Installation Steps

1. **Clone the repository:**

```bash
git clone <your-repo-url>
cd ivr_api_script
```

2. **Install dependencies:**

```bash
go mod download
```

3. **Set up environment variables:**

```bash
cp .env.example .env
```

4. **Configure your `.env` file:**

```env
PORT=8080
GIN_MODE=debug
IVR_PROVIDER_API_KEY=your_api_key
IVR_PROVIDER_API_SECRET=your_secret
IVR_PROVIDER_BASE_URL=https://api.provider.com
QI_TEAM_PHONE=+917905252436
```

5. **Run the application:**

```bash
go run cmd/server/main.go
```

## API Usage Examples

### Example 1: Initiate a Call

**Using cURL:**

```bash
curl -X POST http://localhost:8080/api/v1/calls/initiate \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "+919876543210",
    "callback_url": "https://yourapp.com/callbacks/ivr"
  }'
```

**Response:**

```json
{
  "call_id": "call_1733739600000000000",
  "phone_number": "+919876543210",
  "status": "initiated",
  "message": "Call initiated successfully"
}
```

### Example 2: Using JavaScript/Node.js

```javascript
const axios = require("axios");

async function initiateCall(phoneNumber) {
  try {
    const response = await axios.post(
      "http://localhost:8080/api/v1/calls/initiate",
      {
        phone_number: phoneNumber,
        callback_url: "https://yourapp.com/callbacks/ivr",
      }
    );

    console.log("Call initiated:", response.data);
    return response.data;
  } catch (error) {
    console.error("Error:", error.response.data);
  }
}

// Usage
initiateCall("+919876543210");
```

### Example 3: Using Python

```python
import requests

def initiate_call(phone_number):
    url = "http://localhost:8080/api/v1/calls/initiate"
    payload = {
        "phone_number": phone_number,
        "callback_url": "https://yourapp.com/callbacks/ivr"
    }

    response = requests.post(url, json=payload)

    if response.status_code == 200:
        print("Call initiated:", response.json())
        return response.json()
    else:
        print("Error:", response.json())
        return None

# Usage
initiate_call("+919876543210")
```

### Example 4: Get IVR Configuration

```bash
curl -X GET http://localhost:8080/api/v1/config/ivr
```

### Example 5: Handle Callbacks (Webhook Endpoint)

When the IVR provider sends callbacks, your application should be ready to receive them:

```javascript
// Express.js example
app.post("/callbacks/ivr", (req, res) => {
  const { call_id, event, digit_input } = req.body;

  console.log(`Call ${call_id}: Event ${event}, Digit: ${digit_input}`);

  // Process the callback
  // You can forward this to the Q&I API

  res.json({ status: "success" });
});
```

## Integration with IVR Providers

### Twilio Integration

**Step 1: Configure Twilio Credentials**

```env
IVR_PROVIDER_BASE_URL=https://api.twilio.com/2010-04-01
IVR_PROVIDER_API_KEY=ACxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
IVR_PROVIDER_API_SECRET=your_auth_token
```

**Step 2: Modify the service (if needed)**

The `ivr_service.go` includes a template for calling IVR providers. For Twilio, you might need to adjust the payload:

```go
// In internal/service/ivr_service.go, modify callIVRProvider method
payload := map[string]interface{}{
    "To":   phoneNumber,
    "From": s.config.TwilioPhoneNumber,
    "Url":  "http://yourserver.com/twiml", // TwiML instructions
}
```

### Exotel Integration

**Step 1: Configure Exotel Credentials**

```env
IVR_PROVIDER_BASE_URL=https://api.exotel.com/v1/Accounts/YOUR_SID
IVR_PROVIDER_API_KEY=your_api_key
IVR_PROVIDER_API_SECRET=your_api_token
```

**Step 2: Adjust API calls for Exotel format**

Exotel typically requires different parameter names. Update the payload in `callIVRProvider`:

```go
payload := map[string]interface{}{
    "From": s.config.ExotelVirtualNumber,
    "To":   phoneNumber,
    "Url":  callbackURL,
}
```

## Error Handling

### HTTP Status Codes

| Status Code | Description                 |
| ----------- | --------------------------- |
| 200         | Success                     |
| 400         | Bad Request (invalid input) |
| 500         | Internal Server Error       |

### Error Response Format

```json
{
  "error": "Failed to initiate call",
  "message": "Invalid phone number format"
}
```

### Handling Errors in Your Code

```go
response, err := h.ivrService.InitiateCall(req.PhoneNumber, req.CallbackURL)
if err != nil {
    c.JSON(http.StatusInternalServerError, models.ErrorResponse{
        Error:   "Failed to initiate call",
        Message: err.Error(),
    })
    return
}
```

## Best Practices

### 1. Phone Number Validation

Always validate phone numbers before making calls:

```go
func isValidPhoneNumber(phone string) bool {
    if len(phone) < 10 || phone[0] != '+' {
        return false
    }
    // Add more validation as needed
    return true
}
```

### 2. Use Environment Variables

Never hardcode sensitive information:

```go
// Bad
apiKey := "sk_test_123456789"

// Good
apiKey := os.Getenv("IVR_PROVIDER_API_KEY")
```

### 3. Implement Rate Limiting

For production, add rate limiting to prevent abuse:

```go
import "github.com/gin-gonic/gin"
import "golang.org/x/time/rate"

func rateLimitMiddleware() gin.HandlerFunc {
    limiter := rate.NewLimiter(10, 100) // 10 requests/sec, burst of 100

    return func(c *gin.Context) {
        if !limiter.Allow() {
            c.JSON(429, gin.H{"error": "Too many requests"})
            c.Abort()
            return
        }
        c.Next()
    }
}
```

### 4. Add Logging

Use structured logging for better debugging:

```go
import "log"

log.Printf("Initiating call to %s with ID %s", phoneNumber, callID)
```

### 5. Database Integration (Optional)

For production, store call records in a database:

```go
type CallRecord struct {
    ID          string
    PhoneNumber string
    Status      string
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

// Save to database
func (s *IVRService) saveCallRecord(record *CallRecord) error {
    // Your database logic here
    return nil
}
```

## Testing

### Unit Testing Example

Create a file `internal/service/ivr_service_test.go`:

```go
package service

import (
    "testing"
    "github.com/qandi/ivr-calling-api/internal/config"
)

func TestInitiateCall(t *testing.T) {
    cfg := &config.Config{
        IVRProviderAPIKey: "test_key",
    }

    service := NewIVRService(cfg)

    response, err := service.InitiateCall("+919876543210", "http://callback.url")

    if err != nil {
        t.Errorf("Expected no error, got %v", err)
    }

    if response.Status != "initiated" {
        t.Errorf("Expected status 'initiated', got %s", response.Status)
    }
}
```

Run tests:

```bash
go test ./...
```

### Integration Testing

Use tools like Postman or write integration tests:

```bash
# Test health endpoint
curl http://localhost:8080/health

# Test call initiation
curl -X POST http://localhost:8080/api/v1/calls/initiate \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "+919876543210"}'
```

## Deployment

### Building for Production

```bash
# Build binary
go build -o bin/ivr-api cmd/server/main.go

# Run in production mode
GIN_MODE=release ./bin/ivr-api
```

### Docker Deployment

Create a `Dockerfile`:

```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o ivr-api cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/
COPY --from=builder /app/ivr-api .
COPY --from=builder /app/.env .

EXPOSE 8080
CMD ["./ivr-api"]
```

Build and run:

```bash
docker build -t qandi-ivr-api .
docker run -p 8080:8080 --env-file .env qandi-ivr-api
```

### Environment-Specific Configuration

**Development:**

```env
GIN_MODE=debug
PORT=8080
```

**Production:**

```env
GIN_MODE=release
PORT=80
```

### Health Checks

Configure health checks for your deployment platform:

```bash
# Health check endpoint
curl http://localhost:8080/health
```

## Extending the API

### Adding New Actions

To add a new IVR action, modify `internal/models/ivr_config.go`:

```go
{
    Key:     "4",
    Message: "To schedule a demo, press 4",
    Action:  "schedule_demo",
}
```

Then handle it in `internal/service/ivr_service.go`:

```go
case "schedule_demo":
    fmt.Printf("Scheduling demo for call %s\n", callID)
    // Add your logic
```

### Adding Authentication

Add JWT authentication middleware:

```go
func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")

        if token == "" {
            c.JSON(401, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // Validate token
        c.Next()
    }
}

// Apply to routes
calls.POST("/initiate", authMiddleware(), callHandler.InitiateCall)
```

## Troubleshooting

### Common Issues

**Issue: "Connection refused"**

- Check if the server is running
- Verify the PORT environment variable

**Issue: "Invalid phone number"**

- Ensure phone numbers start with +
- Include country code

**Issue: "IVR provider error"**

- Verify API credentials
- Check IVR provider status
- Review API logs

## Support and Resources

- **Documentation**: `/docs` directory
- **API Reference**: See `docs/API.md`
- **Issue Tracker**: GitHub Issues
- **Email**: support@qandi.com

## Conclusion

This developer guide provides a comprehensive overview of the Q&I IVR Calling API. For specific questions or contributions, please refer to the README.md or contact the development team.

---

**Last Updated**: December 9, 2025  
**Version**: 1.0.0
