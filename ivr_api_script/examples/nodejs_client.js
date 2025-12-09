const axios = require('axios');

const BASE_URL = 'http://localhost:8080';

/**
 * Example client for Q&I IVR API
 * Usage: node examples/nodejs_client.js +919876543210
 */

class IVRClient {
  constructor(baseURL) {
    this.baseURL = baseURL;
    this.client = axios.create({
      baseURL: this.baseURL,
      headers: {
        'Content-Type': 'application/json'
      }
    });
  }

  /**
   * Check API health
   */
  async checkHealth() {
    try {
      const response = await this.client.get('/health');
      console.log('Health Check:', response.data);
      return response.data;
    } catch (error) {
      console.error('Health check failed:', error.message);
      throw error;
    }
  }

  /**
   * Get IVR configuration
   */
  async getIVRConfig() {
    try {
      const response = await this.client.get('/api/v1/config/ivr');
      console.log('IVR Configuration:');
      console.log(JSON.stringify(response.data, null, 2));
      return response.data;
    } catch (error) {
      console.error('Failed to get IVR config:', error.message);
      throw error;
    }
  }

  /**
   * Initiate an IVR call
   * @param {string} phoneNumber - Phone number in E.164 format
   * @param {string} callbackURL - Optional callback URL
   */
  async initiateCall(phoneNumber, callbackURL = null) {
    try {
      const payload = {
        phone_number: phoneNumber
      };

      if (callbackURL) {
        payload.callback_url = callbackURL;
      }

      const response = await this.client.post('/api/v1/calls/initiate', payload);
      
      console.log('\nCall Initiated Successfully!');
      console.log('  Call ID:', response.data.call_id);
      console.log('  Phone Number:', response.data.phone_number);
      console.log('  Status:', response.data.status);
      console.log('  Message:', response.data.message);
      
      return response.data;
    } catch (error) {
      if (error.response) {
        console.error('Failed to initiate call:', error.response.data);
      } else {
        console.error('Failed to initiate call:', error.message);
      }
      throw error;
    }
  }

  /**
   * Send a callback (for testing)
   * @param {string} callID - Call ID
   * @param {string} event - Event type
   * @param {string} digitInput - Digit pressed by user
   */
  async sendCallback(callID, event, digitInput = null) {
    try {
      const payload = {
        call_id: callID,
        event: event,
        timestamp: new Date().toISOString()
      };

      if (digitInput) {
        payload.digit_input = digitInput;
      }

      const response = await this.client.post('/api/v1/callbacks/ivr', payload);
      console.log('Callback sent:', response.data);
      return response.data;
    } catch (error) {
      console.error('Failed to send callback:', error.message);
      throw error;
    }
  }
}

// Main execution
async function main() {
  const args = process.argv.slice(2);
  
  if (args.length === 0) {
    console.log('Usage: node examples/nodejs_client.js <phone_number>');
    console.log('Example: node examples/nodejs_client.js +919876543210');
    process.exit(1);
  }

  const phoneNumber = args[0];
  const client = new IVRClient(BASE_URL);

  try {
    // 1. Check health
    console.log('\n1. Checking API Health...');
    console.log('='.repeat(50));
    await client.checkHealth();

    // 2. Get IVR config
    console.log('\n2. Getting IVR Configuration...');
    console.log('='.repeat(50));
    await client.getIVRConfig();

    // 3. Initiate call
    console.log('\n3. Initiating Call...');
    console.log('='.repeat(50));
    const callResponse = await client.initiateCall(
      phoneNumber,
      'https://yourapp.com/callback'
    );

    console.log('\n✓ All operations completed successfully!');
  } catch (error) {
    console.error('\n✗ Error occurred:', error.message);
    process.exit(1);
  }
}

// Run if called directly
if (require.main === module) {
  main();
}

// Export for use as a module
module.exports = IVRClient;
