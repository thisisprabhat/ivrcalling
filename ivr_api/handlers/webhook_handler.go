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

	log.Printf("=== VOICE WEBHOOK CALLED ===")
	log.Printf("Call ID: %s, Language: %s", callIDStr, language)

	if language == "" {
		language = "en"
	}

	// Get call details
	var customerName string
	var campaign models.Campaign
	useDynamicIVR := false

	if callIDStr != "" {
		callObjID, err := primitive.ObjectIDFromHex(callIDStr)
		if err == nil {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			var call models.Call
			err = h.db.Collection("calls").FindOne(ctx, bson.M{"_id": callObjID}).Decode(&call)
			if err == nil {
				customerName = call.CustomerName
				log.Printf("Found call record - Customer: %s, Campaign ID: %s", customerName, call.CampaignID.Hex())

				// Get campaign details for dynamic IVR
				err = h.db.Collection("campaigns").FindOne(ctx, bson.M{"_id": call.CampaignID}).Decode(&campaign)
				if err == nil {
					log.Printf("Found campaign - Name: %s, IntroText: '%s', Actions count: %d", campaign.Name, campaign.IntroText, len(campaign.Actions))
					if campaign.IntroText != "" || len(campaign.Actions) > 0 {
						useDynamicIVR = true
						log.Printf("✓ USING DYNAMIC IVR - Campaign: %s", campaign.Name)
						for i, action := range campaign.Actions {
							log.Printf("  Action %d: Type=%s, Input=%s, Message=%s, Phone=%s",
								i+1, action.ActionType, action.ActionInput, action.Message, action.ForwardPhone)
						}
					} else {
						log.Printf("✗ Campaign found but no intro_text or actions - using legacy IVR")
					}
				} else {
					log.Printf("✗ Failed to fetch campaign: %v", err)
				}
			} else {
				log.Printf("✗ Failed to fetch call record: %v", err)
			}
		} else {
			log.Printf("✗ Invalid call ID format: %s", callIDStr)
		}
	} else {
		log.Printf("✗ No call_id provided in webhook")
	}

	// Generate TwiML response
	generator := services.NewTwiMLGenerator(language)
	var twiml string

	if useDynamicIVR {
		log.Printf("Generating dynamic welcome TwiML...")
		twiml = generator.GenerateDynamicWelcome(customerName, &campaign)
	} else {
		log.Printf("Generating legacy welcome TwiML...")
		twiml = generator.GenerateWelcome(customerName)
	}

	log.Printf("Sending TwiML response (length: %d bytes)", len(twiml))
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

	log.Printf("=== GATHER WEBHOOK CALLED ===")
	log.Printf("CallSid: %s, Digits pressed: %s", input.CallSid, input.Digits)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get call by Twilio SID to determine language and campaign
	var call models.Call
	language := "en"
	var campaign models.Campaign
	useDynamicIVR := false

	err := h.db.Collection("calls").FindOne(ctx, bson.M{"twilio_call_sid": input.CallSid}).Decode(&call)
	if err == nil {
		language = call.Language
		log.Printf("Found call - ID: %s, Campaign ID: %s", call.ID.Hex(), call.CampaignID.Hex())

		// Log user input
		callLog := models.CallLog{
			CallID:    call.ID,
			Event:     "input_received",
			UserInput: input.Digits,
			Details:   fmt.Sprintf("User pressed: %s", input.Digits),
			CreatedAt: time.Now(),
		}
		h.db.Collection("call_logs").InsertOne(ctx, callLog)

		// Get campaign for dynamic IVR
		err = h.db.Collection("campaigns").FindOne(ctx, bson.M{"_id": call.CampaignID}).Decode(&campaign)
		if err == nil && (campaign.IntroText != "" || len(campaign.Actions) > 0) {
			useDynamicIVR = true
			log.Printf("✓ USING DYNAMIC IVR for gather - Campaign: %s, Actions: %d", campaign.Name, len(campaign.Actions))
		}
	} else {
		log.Printf("✗ Failed to find call by SID: %v", err)
	}

	generator := services.NewTwiMLGenerator(language)
	var twiml string

	if useDynamicIVR {
		log.Printf("Processing dynamic IVR input: %s", input.Digits)
		// Handle dynamic IVR based on campaign actions
		if input.Digits == "0" {
			// Repeat menu
			log.Printf("User pressed 0 - repeating menu")
			twiml = generator.GenerateDynamicWelcome("", &campaign)
		} else if len(campaign.Actions) == 0 {
			// Campaign has intro_text but no actions - just repeat the intro
			log.Printf("No actions defined - repeating intro")
			twiml = generator.GenerateDynamicWelcome("", &campaign)
		} else {
			// Find matching action
			var matchedAction *models.IVRAction
			for i := range campaign.Actions {
				log.Printf("Checking action %d: input='%s' vs pressed='%s'", i, campaign.Actions[i].ActionInput, input.Digits)
				if campaign.Actions[i].ActionInput == input.Digits {
					matchedAction = &campaign.Actions[i]
					log.Printf("✓ MATCHED ACTION: Type=%s, Message=%s, Phone=%s",
						matchedAction.ActionType, matchedAction.Message, matchedAction.ForwardPhone)
					break
				}
			}

			if matchedAction != nil {
				// Execute the matched action
				log.Printf("Executing action: %s", matchedAction.ActionType)
				twiml = generator.GenerateDynamicResponse(matchedAction, &campaign)

				// Log action execution
				if !call.ID.IsZero() {
					eventType := fmt.Sprintf("action_%s_executed", matchedAction.ActionType)
					details := fmt.Sprintf("User pressed %s - Action type: %s", input.Digits, matchedAction.ActionType)
					h.createCallLog(call.ID, eventType, details)
				}
			} else {
				// Invalid input - repeat the menu
				log.Printf("✗ No matching action found for input: %s - repeating menu", input.Digits)
				twiml = generator.GenerateDynamicWelcome("", &campaign)
			}
		}
	} else {
		// Use legacy static menu handling
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
