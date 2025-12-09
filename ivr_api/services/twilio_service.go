package services

import (
	"fmt"
	"log"

	"github.com/prabhatkumar/ivrcalling/config"
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioService struct {
	client      *twilio.RestClient
	phoneNumber string
	webhookURL  string
}

func NewTwilioService(cfg *config.Config) *TwilioService {
	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: cfg.TwilioAccountSID,
		Password: cfg.TwilioAuthToken,
	})

	return &TwilioService{
		client:      client,
		phoneNumber: cfg.TwilioPhoneNumber,
		webhookURL:  cfg.WebhookBaseURL,
	}
}

// MakeCall initiates an outbound IVR call
func (s *TwilioService) MakeCall(toNumber string, language string, callID string) (*twilioApi.ApiV2010Call, error) {
	// Construct webhook URL with call ID and language
	statusCallbackURL := fmt.Sprintf("%s/api/webhook/status", s.webhookURL)
	voiceURL := fmt.Sprintf("%s/api/webhook/voice?call_id=%s&language=%s", s.webhookURL, callID, language)

	log.Printf("=== INITIATING TWILIO CALL ===")
	log.Printf("To: %s", toNumber)
	log.Printf("From: %s", s.phoneNumber)
	log.Printf("Voice URL: %s", voiceURL)
	log.Printf("Status Callback URL: %s", statusCallbackURL)

	params := &twilioApi.CreateCallParams{}
	params.SetTo(toNumber)
	params.SetFrom(s.phoneNumber)
	params.SetUrl(voiceURL)
	params.SetMethod("POST")
	params.SetStatusCallback(statusCallbackURL)
	params.SetStatusCallbackMethod("POST")
	params.SetStatusCallbackEvent([]string{"initiated", "ringing", "answered", "completed"})

	call, err := s.client.Api.CreateCall(params)
	if err != nil {
		log.Printf("✗ Twilio call creation failed: %v", err)
		return nil, fmt.Errorf("failed to create call: %w", err)
	}

	log.Printf("✓ Twilio call created - SID: %s", *call.Sid)
	return call, nil
}

// GetCallDetails retrieves call information from Twilio
func (s *TwilioService) GetCallDetails(callSid string) (*twilioApi.ApiV2010Call, error) {
	call, err := s.client.Api.FetchCall(callSid, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch call: %w", err)
	}

	return call, nil
}
