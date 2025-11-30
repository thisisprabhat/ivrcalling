# IVR Calling System

A complete IVR (Interactive Voice Response) calling system with a Go backend API and React frontend for managing marketing campaigns and automated calls using Twilio.

## 🚀 Features

### Backend (Go + Gin + MongoDB)
- **Campaign Management**: Create, update, delete, and list marketing campaigns
- **Bulk Call Initiation**: Send automated calls to multiple contacts
- **Multi-language Support**: English, Spanish, French, German, and Hindi
- **Twilio Integration**: Handle IVR flows with voice webhooks
- **Call Tracking**: Real-time status updates and detailed call logs
- **RESTful API**: Comprehensive API with proper error handling

### Frontend (React + Vite + TailwindCSS)
- **Interactive Dashboard**: Overview of campaigns and call statistics
- **Campaign Management**: Full CRUD operations with an intuitive UI
- **Bulk Call Interface**: Upload CSV or manually enter contacts
- **Real-time Monitoring**: Auto-refreshing call status and logs
- **Call Details**: View complete call timeline with user interactions
- **Responsive Design**: Works on desktop and mobile devices

## 📋 Prerequisites

- **Go** 1.23+ 
- **Node.js** 18+
- **MongoDB** 4.4+ (local or Atlas)
- **Twilio Account** (for making calls)

## 🛠️ Installation

### 1. Clone the Repository

```bash
git clone <repository-url>
cd ivrcalling
```

### 2. Backend Setup

```bash
cd ivr_api

# Install Go dependencies
go mod download

# Copy environment template
cp .env.example .env

# Edit .env with your credentials
nano .env
```

**Required Environment Variables:**

```env
# MongoDB
MONGODB_URI=mongodb://localhost:27017
MONGODB_DATABASE=ivr_calling_db

# Twilio
TWILIO_ACCOUNT_SID=your_account_sid
TWILIO_AUTH_TOKEN=your_auth_token
TWILIO_PHONE_NUMBER=+1234567890

# Server
PORT=8080
ENVIRONMENT=development
BASE_URL=http://localhost:8080
```

### 3. Frontend Setup

```bash
cd ../ivr_frontend

# Install dependencies
npm install
```

## 🚀 Quick Start

### Option 1: Using the Start Script (Recommended)

```bash
# Make the script executable
chmod +x start.sh

# Start both backend and frontend
./start.sh
```

This will:
- Start the backend API on `http://localhost:8080`
- Start the frontend on `http://localhost:3000`
- Install frontend dependencies if needed

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

## 🌐 Access the Application

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080/api
- **API Documentation**: http://localhost:8080/docs
- **Health Check**: http://localhost:8080/api/health

## 📱 Usage

### Creating a Campaign

1. Navigate to the **Campaigns** page
2. Click **Create Campaign**
3. Fill in:
   - Campaign name
   - Description
   - Default language
   - Active status
4. Click **Create**

### Initiating Bulk Calls

1. Go to a campaign's detail page
2. Click **Initiate Calls**
3. Either:
   - **Upload CSV**: Use the template format (phone_number, name)
   - **Manual Entry**: Add contacts one by one
4. Select language (optional)
5. Click **Initiate Bulk Calls**

### Monitoring Calls

1. View calls from the campaign detail page
2. Click **View Details** on any call to see:
   - Call status and duration
   - Complete timeline of events
   - User interactions (button presses)
   - Error messages (if any)

### Dashboard Overview

The dashboard shows:
- Total campaigns and active campaigns
- Call statistics (total, pending, completed, failed)
- Success rate percentage
- Recent calls table
- Active campaigns list

## 📊 API Endpoints

### Campaigns
- `GET /api/campaigns` - List all campaigns
- `POST /api/campaigns` - Create campaign
- `GET /api/campaigns/:id` - Get campaign details
- `PUT /api/campaigns/:id` - Update campaign
- `DELETE /api/campaigns/:id` - Delete campaign
- `GET /api/campaigns/:id/calls` - Get campaign calls with stats

### Calls
- `POST /api/calls/bulk` - Initiate bulk calls
- `GET /api/calls/:id` - Get call status with logs

### System
- `GET /api/health` - Health check
- `GET /api/languages` - Get supported languages

### Webhooks (Twilio)
- `POST /api/webhook/voice` - Initial call webhook
- `POST /api/webhook/gather` - Handle user input
- `POST /api/webhook/status` - Call status updates
- `POST /api/webhook/optout` - Opt-out confirmation

## 🎨 Frontend Architecture

```
ivr_frontend/
├── src/
│   ├── components/       # Reusable UI components
│   │   ├── Sidebar.jsx
│   │   ├── CampaignForm.jsx
│   │   ├── BulkCallForm.jsx
│   │   └── CallDetailsModal.jsx
│   ├── pages/           # Page components
│   │   ├── DashboardPage.jsx
│   │   ├── CampaignsPage.jsx
│   │   └── CampaignCallsPage.jsx
│   ├── services/        # API integration
│   │   └── api.js
│   ├── utils/           # Helper functions
│   │   └── helpers.js
│   ├── App.jsx          # Main app with routing
│   └── main.jsx         # Entry point
├── index.html
├── package.json
└── vite.config.js
```

## 🔧 Development

### Backend Development

```bash
cd ivr_api

# Run with hot reload (install air first)
go install github.com/cosmtrek/air@latest
air

# Run tests
go test ./...

# Build for production
go build -o bin/ivr-api
```

### Frontend Development

```bash
cd ivr_frontend

# Development server with hot reload
npm run dev

# Build for production
npm run build

# Preview production build
npm run preview
```

## 🧪 Testing

### Testing with Sample Data

1. Create a test campaign
2. Use test phone numbers (Twilio provides test credentials)
3. Monitor call logs in real-time

### CSV Format

```csv
phone_number,name
+1234567890,John Doe
+0987654321,Jane Smith
```

## 🌍 Supported Languages

- **en** - English (US)
- **es** - Spanish (Spain)
- **fr** - French (France)
- **de** - German (Germany)
- **hi** - Hindi (India)

## 📞 IVR Flow

1. **Welcome Message**: Personalized greeting
2. **Main Menu**:
   - Press 1: Product information
   - Press 2: Special offers
   - Press 3: Opt out
   - Press 9: Repeat menu
3. **Response Handling**: Based on user input
4. **Call Logging**: All interactions tracked

## 🚢 Deployment

### Backend Deployment

1. Set environment variables
2. Build: `go build -o bin/ivr-api`
3. Run: `./bin/ivr-api`

### Frontend Deployment

1. Build: `npm run build`
2. Serve the `dist` folder with any static server

### Environment Variables for Production

```env
ENVIRONMENT=production
BASE_URL=https://your-domain.com
# ... other production configs
```

## 🛡️ Security Considerations

- Implement authentication (JWT recommended)
- Add rate limiting for API endpoints
- Validate Twilio webhook signatures
- Use HTTPS in production
- Sanitize user inputs
- Implement proper error handling

## 📝 License

MIT

## 👥 Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## 🐛 Troubleshooting

### Backend Issues

- **MongoDB Connection**: Ensure MongoDB is running
- **Twilio Errors**: Verify credentials and phone number format
- **Port Already in Use**: Change PORT in .env

### Frontend Issues

- **API Connection**: Check if backend is running on port 8080
- **CORS Errors**: Verify CORS configuration in backend
- **Build Errors**: Clear node_modules and reinstall

## 📚 Documentation

- [API Documentation](ivr_api/docs/API_DOCUMENTATION.md)
- [MongoDB Setup](ivr_api/docs/MONGODB_SETUP.md)
- [Quick Start Guide](ivr_api/docs/QUICKSTART.md)
- [Feature Checklist](ivr_api/docs/FEATURE_CHECKLIST.md)

## 🙏 Acknowledgments

- Twilio for voice API
- Gin framework
- React team
- TailwindCSS
- MongoDB

## 📧 Support

For issues and questions, please create an issue in the repository.

---

**Built with ❤️ using Go, React, and Twilio**
