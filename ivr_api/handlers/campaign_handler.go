package handlers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prabhatkumar/ivrcalling/database"
	"github.com/prabhatkumar/ivrcalling/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CampaignHandler struct {
	db *database.MongoDB
}

func NewCampaignHandler(db *database.MongoDB) *CampaignHandler {
	return &CampaignHandler{db: db}
}

// CreateCampaign creates a new campaign
func (h *CampaignHandler) CreateCampaign(c *gin.Context) {
	var campaign models.Campaign
	if err := c.ShouldBindJSON(&campaign); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Printf("=== CREATING CAMPAIGN ===")
	log.Printf("Name: %s", campaign.Name)
	log.Printf("Description: %s", campaign.Description)
	log.Printf("Language: %s", campaign.Language)
	log.Printf("IntroText: %s", campaign.IntroText)
	log.Printf("Actions count: %d", len(campaign.Actions))
	for i, action := range campaign.Actions {
		log.Printf("  Action %d: Type=%s, Input=%s, Message=%s, Phone=%s",
			i+1, action.ActionType, action.ActionInput, action.Message, action.ForwardPhone)
	}

	// Validate required fields
	if campaign.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Campaign name is required"})
		return
	}

	if campaign.Description == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Description is required"})
		return
	}

	if campaign.IntroText == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Intro text is required"})
		return
	}

	// Set timestamps
	campaign.CreatedAt = time.Now()
	campaign.UpdatedAt = time.Now()

	// Set defaults
	if campaign.Language == "" {
		campaign.Language = "en"
	}

	// Initialize actions array if nil
	if campaign.Actions == nil {
		campaign.Actions = []models.IVRAction{}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := h.db.Collection("campaigns").InsertOne(ctx, campaign)
	if err != nil {
		log.Printf("Failed to insert campaign into database: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create campaign"})
		return
	}

	campaign.ID = result.InsertedID.(primitive.ObjectID)
	log.Printf("✓ Campaign created successfully with ID: %s", campaign.ID.Hex())
	c.JSON(http.StatusCreated, campaign)
}

// GetCampaign retrieves a specific campaign
func (h *CampaignHandler) GetCampaign(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid campaign ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var campaign models.Campaign
	err = h.db.Collection("campaigns").FindOne(ctx, bson.M{"_id": objID}).Decode(&campaign)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}

	c.JSON(http.StatusOK, campaign)
}

// ListCampaigns retrieves all campaigns
func (h *CampaignHandler) ListCampaigns(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := h.db.Collection("campaigns").Find(ctx, bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve campaigns"})
		return
	}
	defer cursor.Close(ctx)

	var campaigns []models.Campaign
	if err = cursor.All(ctx, &campaigns); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode campaigns"})
		return
	}

	if campaigns == nil {
		campaigns = []models.Campaign{}
	}

	c.JSON(http.StatusOK, campaigns)
}

// UpdateCampaign updates an existing campaign
func (h *CampaignHandler) UpdateCampaign(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid campaign ID"})
		return
	}

	var updateData map[string]interface{}
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add updated timestamp
	updateData["updated_at"] = time.Now()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := h.db.Collection("campaigns").UpdateOne(
		ctx,
		bson.M{"_id": objID},
		bson.M{"$set": updateData},
	)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update campaign"})
		return
	}

	if result.MatchedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}

	// Fetch updated campaign
	var campaign models.Campaign
	err = h.db.Collection("campaigns").FindOne(ctx, bson.M{"_id": objID}).Decode(&campaign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated campaign"})
		return
	}

	c.JSON(http.StatusOK, campaign)
}

// DeleteCampaign deletes a campaign
func (h *CampaignHandler) DeleteCampaign(c *gin.Context) {
	id := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid campaign ID"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := h.db.Collection("campaigns").DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete campaign"})
		return
	}

	if result.DeletedCount == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Campaign not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Campaign deleted successfully"})
}
