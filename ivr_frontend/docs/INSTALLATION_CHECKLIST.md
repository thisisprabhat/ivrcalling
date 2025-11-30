# ✅ Installation & Verification Checklist

Use this checklist to verify your IVR Calling System is set up correctly.

## 📦 Prerequisites Check

- [ ] **Go 1.23+** installed
  ```bash
  go version
  # Should show: go version go1.23.x or higher
  ```

- [ ] **Node.js 18+** installed
  ```bash
  node --version
  # Should show: v18.x.x or higher
  ```

- [ ] **MongoDB** running
  ```bash
  # Local MongoDB
  mongosh
  # Or check MongoDB Atlas connection
  ```

- [ ] **Twilio Account** ready
  - [ ] Account SID available
  - [ ] Auth Token available
  - [ ] Phone number purchased

---

## 🔧 Backend Setup

### 1. Navigate to Backend
```bash
cd ivr_api
```

### 2. Verify Dependencies
- [ ] Check `go.mod` exists
- [ ] Run:
  ```bash
  go mod download
  ```

### 3. Configure Environment
- [ ] Copy `.env.example` to `.env`
  ```bash
  cp .env.example .env
  ```

- [ ] Edit `.env` with your credentials:
  - [ ] `MONGODB_URI` - Your MongoDB connection string
  - [ ] `MONGODB_DATABASE` - Database name (e.g., `ivr_calling_db`)
  - [ ] `TWILIO_ACCOUNT_SID` - Your Twilio Account SID
  - [ ] `TWILIO_AUTH_TOKEN` - Your Twilio Auth Token
  - [ ] `TWILIO_PHONE_NUMBER` - Your Twilio phone number (E.164 format)
  - [ ] `PORT` - API port (default: 8080)
  - [ ] `BASE_URL` - API base URL (default: http://localhost:8080)

### 4. Verify CORS Configuration
- [ ] Check `routes/routes.go` has CORS middleware
- [ ] Verify allowed origins include `http://localhost:3000`

### 5. Test Backend
- [ ] Start the server:
  ```bash
  go run main.go
  ```

- [ ] Should see:
  ```
  Starting IVR Calling System on port 8080
  ```

- [ ] Test health endpoint:
  ```bash
  curl http://localhost:8080/api/health
  # Should return: {"service":"IVR Calling System","status":"healthy"}
  ```

- [ ] Test languages endpoint:
  ```bash
  curl http://localhost:8080/api/languages
  # Should return: {"languages":["en","es","fr","de","hi"]}
  ```

---

## 🎨 Frontend Setup

### 1. Navigate to Frontend
```bash
cd ../ivr_frontend
```

### 2. Install Dependencies
- [ ] Run:
  ```bash
  npm install
  ```

- [ ] Verify `node_modules/` created
- [ ] Verify `package-lock.json` created

### 3. Verify Configuration Files
- [ ] `package.json` exists
- [ ] `vite.config.js` exists
- [ ] `tailwind.config.js` exists
- [ ] `postcss.config.js` exists
- [ ] `.env` exists with `VITE_API_URL=http://localhost:8080/api`

### 4. Verify Source Files
- [ ] `src/App.jsx` exists
- [ ] `src/main.jsx` exists
- [ ] `src/index.css` exists
- [ ] `src/components/` directory exists with 4 files
- [ ] `src/pages/` directory exists with 3 files
- [ ] `src/services/api.js` exists
- [ ] `src/utils/helpers.js` exists

### 5. Test Frontend
- [ ] Start dev server:
  ```bash
  npm run dev
  ```

- [ ] Should see:
  ```
  VITE vX.X.X  ready in XXX ms
  
  ➜  Local:   http://localhost:3000/
  ```

- [ ] Open browser to `http://localhost:3000`
- [ ] Should see the Dashboard page

---

## 🚀 Quick Start Verification

### Option 1: Use Start Script

- [ ] Make script executable:
  ```bash
  chmod +x start.sh
  ```

- [ ] Run:
  ```bash
  ./start.sh
  ```

- [ ] Should see both services starting
- [ ] Backend on port 8080
- [ ] Frontend on port 3000

### Option 2: Manual Start

**Terminal 1 - Backend:**
```bash
cd ivr_api
go run main.go
```

**Terminal 2 - Frontend:**
```bash
cd ivr_frontend
npm run dev
```

---

## 🧪 Functionality Tests

### Test 1: Dashboard Access
- [ ] Open `http://localhost:3000`
- [ ] Dashboard loads without errors
- [ ] See "Dashboard" heading
- [ ] See 6 statistic cards
- [ ] No console errors

### Test 2: Create Campaign
- [ ] Click "Campaigns" in sidebar
- [ ] Click "Create Campaign" button
- [ ] Modal opens
- [ ] Fill in:
  - Name: "Test Campaign"
  - Description: "Testing"
  - Language: English
  - Active: ✓
- [ ] Click "Create"
- [ ] Campaign appears in grid
- [ ] No errors

### Test 3: View Campaign Calls
- [ ] Click on created campaign card
- [ ] Calls page loads
- [ ] See campaign name in header
- [ ] See statistics cards (all showing 0)
- [ ] See empty calls table
- [ ] Back button works

### Test 4: Initiate Calls (Dry Run)
- [ ] Click "Initiate Calls" button
- [ ] Modal opens
- [ ] See CSV upload area
- [ ] See contacts list
- [ ] Add a test contact:
  - Phone: +1234567890
  - Name: Test User
- [ ] See validation (don't submit unless you want to test real calls)
- [ ] Can close modal

### Test 5: API Integration
- [ ] Open browser DevTools (F12)
- [ ] Go to Network tab
- [ ] Refresh Dashboard
- [ ] See API calls:
  - `GET /api/campaigns`
  - Status: 200
  - Response: JSON array
- [ ] No CORS errors
- [ ] No 404 errors

---

## 📁 File Structure Verification

### Root Directory
```
ivrcalling/
├── ivr_api/
├── ivr_frontend/
├── start.sh (executable)
├── README.md
└── PROJECT_SUMMARY.md
```

### Backend (ivr_api/)
```
ivr_api/
├── main.go
├── go.mod
├── go.sum
├── .env
├── .env.example
├── config/
│   └── config.go
├── database/
│   └── database.go
├── handlers/
│   ├── call_handler.go
│   ├── campaign_handler.go
│   ├── docs_handler.go
│   └── webhook_handler.go
├── models/
│   └── models.go
├── routes/
│   └── routes.go (with CORS)
├── services/
│   ├── language_service.go
│   ├── twilio_service.go
│   └── twiml_service.go
└── docs/
    └── (documentation files)
```

### Frontend (ivr_frontend/)
```
ivr_frontend/
├── src/
│   ├── components/
│   │   ├── Sidebar.jsx
│   │   ├── CampaignForm.jsx
│   │   ├── BulkCallForm.jsx
│   │   └── CallDetailsModal.jsx
│   ├── pages/
│   │   ├── DashboardPage.jsx
│   │   ├── CampaignsPage.jsx
│   │   └── CampaignCallsPage.jsx
│   ├── services/
│   │   └── api.js
│   ├── utils/
│   │   └── helpers.js
│   ├── App.jsx
│   ├── main.jsx
│   └── index.css
├── index.html
├── package.json
├── vite.config.js
├── tailwind.config.js
├── postcss.config.js
├── .env
├── README.md
├── QUICKSTART.md
├── FEATURES.md
└── UI_GUIDE.md
```

---

## 🔍 Troubleshooting

### Backend Issues

**Port 8080 already in use**
- [ ] Change `PORT` in `.env`
- [ ] Update frontend `.env` with new port

**MongoDB connection failed**
- [ ] Verify MongoDB is running
- [ ] Check connection string format
- [ ] Test connection with mongosh

**Twilio errors**
- [ ] Verify Account SID and Auth Token
- [ ] Check phone number format (+1234567890)
- [ ] Ensure Twilio account has credits

**CORS errors**
- [ ] Verify CORS middleware in `routes/routes.go`
- [ ] Check allowed origins includes frontend URL
- [ ] Restart backend after changes

### Frontend Issues

**npm install fails**
- [ ] Delete `node_modules/` and `package-lock.json`
- [ ] Run `npm install` again
- [ ] Check Node version (needs 18+)

**Page shows blank**
- [ ] Check browser console for errors
- [ ] Verify API is running
- [ ] Check `.env` has correct API URL

**API calls fail**
- [ ] Verify backend is running on correct port
- [ ] Check Network tab in DevTools
- [ ] Look for CORS or 404 errors

**Styles not loading**
- [ ] Verify Tailwind is configured
- [ ] Check `index.css` imports Tailwind
- [ ] Restart dev server

---

## ✨ Success Criteria

Your installation is successful if:

- [ ] Backend starts without errors
- [ ] Frontend loads in browser
- [ ] Dashboard shows UI correctly
- [ ] Can create a campaign
- [ ] Can navigate between pages
- [ ] API calls work (check Network tab)
- [ ] No console errors
- [ ] CORS is working
- [ ] All routes accessible

---

## 📞 Test with Real Calls (Optional)

⚠️ **Warning**: This will use Twilio credits

- [ ] Ensure Twilio account has credits
- [ ] Create a campaign
- [ ] Add your own phone number (for testing)
- [ ] Initiate call
- [ ] Answer phone and test IVR menu
- [ ] Check call appears in dashboard
- [ ] View call details and logs

---

## 🎉 You're Ready!

If all checks pass, you have a fully functional IVR Calling System!

### Next Steps:
1. **Read Documentation**: Check all README and guide files
2. **Explore Features**: Try all UI features
3. **Customize**: Modify for your needs
4. **Deploy**: When ready, deploy to production

### Quick Links:
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080/api
- API Docs: http://localhost:8080/docs
- Health Check: http://localhost:8080/api/health

---

**Happy Calling! 📞🎉**
