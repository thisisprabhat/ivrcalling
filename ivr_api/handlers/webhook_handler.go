package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prabhatkumar/ivrcalling/database"
	"github.com/prabhatkumar/ivrcalling/models"
	"github.com/prabhatkumar/ivrcalling/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WebhookHandler struct {
	db *database.MongoDB
}

func NewWebhookHandler(db *database.MongoDB) *WebhookHandler {
	return &WebhookHandler{db: db}
}

// HandleVoiceWebhook handles initial voice webhook from Twilio
func (h *WebhookHandler) HandleVoiceWebhook(c *gin.Context) {
	callIDStr := c.Query("call_id")
	language := c.Query("language")

	if language == "" {
		language = "en"
	}

	// Get call details
	var customerName string
	if callIDStr != "" {
		callObjID, err := primitive.ObjectIDFromHex(callIDStr)
		if err == nil {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			var call models.Call
			err = h.db.Collection("calls").FindOne(ctx, bson.M{"_id": callObjID}).Decode(&call)
			if err == nil {
				customerName = call.CustomerName
			}
		}
	}

	// Generate TwiML response
	generator := services.NewTwiMLGenerator(language)
	twiml := generator.GenerateWelcome(customerName)

	c.Header("Content-Type", "text/xml")
	c.String(http.StatusOK, twiml)
}

// HandleGatherWebhook handles digit gathering from IVR menu
func (h *WebhookHandler) HandleGatherWebhook(c *gin.Context) {
	var input models.IVRInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get call by Twilio SID to determine language
	var call models.Call
	language := "en"
	err := h.db.Collection("calls").FindOne(ctx, bson.M{"twilio_call_sid": input.CallSid}).Decode(&call)
	if err == nil {
		language = call.Language

		// Log user input
		callLog := models.CallLog{
			CallID:    call.ID,
			Event:     "input_received",
			UserInput: input.Digits,
			Details:   fmt.Sprintf("User pressed: %s", input.Digits),
			CreatedAt: time.Now(),
		}
		h.db.Collection("call_logs").InsertOne(ctx, callLog)
	}

	generator := services.NewTwiMLGenerator(language)
	var twiml string

	// Handle menu options
	switch input.Digits {
	case "1":
		// Product information
		twiml = generator.GenerateProductInfo()
		if !call.ID.IsZero() {
			h.createCallLog(call.ID, "product_info_requested", "User requested product information")
		}
	case "2":
		// Special offers
		twiml = generator.GenerateOfferDetails()
		if !call.ID.IsZero() {
			h.createCallLog(call.ID, "offer_requested", "User requested offer details")
		}
	case "3":
		// Opt out
		twiml = generator.GenerateOptOut()
		if !call.ID.IsZero() {
			h.createCallLog(call.ID, "opt_out_requested", "User requested to opt out")
		}
	case "0":
		// Return to main menu
		twiml = generator.GenerateMainMenu()
	case "9":
		// Repeat menu
		twiml = generator.GenerateMainMenu()
	default:
		// Invalid input
		twiml = generator.GenerateInvalidInput()
	}

	c.Header("Content-Type", "text/xml")
	c.String(http.StatusOK, twiml)
}

// HandleOptOutConfirm handles opt-out confirmation
func (h *WebhookHandler) HandleOptOutConfirm(c *gin.Context) {
	var input models.IVRInput
	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get call by Twilio SID
	var call models.Call
	language := "en"
	err := h.db.Collection("calls").FindOne(ctx, bson.M{"twilio_call_sid": input.CallSid}).Decode(&call)
	if err == nil {
		language = call.Language

		if input.Digits == "1" {
			// Mark customer as opted out (you can add an opt-out table)
			h.createCallLog(call.ID, "opted_out", "User confirmed opt-out")
		}
	}

	generator := services.NewTwiMLGenerator(language)
	var twiml string

	if input.Digits == "1" {
		twiml = generator.GenerateOptOutConfirm()
	} else {
		twiml = generator.GenerateMainMenu()
	}

	c.Header("Content-Type", "text/xml")
	c.String(http.StatusOK, twiml)
}

func (h *WebhookHandler) createCallLog(callID primitive.ObjectID, event, details string) {
	callLog := models.CallLog{
		CallID:    callID,
		Event:     event,
		Details:   details,
		CreatedAt: time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if _, err := h.db.Collection("call_logs").InsertOne(ctx, callLog); err != nil {
		log.Printf("Failed to create call log: %v", err)
	}
}
