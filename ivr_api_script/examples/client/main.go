package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const baseURL = "http://localhost:8080"

type CallRequest struct {
	PhoneNumber string `json:"phone_number"`
	CallbackURL string `json:"callback_url,omitempty"`
}

type CallResponse struct {
	CallID      string `json:"call_id"`
	PhoneNumber string `json:"phone_number"`
	Status      string `json:"status"`
	Message     string `json:"message"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run examples/client/main.go <phone_number>")
		fmt.Println("Example: go run examples/client/main.go +919876543210")
		os.Exit(1)
	}

	phoneNumber := os.Args[1]

	// 1. Check API health
	fmt.Println("1. Checking API health...")
	checkHealth()

	// 2. Get IVR configuration
	fmt.Println("\n2. Getting IVR configuration...")
	getIVRConfig()

	// 3. Initiate a call
	fmt.Printf("\n3. Initiating call to %s...\n", phoneNumber)
	initiateCall(phoneNumber)
}

func checkHealth() {
	resp, err := http.Get(baseURL + "/health")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	fmt.Printf("Response: %s\n", string(body))
}

func getIVRConfig() {
	resp, err := http.Get(baseURL + "/api/v1/config/ivr")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// Pretty print JSON
	var prettyJSON bytes.Buffer
	json.Indent(&prettyJSON, body, "", "  ")
	fmt.Printf("Response:\n%s\n", prettyJSON.String())
}

func initiateCall(phoneNumber string) {
	reqBody := CallRequest{
		PhoneNumber: phoneNumber,
		CallbackURL: "https://yourapp.com/callback",
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		fmt.Printf("Error marshaling request: %v\n", err)
		return
	}

	resp, err := http.Post(
		baseURL+"/api/v1/calls/initiate",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode == http.StatusOK {
		var callResp CallResponse
		json.Unmarshal(body, &callResp)
		fmt.Printf("Success!\n")
		fmt.Printf("  Call ID: %s\n", callResp.CallID)
		fmt.Printf("  Status: %s\n", callResp.Status)
		fmt.Printf("  Message: %s\n", callResp.Message)
	} else {
		fmt.Printf("Error: %s\n", string(body))
	}
}
