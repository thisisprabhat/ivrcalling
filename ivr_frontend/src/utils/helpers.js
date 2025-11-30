export const formatPhoneNumber = (phoneNumber) => {
    // Remove all non-numeric characters except +
    const cleaned = phoneNumber.replace(/[^\d+]/g, '');
    return cleaned;
};

export const validatePhoneNumber = (phoneNumber) => {
    // E.164 format validation
    const e164Regex = /^\+[1-9]\d{1,14}$/;
    return e164Regex.test(phoneNumber);
};

export const formatDate = (dateString) => {
    const date = new Date(dateString);
    return date.toLocaleString();
};

export const getStatusColor = (status) => {
    const colors = {
        pending: 'bg-yellow-100 text-yellow-800',
        initiated: 'bg-blue-100 text-blue-800',
        'in-progress': 'bg-indigo-100 text-indigo-800',
        completed: 'bg-green-100 text-green-800',
        failed: 'bg-red-100 text-red-800',
    };
    return colors[status] || 'bg-gray-100 text-gray-800';
};

export const getStatusBadgeColor = (status) => {
    const colors = {
        pending: 'yellow',
        initiated: 'blue',
        'in-progress': 'indigo',
        completed: 'green',
        failed: 'red',
    };
    return colors[status] || 'gray';
};

export const parseCSV = (csvText) => {
    const lines = csvText.trim().split('\n');
    const contacts = [];

    // Skip header if exists
    const startIndex = lines[0].toLowerCase().includes('phone') ? 1 : 0;

    for (let i = startIndex; i < lines.length; i++) {
        const line = lines[i].trim();
        if (!line) continue;

        const parts = line.split(',').map(p => p.trim().replace(/['"]/g, ''));

        if (parts.length >= 1) {
            contacts.push({
                phone_number: formatPhoneNumber(parts[0]),
                name: parts[1] || '',
            });
        }
    }

    return contacts;
};

export const downloadCSVTemplate = () => {
    const csvContent = 'phone_number,name\n+1234567890,John Doe\n+0987654321,Jane Smith';
    const blob = new Blob([csvContent], { type: 'text/csv' });
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = 'contacts_template.csv';
    a.click();
    window.URL.revokeObjectURL(url);
};

export const getLanguageName = (code) => {
    const languages = {
        en: 'English',
        es: 'Spanish',
        fr: 'French',
        de: 'German',
        hi: 'Hindi',
    };
    return languages[code] || code;
};
