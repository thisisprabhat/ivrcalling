# Dynamic IVR Implementation Guide

## Overview

The IVR system has been enhanced to support dynamic, campaign-based interactive voice response flows. Campaigns can now define custom actions that control what happens when users press different keys during a call.

## Features Implemented

### 1. Campaign Configuration

Campaigns now support the following new fields:

- **Intro Text** (Required): Custom message played at the start of each call
- **IVR Actions**: Dynamic menu options configured per campaign
  - Information actions: Play text-to-speech or audio files
  - Forward actions: Transfer calls to specific phone numbers

### 2. Action Types

#### Information Action

Plays a message to the caller using either:

- **Text-to-Speech**: Enter any text that will be converted to speech
- **Audio File**: Provide a URL to a pre-recorded audio file (must start with http:// or https://)

**Fields:**

- `action_type`: "information"
- `action_input`: Key press (1-9)
- `message`: Text or URL to audio file

**Example:**

```json
{
  "action_type": "information",
  "action_input": "1",
  "message": "Thank you for your interest in our products. We offer a wide range of solutions..."
}
```

#### Forward Action

Forwards the call to another phone number.

**Fields:**

- `action_type`: "forward"
- `action_input`: Key press (1-9)
- `forward_phone`: Phone number to forward to (E.164 format recommended)
- `message`: Optional message to play before forwarding

**Example:**

```json
{
  "action_type": "forward",
  "action_input": "2",
  "forward_phone": "+1234567890",
  "message": "Connecting you to our sales team"
}
```

### 3. Backend Changes

#### Models (`ivr_api/models/models.go`)

- Added `IVRAction` struct for action configuration
- Extended `Campaign` model with:
  - `IntroText`: Intro message played at call start
  - `Actions`: Array of IVR actions

#### Campaign Handler (`ivr_api/handlers/campaign_handler.go`)

- Added validation for required fields (name, description, intro_text)
- Initializes empty actions array if not provided

#### TwiML Service (`ivr_api/services/twiml_service.go`)

New methods:

- `GenerateDynamicWelcome()`: Creates welcome TwiML with campaign intro and dynamic menu
- `GenerateDynamicResponse()`: Generates TwiML based on action configuration
- `GeneratePlayAudio()`: Plays audio file from URL
- `GenerateTextToSpeech()`: Converts text to speech
- `GenerateForward()`: Forwards call to specified number
- `buildMenuFromActions()`: Builds IVR menu from campaign actions

#### Webhook Handler (`ivr_api/handlers/webhook_handler.go`)

- `HandleVoiceWebhook()`: Detects if campaign has dynamic actions and uses appropriate flow
- `HandleGatherWebhook()`: Processes user input based on campaign actions
- Maintains backward compatibility with legacy static menu

### 4. Frontend Changes

#### CampaignForm Component (`ivr_frontend/src/components/CampaignForm.jsx`)

Enhanced UI with:

- **Intro Text field**: Required textarea for call introduction
- **IVR Actions section**:
  - Add/remove actions dynamically
  - Configure action type (Information/Forward)
  - Set key press for each action
  - Input message or phone number based on action type
- **Validation**: Ensures all required fields are filled
- **Better UX**: Scrollable modal for longer forms

## Usage Guide

### Creating a Dynamic IVR Campaign

1. **Navigate to Campaigns Page**

   - Click "Create Campaign" button

2. **Fill Basic Information**

   - **Campaign Name** (required): Give your campaign a descriptive name
   - **Description** (required): Describe the campaign purpose
   - **Default Language** (required): Select the language for IVR prompts
   - **Intro Text** (required): Write the message that plays when call starts
   - **Campaign is active**: Check to make campaign active

3. **Add IVR Actions**

   - Click "Add Action" button
   - For each action:
     - Select **Action Type**: Information or Forward
     - Enter **Key Press**: Single digit (1-9) users will press
     - For Information: Enter message text or audio URL
     - For Forward: Enter phone number to forward to

4. **Save Campaign**
   - Click "Create" or "Update" button
   - Campaign is now ready for calls

### Example Campaign Configuration

**Scenario**: Product inquiry campaign

**Campaign Details:**

- Name: "Product Inquiry 2025"
- Description: "Handle product inquiries and route to sales"
- Language: English
- Intro Text: "Welcome to ABC Company. We're glad you called."

**Actions:**

1. Action 1 (Information):

   - Key: 1
   - Message: "Our flagship product is the XYZ solution which offers advanced features including real-time analytics, cloud integration, and 24/7 support."

2. Action 2 (Information with Audio):

   - Key: 2
   - Message: "https://example.com/audio/special-offer.mp3"

3. Action 3 (Forward to Sales):

   - Key: 3
   - Phone: "+11234567890"
   - Message: "Connecting you to our sales team. Please hold."

4. Action 0 (Built-in):
   - Automatically repeats the menu

## IVR Call Flow

1. **Call Initiated**: System retrieves campaign configuration
2. **Welcome**: Plays greeting + intro text + menu built from actions
3. **User Input**: Caller presses a key
4. **Action Execution**:
   - Information: Plays message/audio → Returns to menu
   - Forward: Plays optional message → Transfers call
   - Invalid: Plays error message → Returns to menu
5. **Logging**: All interactions logged to call_logs collection

## API Examples

### Create Campaign with Actions

```bash
POST /api/campaigns
Content-Type: application/json

{
  "name": "Summer Sale 2025",
  "description": "Promotional campaign for summer",
  "language": "en",
  "intro_text": "Welcome to our exclusive summer sale event!",
  "is_active": true,
  "actions": [
    {
      "action_type": "information",
      "action_input": "1",
      "message": "Get 50% off on all summer collections. Visit our website for more details."
    },
    {
      "action_type": "forward",
      "action_input": "2",
      "forward_phone": "+11234567890",
      "message": "Let me connect you to our customer service team"
    }
  ]
}
```

### Update Campaign

```bash
PUT /api/campaigns/{campaign_id}
Content-Type: application/json

{
  "intro_text": "Updated intro message",
  "actions": [
    {
      "action_type": "information",
      "action_input": "1",
      "message": "New product information"
    }
  ]
}
```

## Database Schema

### Campaign Document

```javascript
{
  "_id": ObjectId("..."),
  "name": "Campaign Name",
  "description": "Campaign description",
  "language": "en",
  "intro_text": "Welcome message",
  "actions": [
    {
      "action_type": "information|forward",
      "action_input": "1",
      "message": "Text or URL",
      "forward_phone": "+1234567890"
    }
  ],
  "is_active": true,
  "created_at": ISODate("..."),
  "updated_at": ISODate("...")
}
```

## Backward Compatibility

The system maintains full backward compatibility:

- Campaigns without actions use the legacy static IVR menu
- Existing campaigns continue to work without modification
- Legacy TwiML generation methods remain available

## Best Practices

1. **Keep intro text concise**: 2-3 sentences maximum
2. **Limit actions**: 3-5 actions for best user experience
3. **Use clear descriptions**: When building menus, first words of message are used
4. **Test audio URLs**: Ensure audio files are accessible and in supported format
5. **Phone number format**: Use E.164 format (+[country][number])
6. **Action keys**: Use sequential numbers (1, 2, 3...) for intuitive navigation
7. **Always provide fallback**: Key 0 automatically repeats menu

## Troubleshooting

### Issue: Actions not working

- Verify campaign has `actions` array with at least one action
- Check `action_input` is a single digit (1-9)
- Ensure campaign is active

### Issue: Audio not playing

- Confirm URL starts with http:// or https://
- Verify audio file is publicly accessible
- Check audio format is supported by Twilio (MP3, WAV)

### Issue: Forward not working

- Validate phone number format
- Ensure number is E.164 formatted
- Check Twilio account permissions

## Testing

1. Create a test campaign with both action types
2. Initiate a call through the system
3. Test each key press to verify:
   - Correct message plays
   - Call forwards properly
   - Menu repeats on key 0
   - Invalid input handled gracefully

## Future Enhancements

Potential improvements:

- Multi-level menus (sub-menus)
- Time-based routing
- Conditional actions based on caller data
- Voice recording actions
- SMS follow-up actions
- Analytics dashboard for action usage
