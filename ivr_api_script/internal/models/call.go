package models

// CallRequest represents a request to initiate an IVR call
type CallRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required" example:"+919876543210"`
	CallbackURL string `json:"callback_url,omitempty" example:"https://yourapp.com/callback"`
}

// CallResponse represents the response after initiating a call
type CallResponse struct {
	CallID      string `json:"call_id" example:"call_123456"`
	PhoneNumber string `json:"phone_number" example:"+919876543210"`
	Status      string `json:"status" example:"initiated"`
	Message     string `json:"message" example:"Call initiated successfully"`
}

// CallbackRequest represents the callback data from IVR provider
type CallbackRequest struct {
	CallID     string `json:"call_id" example:"call_123456"`
	Event      string `json:"event" example:"call_answered"`
	DigitInput string `json:"digit_input,omitempty" example:"1"`
	Timestamp  string `json:"timestamp" example:"2025-12-09T10:30:00Z"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error" example:"Invalid phone number format"`
	Message string `json:"message" example:"Phone number must start with +"`
}

// HealthResponse represents health check response
type HealthResponse struct {
	Status  string `json:"status" example:"healthy"`
	Version string `json:"version" example:"1.0.0"`
}
