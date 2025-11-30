import React, { useState, useEffect } from 'react';
import { campaignService } from '../services/api';
import { Plus, Edit, Trash2, Phone, ToggleLeft, ToggleRight } from 'lucide-react';
import { getLanguageName, formatDate } from '../utils/helpers';
import CampaignForm from '../components/CampaignForm';
import { useNavigate } from 'react-router-dom';

const CampaignsPage = () => {
    const [campaigns, setCampaigns] = useState([]);
    const [loading, setLoading] = useState(true);
    const [showForm, setShowForm] = useState(false);
    const [editingCampaign, setEditingCampaign] = useState(null);
    const navigate = useNavigate();

    useEffect(() => {
        loadCampaigns();
    }, []);

    const loadCampaigns = async () => {
        try {
            setLoading(true);
            const data = await campaignService.getAllCampaigns();
            setCampaigns(data);
        } catch (error) {
            console.error('Failed to load campaigns:', error);
            alert('Failed to load campaigns');
        } finally {
            setLoading(false);
        }
    };

    const handleDelete = async (id) => {
        if (!confirm('Are you sure you want to delete this campaign?')) return;

        try {
            await campaignService.deleteCampaign(id);
            loadCampaigns();
        } catch (error) {
            console.error('Failed to delete campaign:', error);
            alert('Failed to delete campaign');
        }
    };

    const handleToggleActive = async (campaign) => {
        try {
            await campaignService.updateCampaign(campaign.id, {
                is_active: !campaign.is_active,
            });
            loadCampaigns();
        } catch (error) {
            console.error('Failed to update campaign:', error);
            alert('Failed to update campaign');
        }
    };

    const handleFormSubmit = async (campaignData) => {
        try {
            if (editingCampaign) {
                await campaignService.updateCampaign(editingCampaign.id, campaignData);
            } else {
                await campaignService.createCampaign(campaignData);
            }
            setShowForm(false);
            setEditingCampaign(null);
            loadCampaigns();
        } catch (error) {
            console.error('Failed to save campaign:', error);
            throw error;
        }
    };

    const handleEdit = (campaign) => {
        setEditingCampaign(campaign);
        setShowForm(true);
    };

    const handleCreate = () => {
        setEditingCampaign(null);
        setShowForm(true);
    };

    if (loading) {
        return (
            <div className="flex items-center justify-center h-64">
                <div className="text-gray-500">Loading campaigns...</div>
            </div>
        );
    }

    return (
        <div className="space-y-6">
            {/* Header */}
            <div className="flex justify-between items-center">
                <div>
                    <h1 className="text-3xl font-bold text-gray-900">Campaigns</h1>
                    <p className="text-gray-600 mt-1">Manage your marketing campaigns</p>
                </div>
                <button
                    onClick={handleCreate}
                    className="flex items-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition"
                >
                    <Plus size={20} />
                    Create Campaign
                </button>
            </div>

            {/* Campaign Form Modal */}
            {showForm && (
                <CampaignForm
                    campaign={editingCampaign}
                    onSubmit={handleFormSubmit}
                    onCancel={() => {
                        setShowForm(false);
                        setEditingCampaign(null);
                    }}
                />
            )}

            {/* Campaigns Grid */}
            {campaigns.length === 0 ? (
                <div className="text-center py-12 bg-white rounded-lg shadow">
                    <p className="text-gray-500">No campaigns yet. Create your first campaign!</p>
                </div>
            ) : (
                <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                    {campaigns.map((campaign) => (
                        <div
                            key={campaign.id}
                            className="bg-white rounded-lg shadow-md hover:shadow-lg transition p-6"
                        >
                            {/* Campaign Header */}
                            <div className="flex justify-between items-start mb-4">
                                <div className="flex-1">
                                    <h3 className="text-xl font-semibold text-gray-900 mb-1">
                                        {campaign.name}
                                    </h3>
                                    <p className="text-sm text-gray-500 line-clamp-2">
                                        {campaign.description}
                                    </p>
                                </div>
                                <button
                                    onClick={() => handleToggleActive(campaign)}
                                    className="ml-2"
                                    title={campaign.is_active ? 'Active' : 'Inactive'}
                                >
                                    {campaign.is_active ? (
                                        <ToggleRight size={28} className="text-green-600" />
                                    ) : (
                                        <ToggleLeft size={28} className="text-gray-400" />
                                    )}
                                </button>
                            </div>

                            {/* Campaign Details */}
                            <div className="space-y-2 mb-4">
                                <div className="flex items-center gap-2 text-sm">
                                    <span className="text-gray-500">Language:</span>
                                    <span className="font-medium text-gray-700">
                                        {getLanguageName(campaign.language)}
                                    </span>
                                </div>
                                <div className="flex items-center gap-2 text-sm">
                                    <span className="text-gray-500">Status:</span>
                                    <span
                                        className={`px-2 py-1 rounded-full text-xs font-medium ${campaign.is_active
                                                ? 'bg-green-100 text-green-800'
                                                : 'bg-gray-100 text-gray-800'
                                            }`}
                                    >
                                        {campaign.is_active ? 'Active' : 'Inactive'}
                                    </span>
                                </div>
                                <div className="text-xs text-gray-500">
                                    Created: {formatDate(campaign.created_at)}
                                </div>
                            </div>

                            {/* Actions */}
                            <div className="flex gap-2 pt-4 border-t">
                                <button
                                    onClick={() => navigate(`/campaigns/${campaign.id}/calls`)}
                                    className="flex-1 flex items-center justify-center gap-2 px-3 py-2 bg-blue-50 text-blue-600 rounded hover:bg-blue-100 transition"
                                >
                                    <Phone size={16} />
                                    <span className="text-sm font-medium">View Calls</span>
                                </button>
                                <button
                                    onClick={() => handleEdit(campaign)}
                                    className="px-3 py-2 bg-gray-50 text-gray-600 rounded hover:bg-gray-100 transition"
                                    title="Edit"
                                >
                                    <Edit size={16} />
                                </button>
                                <button
                                    onClick={() => handleDelete(campaign.id)}
                                    className="px-3 py-2 bg-red-50 text-red-600 rounded hover:bg-red-100 transition"
                                    title="Delete"
                                >
                                    <Trash2 size={16} />
                                </button>
                            </div>
                        </div>
                    ))}
                </div>
            )}
        </div>
    );
};

export default CampaignsPage;
