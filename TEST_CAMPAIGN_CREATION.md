# Complete Test Guide - Campaign Creation & Call Flow

## Step-by-Step Testing Process

### 1. Start Backend (WITH LOGS)
```powershell
cd d:\Projects\AIPROJECTS\ivrcalling\ivr_api
.\ivr_api.exe
```

**Watch for:**
- Server starts on port 8080
- Database connection successful

---

### 2. Start Frontend
```powershell
cd d:\Projects\AIPROJECTS\ivrcalling\ivr_frontend
npm run dev
```

**Expected:** Server runs on http://localhost:5173

---

### 3. Create a New Campaign

**IMPORTANT:** Use these EXACT values to test:

**Basic Information:**
- **Name:** `Dynamic IVR Test`
- **Description:** `Testing the dynamic IVR system`
- **Language:** `English`
- **Intro Text:** `Thank you for calling. We have special offers for you today.`
- **Active:** ✓ (checked)

**IVR Actions:**

**Action 1 (Information):**
- Type: `Information`
- Key Press: `1`
- Message: `Get fifty percent off all items today`

**Action 2 (Forward):**
- Type: `Forward`
- Key Press: `2`
- Phone: `+1234567890`
- Message: (leave empty or add custom like "connecting to sales")

**Action 3 (Information - Optional):**
- Type: `Information`
- Key Press: `3`
- Message: `Visit our website for more details`

---

### 4. Check Backend Logs After Creation

**You MUST see these logs:**
```
=== CREATING CAMPAIGN ===
Name: Dynamic IVR Test
Description: Testing the dynamic IVR system
Language: en
IntroText: Thank you for calling. We have special offers for you today.
Actions count: 2 (or 3 if you added action 3)
  Action 1: Type=information, Input=1, Message=Get fifty percent off all items today, Phone=
  Action 2: Type=forward, Input=2, Message=, Phone=+1234567890
✓ Campaign created successfully with ID: [some-id]
```

**If you DON'T see these logs:**
- ❌ Backend is not receiving the request
- ❌ CORS issue
- ❌ API endpoint mismatch

---

### 5. Make a Test Call

**From the campaign list:**
1. Click the phone icon on your new campaign
2. Add a test contact:
   - Phone: Your test number (or Twilio test number)
   - Name: `Test User`
3. Click "Initiate Calls"

---

### 6. Check Backend Logs During Call

**You MUST see (in order):**

```
=== BULK CALL REQUEST ===
Campaign ID: [id]
Language: en
Contact count: 1

Campaign found: Dynamic IVR Test, Active: true, IntroText: Thank you for calling..., Actions: 2

Call record created - ID: [call-id], Phone: [phone], Name: Test User

=== INITIATING TWILIO CALL ===
To: [phone]
From: [your-twilio-number]
Voice URL: http://[your-webhook]/api/webhook/voice?call_id=[call-id]&language=en
✓ Twilio call created - SID: [twilio-sid]
```

**Then when Twilio calls your webhook:**

```
=== VOICE WEBHOOK CALLED ===
Call ID: [call-id], Language: en
Found call record - Customer: Test User, Campaign ID: [campaign-id]
Found campaign - Name: Dynamic IVR Test, IntroText: 'Thank you for calling...', Actions count: 2
✓ USING DYNAMIC IVR - Campaign: Dynamic IVR Test
  Action 1: Type=information, Input=1, Message=Get fifty percent off all items today, Phone=
  Action 2: Type=forward, Input=2, Message=, Phone=+1234567890

★★★ USING DYNAMIC IVR FLOW ★★★

=== GENERATING DYNAMIC WELCOME TwiML ===
Campaign Name: Dynamic IVR Test
Greeting: Hello Test User, welcome to our marketing campaign
Intro Text: Thank you for calling. We have special offers for you today.

=== BUILDING MENU FROM 2 ACTIONS ===
Action 1: Type=information, Input=1, Message='Get fifty percent off all items today'
  → Info action with message: Press 1 for Get fifty percent off all items today
Action 2: Type=forward, Input=2, Message='', Phone=+1234567890
  → Forward with default message: Press 2 to speak with an agent
=== FINAL MENU TEXT: Press 1 for Get fifty percent off all items today. Press 2 to speak with an agent. Press 0 to repeat this menu ===

Sending TwiML response (length: [number] bytes)
```

---

### 7. What You Should Hear on the Call

**Exactly this sequence:**
1. `"Hello Test User, welcome to our marketing campaign"`
2. `"Thank you for calling. We have special offers for you today."`
3. `"Press 1 for Get fifty percent off all items today. Press 2 to speak with an agent. Press 0 to repeat this menu."`

---

## Common Issues and Solutions

### Issue 1: "press any key to execute your code"

**Root Causes:**
1. ❌ Using an OLD campaign (created before fixes)
   - **Solution:** DELETE old campaigns, create NEW one
   
2. ❌ Campaign has NO actions or EMPTY intro_text
   - **Solution:** Check campaign in database, verify fields
   
3. ❌ Dynamic IVR not triggered (see logs showing `LEGACY STATIC IVR`)
   - **Solution:** Check if intro_text and actions are properly saved

### Issue 2: No logs appearing

**Check:**
- Backend is running (not crashed)
- Correct port (8080)
- No errors in terminal

### Issue 3: Campaign creation fails

**Check frontend console (F12):**
- Network tab for API errors
- Console for JavaScript errors

**Check backend logs:**
- Validation error messages
- Database connection issues

---

## Debugging Checklist

Before reporting an issue, verify:

- [ ] Backend is running and shows startup logs
- [ ] Frontend is running on localhost:5173
- [ ] All old campaigns are deleted
- [ ] New campaign created with EXACT values above
- [ ] Backend shows campaign creation logs
- [ ] Campaign appears in frontend list
- [ ] Campaign is marked as ACTIVE (green toggle)
- [ ] Test call initiated successfully
- [ ] Backend shows webhook call logs
- [ ] Logs show "★★★ USING DYNAMIC IVR FLOW ★★★"

---

## What to Share If Still Having Issues

1. **Backend logs from campaign creation**
2. **Backend logs from making the call**
3. **Frontend console errors (if any)**
4. **Screenshot of the campaign form** (showing all fields filled)
5. **What you actually hear** on the phone call

This will help identify the exact issue!
