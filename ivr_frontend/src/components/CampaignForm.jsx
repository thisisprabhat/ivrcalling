import React, { useState, useEffect } from 'react';
import { X, Plus, Trash2 } from 'lucide-react';
import { systemService } from '../services/api';

const CampaignForm = ({ campaign, onSubmit, onCancel }) => {
    const [formData, setFormData] = useState({
        name: '',
        description: '',
        language: 'en',
        intro_text: '',
        actions: [],
        is_active: true,
    });
    const [languages, setLanguages] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');

    useEffect(() => {
        loadLanguages();
        if (campaign) {
            setFormData({
                name: campaign.name,
                description: campaign.description,
                language: campaign.language,
                intro_text: campaign.intro_text || '',
                actions: campaign.actions || [],
                is_active: campaign.is_active,
            });
        }
    }, [campaign]);

    const loadLanguages = async () => {
        try {
            const data = await systemService.getLanguages();
            setLanguages(data.languages || []);
        } catch (error) {
            console.error('Failed to load languages:', error);
        }
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');

        // Validation
        if (!formData.name.trim()) {
            setError('Campaign name is required');
            return;
        }

        if (!formData.description.trim()) {
            setError('Description is required');
            return;
        }

        if (!formData.intro_text.trim()) {
            setError('Intro text is required');
            return;
        }

        setLoading(true);

        try {
            await onSubmit(formData);
        } catch (err) {
            setError(err.response?.data?.error || 'Failed to save campaign');
        } finally {
            setLoading(false);
        }
    };

    const handleChange = (e) => {
        const { name, value, type, checked } = e.target;
        setFormData((prev) => ({
            ...prev,
            [name]: type === 'checkbox' ? checked : value,
        }));
    };

    const handleAddAction = () => {
        setFormData((prev) => ({
            ...prev,
            actions: [
                ...prev.actions,
                {
                    action_type: 'information',
                    action_input: '',
                    message: '',
                    forward_phone: '',
                },
            ],
        }));
    };

    const handleRemoveAction = (index) => {
        setFormData((prev) => ({
            ...prev,
            actions: prev.actions.filter((_, i) => i !== index),
        }));
    };

    const handleActionChange = (index, field, value) => {
        setFormData((prev) => ({
            ...prev,
            actions: prev.actions.map((action, i) =>
                i === index ? { ...action, [field]: value } : action
            ),
        }));
    };

    return (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 p-4">
            <div className="bg-white rounded-lg shadow-xl max-w-3xl w-full max-h-[90vh] flex flex-col">
                {/* Header */}
                <div className="flex justify-between items-center p-6 border-b bg-white rounded-t-lg flex-shrink-0">
                    <h2 className="text-xl font-semibold">
                        {campaign ? 'Edit Campaign' : 'Create Campaign'}
                    </h2>
                    <button
                        onClick={onCancel}
                        className="text-gray-400 hover:text-gray-600"
                        type="button"
                    >
                        <X size={24} />
                    </button>
                </div>

                {/* Form */}
                <form onSubmit={handleSubmit} className="flex flex-col flex-1 overflow-hidden">
                    <div className="p-6 space-y-6 overflow-y-auto flex-1">
                    {error && (
                        <div className="p-3 bg-red-50 border border-red-200 text-red-700 rounded">
                            {error}
                        </div>
                    )}

                    {/* Basic Information */}
                    <div className="space-y-4">
                        <h3 className="text-lg font-semibold text-gray-900">Basic Information</h3>
                        
                        <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">
                                Campaign Name <span className="text-red-500">*</span>
                            </label>
                            <input
                                type="text"
                                name="name"
                                value={formData.name}
                                onChange={handleChange}
                                required
                                className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                                placeholder="Summer Sale 2025"
                            />
                        </div>

                        <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">
                                Description <span className="text-red-500">*</span>
                            </label>
                            <textarea
                                name="description"
                                value={formData.description}
                                onChange={handleChange}
                                required
                                rows={2}
                                className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                                placeholder="Promotional campaign for summer products"
                            />
                        </div>

                        <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">
                                Default Language <span className="text-red-500">*</span>
                            </label>
                            <select
                                name="language"
                                value={formData.language}
                                onChange={handleChange}
                                required
                                className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                            >
                                {languages.map((lang) => (
                                    <option key={lang} value={lang}>
                                        {lang.toUpperCase()} - {getLanguageFullName(lang)}
                                    </option>
                                ))}
                            </select>
                        </div>

                        <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">
                                Intro Text <span className="text-red-500">*</span>
                            </label>
                            <textarea
                                name="intro_text"
                                value={formData.intro_text}
                                onChange={handleChange}
                                required
                                rows={3}
                                className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                                placeholder="Welcome to our service. We have exclusive offers for you today."
                            />
                            <p className="text-xs text-gray-500 mt-1">
                                This message will be played at the start of the call
                            </p>
                        </div>

                        <div className="flex items-center gap-2">
                            <input
                                type="checkbox"
                                name="is_active"
                                id="is_active"
                                checked={formData.is_active}
                                onChange={handleChange}
                                className="w-4 h-4 text-blue-600 border-gray-300 rounded focus:ring-blue-500"
                            />
                            <label htmlFor="is_active" className="text-sm font-medium text-gray-700">
                                Campaign is active
                            </label>
                        </div>
                    </div>

                    {/* IVR Actions */}
                    <div className="space-y-4 border-t pt-6">
                        <div className="flex justify-between items-center">
                            <div>
                                <h3 className="text-lg font-semibold text-gray-900">IVR Actions</h3>
                                <p className="text-sm text-gray-500">Define what happens when users press keys</p>
                            </div>
                            <button
                                type="button"
                                onClick={handleAddAction}
                                className="flex items-center gap-2 px-3 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition text-sm"
                            >
                                <Plus size={16} />
                                Add Action
                            </button>
                        </div>

                        {formData.actions.length === 0 ? (
                            <div className="text-center py-8 bg-gray-50 rounded-lg">
                                <p className="text-gray-500 text-sm">
                                    No actions defined. Click "Add Action" to create IVR menu options.
                                </p>
                            </div>
                        ) : (
                            <div className="space-y-4">
                                {formData.actions.map((action, index) => (
                                    <div
                                        key={index}
                                        className="p-4 border border-gray-200 rounded-lg space-y-3 bg-gray-50"
                                    >
                                        <div className="flex justify-between items-center">
                                            <h4 className="font-medium text-gray-900">
                                                Action {index + 1}
                                            </h4>
                                            <button
                                                type="button"
                                                onClick={() => handleRemoveAction(index)}
                                                className="text-red-600 hover:text-red-700"
                                            >
                                                <Trash2 size={18} />
                                            </button>
                                        </div>

                                        <div className="grid grid-cols-2 gap-3">
                                            <div>
                                                <label className="block text-sm font-medium text-gray-700 mb-1">
                                                    Action Type
                                                </label>
                                                <select
                                                    value={action.action_type}
                                                    onChange={(e) =>
                                                        handleActionChange(index, 'action_type', e.target.value)
                                                    }
                                                    className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                                                >
                                                    <option value="information">Information</option>
                                                    <option value="forward">Forward</option>
                                                </select>
                                            </div>

                                            <div>
                                                <label className="block text-sm font-medium text-gray-700 mb-1">
                                                    Key Press
                                                </label>
                                                <input
                                                    type="text"
                                                    value={action.action_input}
                                                    onChange={(e) =>
                                                        handleActionChange(index, 'action_input', e.target.value)
                                                    }
                                                    maxLength={1}
                                                    pattern="[0-9]"
                                                    className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                                                    placeholder="1-9"
                                                />
                                            </div>
                                        </div>

                                        {action.action_type === 'information' ? (
                                            <div>
                                                <label className="block text-sm font-medium text-gray-700 mb-1">
                                                    Message (Text or Audio URL)
                                                </label>
                                                <textarea
                                                    value={action.message}
                                                    onChange={(e) =>
                                                        handleActionChange(index, 'message', e.target.value)
                                                    }
                                                    rows={2}
                                                    className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                                                    placeholder="Enter text to speak or URL to audio file (http://...)"
                                                />
                                                <p className="text-xs text-gray-500 mt-1">
                                                    Use text for speech or provide an audio file URL
                                                </p>
                                            </div>
                                        ) : (
                                            <div>
                                                <label className="block text-sm font-medium text-gray-700 mb-1">
                                                    Forward to Phone Number
                                                </label>
                                                <input
                                                    type="tel"
                                                    value={action.forward_phone}
                                                    onChange={(e) =>
                                                        handleActionChange(index, 'forward_phone', e.target.value)
                                                    }
                                                    className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                                                    placeholder="+1234567890"
                                                />
                                                <p className="text-xs text-gray-500 mt-1">
                                                    Call will be forwarded to this number
                                                </p>
                                            </div>
                                        )}
                                    </div>
                                ))}
                            </div>
                        )}
                    </div>

                    </div>

                    {/* Actions */}
                    <div className="flex gap-3 p-6 border-t bg-white flex-shrink-0 rounded-b-lg">
                        <button
                            type="button"
                            onClick={onCancel}
                            className="flex-1 px-4 py-2 border border-gray-300 text-gray-700 rounded-lg hover:bg-gray-50 transition"
                        >
                            Cancel
                        </button>
                        <button
                            type="submit"
                            disabled={loading}
                            className="flex-1 px-4 py-2 bg-blue-600 text-white rounded-lg hover:bg-blue-700 transition disabled:opacity-50"
                        >
                            {loading ? 'Saving...' : campaign ? 'Update' : 'Create'}
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
};

const getLanguageFullName = (code) => {
    const names = {
        en: 'English',
        es: 'Spanish',
        fr: 'French',
        de: 'German',
        hi: 'Hindi',
    };
    return names[code] || code;
};

export default CampaignForm;
