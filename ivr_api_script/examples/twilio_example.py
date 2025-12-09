#!/usr/bin/env python3
"""
Twilio IVR API Client Example
Demonstrates how to use the Q&I IVR API with Twilio integration

Usage: python examples/twilio_example.py +919876543210
"""

import sys
import requests
import json
import time

BASE_URL = "http://localhost:8080"


class TwilioIVRClient:
    def __init__(self, base_url):
        self.base_url = base_url

    def check_health(self):
        """Check if the API is running"""
        print("🔍 Checking API health...")
        try:
            response = requests.get(f"{self.base_url}/health")
            response.raise_for_status()
            data = response.json()
            print(f"✅ API is {data['status']} (version {data['version']})\n")
            return True
        except requests.exceptions.RequestException as e:
            print(f"❌ Health check failed: {e}\n")
            return False

    def get_ivr_config(self):
        """Get the IVR configuration"""
        print("📋 Fetching IVR configuration...")
        try:
            response = requests.get(f"{self.base_url}/api/v1/config/ivr")
            response.raise_for_status()
            config = response.json()

            print("\n📢 Intro Message:")
            print(f"   {config['intro_text'][:100]}...\n")

            print("🎯 Available Actions:")
            for action in config['actions']:
                print(f"   [{action['key']}] {action['message']}")

            print(f"\n👋 End Message:")
            print(f"   {config['end_message']}\n")

            return config
        except requests.exceptions.RequestException as e:
            print(f"❌ Failed to get config: {e}\n")
            return None

    def initiate_call(self, phone_number, callback_url=None):
        """Initiate an IVR call via Twilio"""
        print(f"📞 Initiating call to {phone_number}...")

        payload = {
            "phone_number": phone_number
        }

        if callback_url:
            payload["callback_url"] = callback_url

        try:
            response = requests.post(
                f"{self.base_url}/api/v1/calls/initiate",
                json=payload,
                timeout=30
            )

            if response.status_code == 200:
                result = response.json()
                print("\n✅ Call initiated successfully!")
                print(f"   📱 Phone Number: {result['phone_number']}")
                print(f"   🆔 Call SID: {result['call_id']}")
                print(f"   📊 Status: {result['status']}")
                print(f"   💬 Message: {result['message']}\n")

                print("🎤 The recipient will hear:")
                print("   1. Welcome message about Q&I")
                print("   2. Menu with 3 options:")
                print("      [1] Talk to Q&I team")
                print("      [2] Learn more about Q&I")
                print("      [3] Repeat the message")
                print("   3. Thank you message\n")

                return result
            else:
                error = response.json()
                print(f"❌ Failed to initiate call:")
                print(f"   Error: {error.get('error', 'Unknown error')}")
                print(f"   Message: {error.get('message', 'No details')}\n")
                return None

        except requests.exceptions.Timeout:
            print("❌ Request timed out. The call may still be processing.\n")
            return None
        except requests.exceptions.RequestException as e:
            print(f"❌ Request failed: {e}\n")
            return None

    def simulate_callback(self, call_id, event, digit=None):
        """Simulate a callback from Twilio (for testing)"""
        print(f"🔔 Simulating callback: {event}")

        payload = {
            "call_id": call_id,
            "event": event,
            "timestamp": time.strftime("%Y-%m-%dT%H:%M:%SZ", time.gmtime())
        }

        if digit:
            payload["digit_input"] = digit

        try:
            response = requests.post(
                f"{self.base_url}/api/v1/callbacks/ivr",
                json=payload
            )

            if response.status_code == 200:
                print(f"✅ Callback processed\n")
            else:
                print(f"❌ Callback failed\n")

        except requests.exceptions.RequestException as e:
            print(f"❌ Callback error: {e}\n")


def print_header():
    """Print a nice header"""
    print("\n" + "="*60)
    print("  Q&I IVR API - Twilio Integration Example")
    print("="*60 + "\n")


def print_footer():
    """Print a footer"""
    print("\n" + "="*60)
    print("  Test Complete!")
    print("="*60 + "\n")


def main():
    """Main execution"""
    if len(sys.argv) < 2:
        print("Usage: python examples/twilio_example.py <phone_number>")
        print("Example: python examples/twilio_example.py +919876543210")
        print("\nMake sure:")
        print("  1. Server is running (go run cmd/server/main.go)")
        print("  2. Twilio credentials are configured in .env")
        print("  3. Phone number is in E.164 format (+[country][number])")
        sys.exit(1)

    phone_number = sys.argv[1]

    # Validate phone number format
    if not phone_number.startswith('+'):
        print("❌ Error: Phone number must start with + (E.164 format)")
        print("Example: +919876543210")
        sys.exit(1)

    print_header()

    client = TwilioIVRClient(BASE_URL)

    # Step 1: Health check
    if not client.check_health():
        print("⚠️  Server might not be running. Start it with:")
        print("   go run cmd/server/main.go")
        sys.exit(1)

    # Step 2: Get IVR configuration
    config = client.get_ivr_config()
    if not config:
        print("⚠️  Could not fetch IVR configuration")
        sys.exit(1)

    # Step 3: Ask for confirmation
    print(f"🚀 Ready to call {phone_number}")
    response = input("   Continue? (y/n): ")

    if response.lower() != 'y':
        print("\n❌ Call cancelled\n")
        sys.exit(0)

    print()

    # Step 4: Initiate the call
    result = client.initiate_call(
        phone_number,
        callback_url=f"{BASE_URL}/api/v1/callbacks/ivr"
    )

    if result:
        call_id = result.get('call_id')

        # Optional: Simulate callbacks (for testing webhook handling)
        if call_id and input("   Simulate callbacks for testing? (y/n): ").lower() == 'y':
            print()
            time.sleep(1)
            client.simulate_callback(call_id, "call_answered")
            time.sleep(1)
            client.simulate_callback(call_id, "digit_pressed", "1")
            time.sleep(1)
            client.simulate_callback(call_id, "call_completed")

    print_footer()

    print("📚 Next steps:")
    print("   - Check Twilio console for call status")
    print("   - Monitor server logs for webhook callbacks")
    print("   - Review docs/TWILIO_SETUP.md for more details\n")


if __name__ == "__main__":
    main()
