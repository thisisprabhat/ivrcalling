# Frontend Quick Start Guide

## 🎯 Overview

This React-based frontend provides a complete GUI for managing IVR campaigns and calls. Built with modern technologies for a responsive, real-time experience.

## 🏗️ Tech Stack

- **React 18** - UI library
- **Vite** - Build tool and dev server
- **React Router v6** - Client-side routing
- **Axios** - HTTP client
- **TailwindCSS** - Utility-first CSS
- **Lucide React** - Beautiful icons

## ⚡ Getting Started

### 1. Install Dependencies

```bash
npm install
```

### 2. Configure Environment

The `.env` file is already configured for local development:

```env
VITE_API_URL=http://localhost:8080/api
```

### 3. Start Development Server

```bash
npm run dev
```

The app will be available at `http://localhost:3000`

## 📱 Features

### Dashboard
- **Real-time Statistics**: Auto-refreshing stats every 10 seconds
- **Campaign Overview**: See all active campaigns at a glance
- **Recent Calls**: Monitor the latest call activity
- **Success Metrics**: Track overall performance

### Campaign Management
- **Create Campaigns**: Easy form with language selection
- **Edit Campaigns**: Update campaign details
- **Toggle Active Status**: Enable/disable campaigns quickly
- **Delete Campaigns**: Remove campaigns with confirmation

### Bulk Calls
- **Manual Entry**: Add contacts one by one
- **CSV Upload**: Import multiple contacts at once
- **CSV Template**: Download template for proper format
- **Language Override**: Choose different language per call batch
- **Phone Validation**: E.164 format validation

### Call Monitoring
- **Real-time Updates**: Auto-refresh every 5 seconds
- **Status Tracking**: Visual indicators for call states
- **Call Details**: Complete timeline with events
- **User Interactions**: See what buttons users pressed
- **Error Tracking**: View error messages for failed calls

## 🎨 UI Components

### Reusable Components

- **Sidebar**: Navigation with active route highlighting
- **CampaignForm**: Modal form for creating/editing campaigns
- **BulkCallForm**: Modal for initiating bulk calls with CSV support
- **CallDetailsModal**: Detailed call information with timeline

### Pages

- **DashboardPage**: Overview with statistics and recent activity
- **CampaignsPage**: Campaign listing with CRUD operations
- **CampaignCallsPage**: Campaign-specific calls with statistics

## 🔄 Data Flow

### API Integration

All API calls go through the centralized `services/api.js`:

```javascript
import { campaignService, callService, systemService } from './services/api';

// Example: Get all campaigns
const campaigns = await campaignService.getAllCampaigns();

// Example: Initiate bulk calls
await callService.initiateBulkCalls({
  campaign_id: '123',
  language: 'en',
  contacts: [...]
});
```

### Auto-refresh Pattern

Pages automatically refresh data at intervals:

```javascript
useEffect(() => {
  loadData();
  const interval = setInterval(loadData, 5000); // Refresh every 5s
  return () => clearInterval(interval);
}, []);
```

## 📋 CSV Format

### Template

Download the template or use this format:

```csv
phone_number,name
+1234567890,John Doe
+44123456789,Jane Smith
+91987654321,Raj Kumar
```

### Rules

- Phone numbers must be in E.164 format (`+[country code][number]`)
- Name is optional but recommended
- First row can be a header (will be auto-detected)

## 🎯 Usage Examples

### Creating a Campaign

1. Click "Create Campaign" button
2. Fill in:
   ```
   Name: Summer Sale 2025
   Description: Promotional campaign
   Language: English
   ✓ Campaign is active
   ```
3. Click "Create"

### Initiating Calls

1. Navigate to campaign
2. Click "Initiate Calls"
3. Option A - Upload CSV:
   - Click "Upload CSV file"
   - Select your CSV
4. Option B - Manual:
   - Click "+ Add Contact"
   - Enter phone number (e.g., +1234567890)
   - Enter name
5. Click "Initiate X Call(s)"

### Viewing Call Details

1. Find call in table
2. Click "View Details"
3. See:
   - Call information (customer, phone, language)
   - Current status
   - Duration
   - Complete timeline of events
   - User inputs (button presses)

## 🔧 Development

### Project Structure

```
src/
├── components/          # Reusable UI components
│   ├── Sidebar.jsx
│   ├── CampaignForm.jsx
│   ├── BulkCallForm.jsx
│   └── CallDetailsModal.jsx
├── pages/              # Route-based pages
│   ├── DashboardPage.jsx
│   ├── CampaignsPage.jsx
│   └── CampaignCallsPage.jsx
├── services/           # API integration layer
│   └── api.js
├── utils/              # Helper functions
│   └── helpers.js
├── App.jsx             # Main app with routing
├── main.jsx            # React entry point
└── index.css           # Global styles + Tailwind
```

### Adding a New Feature

1. **Create Component** (if needed):
   ```bash
   touch src/components/NewComponent.jsx
   ```

2. **Add API Service** (if needed):
   ```javascript
   // In services/api.js
   export const newService = {
     getData: async () => {
       const response = await api.get('/new-endpoint');
       return response.data;
     }
   };
   ```

3. **Create Page** (if needed):
   ```bash
   touch src/pages/NewPage.jsx
   ```

4. **Add Route**:
   ```javascript
   // In App.jsx
   <Route path="/new" element={<NewPage />} />
   ```

5. **Update Sidebar** (if needed):
   ```javascript
   // In Sidebar.jsx
   { path: '/new', icon: <Icon />, label: 'New' }
   ```

### Building for Production

```bash
# Create optimized production build
npm run build

# Preview production build locally
npm run preview
```

Output will be in `dist/` directory.

## 🎨 Styling

### TailwindCSS Utilities

Common patterns used:

```javascript
// Card
className="bg-white rounded-lg shadow p-6"

// Button Primary
className="px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700"

// Button Secondary
className="px-4 py-2 border border-gray-300 rounded-lg hover:bg-gray-50"

// Status Badge
className="px-2 py-1 bg-green-100 text-green-800 rounded-full text-xs"

// Input Field
className="px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500"
```

### Status Colors

- Pending: Yellow (`bg-yellow-100 text-yellow-800`)
- Initiated: Blue (`bg-blue-100 text-blue-800`)
- In Progress: Indigo (`bg-indigo-100 text-indigo-800`)
- Completed: Green (`bg-green-100 text-green-800`)
- Failed: Red (`bg-red-100 text-red-800`)

## 🐛 Troubleshooting

### Common Issues

**API Connection Failed**
- Ensure backend is running on port 8080
- Check CORS configuration in backend
- Verify `VITE_API_URL` in `.env`

**Build Errors**
```bash
# Clear cache and reinstall
rm -rf node_modules package-lock.json
npm install
```

**Module Not Found**
```bash
# Install missing dependencies
npm install
```

**Hot Reload Not Working**
- Restart dev server
- Clear browser cache
- Check for errors in console

### Debug Mode

Open browser DevTools:
- **Console**: See API errors and logs
- **Network**: Monitor API calls
- **React DevTools**: Inspect component state

## 📊 Performance

### Auto-refresh Intervals

- Dashboard: 10 seconds
- Campaign Calls: 5 seconds

To modify:
```javascript
const interval = setInterval(loadData, 5000); // Change 5000 to desired ms
```

### Optimization Tips

1. **Pagination**: Add pagination for large datasets
2. **Debouncing**: Add for search inputs
3. **Lazy Loading**: For large lists
4. **Memoization**: Use `useMemo` for expensive calculations

## 🔐 Security Notes

For production:

1. **API Key**: Add authentication headers
2. **Input Validation**: Already includes phone number validation
3. **XSS Protection**: React handles by default
4. **HTTPS**: Use in production
5. **Environment Variables**: Don't commit sensitive data

## 📚 Resources

- [React Documentation](https://react.dev)
- [Vite Guide](https://vitejs.dev/guide/)
- [TailwindCSS Docs](https://tailwindcss.com/docs)
- [React Router](https://reactrouter.com)
- [Axios](https://axios-http.com/docs/intro)

## 🚀 Next Steps

1. Add user authentication
2. Implement role-based access
3. Add data export functionality
4. Create advanced analytics charts
5. Add notification system
6. Implement WebSocket for real-time updates

---

**Happy Coding! 🎉**
