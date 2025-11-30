# API Reference Documentation

## Overview

The IVR Calling System API provides RESTful endpoints for managing marketing campaigns and initiating IVR calls using Twilio.

**Base URL:** `http://localhost:8080/api`

**Content-Type:** `application/json`

---

## Authentication

> **Note:** This version does not include authentication. For production, implement JWT or API key authentication.

---

## Core Concepts

### Campaign
A campaign represents a marketing initiative with associated calls. Each campaign has:
- Unique identifier
- Name and description
- Default language
- Active/inactive status

### Call
A call represents an individual phone interaction with:
- Link to parent campaign
- Contact information (phone, name)
- Status tracking
- Language preference
- Call logs and duration

### Call Statuses
- `pending` - Call created but not initiated
- `initiated` - Call sent to Twilio
- `in-progress` - Call is active
- `completed` - Call finished successfully
- `failed` - Call failed, busy, or no answer

---

## Endpoints

## System Endpoints

### Health Check

Check if the API is running.

```http
GET /api/health
```

#### Response

```json
{
  "status": "healthy",
  "service": "IVR Calling System"
}
```

---

### Get Supported Languages

Retrieve list of supported languages.

```http
GET /api/languages
```

#### Response

```json
{
  "languages": ["en", "es", "fr", "de", "hi"]
}
```

#### Language Codes

| Code | Language | Region |
|------|----------|--------|
| en   | English  | US     |
| es   | Spanish  | Spain  |
| fr   | French   | France |
| de   | German   | Germany|
| hi   | Hindi    | India  |

---

## Campaign Management

### Create Campaign

Create a new marketing campaign.

```http
POST /api/campaigns
```

#### Request Body

```json
{
  "name": "Summer Sale 2025",
  "description": "Promotional campaign for summer products",
  "language": "en",
  "is_active": true
}
```

#### Parameters

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| name | string | Yes | Campaign name |
| description | string | No | Campaign description |
| language | string | No | Default language (en, es, fr, de, hi). Default: "en" |
| is_active | boolean | No | Active status. Default: true |

#### Response (201 Created)

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

---

### List All Campaigns

Retrieve all campaigns.

```http
GET /api/campaigns
```

#### Response (200 OK)

```json
[
  {
    "id": 1,
    "name": "Summer Sale 2025",
    "description": "Promotional campaign for summer products",
    "language": "en",
    "is_active": true,
    "created_at": "2025-11-30T10:00:00Z",
    "updated_at": "2025-11-30T10:00:00Z"
  },
  {
    "id": 2,
    "name": "Black Friday Deals",
    "description": "Annual Black Friday promotion",
    "language": "es",
    "is_active": false,
    "created_at": "2025-11-15T08:00:00Z",
    "updated_at": "2025-11-20T09:00:00Z"
  }
]
```

---

### Get Campaign Details

Retrieve a specific campaign by ID.

```http
GET /api/campaigns/{id}
```

#### Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | integer | Campaign ID |

#### Response (200 OK)

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

#### Error Response (404 Not Found)

```json
{
  "error": "Campaign not found"
}
```

---

### Update Campaign

Update an existing campaign.

```http
PUT /api/campaigns/{id}
```

#### Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | integer | Campaign ID |

#### Request Body

```json
{
  "name": "Updated Summer Sale",
  "is_active": false
}
```

> **Note:** Only include fields you want to update.

#### Response (200 OK)

```json
{
  "id": 1,
  "name": "Updated Summer Sale",
  "description": "Promotional campaign for summer products",
  "language": "en",
  "is_active": false,
  "created_at": "2025-11-30T10:00:00Z",
  "updated_at": "2025-11-30T12:00:00Z"
}
```

---

### Delete Campaign

Delete a campaign.

```http
DELETE /api/campaigns/{id}
```

#### Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | integer | Campaign ID |

#### Response (200 OK)

```json
{
  "message": "Campaign deleted successfully"
}
```

---

## Call Management

### Initiate Bulk Calls

**Main Endpoint** - Initiate calls to multiple contacts.

```http
POST /api/calls/bulk
```

#### Request Body

```json
{
  "campaign_id": 1,
  "language": "en",
  "contacts": [
    {
      "phone_number": "+1234567890",
      "name": "John Doe"
    },
    {
      "phone_number": "+44123456789",
      "name": "Jane Smith"
    },
    {
      "phone_number": "+91987654321",
      "name": "Raj Kumar"
    }
  ]
}
```

#### Parameters

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| campaign_id | integer | Yes | ID of the campaign |
| language | string | No | Language code (overrides campaign default) |
| contacts | array | Yes | Array of contact objects (min: 1) |
| contacts[].phone_number | string | Yes | Phone number in E.164 format (+country code + number) |
| contacts[].name | string | No | Customer name for personalization |

#### Phone Number Format

Use E.164 format: `+[country code][number]`

Examples:
- US: `+11234567890`
- UK: `+441234567890`
- India: `+919876543210`
- Spain: `+34612345678`

#### Response (200 OK)

```json
{
  "message": "Bulk calls initiated",
  "success_count": 3,
  "fail_count": 0,
  "call_ids": [1, 2, 3]
}
```

#### Response Fields

| Field | Type | Description |
|-------|------|-------------|
| message | string | Status message |
| success_count | integer | Number of successfully initiated calls |
| fail_count | integer | Number of failed call attempts |
| call_ids | array | IDs of created call records |

#### Error Responses

**400 Bad Request** - Invalid request body
```json
{
  "error": "Key: 'BulkCallRequest.Contacts' Error:Field validation for 'Contacts' failed on the 'required' tag"
}
```

**404 Not Found** - Campaign not found
```json
{
  "error": "Campaign not found"
}
```

**400 Bad Request** - Inactive campaign
```json
{
  "error": "Campaign is not active"
}
```

---

### Get Call Status

Retrieve detailed information about a specific call.

```http
GET /api/calls/{id}
```

#### Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | integer | Call ID |

#### Response (200 OK)

```json
{
  "id": 1,
  "campaign_id": 1,
  "phone_number": "+1234567890",
  "customer_name": "John Doe",
  "status": "completed",
  "twilio_call_sid": "CA1234567890abcdef1234567890abcdef",
  "language": "en",
  "duration": 45,
  "error_message": "",
  "created_at": "2025-11-30T10:00:00Z",
  "updated_at": "2025-11-30T10:01:00Z",
  "call_logs": [
    {
      "id": 1,
      "call_id": 1,
      "event": "initiated",
      "details": "Call initiated to +1234567890",
      "user_input": "",
      "created_at": "2025-11-30T10:00:00Z"
    },
    {
      "id": 2,
      "call_id": 1,
      "event": "answered",
      "details": "Call status: answered",
      "user_input": "",
      "created_at": "2025-11-30T10:00:15Z"
    },
    {
      "id": 3,
      "call_id": 1,
      "event": "input_received",
      "details": "User pressed: 1",
      "user_input": "1",
      "created_at": "2025-11-30T10:00:30Z"
    },
    {
      "id": 4,
      "call_id": 1,
      "event": "completed",
      "details": "Call status: completed",
      "user_input": "",
      "created_at": "2025-11-30T10:00:45Z"
    }
  ]
}
```

#### Call Events

| Event | Description |
|-------|-------------|
| initiated | Call sent to Twilio |
| ringing | Phone is ringing |
| answered | Call was answered |
| input_received | User pressed a digit |
| product_info_requested | User requested product info (pressed 1) |
| offer_requested | User requested offers (pressed 2) |
| opt_out_requested | User requested opt-out (pressed 3) |
| opted_out | User confirmed opt-out |
| completed | Call finished |
| failed | Call failed |

---

### Get Campaign Calls

Retrieve all calls for a specific campaign with statistics.

```http
GET /api/campaigns/{id}/calls
```

#### Path Parameters

| Parameter | Type | Description |
|-----------|------|-------------|
| id | integer | Campaign ID |

#### Response (200 OK)

```json
{
  "calls": [
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
      "updated_at": "2025-11-30T10:01:00Z"
    },
    {
      "id": 2,
      "campaign_id": 1,
      "phone_number": "+44123456789",
      "customer_name": "Jane Smith",
      "status": "in-progress",
      "twilio_call_sid": "CA0987654321fedcba",
      "language": "en",
      "duration": 0,
      "created_at": "2025-11-30T10:02:00Z",
      "updated_at": "2025-11-30T10:02:30Z"
    }
  ],
  "stats": {
    "total": 10,
    "pending": 2,
    "initiated": 3,
    "completed": 4,
    "failed": 1
  }
}
```

---

## Webhook Endpoints

> **Note:** These endpoints are called by Twilio and should not be invoked directly.

### Voice Webhook

Initial webhook when call is answered.

```http
POST /api/webhook/voice?call_id={id}&language={lang}
```

Returns TwiML response with welcome message and main menu.

---

### Gather Webhook

Handles user input from IVR menu.

```http
POST /api/webhook/gather
```

Returns appropriate TwiML based on digit pressed.

---

### Status Webhook

Receives call status updates from Twilio.

```http
POST /api/webhook/status
```

Updates call status in database.

---

### Opt-Out Webhook

Handles opt-out confirmation.

```http
POST /api/webhook/optout
```

---

## Code Examples

### cURL Examples

**Create Campaign:**
```bash
curl -X POST http://localhost:8080/api/campaigns \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Holiday Sale",
    "language": "es",
    "is_active": true
  }'
```

**Initiate Calls:**
```bash
curl -X POST http://localhost:8080/api/calls/bulk \
  -H "Content-Type: application/json" \
  -d '{
    "campaign_id": 1,
    "contacts": [
      {"phone_number": "+1234567890", "name": "Alice"},
      {"phone_number": "+0987654321", "name": "Bob"}
    ]
  }'
```

**Get Call Status:**
```bash
curl http://localhost:8080/api/calls/1
```

---

### JavaScript/Node.js Example

```javascript
const axios = require('axios');

const API_BASE = 'http://localhost:8080/api';

// Create campaign
async function createCampaign() {
  const response = await axios.post(`${API_BASE}/campaigns`, {
    name: 'Spring Collection 2025',
    description: 'New spring products launch',
    language: 'en',
    is_active: true
  });
  return response.data;
}

// Initiate bulk calls
async function initiateCalls(campaignId, contacts) {
  const response = await axios.post(`${API_BASE}/calls/bulk`, {
    campaign_id: campaignId,
    language: 'en',
    contacts: contacts
  });
  return response.data;
}

// Usage
(async () => {
  const campaign = await createCampaign();
  console.log('Campaign created:', campaign);

  const result = await initiateCalls(campaign.id, [
    { phone_number: '+1234567890', name: 'Customer One' },
    { phone_number: '+0987654321', name: 'Customer Two' }
  ]);
  console.log('Calls initiated:', result);
})();
```

---

### Python Example

```python
import requests

API_BASE = 'http://localhost:8080/api'

# Create campaign
def create_campaign():
    response = requests.post(f'{API_BASE}/campaigns', json={
        'name': 'Autumn Promotion',
        'description': 'Fall season sale',
        'language': 'en',
        'is_active': True
    })
    return response.json()

# Initiate bulk calls
def initiate_calls(campaign_id, contacts):
    response = requests.post(f'{API_BASE}/calls/bulk', json={
        'campaign_id': campaign_id,
        'language': 'en',
        'contacts': contacts
    })
    return response.json()

# Usage
campaign = create_campaign()
print(f'Campaign created: {campaign}')

result = initiate_calls(campaign['id'], [
    {'phone_number': '+1234567890', 'name': 'Alice Johnson'},
    {'phone_number': '+0987654321', 'name': 'Bob Williams'}
])
print(f'Calls initiated: {result}')
```

---

## Error Codes

| HTTP Code | Description |
|-----------|-------------|
| 200 | OK - Request successful |
| 201 | Created - Resource created |
| 400 | Bad Request - Invalid input |
| 404 | Not Found - Resource not found |
| 500 | Internal Server Error - Server error |

---

## Rate Limits

> **Note:** This version does not implement rate limiting. For production, consider implementing:
- Per-IP rate limits
- Per-campaign call limits
- Daily/hourly quotas

---

## Best Practices

1. **Phone Numbers**: Always use E.164 format
2. **Bulk Calls**: Limit to 100 contacts per request
3. **Error Handling**: Check `fail_count` in bulk call responses
4. **Polling**: Use webhooks instead of polling for call status
5. **Language**: Specify language at campaign level for consistency
6. **Testing**: Use Twilio test credentials during development

---

## Support

For API issues or questions:
- Check this documentation
- Review the main README.md
- Visit Twilio documentation: https://www.twilio.com/docs

---

**Last Updated:** November 30, 2025
