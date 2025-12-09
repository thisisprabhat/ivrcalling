#!/usr/bin/env python3
"""
Q&I IVR API Client Example
Usage: python examples/python_client.py +919876543210
"""

import sys
import requests
import json

BASE_URL = "http://localhost:8080"


def check_health():
    """Check API health"""
    print("1. Checking API health...")
    response = requests.get(f"{BASE_URL}/health")
    print(f"Response: {response.json()}\n")


def get_ivr_config():
    """Get IVR configuration"""
    print("2. Getting IVR configuration...")
    response = requests.get(f"{BASE_URL}/api/v1/config/ivr")
    config = response.json()
    print(json.dumps(config, indent=2))
    print()


def initiate_call(phone_number):
    """Initiate an IVR call"""
    print(f"3. Initiating call to {phone_number}...")

    payload = {
        "phone_number": phone_number,
        "callback_url": "https://yourapp.com/callback"
    }

    response = requests.post(
        f"{BASE_URL}/api/v1/calls/initiate",
        json=payload
    )

    if response.status_code == 200:
        result = response.json()
        print("Success!")
        print(f"  Call ID: {result['call_id']}")
        print(f"  Status: {result['status']}")
        print(f"  Message: {result['message']}")
    else:
        print(f"Error: {response.json()}")


def main():
    if len(sys.argv) < 2:
        print("Usage: python examples/python_client.py <phone_number>")
        print("Example: python examples/python_client.py +919876543210")
        sys.exit(1)

    phone_number = sys.argv[1]

    try:
        check_health()
        get_ivr_config()
        initiate_call(phone_number)
    except requests.exceptions.ConnectionError:
        print("Error: Could not connect to API server.")
        print("Make sure the server is running on http://localhost:8080")
    except Exception as e:
        print(f"Error: {e}")


if __name__ == "__main__":
    main()
