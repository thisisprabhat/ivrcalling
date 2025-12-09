package service

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/qandi/ivr-calling-api/internal/config"
	"github.com/qandi/ivr-calling-api/internal/models"
)

const twilioAPIBaseURL = "https://api.twilio.com/2010-04-01"

type TwilioService struct {
	config     *config.Config
	httpClient *http.Client
	ivrConfig  models.IVRConfig
}

func NewTwilioService(cfg *config.Config) *TwilioService {
	return &TwilioService{
		config: cfg,
		httpClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		ivrConfig: models.GetIVRConfig(),
	}
}

// InitiateCall initiates an IVR call using Twilio
func (s *TwilioService) InitiateCall(phoneNumber, callbackURL string) (*models.CallResponse, error) {
	// Validate phone number
	if !isValidPhoneNumber(phoneNumber) {
		return nil, fmt.Errorf("invalid phone number format")
	}

	// Validate Twilio configuration
	if s.config.TwilioAccountSID == "" || s.config.TwilioAuthToken == "" {
		return nil, fmt.Errorf("Twilio credentials not configured")
	}

	if s.config.TwilioPhoneNumber == "" {
		return nil, fmt.Errorf("Twilio phone number not configured")
	}

	// Prepare TwiML URL for IVR flow
	twimlURL := fmt.Sprintf("%s/api/v1/twiml/welcome", s.config.ServerBaseURL)

	// Prepare Twilio API request
	apiURL := fmt.Sprintf("%s/Accounts/%s/Calls.json", twilioAPIBaseURL, s.config.TwilioAccountSID)

	data := url.Values{}
	data.Set("To", phoneNumber)
	data.Set("From", s.config.TwilioPhoneNumber)
	data.Set("Url", twimlURL)

	if callbackURL != "" {
		data.Set("StatusCallback", callbackURL)
		data.Set("StatusCallbackEvent", "initiated,ringing,answered,completed")
	}

	// Create HTTP request
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	// Add authentication
	auth := base64.StdEncoding.EncodeToString([]byte(s.config.TwilioAccountSID + ":" + s.config.TwilioAuthToken))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Make the API call
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to call Twilio API: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Twilio API error (status %d): %s", resp.StatusCode, string(body))
	}

	// Parse Twilio response
	var twilioResp struct {
		SID    string `json:"sid"`
		Status string `json:"status"`
		To     string `json:"to"`
	}

	if err := json.Unmarshal(body, &twilioResp); err != nil {
		return nil, fmt.Errorf("failed to parse Twilio response: %w", err)
	}

	return &models.CallResponse{
		CallID:      twilioResp.SID,
		PhoneNumber: phoneNumber,
		Status:      twilioResp.Status,
		Message:     "Call initiated successfully via Twilio",
	}, nil
}

// HandleCallback processes callbacks from Twilio
func (s *TwilioService) HandleCallback(callback *models.CallbackRequest) error {
	fmt.Printf("Received Twilio callback for call %s: Event=%s, Digit=%s\n",
		callback.CallID, callback.Event, callback.DigitInput)

	// Handle different events
	switch callback.Event {
	case "call_answered", "answered":
		fmt.Println("Call was answered")
	case "call_completed", "completed":
		fmt.Println("Call completed")
	case "digit_pressed":
		s.handleDigitInput(callback.CallID, callback.DigitInput)
	case "call_failed", "failed":
		fmt.Println("Call failed")
	}

	return nil
}

// GetIVRConfig returns the current IVR configuration
func (s *TwilioService) GetIVRConfig() models.IVRConfig {
	return s.ivrConfig
}

// GenerateWelcomeTwiML generates TwiML for the welcome message
func (s *TwilioService) GenerateWelcomeTwiML() string {
	twiml := `<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice">` + s.ivrConfig.IntroText + `</Say>
    <Gather numDigits="1" action="` + s.config.ServerBaseURL + `/api/v1/twiml/handle-input" method="POST" timeout="10">
        <Say voice="alice">`

	for _, action := range s.ivrConfig.Actions {
		twiml += action.Message + ". "
	}

	twiml += `</Say>
    </Gather>
    <Say voice="alice">We did not receive any input. Goodbye!</Say>
</Response>`

	return twiml
}

// GenerateHandleInputTwiML generates TwiML based on user's digit input
func (s *TwilioService) GenerateHandleInputTwiML(digit string) string {
	for _, action := range s.ivrConfig.Actions {
		if action.Key == digit {
			switch action.Action {
			case "forward":
				// Forward the call to Q&I team
				return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice">Connecting you to the Q and I team. Please wait.</Say>
    <Dial>%s</Dial>
    <Say voice="alice">%s</Say>
</Response>`, action.ForwardTo, s.ivrConfig.EndMessage)

			case "inform":
				// Provide information
				return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice">%s</Say>
    <Say voice="alice">%s</Say>
</Response>`, action.Description, s.ivrConfig.EndMessage)

			case "repeat":
				// Repeat the welcome message
				return s.GenerateWelcomeTwiML()
			}
		}
	}

	// Invalid input
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice">Invalid input. %s</Say>
</Response>`, s.ivrConfig.EndMessage)
}

// handleDigitInput processes digit input from the caller
func (s *TwilioService) handleDigitInput(callID, digit string) {
	for _, action := range s.ivrConfig.Actions {
		if action.Key == digit {
			switch action.Action {
			case "forward":
				fmt.Printf("Forwarding call %s to %s\n", callID, action.ForwardTo)
			case "inform":
				fmt.Printf("Playing information: %s\n", action.Description)
			case "repeat":
				fmt.Printf("Repeating message for call %s\n", callID)
			}
			return
		}
	}
	fmt.Printf("Invalid digit pressed: %s\n", digit)
}

// Helper functions
func isValidPhoneNumber(phone string) bool {
	// Basic validation - should start with + and contain only digits
	if len(phone) < 10 || phone[0] != '+' {
		return false
	}
	return true
}
