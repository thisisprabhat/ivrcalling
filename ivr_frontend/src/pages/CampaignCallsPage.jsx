import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';
import { campaignService, callService } from '../services/api';
import { ArrowLeft, Phone, TrendingUp, TrendingDown, Clock, XCircle } from 'lucide-react';
import { getStatusColor, formatDate } from '../utils/helpers';
import BulkCallForm from '../components/BulkCallForm';
import CallDetailsModal from '../components/CallDetailsModal';

const CampaignCallsPage = () => {
    const { id } = useParams();
    const navigate = useNavigate();
    const [campaign, setCampaign] = useState(null);
    const [calls, setCalls] = useState([]);
    const [stats, setStats] = useState(null);
    const [loading, setLoading] = useState(true);
    const [showBulkForm, setShowBulkForm] = useState(false);
    const [selectedCall, setSelectedCall] = useState(null);

    useEffect(() => {
        loadData();
        // Poll for updates every 5 seconds
        const interval = setInterval(loadData, 5000);
        return () => clearInterval(interval);
    }, [id]);

    const loadData = async () => {
        try {
            const [campaignData, callsData] = await Promise.all([
                campaignService.getCampaign(id),
                campaignService.getCampaignCalls(id),
            ]);
            setCampaign(campaignData);
            setCalls(callsData.calls);
            setStats(callsData.stats);
        } catch (error) {
            console.error('Failed to load data:', error);
        } finally {
            setLoading(false);
        }
    };

    const handleBulkCallSubmit = async (callData) => {
        try {
            await callService.initiateBulkCalls(callData);
            setShowBulkForm(false);
            loadData();
        } catch (error) {
            console.error('Failed to initiate calls:', error);
            throw error;
        }
    };

    const handleCallClick = async (call) => {
        try {
            const callDetails = await callService.getCallStatus(call.id);
            setSelectedCall(callDetails);
        } catch (error) {
            console.error('Failed to load call details:', error);
            alert('Failed to load call details');
        }
    };

    if (loading) {
        return (
            <div className="flex items-center justify-center h-64">
                <div className="text-gray-500">Loading...</div>
            </div>
        );
    }

    return (
        <div className="space-y-6">
            {/* Header */}
            <div className="flex items-center gap-4">
                <button
                    onClick={() => navigate('/campaigns')}
                    className="p-2 hover:bg-gray-100 rounded-lg transition"
                >
                    <ArrowLeft size={24} />
                </button>
                <div className="flex-1">
                    <h1 className="text-3xl font-bold text-gray-900">{campaign?.name}</h1>
                    <p className="text-gray-600 mt-1">Campaign Calls & Statistics</p>
                </div>
                <button
                    onClick={() => setShowBulkForm(true)}
                    className="flex items-center gap-2 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition"
                >
                    <Phone size={20} />
                    Initiate Calls
                </button>
            </div>

            {/* Stats Cards */}
            {stats && (
                <div className="grid grid-cols-1 md:grid-cols-5 gap-4">
                    <StatCard
                        title="Total Calls"
                        value={stats.total}
                        icon={<Phone size={24} />}
                        color="blue"
                    />
                    <StatCard
                        title="Pending"
                        value={stats.pending}
                        icon={<Clock size={24} />}
                        color="yellow"
                    />
                    <StatCard
                        title="In Progress"
                        value={stats.initiated + (stats.in_progress || 0)}
                        icon={<TrendingUp size={24} />}
                        color="indigo"
                    />
                    <StatCard
                        title="Completed"
                        value={stats.completed}
                        icon={<TrendingUp size={24} />}
                        color="green"
                    />
                    <StatCard
                        title="Failed"
                        value={stats.failed}
                        icon={<XCircle size={24} />}
                        color="red"
                    />
                </div>
            )}

            {/* Calls Table */}
            <div className="bg-white rounded-lg shadow overflow-hidden">
                <div className="px-6 py-4 border-b">
                    <h2 className="text-lg font-semibold">All Calls</h2>
                </div>
                <div className="overflow-x-auto">
                    {calls.length === 0 ? (
                        <div className="text-center py-12 text-gray-500">
                            No calls yet. Click "Initiate Calls" to get started.
                        </div>
                    ) : (
                        <table className="w-full">
                            <thead className="bg-gray-50">
                                <tr>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                                        Contact
                                    </th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                                        Phone Number
                                    </th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                                        Status
                                    </th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                                        Duration
                                    </th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                                        Created
                                    </th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">
                                        Actions
                                    </th>
                                </tr>
                            </thead>
                            <tbody className="divide-y divide-gray-200">
                                {calls.map((call) => (
                                    <tr key={call.id} className="hover:bg-gray-50">
                                        <td className="px-6 py-4 whitespace-nowrap">
                                            <div className="text-sm font-medium text-gray-900">
                                                {call.customer_name || 'N/A'}
                                            </div>
                                        </td>
                                        <td className="px-6 py-4 whitespace-nowrap">
                                            <div className="text-sm text-gray-900">{call.phone_number}</div>
                                        </td>
                                        <td className="px-6 py-4 whitespace-nowrap">
                                            <span
                                                className={`px-2 py-1 inline-flex text-xs leading-5 font-semibold rounded-full ${getStatusColor(
                                                    call.status
                                                )}`}
                                            >
                                                {call.status}
                                            </span>
                                        </td>
                                        <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                                            {call.duration ? `${call.duration}s` : '-'}
                                        </td>
                                        <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                                            {formatDate(call.created_at)}
                                        </td>
                                        <td className="px-6 py-4 whitespace-nowrap text-sm">
                                            <button
                                                onClick={() => handleCallClick(call)}
                                                className="text-blue-600 hover:text-blue-900"
                                            >
                                                View Details
                                            </button>
                                        </td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    )}
                </div>
            </div>

            {/* Bulk Call Form Modal */}
            {showBulkForm && (
                <BulkCallForm
                    campaignId={id}
                    campaignLanguage={campaign.language}
                    onSubmit={handleBulkCallSubmit}
                    onCancel={() => setShowBulkForm(false)}
                />
            )}

            {/* Call Details Modal */}
            {selectedCall && (
                <CallDetailsModal
                    call={selectedCall}
                    onClose={() => setSelectedCall(null)}
                />
            )}
        </div>
    );
};

const StatCard = ({ title, value, icon, color }) => {
    const colorClasses = {
        blue: 'bg-blue-50 text-blue-600',
        yellow: 'bg-yellow-50 text-yellow-600',
        indigo: 'bg-indigo-50 text-indigo-600',
        green: 'bg-green-50 text-green-600',
        red: 'bg-red-50 text-red-600',
    };

    return (
        <div className="bg-white rounded-lg shadow p-6">
            <div className="flex items-center gap-4">
                <div className={`p-3 rounded-lg ${colorClasses[color]}`}>{icon}</div>
                <div>
                    <p className="text-sm text-gray-600">{title}</p>
                    <p className="text-2xl font-bold text-gray-900">{value}</p>
                </div>
            </div>
        </div>
    );
};

export default CampaignCallsPage;
