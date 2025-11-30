import React from 'react';
import { X, Phone, Clock, User, Globe } from 'lucide-react';
import { getStatusColor, formatDate, getLanguageName } from '../utils/helpers';

const CallDetailsModal = ({ call, onClose }) => {
    return (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 overflow-y-auto">
            <div className="bg-white rounded-lg shadow-xl max-w-3xl w-full mx-4 my-8">
                {/* Header */}
                <div className="flex justify-between items-center p-6 border-b">
                    <h2 className="text-xl font-semibold">Call Details</h2>
                    <button onClick={onClose} className="text-gray-400 hover:text-gray-600">
                        <X size={24} />
                    </button>
                </div>

                {/* Content */}
                <div className="p-6 space-y-6">
                    {/* Call Information */}
                    <div className="grid grid-cols-2 gap-4">
                        <InfoItem
                            icon={<User size={20} />}
                            label="Customer"
                            value={call.customer_name || 'N/A'}
                        />
                        <InfoItem
                            icon={<Phone size={20} />}
                            label="Phone Number"
                            value={call.phone_number}
                        />
                        <InfoItem
                            icon={<Globe size={20} />}
                            label="Language"
                            value={getLanguageName(call.language)}
                        />
                        <InfoItem
                            icon={<Clock size={20} />}
                            label="Duration"
                            value={call.duration ? `${call.duration} seconds` : 'N/A'}
                        />
                    </div>

                    {/* Status */}
                    <div>
                        <label className="block text-sm font-medium text-gray-700 mb-2">
                            Status
                        </label>
                        <span
                            className={`px-3 py-1 inline-flex text-sm font-semibold rounded-full ${getStatusColor(
                                call.status
                            )}`}
                        >
                            {call.status}
                        </span>
                    </div>

                    {/* Error Message */}
                    {call.error_message && (
                        <div className="p-3 bg-red-50 border border-red-200 rounded">
                            <p className="text-sm font-medium text-red-800">Error:</p>
                            <p className="text-sm text-red-700 mt-1">{call.error_message}</p>
                        </div>
                    )}

                    {/* Twilio Call SID */}
                    {call.twilio_call_sid && (
                        <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">
                                Twilio Call SID
                            </label>
                            <code className="block px-3 py-2 bg-gray-100 rounded text-sm font-mono">
                                {call.twilio_call_sid}
                            </code>
                        </div>
                    )}

                    {/* Timestamps */}
                    <div className="grid grid-cols-2 gap-4 text-sm">
                        <div>
                            <label className="block font-medium text-gray-700 mb-1">
                                Created At
                            </label>
                            <p className="text-gray-600">{formatDate(call.created_at)}</p>
                        </div>
                        <div>
                            <label className="block font-medium text-gray-700 mb-1">
                                Updated At
                            </label>
                            <p className="text-gray-600">{formatDate(call.updated_at)}</p>
                        </div>
                    </div>

                    {/* Call Logs */}
                    {call.call_logs && call.call_logs.length > 0 && (
                        <div>
                            <label className="block text-sm font-medium text-gray-700 mb-3">
                                Call Timeline
                            </label>
                            <div className="space-y-3">
                                {call.call_logs.map((log, index) => (
                                    <div
                                        key={log.id}
                                        className="flex gap-4 pb-3 border-b last:border-b-0"
                                    >
                                        <div className="flex-shrink-0 w-2 h-2 mt-2 rounded-full bg-blue-500"></div>
                                        <div className="flex-1">
                                            <div className="flex justify-between items-start">
                                                <div>
                                                    <p className="font-medium text-gray-900">{log.event}</p>
                                                    <p className="text-sm text-gray-600 mt-1">{log.details}</p>
                                                    {log.user_input && (
                                                        <p className="text-sm text-blue-600 mt-1">
                                                            User Input: {log.user_input}
                                                        </p>
                                                    )}
                                                </div>
                                                <p className="text-xs text-gray-500 whitespace-nowrap ml-4">
                                                    {formatDate(log.created_at)}
                                                </p>
                                            </div>
                                        </div>
                                    </div>
                                ))}
                            </div>
                        </div>
                    )}
                </div>

                {/* Footer */}
                <div className="flex justify-end p-6 border-t">
                    <button
                        onClick={onClose}
                        className="px-4 py-2 bg-gray-100 text-gray-700 rounded-lg hover:bg-gray-200 transition"
                    >
                        Close
                    </button>
                </div>
            </div>
        </div>
    );
};

const InfoItem = ({ icon, label, value }) => (
    <div className="flex items-center gap-3">
        <div className="text-gray-400">{icon}</div>
        <div>
            <p className="text-sm text-gray-500">{label}</p>
            <p className="font-medium text-gray-900">{value}</p>
        </div>
    </div>
);

export default CallDetailsModal;
