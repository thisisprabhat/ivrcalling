package api

import (
	"github.com/gin-gonic/gin"
	"github.com/qandi/ivr-calling-api/internal/handlers"
)

func SetupRoutes(router *gin.Engine, callHandler *handlers.CallHandler, twimlHandler *handlers.TwiMLHandler) {
	// Health check
	router.GET("/health", callHandler.HealthCheck)

	// API v1 routes
	v1 := router.Group("/api/v1")
	{
		// Call routes
		calls := v1.Group("/calls")
		{
			calls.POST("/initiate", callHandler.InitiateCall)
		}

		// Callback routes
		callbacks := v1.Group("/callbacks")
		{
			callbacks.POST("/ivr", callHandler.HandleCallback)
		}

		// Configuration routes
		config := v1.Group("/config")
		{
			config.GET("/ivr", callHandler.GetIVRConfig)
		}

		// TwiML routes (for Twilio to call)
		twiml := v1.Group("/twiml")
		{
			twiml.GET("/welcome", twimlHandler.WelcomeMessage)
			twiml.POST("/welcome", twimlHandler.WelcomeMessage)
			twiml.POST("/handle-input", twimlHandler.HandleInput)
		}
	}
}
