package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// This is a manual integration test - run with: go run verify_test.go
// Requires MongoDB to be running

var baseURL = "http://localhost:8080/api"

func main() {
	fmt.Println("🔍 IVR Calling System - Feature Verification")
	fmt.Println("===========================================\n")

	// Check if server is running
	if !checkServerHealth() {
		fmt.Println("❌ Server is not running. Please start it first with: go run main.go")
		os.Exit(1)
	}

	passed := 0
	failed := 0

	// Test 1: Health Check
	if testHealthCheck() {
		passed++
	} else {
		failed++
	}

	// Test 2: Get Supported Languages
	if testGetLanguages() {
		passed++
	} else {
		failed++
	}

	// Test 3: Create Campaign
	campaignID := testCreateCampaign()
	if campaignID != "" {
		passed++
	} else {
		failed++
	}

	// Test 4: List Campaigns
	if testListCampaigns() {
		passed++
	} else {
		failed++
	}

	// Test 5: Get Campaign
	if campaignID != "" && testGetCampaign(campaignID) {
		passed++
	} else {
		failed++
	}

	// Test 6: Update Campaign
	if campaignID != "" && testUpdateCampaign(campaignID) {
		passed++
	} else {
		failed++
	}

	// Test 7: Bulk Call Request Structure (without actual Twilio call)
	if testBulkCallStructure() {
		passed++
	} else {
		failed++
	}

	// Test 8: Delete Campaign
	if campaignID != "" && testDeleteCampaign(campaignID) {
		passed++
	} else {
		failed++
	}

	// Summary
	fmt.Println("\n===========================================")
	fmt.Printf("✅ Passed: %d\n", passed)
	fmt.Printf("❌ Failed: %d\n", failed)
	fmt.Printf("📊 Total:  %d\n", passed+failed)
	fmt.Println("===========================================\n")

	if failed == 0 {
		fmt.Println("🎉 All features verified successfully!")
		os.Exit(0)
	} else {
		fmt.Println("⚠️  Some tests failed. Please check the output above.")
		os.Exit(1)
	}
}

func checkServerHealth() bool {
	resp, err := http.Get(baseURL + "/health")
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	return resp.StatusCode == 200
}

func testHealthCheck() bool {
	fmt.Print("Test 1: Health Check Endpoint... ")
	resp, err := http.Get(baseURL + "/health")
	if err != nil {
		fmt.Println("❌ FAILED:", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("❌ FAILED: Expected 200, got %d\n", resp.StatusCode)
		return false
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	if result["status"] != "healthy" {
		fmt.Println("❌ FAILED: Status not healthy")
		return false
	}

	fmt.Println("✅ PASSED")
	return true
}

func testGetLanguages() bool {
	fmt.Print("Test 2: Get Supported Languages... ")
	resp, err := http.Get(baseURL + "/languages")
	if err != nil {
		fmt.Println("❌ FAILED:", err)
		return false
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	languages, ok := result["languages"].([]interface{})
	if !ok || len(languages) != 5 {
		fmt.Printf("❌ FAILED: Expected 5 languages, got %v\n", languages)
		return false
	}

	fmt.Println("✅ PASSED")
	return true
}

func testCreateCampaign() string {
	fmt.Print("Test 3: Create Campaign... ")

	campaign := map[string]interface{}{
		"name":        "Test Campaign " + time.Now().Format("15:04:05"),
		"description": "Automated test campaign",
		"language":    "en",
		"is_active":   true,
	}

	jsonData, _ := json.Marshal(campaign)
	resp, err := http.Post(baseURL+"/campaigns", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("❌ FAILED:", err)
		return ""
	}
	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("❌ FAILED: Expected 201, got %d - %s\n", resp.StatusCode, string(body))
		return ""
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)

	campaignID, ok := result["id"].(string)
	if !ok {
		fmt.Println("❌ FAILED: No campaign ID returned")
		return ""
	}

	fmt.Printf("✅ PASSED (ID: %s)\n", campaignID[:8]+"...")
	return campaignID
}

func testListCampaigns() bool {
	fmt.Print("Test 4: List Campaigns... ")
	resp, err := http.Get(baseURL + "/campaigns")
	if err != nil {
		fmt.Println("❌ FAILED:", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("❌ FAILED: Expected 200, got %d\n", resp.StatusCode)
		return false
	}

	var campaigns []map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&campaigns)

	fmt.Printf("✅ PASSED (Found %d campaigns)\n", len(campaigns))
	return true
}

func testGetCampaign(campaignID string) bool {
	fmt.Print("Test 5: Get Campaign by ID... ")
	resp, err := http.Get(baseURL + "/campaigns/" + campaignID)
	if err != nil {
		fmt.Println("❌ FAILED:", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("❌ FAILED: Expected 200, got %d\n", resp.StatusCode)
		return false
	}

	var campaign map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&campaign)

	if campaign["id"] != campaignID {
		fmt.Println("❌ FAILED: Campaign ID mismatch")
		return false
	}

	fmt.Println("✅ PASSED")
	return true
}

func testUpdateCampaign(campaignID string) bool {
	fmt.Print("Test 6: Update Campaign... ")

	update := map[string]interface{}{
		"name":      "Updated Test Campaign",
		"is_active": false,
	}

	jsonData, _ := json.Marshal(update)
	req, _ := http.NewRequest("PUT", baseURL+"/campaigns/"+campaignID, bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ FAILED:", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("❌ FAILED: Expected 200, got %d - %s\n", resp.StatusCode, string(body))
		return false
	}

	fmt.Println("✅ PASSED")
	return true
}

func testBulkCallStructure() bool {
	fmt.Print("Test 7: Bulk Call Request Validation... ")

	// Test with invalid campaign ID
	bulkCall := map[string]interface{}{
		"campaign_id": "invalid_id",
		"contacts": []map[string]string{
			{"phone_number": "+1234567890", "name": "Test User"},
		},
	}

	jsonData, _ := json.Marshal(bulkCall)
	resp, err := http.Post(baseURL+"/calls/bulk", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("❌ FAILED:", err)
		return false
	}
	defer resp.Body.Close()

	// Should fail with invalid campaign ID
	if resp.StatusCode != 400 {
		fmt.Printf("❌ FAILED: Expected 400 for invalid campaign, got %d\n", resp.StatusCode)
		return false
	}

	fmt.Println("✅ PASSED (Validation working)")
	return true
}

func testDeleteCampaign(campaignID string) bool {
	fmt.Print("Test 8: Delete Campaign... ")

	req, _ := http.NewRequest("DELETE", baseURL+"/campaigns/"+campaignID, nil)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ FAILED:", err)
		return false
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("❌ FAILED: Expected 200, got %d\n", resp.StatusCode)
		return false
	}

	fmt.Println("✅ PASSED")
	return true
}
