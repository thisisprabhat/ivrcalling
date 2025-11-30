# API Documentation with Scalar

The IVR Calling System now includes beautiful, interactive API documentation powered by Scalar.

## 🚀 Accessing the Documentation

Once your server is running, visit:

```
http://localhost:8080/docs
```

## 📖 Features

### Interactive API Explorer
- **Try It Out**: Test all API endpoints directly from the browser
- **Request Examples**: See cURL, JavaScript, and Python examples
- **Response Previews**: View example responses for all endpoints
- **Schema Browser**: Explore data models and structures

### Modern UI
- **Dark Mode**: Beautiful dark theme by default
- **Search**: Quick search with `Cmd+K` (Mac) or `Ctrl+K` (Windows/Linux)
- **Sidebar Navigation**: Easy navigation between endpoints
- **Responsive**: Works on desktop, tablet, and mobile

### Complete Coverage
- ✅ **Health & System** endpoints
- ✅ **Campaign Management** (CRUD operations)
- ✅ **Call Management** (bulk calls, status tracking)
- ✅ **Language Support** information
- ✅ **Webhook Endpoints** (Twilio callbacks)

## 📋 What's Documented

### Campaign Management
- `POST /api/campaigns` - Create a new campaign
- `GET /api/campaigns` - List all campaigns
- `GET /api/campaigns/:id` - Get campaign details
- `PUT /api/campaigns/:id` - Update campaign
- `DELETE /api/campaigns/:id` - Delete campaign
- `GET /api/campaigns/:id/calls` - Get campaign calls with statistics

### Call Management
- `POST /api/calls/bulk` - Initiate bulk calls (main feature)
- `GET /api/calls/:id` - Get call status and logs

### System Endpoints
- `GET /api/health` - Health check
- `GET /api/languages` - Supported languages

### Webhooks (Internal/Twilio)
- `POST /api/webhook/voice` - Initial call webhook
- `POST /api/webhook/gather` - User input handling
- `POST /api/webhook/status` - Call status updates
- `POST /api/webhook/optout` - Opt-out confirmation

## 🎨 Customization

The documentation includes:
- **Purple theme** matching modern API doc standards
- **Phone emoji (📞)** as favicon
- **Searchable** with keyboard shortcuts
- **Dark mode** enabled by default
- **Request/Response examples** in multiple languages

## 📝 OpenAPI Specification

The API follows OpenAPI 3.0 specification. The raw YAML file is available at:

```
http://localhost:8080/docs/swagger.yaml
```

You can import this file into:
- Postman
- Insomnia
- Any OpenAPI-compatible tool

## 🔧 Files Added

```
/docs/swagger.yaml                  # OpenAPI 3.0 specification
/handlers/docs_handler.go           # Scalar documentation handler
```

## 🌐 Production Deployment

When deploying to production:

1. **Update server URL** in `docs/swagger.yaml`:
```yaml
servers:
  - url: https://api.yourcompany.com
    description: Production server
```

2. **Add authentication** if needed (API keys, JWT, etc.)

3. **Enable HTTPS** for security

4. **Consider rate limiting** to protect the API

## 💡 Usage Examples

The documentation includes complete examples for:

### cURL
```bash
curl -X POST http://localhost:8080/api/campaigns \
  -H "Content-Type: application/json" \
  -d '{"name": "Summer Sale", "description": "Q3 Campaign"}'
```

### JavaScript (fetch)
```javascript
fetch('http://localhost:8080/api/campaigns', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({
    name: 'Summer Sale',
    description: 'Q3 Campaign'
  })
})
```

### Python (requests)
```python
import requests

response = requests.post(
    'http://localhost:8080/api/campaigns',
    json={'name': 'Summer Sale', 'description': 'Q3 Campaign'}
)
```

## 🎯 Testing the API

1. **Start MongoDB**:
```bash
brew services start mongodb-community
```

2. **Configure environment**:
```bash
cp .env.example .env
# Edit .env with your Twilio credentials
```

3. **Run the server**:
```bash
go run main.go
```

4. **Open documentation**:
```
http://localhost:8080/docs
```

5. **Try the API**:
   - Click any endpoint
   - Click "Try It Out"
   - Fill in the parameters
   - Click "Send"

## 📚 Additional Resources

- **API Documentation**: `/docs` endpoint
- **OpenAPI Spec**: `/docs/swagger.yaml`
- **Code Documentation**: See `/docs/API_DOCUMENTATION.md`
- **MongoDB Setup**: See `/docs/MONGODB_SETUP.md`
- **Feature Checklist**: See `/docs/FEATURE_CHECKLIST.md`

## 🐛 Troubleshooting

### Documentation not loading?
1. Ensure server is running: `go run main.go`
2. Check `docs/swagger.yaml` exists
3. Verify port 8080 is not blocked

### Swagger YAML not found?
Make sure you're running the server from the project root directory.

### Want to customize?
Edit `handlers/docs_handler.go` to change theme, colors, or behavior.

---

**Happy API Testing! 🚀**
