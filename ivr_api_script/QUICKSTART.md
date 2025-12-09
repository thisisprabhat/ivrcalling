# Quick Start Guide

## Running the Application

1. **Install dependencies:**

```bash
go mod download
```

2. **Set up environment:**

```bash
cp .env.example .env
# Edit .env with your credentials
```

3. **Run the server:**

```bash
go run cmd/server/main.go
```

## Testing the API

### Using cURL

**Health Check:**

```bash
curl http://localhost:8080/health
```

**Initiate a Call:**

```bash
curl -X POST http://localhost:8080/api/v1/calls/initiate \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "+919876543210",
    "callback_url": "https://yourapp.com/callback"
  }'
```

**Get IVR Config:**

```bash
curl http://localhost:8080/api/v1/config/ivr
```

### Using Postman

Import the Postman collection from `docs/postman_collection.json` to test all endpoints easily.

## Project Structure

```
.
├── cmd/server/main.go          # Entry point
├── internal/
│   ├── api/routes.go           # API routes
│   ├── config/config.go        # Configuration
│   ├── handlers/               # HTTP handlers
│   ├── models/                 # Data models
│   └── service/                # Business logic
├── docs/                       # Documentation
├── .env.example               # Environment template
└── go.mod                     # Dependencies
```

## Next Steps

- Read the [Developer Guide](docs/DEVELOPER_GUIDE.md) for detailed information
- Check [API Documentation](docs/API.md) for endpoint details
- Configure your IVR provider credentials
- Implement your callback endpoint

## Common Commands

```bash
# Run the server
go run cmd/server/main.go

# Build binary
go build -o bin/ivr-api cmd/server/main.go

# Run tests
go test ./...

# Format code
go fmt ./...

# Check for issues
go vet ./...
```

## Support

For help, see:

- [README.md](README.md)
- [Developer Guide](docs/DEVELOPER_GUIDE.md)
- [API Documentation](docs/API.md)
