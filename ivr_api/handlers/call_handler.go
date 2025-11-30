package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prabhatkumar/ivrcalling/database"
	"github.com/prabhatkumar/ivrcalling/models"
	"github.com/prabhatkumar/ivrcalling/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CallHandler struct {
	db            *database.MongoDB
	twilioService *services.TwilioService
}

func NewCallHandler(db *database.MongoDB, twilioService *services.TwilioService) *CallHandler {
	return &CallHandler{
		db:            db,
		twilioService: twilioService,
	}
}

// InitiateBulkCalls handles the bulk call endpoint
func (h *CallHandler) InitiateBulkCalls(c *gin.Context) {
	var request models.BulkCallRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Convert campaign ID string to ObjectID
	campaignObjID, err := primitive.ObjectIDFromHex(request.CampaignID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid campaign ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Verify campaign exists
	var campaign models.Campaign
	err = h.db.Collection("campaigns").FindOne(ctx, bson.M{"_id": campaignObjID}).Decode(&campaign)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}

	if !campaign.IsActive {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campaign is not active"})
		return
	}

	// Determine language
	language := request.Language
	if language == "" {
		language = campaign.Language
	}

	// Create calls and initiate them
	var successCount, failCount int
	var callIDs []string

	for _, contact := range request.Contacts {
		// Create call record
		call := models.Call{
			CampaignID:   campaignObjID,
			PhoneNumber:  contact.PhoneNumber,
			CustomerName: contact.Name,
			Status:       "pending",
			Language:     language,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		}

		callCtx, callCancel := context.WithTimeout(context.Background(), 5*time.Second)
		result, err := h.db.Collection("calls").InsertOne(callCtx, call)
		callCancel()

		if err != nil {
			log.Printf("Failed to create call record: %v", err)
			failCount++
			continue
		}

		call.ID = result.InsertedID.(primitive.ObjectID)
		callIDs = append(callIDs, call.ID.Hex())

		// Initiate Twilio call
		twilioCall, err := h.twilioService.MakeCall(contact.PhoneNumber, language, call.ID.Hex())
		if err != nil {
			log.Printf("Failed to initiate call for %s: %v", contact.PhoneNumber, err)

			// Update call status to failed
			updateCtx, updateCancel := context.WithTimeout(context.Background(), 5*time.Second)
			h.db.Collection("calls").UpdateOne(
				updateCtx,
				bson.M{"_id": call.ID},
				bson.M{"$set": bson.M{
					"status":        "failed",
					"error_message": err.Error(),
					"updated_at":    time.Now(),
				}},
			)
			updateCancel()

			failCount++
			continue
		}

		// Update call with Twilio SID
		updateCtx, updateCancel := context.WithTimeout(context.Background(), 5*time.Second)
		h.db.Collection("calls").UpdateOne(
			updateCtx,
			bson.M{"_id": call.ID},
			bson.M{"$set": bson.M{
				"status":          "initiated",
				"twilio_call_sid": *twilioCall.Sid,
				"updated_at":      time.Now(),
			}},
		)
		updateCancel()

		// Create call log
		callLog := models.CallLog{
			CallID:    call.ID,
			Event:     "initiated",
			Details:   fmt.Sprintf("Call initiated to %s", contact.PhoneNumber),
			CreatedAt: time.Now(),
		}

		logCtx, logCancel := context.WithTimeout(context.Background(), 5*time.Second)
		h.db.Collection("call_logs").InsertOne(logCtx, callLog)
		logCancel()

		successCount++
	}

	c.JSON(http.StatusOK, gin.H{
		"message":       "Bulk calls initiated",
		"success_count": successCount,
		"fail_count":    failCount,
		"call_ids":      callIDs,
	})
}

// GetCallStatus retrieves the status of a specific call
func (h *CallHandler) GetCallStatus(c *gin.Context) {
	callID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(callID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid call ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var call models.Call
	err = h.db.Collection("calls").FindOne(ctx, bson.M{"_id": objID}).Decode(&call)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Call not found"})
		return
	}

	// Get call logs
	cursor, err := h.db.Collection("call_logs").Find(ctx, bson.M{"call_id": objID})
	if err == nil {
		var callLogs []models.CallLog
		cursor.All(ctx, &callLogs)

		// Return call with logs
		c.JSON(http.StatusOK, gin.H{
			"id":              call.ID,
			"campaign_id":     call.CampaignID,
			"phone_number":    call.PhoneNumber,
			"customer_name":   call.CustomerName,
			"status":          call.Status,
			"twilio_call_sid": call.TwilioCallSID,
			"language":        call.Language,
			"duration":        call.Duration,
			"error_message":   call.ErrorMessage,
			"created_at":      call.CreatedAt,
			"updated_at":      call.UpdatedAt,
			"call_logs":       callLogs,
		})
		return
	}

	c.JSON(http.StatusOK, call)
}

// GetCampaignCalls retrieves all calls for a specific campaign
func (h *CallHandler) GetCampaignCalls(c *gin.Context) {
	campaignID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(campaignID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid campaign ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := h.db.Collection("calls").Find(ctx, bson.M{"campaign_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve calls"})
		return
	}
	defer cursor.Close(ctx)

	var calls []models.Call
	if err = cursor.All(ctx, &calls); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode calls"})
		return
	}

	if calls == nil {
		calls = []models.Call{}
	}

	// Calculate statistics
	total, _ := h.db.Collection("calls").CountDocuments(ctx, bson.M{"campaign_id": objID})
	pending, _ := h.db.Collection("calls").CountDocuments(ctx, bson.M{"campaign_id": objID, "status": "pending"})
	initiated, _ := h.db.Collection("calls").CountDocuments(ctx, bson.M{"campaign_id": objID, "status": "initiated"})
	completed, _ := h.db.Collection("calls").CountDocuments(ctx, bson.M{"campaign_id": objID, "status": "completed"})
	failed, _ := h.db.Collection("calls").CountDocuments(ctx, bson.M{"campaign_id": objID, "status": "failed"})

	c.JSON(http.StatusOK, gin.H{
		"calls": calls,
		"stats": gin.H{
			"total":     total,
			"pending":   pending,
			"initiated": initiated,
			"completed": completed,
			"failed":    failed,
		},
	})
}

// HandleStatusWebhook handles Twilio status callbacks
func (h *CallHandler) HandleStatusWebhook(c *gin.Context) {
	var statusUpdate models.CallStatusUpdate
	if err := c.ShouldBind(&statusUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Find call by Twilio SID
	var call models.Call
	err := h.db.Collection("calls").FindOne(ctx, bson.M{"twilio_call_sid": statusUpdate.CallSid}).Decode(&call)
	if err != nil {
		log.Printf("Call not found for SID: %s", statusUpdate.CallSid)
		c.JSON(http.StatusNotFound, gin.H{"error": "Call not found"})
		return
	}

	// Map Twilio status to our status
	var newStatus string
	switch statusUpdate.CallStatus {
	case "queued", "ringing":
		newStatus = "initiated"
	case "in-progress":
		newStatus = "in-progress"
	case "completed":
		newStatus = "completed"
	case "failed", "busy", "no-answer":
		newStatus = "failed"
	default:
		newStatus = call.Status
	}

	// Update call status
	updateData := bson.M{
		"status":     newStatus,
		"updated_at": time.Now(),
	}

	if statusUpdate.CallDuration != "" {
		if duration, err := strconv.Atoi(statusUpdate.CallDuration); err == nil {
			updateData["duration"] = duration
		}
	}

	h.db.Collection("calls").UpdateOne(ctx, bson.M{"_id": call.ID}, bson.M{"$set": updateData})

	// Create call log
	callLog := models.CallLog{
		CallID:    call.ID,
		Event:     statusUpdate.CallStatus,
		Details:   fmt.Sprintf("Call status: %s", statusUpdate.CallStatus),
		CreatedAt: time.Now(),
	}
	h.db.Collection("call_logs").InsertOne(ctx, callLog)

	c.XML(http.StatusOK, []byte("<Response></Response>"))
}
