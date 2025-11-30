import React, { useState, useEffect } from 'react';
import { X, Upload, Plus, Trash2, Download } from 'lucide-react';
import { systemService } from '../services/api';
import { parseCSV, validatePhoneNumber, downloadCSVTemplate, getLanguageName } from '../utils/helpers';

const BulkCallForm = ({ campaignId, campaignLanguage, onSubmit, onCancel }) => {
    const [contacts, setContacts] = useState([{ phone_number: '', name: '' }]);
    const [language, setLanguage] = useState(campaignLanguage || 'en');
    const [languages, setLanguages] = useState([]);
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');

    useEffect(() => {
        loadLanguages();
    }, []);

    const loadLanguages = async () => {
        try {
            const data = await systemService.getLanguages();
            setLanguages(data.languages || []);
        } catch (error) {
            console.error('Failed to load languages:', error);
        }
    };

    const handleAddContact = () => {
        setContacts([...contacts, { phone_number: '', name: '' }]);
    };

    const handleRemoveContact = (index) => {
        setContacts(contacts.filter((_, i) => i !== index));
    };

    const handleContactChange = (index, field, value) => {
        const updated = [...contacts];
        updated[index][field] = value;
        setContacts(updated);
    };

    const handleFileUpload = (e) => {
        const file = e.target.files[0];
        if (!file) return;

        const reader = new FileReader();
        reader.onload = (event) => {
            try {
                const csvContacts = parseCSV(event.target.result);
                setContacts(csvContacts);
                setError('');
            } catch (err) {
                setError('Failed to parse CSV file');
            }
        };
        reader.readAsText(file);
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError('');

        // Validate contacts
        const validContacts = contacts.filter((c) => c.phone_number.trim());
        if (validContacts.length === 0) {
            setError('Please add at least one contact');
            return;
        }

        // Validate phone numbers
        const invalidNumbers = validContacts.filter(
            (c) => !validatePhoneNumber(c.phone_number)
        );
        if (invalidNumbers.length > 0) {
            setError(
                `Invalid phone number format. Use E.164 format (e.g., +1234567890). Found ${invalidNumbers.length} invalid number(s).`
            );
            return;
        }

        setLoading(true);

        try {
            await onSubmit({
                campaign_id: campaignId,
                language: language,
                contacts: validContacts,
            });
        } catch (err) {
            setError(err.response?.data?.error || 'Failed to initiate calls');
        } finally {
            setLoading(false);
        }
    };

    return (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 overflow-y-auto">
            <div className="bg-white rounded-lg shadow-xl max-w-4xl w-full mx-4 my-8">
                {/* Header */}
                <div className="flex justify-between items-center p-6 border-b">
                    <h2 className="text-xl font-semibold">Initiate Bulk Calls</h2>
                    <button onClick={onCancel} className="text-gray-400 hover:text-gray-600">
                        <X size={24} />
                    </button>
                </div>

                {/* Form */}
                <form onSubmit={handleSubmit} className="p-6 space-y-6">
                    {error && (
                        <div className="p-3 bg-red-50 border border-red-200 text-red-700 rounded">
                            {error}
                        </div>
                    )}

                    {/* Language Selection */}
                    <div>
                        <label className="block text-sm font-medium text-gray-700 mb-1">
                            Language
                        </label>
                        <select
                            value={language}
                            onChange={(e) => setLanguage(e.target.value)}
                            className="w-full px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                        >
                            {languages.map((lang) => (
                                <option key={lang} value={lang}>
                                    {getLanguageName(lang)}
                                </option>
                            ))}
                        </select>
                    </div>

                    {/* CSV Upload */}
                    <div className="border-2 border-dashed border-gray-300 rounded-lg p-6">
                        <div className="text-center">
                            <Upload className="mx-auto h-12 w-12 text-gray-400" />
                            <div className="mt-4">
                                <label
                                    htmlFor="csv-upload"
                                    className="cursor-pointer text-blue-600 hover:text-blue-700 font-medium"
                                >
                                    Upload CSV file
                                    <input
                                        id="csv-upload"
                                        type="file"
                                        accept=".csv"
                                        onChange={handleFileUpload}
                                        className="hidden"
                                    />
                                </label>
                                <p className="text-sm text-gray-500 mt-1">
                                    or enter contacts manually below
                                </p>
                            </div>
                            <button
                                type="button"
                                onClick={downloadCSVTemplate}
                                className="mt-2 text-sm text-gray-600 hover:text-gray-800 flex items-center gap-1 mx-auto"
                            >
                                <Download size={16} />
                                Download CSV Template
                            </button>
                        </div>
                    </div>

                    {/* Contacts List */}
                    <div className="space-y-3">
                        <div className="flex justify-between items-center">
                            <label className="block text-sm font-medium text-gray-700">
                                Contacts ({contacts.length})
                            </label>
                            <button
                                type="button"
                                onClick={handleAddContact}
                                className="flex items-center gap-1 text-sm text-blue-600 hover:text-blue-700"
                            >
                                <Plus size={16} />
                                Add Contact
                            </button>
                        </div>

                        <div className="max-h-64 overflow-y-auto space-y-2">
                            {contacts.map((contact, index) => (
                                <div key={index} className="flex gap-2">
                                    <input
                                        type="text"
                                        placeholder="+1234567890"
                                        value={contact.phone_number}
                                        onChange={(e) =>
                                            handleContactChange(index, 'phone_number', e.target.value)
                                        }
                                        className="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                                    />
                                    <input
                                        type="text"
                                        placeholder="Name (optional)"
                                        value={contact.name}
                                        onChange={(e) =>
                                            handleContactChange(index, 'name', e.target.value)
                                        }
                                        className="flex-1 px-3 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent"
                                    />
                                    {contacts.length > 1 && (
                                        <button
                                            type="button"
                                            onClick={() => handleRemoveContact(index)}
                                            className="px-3 py-2 text-red-600 hover:bg-red-50 rounded-lg transition"
                                        >
                                            <Trash2 size={20} />
                                        </button>
                                    )}
                                </div>
                            ))}
                        </div>

                        <p className="text-xs text-gray-500">
                            Phone numbers must be in E.164 format (e.g., +1234567890)
                        </p>
                    </div>

                    {/* Actions */}
                    <div className="flex gap-3 pt-4 border-t">
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
                            {loading ? 'Initiating...' : `Initiate ${contacts.filter(c => c.phone_number.trim()).length} Call(s)`}
                        </button>
                    </div>
                </form>
            </div>
        </div>
    );
};

export default BulkCallForm;
