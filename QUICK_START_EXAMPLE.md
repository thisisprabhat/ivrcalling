# Dynamic IVR - Quick Start Example

## Example: E-commerce Support Campaign

This example demonstrates how to set up a complete dynamic IVR campaign for an e-commerce business.

### Campaign Setup

**Step 1: Basic Information**

```
Campaign Name: E-commerce Customer Support 2025
Description: Handle customer inquiries for online orders and support
Default Language: English (en)
Intro Text: Thank you for calling ShopXYZ. We value your business and are here to help you today.
Campaign Active: ✓ Yes
```

**Step 2: Define IVR Actions**

#### Action 1: Order Status (Information - Text-to-Speech)

```
Action Type: Information
Key Press: 1
Message: To check your order status, please visit our website and log in to your account. You can track your order in real-time. For urgent inquiries, press 3 to speak with our team.
```

#### Action 2: Shipping Information (Information - Audio File)

```
Action Type: Information
Key Press: 2
Message: https://yourserver.com/audio/shipping-info.mp3
```

_Audio file should contain: "We offer free shipping on orders over $50. Standard delivery takes 3-5 business days. Express shipping is available for next-day delivery. For more details, visit our shipping page."_

#### Action 3: Speak to Agent (Forward)

```
Action Type: Forward
Key Press: 3
Forward to Phone: +11234567890
Message: Please hold while we connect you to one of our customer service representatives.
```

#### Action 4: Return Policy (Information - Text-to-Speech)

```
Action Type: Information
Key Press: 4
Message: We accept returns within 30 days of purchase. Items must be unused and in original packaging. For return instructions, visit our returns page or press 3 to speak with an agent.
```

#### Action 5: Store Hours (Information - Text-to-Speech)

```
Action Type: Information
Key Press: 5
Message: Our customer support is available Monday through Friday, 9 AM to 6 PM Eastern Time. Our online store is open 24/7. Thank you for calling ShopXYZ.
```

### Expected Call Flow

```
┌─────────────────────────────────────────────────┐
│  CALL INITIATED                                 │
└─────────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────────┐
│  GREETING                                       │
│  "Hello [Customer Name], Thank you for calling  │
│   ShopXYZ..."                                   │
└─────────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────────┐
│  INTRO TEXT (Campaign Specific)                 │
│  "Thank you for calling ShopXYZ. We value your  │
│   business and are here to help you today."     │
└─────────────────────────────────────────────────┘
                    ↓
┌─────────────────────────────────────────────────┐
│  DYNAMIC MENU (Auto-generated from actions)     │
│  "Press 1 for To check your order status..."    │
│  "Press 2 for We offer free shipping on..."     │
│  "Press 3 to speak with an agent."              │
│  "Press 4 for We accept returns within..."      │
│  "Press 5 for Our customer support is..."       │
│  "Press 0 to repeat this menu."                 │
└─────────────────────────────────────────────────┘
                    ↓
        ┌───────────┴───────────┐
        │  USER PRESSES KEY     │
        └───────────┬───────────┘
                    ↓
    ┌───────────────┼───────────────┐
    │               │               │
    ▼               ▼               ▼
┌────────┐    ┌─────────┐    ┌─────────┐
│ Press 1│    │ Press 2 │    │ Press 3 │
│ (Info) │    │ (Audio) │    │(Forward)│
└────┬───┘    └────┬────┘    └────┬────┘
     │             │              │
     ▼             ▼              ▼
 Play Text    Play Audio    Forward Call
     │             │              │
     ▼             ▼              ▼
Return to     Return to      Call Ends
  Menu          Menu
```

### API Request Example

```bash
POST http://localhost:8080/api/campaigns
Content-Type: application/json

{
  "name": "E-commerce Customer Support 2025",
  "description": "Handle customer inquiries for online orders and support",
  "language": "en",
  "intro_text": "Thank you for calling ShopXYZ. We value your business and are here to help you today.",
  "is_active": true,
  "actions": [
    {
      "action_type": "information",
      "action_input": "1",
      "message": "To check your order status, please visit our website and log in to your account. You can track your order in real-time. For urgent inquiries, press 3 to speak with our team."
    },
    {
      "action_type": "information",
      "action_input": "2",
      "message": "https://yourserver.com/audio/shipping-info.mp3"
    },
    {
      "action_type": "forward",
      "action_input": "3",
      "forward_phone": "+11234567890",
      "message": "Please hold while we connect you to one of our customer service representatives."
    },
    {
      "action_type": "information",
      "action_input": "4",
      "message": "We accept returns within 30 days of purchase. Items must be unused and in original packaging. For return instructions, visit our returns page or press 3 to speak with an agent."
    },
    {
      "action_type": "information",
      "action_input": "5",
      "message": "Our customer support is available Monday through Friday, 9 AM to 6 PM Eastern Time. Our online store is open 24/7. Thank you for calling ShopXYZ."
    }
  ]
}
```

### Response Example

```json
{
  "id": "65abc123def456789",
  "name": "E-commerce Customer Support 2025",
  "description": "Handle customer inquiries for online orders and support",
  "language": "en",
  "intro_text": "Thank you for calling ShopXYZ. We value your business and are here to help you today.",
  "actions": [
    {
      "action_type": "information",
      "action_input": "1",
      "message": "To check your order status, please visit our website and log in to your account. You can track your order in real-time. For urgent inquiries, press 3 to speak with our team."
    },
    {
      "action_type": "information",
      "action_input": "2",
      "message": "https://yourserver.com/audio/shipping-info.mp3"
    },
    {
      "action_type": "forward",
      "action_input": "3",
      "forward_phone": "+11234567890",
      "message": "Please hold while we connect you to one of our customer service representatives."
    },
    {
      "action_type": "information",
      "action_input": "4",
      "message": "We accept returns within 30 days of purchase. Items must be unused and in original packaging. For return instructions, visit our returns page or press 3 to speak with an agent."
    },
    {
      "action_type": "information",
      "action_input": "5",
      "message": "Our customer support is available Monday through Friday, 9 AM to 6 PM Eastern Time. Our online store is open 24/7. Thank you for calling ShopXYZ."
    }
  ],
  "is_active": true,
  "created_at": "2025-12-09T10:30:00Z",
  "updated_at": "2025-12-09T10:30:00Z"
}
```

### Using the Frontend

1. **Open the application** in your browser
2. **Navigate to Campaigns** page
3. **Click "Create Campaign"** button
4. **Fill Basic Information:**

   - Enter campaign name
   - Enter description
   - Select language
   - Enter intro text
   - Check "Campaign is active"

5. **Add IVR Actions:**

   - Click "Add Action" for each menu option
   - Select action type
   - Enter key press number
   - Fill in message or phone number
   - Repeat for all actions

6. **Save Campaign:**
   - Review all information
   - Click "Create" button

### Initiating Calls

Once the campaign is created, initiate calls using the Bulk Call form:

```bash
POST http://localhost:8080/api/calls/bulk
Content-Type: application/json

{
  "campaign_id": "65abc123def456789",
  "language": "en",
  "contacts": [
    {
      "phone_number": "+11234567890",
      "name": "John Doe"
    },
    {
      "phone_number": "+10987654321",
      "name": "Jane Smith"
    }
  ]
}
```

### Call Logs

Monitor call interactions in the MongoDB `call_logs` collection:

```javascript
// Example log entries
{
  "call_id": ObjectId("..."),
  "event": "input_received",
  "user_input": "1",
  "details": "User pressed: 1",
  "created_at": ISODate("...")
}

{
  "call_id": ObjectId("..."),
  "event": "action_information_executed",
  "details": "User pressed 1 - Action type: information",
  "created_at": ISODate("...")
}
```

### Tips for Success

1. **Keep messages concise**: Aim for 15-30 seconds per message
2. **Test audio files**: Ensure URLs are publicly accessible
3. **Use clear language**: Avoid jargon or complex terms
4. **Limit options**: 3-5 actions provide the best user experience
5. **Provide agent option**: Always include a way to reach a human
6. **Test thoroughly**: Call the system yourself before going live
7. **Monitor logs**: Review call logs regularly to improve flow

### Common Use Cases

#### Retail/E-commerce

- Order status
- Shipping information
- Returns/exchanges
- Store hours
- Speak to sales

#### Healthcare

- Appointment scheduling
- Insurance information
- Prescription refills
- Emergency line
- Speak to nurse

#### Financial Services

- Account balance
- Recent transactions
- Payment options
- Fraud reporting
- Speak to advisor

#### Real Estate

- Property availability
- Schedule viewing
- Financing options
- Application status
- Speak to agent

### Troubleshooting

**Problem**: Actions not appearing in menu

- **Solution**: Ensure actions array is populated and saved
- **Check**: Campaign has at least one action with action_input set

**Problem**: Audio not playing

- **Solution**: Verify URL is publicly accessible
- **Test**: Open URL in browser to confirm it works
- **Format**: Ensure audio is MP3 or WAV format

**Problem**: Forward not working

- **Solution**: Check phone number format (E.164 recommended)
- **Verify**: Twilio account has permission to call that number
- **Test**: Try with a known working number first

**Problem**: Menu repeating incorrectly

- **Solution**: Ensure key 0 is not used in your actions
- **Reserved**: Key 0 is system-reserved for menu repeat
