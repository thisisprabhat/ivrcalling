package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// IVRAction represents an action in the IVR flow
type IVRAction struct {
	ActionType   string `bson:"action_type" json:"action_type"`                         // "information" or "forward"
	ActionInput  string `bson:"action_input" json:"action_input"`                       // key press (e.g., "1", "2", "3")
	Message      string `bson:"message,omitempty" json:"message,omitempty"`             // text or URL for information type
	ForwardPhone string `bson:"forward_phone,omitempty" json:"forward_phone,omitempty"` // phone number for forward type
}

// Campaign represents a marketing campaign
type Campaign struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	Language    string             `bson:"language" json:"language"`
	IntroText   string             `bson:"intro_text" json:"intro_text"`               // Intro text played at start
	Actions     []IVRAction        `bson:"actions,omitempty" json:"actions,omitempty"` // IVR actions
	IsActive    bool               `bson:"is_active" json:"is_active"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

// Call represents an individual call
type Call struct {
	ID            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CampaignID    primitive.ObjectID `bson:"campaign_id" json:"campaign_id"`
	PhoneNumber   string             `bson:"phone_number" json:"phone_number"`
	CustomerName  string             `bson:"customer_name" json:"customer_name"`
	Status        string             `bson:"status" json:"status"` // pending, initiated, in-progress, completed, failed
	TwilioCallSID string             `bson:"twilio_call_sid" json:"twilio_call_sid"`
	Language      string             `bson:"language" json:"language"`
	Duration      int                `bson:"duration" json:"duration"` // in seconds
	ErrorMessage  string             `bson:"error_message,omitempty" json:"error_message,omitempty"`
	CreatedAt     time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt     time.Time          `bson:"updated_at" json:"updated_at"`
}

// CallLog represents detailed logs for each call
type CallLog struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CallID    primitive.ObjectID `bson:"call_id" json:"call_id"`
	Event     string             `bson:"event" json:"event"` // initiated, answered, input_received, completed, failed
	Details   string             `bson:"details" json:"details"`
	UserInput string             `bson:"user_input,omitempty" json:"user_input,omitempty"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
}

// BulkCallRequest represents the request to initiate bulk calls
type BulkCallRequest struct {
	CampaignID string           `json:"campaign_id" binding:"required"`
	Language   string           `json:"language"`
	Contacts   []ContactRequest `json:"contacts" binding:"required,min=1"`
}

// ContactRequest represents a single contact in bulk call request
type ContactRequest struct {
	PhoneNumber string `json:"phone_number" binding:"required"`
	Name        string `json:"name"`
}

// CallStatusUpdate represents webhook data from Twilio
type CallStatusUpdate struct {
	CallSid      string `form:"CallSid" json:"call_sid"`
	CallStatus   string `form:"CallStatus" json:"call_status"`
	CallDuration string `form:"CallDuration" json:"call_duration"`
	From         string `form:"From" json:"from"`
	To           string `form:"To" json:"to"`
}

// IVRInput represents user input during IVR call
type IVRInput struct {
	CallSid string `form:"CallSid" json:"call_sid"`
	Digits  string `form:"Digits" json:"digits"`
}
