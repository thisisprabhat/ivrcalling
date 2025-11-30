import axios from 'axios';

const API_BASE_URL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api';

const api = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Campaign APIs
export const campaignService = {
  // Get all campaigns
  getAllCampaigns: async () => {
    const response = await api.get('/campaigns');
    return response.data;
  },

  // Get single campaign
  getCampaign: async (id) => {
    const response = await api.get(`/campaigns/${id}`);
    return response.data;
  },

  // Create campaign
  createCampaign: async (campaignData) => {
    const response = await api.post('/campaigns', campaignData);
    return response.data;
  },

  // Update campaign
  updateCampaign: async (id, campaignData) => {
    const response = await api.put(`/campaigns/${id}`, campaignData);
    return response.data;
  },

  // Delete campaign
  deleteCampaign: async (id) => {
    const response = await api.delete(`/campaigns/${id}`);
    return response.data;
  },

  // Get campaign calls
  getCampaignCalls: async (id) => {
    const response = await api.get(`/campaigns/${id}/calls`);
    return response.data;
  },
};

// Call APIs
export const callService = {
  // Initiate bulk calls
  initiateBulkCalls: async (callData) => {
    const response = await api.post('/calls/bulk', callData);
    return response.data;
  },

  // Get call status
  getCallStatus: async (id) => {
    const response = await api.get(`/calls/${id}`);
    return response.data;
  },
};

// System APIs
export const systemService = {
  // Health check
  healthCheck: async () => {
    const response = await api.get('/health');
    return response.data;
  },

  // Get supported languages
  getLanguages: async () => {
    const response = await api.get('/languages');
    return response.data;
  },
};

export default api;
