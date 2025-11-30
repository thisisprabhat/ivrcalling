import React, { useState, useEffect } from 'react';
import { campaignService, systemService } from '../services/api';
import { Phone, TrendingUp, Clock, CheckCircle, XCircle, Activity } from 'lucide-react';
import { useNavigate } from 'react-router-dom';
import { getLanguageName, formatDate } from '../utils/helpers';

const DashboardPage = () => {
    const [campaigns, setCampaigns] = useState([]);
    const [stats, setStats] = useState({
        totalCampaigns: 0,
        activeCampaigns: 0,
        totalCalls: 0,
        completedCalls: 0,
        failedCalls: 0,
        pendingCalls: 0,
    });
    const [recentCalls, setRecentCalls] = useState([]);
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate();

    useEffect(() => {
        loadDashboardData();
        // Refresh every 10 seconds
        const interval = setInterval(loadDashboardData, 10000);
        return () => clearInterval(interval);
    }, []);

    const loadDashboardData = async () => {
        try {
            const campaignsData = await campaignService.getAllCampaigns();
            setCampaigns(campaignsData);

            // Calculate stats
            const activeCampaigns = campaignsData.filter((c) => c.is_active).length;

            // Load calls for all campaigns
            const allCallsData = await Promise.all(
                campaignsData.map((campaign) =>
                    campaignService.getCampaignCalls(campaign.id).catch(() => ({ calls: [], stats: {} }))
                )
            );

            // Aggregate stats
            let totalCalls = 0;
            let completedCalls = 0;
            let failedCalls = 0;
            let pendingCalls = 0;
            const allCalls = [];

            allCallsData.forEach((data) => {
                if (data.stats) {
                    totalCalls += data.stats.total || 0;
                    completedCalls += data.stats.completed || 0;
                    failedCalls += data.stats.failed || 0;
                    pendingCalls += data.stats.pending || 0;
                }
                if (data.calls) {
                    allCalls.push(...data.calls);
                }
            });

            // Sort by created_at and get recent 5
            const recent = allCalls
                .sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
                .slice(0, 5);

            setStats({
                totalCampaigns: campaignsData.length,
                activeCampaigns,
                totalCalls,
                completedCalls,
                failedCalls,
                pendingCalls,
            });
            setRecentCalls(recent);
        } catch (error) {
            console.error('Failed to load dashboard data:', error);
        } finally {
            setLoading(false);
        }
    };

    if (loading) {
        return (
            <div className="flex items-center justify-center h-64">
                <div className="text-gray-500">Loading dashboard...</div>
            </div>
        );
    }

    return (
        <div className="space-y-6">
            {/* Header */}
            <div>
                <h1 className="text-3xl font-bold text-gray-900">Dashboard</h1>
                <p className="text-gray-600 mt-1">Overview of your IVR calling campaigns</p>
            </div>

            {/* Stats Grid */}
            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                <StatCard
                    title="Total Campaigns"
                    value={stats.totalCampaigns}
                    subtitle={`${stats.activeCampaigns} active`}
                    icon={<Activity size={24} />}
                    color="blue"
                />
                <StatCard
                    title="Total Calls"
                    value={stats.totalCalls}
                    subtitle="All time"
                    icon={<Phone size={24} />}
                    color="indigo"
                />
                <StatCard
                    title="Pending Calls"
                    value={stats.pendingCalls}
                    subtitle="Awaiting initiation"
                    icon={<Clock size={24} />}
                    color="yellow"
                />
                <StatCard
                    title="Completed Calls"
                    value={stats.completedCalls}
                    subtitle="Successfully finished"
                    icon={<CheckCircle size={24} />}
                    color="green"
                />
                <StatCard
                    title="Failed Calls"
                    value={stats.failedCalls}
                    subtitle="Unsuccessful attempts"
                    icon={<XCircle size={24} />}
                    color="red"
                />
                <StatCard
                    title="Success Rate"
                    value={
                        stats.totalCalls > 0
                            ? `${Math.round((stats.completedCalls / stats.totalCalls) * 100)}%`
                            : '0%'
                    }
                    subtitle="Overall performance"
                    icon={<TrendingUp size={24} />}
                    color="purple"
                />
            </div>

            {/* Active Campaigns */}
            <div className="bg-white rounded-lg shadow">
                <div className="px-6 py-4 border-b flex justify-between items-center">
                    <h2 className="text-lg font-semibold">Active Campaigns</h2>
                    <button
                        onClick={() => navigate('/campaigns')}
                        className="text-sm text-blue-600 hover:text-blue-700"
                    >
                        View All
                    </button>
                </div>
                <div className="divide-y">
                    {campaigns
                        .filter((c) => c.is_active)
                        .slice(0, 5)
                        .map((campaign) => (
                            <div
                                key={campaign.id}
                                className="px-6 py-4 hover:bg-gray-50 cursor-pointer"
                                onClick={() => navigate(`/campaigns/${campaign.id}/calls`)}
                            >
                                <div className="flex justify-between items-center">
                                    <div>
                                        <h3 className="font-medium text-gray-900">{campaign.name}</h3>
                                        <p className="text-sm text-gray-500 mt-1">
                                            {getLanguageName(campaign.language)} • Created{' '}
                                            {formatDate(campaign.created_at)}
                                        </p>
                                    </div>
                                    <span className="px-3 py-1 bg-green-100 text-green-800 text-xs font-medium rounded-full">
                                        Active
                                    </span>
                                </div>
                            </div>
                        ))}
                    {campaigns.filter((c) => c.is_active).length === 0 && (
                        <div className="px-6 py-8 text-center text-gray-500">
                            No active campaigns
                        </div>
                    )}
                </div>
            </div>

            {/* Recent Calls */}
            <div className="bg-white rounded-lg shadow">
                <div className="px-6 py-4 border-b">
                    <h2 className="text-lg font-semibold">Recent Calls</h2>
                </div>
                <div className="overflow-x-auto">
                    {recentCalls.length === 0 ? (
                        <div className="px-6 py-8 text-center text-gray-500">
                            No calls yet
                        </div>
                    ) : (
                        <table className="w-full">
                            <thead className="bg-gray-50">
                                <tr>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                                        Contact
                                    </th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                                        Phone
                                    </th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                                        Status
                                    </th>
                                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">
                                        Created
                                    </th>
                                </tr>
                            </thead>
                            <tbody className="divide-y divide-gray-200">
                                {recentCalls.map((call) => (
                                    <tr key={call.id} className="hover:bg-gray-50">
                                        <td className="px-6 py-4 whitespace-nowrap">
                                            <div className="text-sm font-medium text-gray-900">
                                                {call.customer_name || 'N/A'}
                                            </div>
                                        </td>
                                        <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                                            {call.phone_number}
                                        </td>
                                        <td className="px-6 py-4 whitespace-nowrap">
                                            <span
                                                className={`px-2 py-1 text-xs font-semibold rounded-full ${call.status === 'completed'
                                                        ? 'bg-green-100 text-green-800'
                                                        : call.status === 'failed'
                                                            ? 'bg-red-100 text-red-800'
                                                            : 'bg-yellow-100 text-yellow-800'
                                                    }`}
                                            >
                                                {call.status}
                                            </span>
                                        </td>
                                        <td className="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                                            {formatDate(call.created_at)}
                                        </td>
                                    </tr>
                                ))}
                            </tbody>
                        </table>
                    )}
                </div>
            </div>
        </div>
    );
};

const StatCard = ({ title, value, subtitle, icon, color }) => {
    const colorClasses = {
        blue: 'bg-blue-50 text-blue-600',
        indigo: 'bg-indigo-50 text-indigo-600',
        yellow: 'bg-yellow-50 text-yellow-600',
        green: 'bg-green-50 text-green-600',
        red: 'bg-red-50 text-red-600',
        purple: 'bg-purple-50 text-purple-600',
    };

    return (
        <div className="bg-white rounded-lg shadow p-6">
            <div className="flex items-center gap-4">
                <div className={`p-3 rounded-lg ${colorClasses[color]}`}>{icon}</div>
                <div className="flex-1">
                    <p className="text-sm text-gray-600">{title}</p>
                    <p className="text-2xl font-bold text-gray-900 mt-1">{value}</p>
                    <p className="text-xs text-gray-500 mt-1">{subtitle}</p>
                </div>
            </div>
        </div>
    );
};

export default DashboardPage;
