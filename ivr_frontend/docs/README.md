# IVR Calling System - Frontend

Interactive React-based GUI for managing IVR campaigns and calls.

## Features

- **Campaign Management**: Create, edit, view, and delete campaigns
- **Bulk Call Initiation**: Send calls to multiple contacts with CSV support
- **Real-time Call Tracking**: Monitor call status and view detailed logs
- **Multi-language Support**: Support for English, Spanish, French, German, and Hindi
- **Responsive Dashboard**: Clean, modern UI with statistics and analytics
- **Campaign Analytics**: View call statistics per campaign

## Tech Stack

- React 18
- Vite
- React Router v6
- Axios
- TailwindCSS
- Lucide Icons

## Getting Started

### Prerequisites

- Node.js 18+ installed
- IVR API running on `http://localhost:8080`

### Installation

```bash
# Install dependencies
npm install

# Start development server
npm run dev
```

The application will be available at `http://localhost:3000`

### Build for Production

```bash
npm run build
npm run preview
```

## Project Structure

```
ivr_frontend/
├── src/
│   ├── components/        # Reusable UI components
│   ├── pages/            # Page components
│   ├── services/         # API services
│   ├── utils/            # Utility functions
│   ├── App.jsx           # Main app component
│   ├── main.jsx          # Entry point
│   └── index.css         # Global styles
├── public/               # Static assets
├── index.html
├── vite.config.js
├── tailwind.config.js
└── package.json
```

## Usage

### Campaign Management

1. Navigate to "Campaigns" page
2. Click "Create Campaign" to add a new campaign
3. Fill in campaign details (name, description, language)
4. Toggle campaign active/inactive status
5. Edit or delete campaigns as needed

### Initiating Calls

1. Select a campaign
2. Click "Initiate Calls"
3. Either:
   - Manually add contacts with phone numbers
   - Upload a CSV file with contact details
4. Select language (optional, defaults to campaign language)
5. Click "Initiate Bulk Calls"

### Monitoring Calls

1. View all calls in the "Calls" page
2. Click on a call to see detailed logs
3. Filter calls by status or campaign
4. View call duration and status updates

## API Integration

The frontend communicates with the backend API at `http://localhost:8080/api`. Ensure the API is running before starting the frontend.

## Environment Variables

Create a `.env` file if needed:

```env
VITE_API_URL=http://localhost:8080/api
```

## License

MIT
