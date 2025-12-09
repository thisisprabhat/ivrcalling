package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qandi/ivr-calling-api/internal/models"
	"github.com/qandi/ivr-calling-api/internal/service"
)

type CallHandler struct {
	twilioService *service.TwilioService
}

func NewCallHandler(twilioService *service.TwilioService) *CallHandler {
	return &CallHandler{
		twilioService: twilioService,
	}
}

// InitiateCall godoc
// @Summary Initiate an IVR call
// @Description Initiates an outbound IVR call to the specified phone number with Q&I information
// @Tags calls
// @Accept json
// @Produce json
// @Param request body models.CallRequest true "Call Request"
// @Success 200 {object} models.CallResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/v1/calls/initiate [post]
func (h *CallHandler) InitiateCall(c *gin.Context) {
	var req models.CallRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid request",
			Message: err.Error(),
		})
		return
	}

	response, err := h.twilioService.InitiateCall(req.PhoneNumber, req.CallbackURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Failed to initiate call",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, response)
}

// HandleCallback godoc
// @Summary Handle IVR callback
// @Description Handles callbacks from the IVR provider (call events, digit inputs, etc.)
// @Tags callbacks
// @Accept json
// @Produce json
// @Param callback body models.CallbackRequest true "Callback Data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} models.ErrorResponse
// @Router /api/v1/callbacks/ivr [post]
func (h *CallHandler) HandleCallback(c *gin.Context) {
	var callback models.CallbackRequest

	if err := c.ShouldBindJSON(&callback); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid callback data",
			Message: err.Error(),
		})
		return
	}

	err := h.twilioService.HandleCallback(&callback)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Error:   "Failed to process callback",
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Callback processed successfully",
	})
}

// GetIVRConfig godoc
// @Summary Get IVR configuration
// @Description Returns the current IVR flow configuration including intro text, actions, and messages
// @Tags configuration
// @Produce json
// @Success 200 {object} models.IVRConfig
// @Router /api/v1/config/ivr [get]
func (h *CallHandler) GetIVRConfig(c *gin.Context) {
	config := h.twilioService.GetIVRConfig()
	c.JSON(http.StatusOK, config)
}

// HealthCheck godoc
// @Summary Health check
// @Description Returns the health status of the API
// @Tags health
// @Produce json
// @Success 200 {object} models.HealthResponse
// @Router /health [get]
func (h *CallHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, models.HealthResponse{
		Status:  "healthy",
		Version: "1.0.0",
	})
}
