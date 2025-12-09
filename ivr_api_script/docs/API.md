# Q&I IVR Calling API - API Documentation

## Base URL

```
http://localhost:8080
```

For production, replace with your deployed server URL.

## Authentication

Currently, the API does not require authentication. For production use, implement API key or JWT-based authentication.

---

## Endpoints

### 1. Health Check

**Endpoint:** `GET /health`

**Description:** Check if the API is running and healthy.

**Response:**

```json
{
  "status": "healthy",
  "version": "1.0.0"
}
```

**Status Codes:**

- `200 OK`: Service is healthy

**Example:**

```bash
curl http://localhost:8080/health
```

---

### 2. Initiate Call

**Endpoint:** `POST /api/v1/calls/initiate`

**Description:** Initiates an outbound IVR call to a specified phone number.

**Request Headers:**

```
Content-Type: application/json
```

**Request Body:**

```json
{
  "phone_number": "+919876543210",
  "callback_url": "https://yourapp.com/callback"
}
```

**Parameters:**

| Field          | Type   | Required | Description                                        |
| -------------- | ------ | -------- | -------------------------------------------------- |
| `phone_number` | string | Yes      | Phone number in E.164 format (e.g., +919876543210) |
| `callback_url` | string | No       | URL to receive callbacks from IVR provider         |

**Response:**

```json
{
  "call_id": "call_1733739600000000000",
  "phone_number": "+919876543210",
  "status": "initiated",
  "message": "Call initiated successfully"
}
```

**Status Codes:**

- `200 OK`: Call initiated successfully
- `400 Bad Request`: Invalid request parameters
- `500 Internal Server Error`: Server error

**Example:**

```bash
curl -X POST http://localhost:8080/api/v1/calls/initiate \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "+919876543210",
    "callback_url": "https://yourapp.com/callback"
  }'
```

**Error Response:**

```json
{
  "error": "Invalid request",
  "message": "phone_number is required"
}
```

---

### 3. IVR Callback Handler

**Endpoint:** `POST /api/v1/callbacks/ivr`

**Description:** Receives callbacks from the IVR provider about call events and user interactions.

**Request Headers:**

```
Content-Type: application/json
```

**Request Body:**

```json
{
  "call_id": "call_1733739600000000000",
  "event": "digit_pressed",
  "digit_input": "1",
  "timestamp": "2025-12-09T10:30:00Z"
}
```

**Parameters:**

| Field         | Type   | Required | Description                                                            |
| ------------- | ------ | -------- | ---------------------------------------------------------------------- |
| `call_id`     | string | Yes      | Unique identifier for the call                                         |
| `event`       | string | Yes      | Event type (call_answered, digit_pressed, call_completed, call_failed) |
| `digit_input` | string | No       | Digit pressed by user (1, 2, 3, etc.)                                  |
| `timestamp`   | string | Yes      | ISO 8601 timestamp of the event                                        |

**Response:**

```json
{
  "status": "success",
  "message": "Callback processed successfully"
}
```

**Status Codes:**

- `200 OK`: Callback processed successfully
- `400 Bad Request`: Invalid callback data
- `500 Internal Server Error`: Processing error

**Example:**

```bash
curl -X POST http://localhost:8080/api/v1/callbacks/ivr \
  -H "Content-Type: application/json" \
  -d '{
    "call_id": "call_1733739600000000000",
    "event": "digit_pressed",
    "digit_input": "1",
    "timestamp": "2025-12-09T10:30:00Z"
  }'
```

---

### 4. Get IVR Configuration

**Endpoint:** `GET /api/v1/config/ivr`

**Description:** Returns the current IVR flow configuration including all messages and actions.

**Response:**

```json
{
  "intro_text": "Welcome to Q&I! We are transforming education with smart digital tools. We help your school digitize teaching and measure true student understanding. Our AI-powered platform provides topic analysis and targeted practice to boost academic performance. With Q&I, teachers get deeper insights, students learn effectively, and your institution achieves measurable growth. Ready to see how Q&I can revolutionize your classrooms?",
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
      "description": "Q&I is an AI-powered educational platform that helps schools digitize teaching and measure student understanding. It provides topic analysis and targeted practice to improve academic performance, giving teachers deeper insights and students more effective learning experiences."
    },
    {
      "key": "3",
      "message": "To hear this message again, press 3",
      "action": "repeat"
    }
  ],
  "end_message": "Thank you for contacting Q&I. We look forward to helping your school achieve success. Goodbye!"
}
```

**Status Codes:**

- `200 OK`: Configuration retrieved successfully

**Example:**

```bash
curl http://localhost:8080/api/v1/config/ivr
```

---

## IVR Flow Details

### Call Flow

When a call is initiated, the following sequence occurs:

1. **Call Initiated**: System dials the provided phone number
2. **Intro Message Played**: Welcome message about Q&I is played
3. **Menu Options Presented**: Three options are presented to the caller
4. **User Input Processing**: System waits for digit input (1, 2, or 3)
5. **Action Execution**: Based on input:
   - **Press 1**: Call is forwarded to Q&I team at +917905252436
   - **Press 2**: Detailed information about Q&I is played
   - **Press 3**: Intro message is repeated
6. **End Message**: Thank you message is played
7. **Call Terminated**: Call ends

### IVR Actions

| Key | Action  | Description                                   |
| --- | ------- | --------------------------------------------- |
| 1   | Forward | Forwards the call to Q&I team phone number    |
| 2   | Inform  | Plays detailed information about Q&I platform |
| 3   | Repeat  | Repeats the intro message                     |

---

## Code Examples

### JavaScript (Node.js with Axios)

```javascript
const axios = require("axios");

const API_BASE_URL = "http://localhost:8080";

// Initiate a call
async function initiateCall(phoneNumber) {
  try {
    const response = await axios.post(`${API_BASE_URL}/api/v1/calls/initiate`, {
      phone_number: phoneNumber,
      callback_url: "https://yourapp.com/callback",
    });

    console.log("Call initiated:", response.data);
    return response.data;
  } catch (error) {
    console.error("Error:", error.response?.data || error.message);
    throw error;
  }
}

// Get IVR configuration
async function getIVRConfig() {
  try {
    const response = await axios.get(`${API_BASE_URL}/api/v1/config/ivr`);
    console.log("IVR Config:", response.data);
    return response.data;
  } catch (error) {
    console.error("Error:", error.response?.data || error.message);
    throw error;
  }
}

// Usage
initiateCall("+919876543210");
getIVRConfig();
```

### Python (with Requests)

```python
import requests

API_BASE_URL = 'http://localhost:8080'

def initiate_call(phone_number, callback_url=None):
    """Initiate an IVR call"""
    url = f'{API_BASE_URL}/api/v1/calls/initiate'
    payload = {
        'phone_number': phone_number
    }

    if callback_url:
        payload['callback_url'] = callback_url

    response = requests.post(url, json=payload)

    if response.status_code == 200:
        print('Call initiated:', response.json())
        return response.json()
    else:
        print('Error:', response.json())
        return None

def get_ivr_config():
    """Get IVR configuration"""
    url = f'{API_BASE_URL}/api/v1/config/ivr'
    response = requests.get(url)

    if response.status_code == 200:
        print('IVR Config:', response.json())
        return response.json()
    else:
        print('Error:', response.status_code)
        return None

# Usage
initiate_call('+919876543210', 'https://yourapp.com/callback')
get_ivr_config()
```

### cURL Examples

**Initiate a call:**

```bash
curl -X POST http://localhost:8080/api/v1/calls/initiate \
  -H "Content-Type: application/json" \
  -d '{
    "phone_number": "+919876543210",
    "callback_url": "https://yourapp.com/callback"
  }'
```

**Get IVR configuration:**

```bash
curl http://localhost:8080/api/v1/config/ivr
```

**Health check:**

```bash
curl http://localhost:8080/health
```

**Send IVR callback:**

```bash
curl -X POST http://localhost:8080/api/v1/callbacks/ivr \
  -H "Content-Type: application/json" \
  -d '{
    "call_id": "call_123",
    "event": "digit_pressed",
    "digit_input": "1",
    "timestamp": "2025-12-09T10:30:00Z"
  }'
```

---

## Webhook Integration

### Setting Up Your Callback Endpoint

When initiating a call, you can provide a `callback_url`. The IVR provider will send HTTP POST requests to this URL with event updates.

**Example Express.js Webhook Handler:**

```javascript
const express = require("express");
const app = express();

app.use(express.json());

app.post("/callback", (req, res) => {
  const { call_id, event, digit_input, timestamp } = req.body;

  console.log(`Received callback:`);
  console.log(`  Call ID: ${call_id}`);
  console.log(`  Event: ${event}`);
  console.log(`  Digit: ${digit_input}`);
  console.log(`  Time: ${timestamp}`);

  // Process the callback based on event type
  switch (event) {
    case "call_answered":
      console.log("Call was answered");
      break;
    case "digit_pressed":
      console.log(`User pressed: ${digit_input}`);
      break;
    case "call_completed":
      console.log("Call completed successfully");
      break;
    case "call_failed":
      console.log("Call failed");
      break;
  }

  res.json({ status: "received" });
});

app.listen(3000, () => {
  console.log("Webhook server listening on port 3000");
});
```

---

## Error Handling

### Error Response Format

All errors follow this format:

```json
{
  "error": "Error type",
  "message": "Detailed error message"
}
```

### Common Errors

| Error                       | Status Code | Description                       | Solution                             |
| --------------------------- | ----------- | --------------------------------- | ------------------------------------ |
| Invalid phone number format | 400         | Phone number doesn't start with + | Use E.164 format: +[country][number] |
| Phone number is required    | 400         | Missing phone_number field        | Include phone_number in request      |
| Failed to initiate call     | 500         | Error calling IVR provider        | Check IVR provider credentials       |
| Invalid callback data       | 400         | Malformed callback payload        | Verify callback payload structure    |

---

## Rate Limits

Currently, there are no enforced rate limits. For production deployment, consider implementing:

- **Per IP**: 100 requests per minute
- **Per API Key**: 1000 requests per hour
- **Global**: 10,000 requests per hour

---

## Best Practices

1. **Phone Number Format**: Always use E.164 format (+[country code][number])
2. **Callback URLs**: Use HTTPS for callback URLs in production
3. **Error Handling**: Always check status codes and handle errors gracefully
4. **Idempotency**: Store call_id to avoid duplicate calls
5. **Logging**: Log all API interactions for debugging
6. **Timeouts**: Set appropriate timeout values for API calls

---

## Support

For API support:

- Email: support@qandi.com
- Documentation: `/docs` directory
- GitHub Issues: [Repository URL]

---

**API Version**: 1.0.0  
**Last Updated**: December 9, 2025
