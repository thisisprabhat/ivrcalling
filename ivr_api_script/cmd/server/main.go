package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/qandi/ivr-calling-api/internal/api"
	"github.com/qandi/ivr-calling-api/internal/config"
	"github.com/qandi/ivr-calling-api/internal/handlers"
	"github.com/qandi/ivr-calling-api/internal/service"
)

// @title Q&I IVR Calling API
// @version 1.0
// @description API for initiating IVR calls with Q&I educational platform information
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email support@qandi.com

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @schemes http https

func main() {
	// Load configuration
	config.LoadConfig()

	// Set Gin mode
	gin.SetMode(config.AppConfig.GinMode)

	// Initialize services
	twilioService := service.NewTwilioService(config.AppConfig)

	// Initialize handlers
	callHandler := handlers.NewCallHandler(twilioService)
	twimlHandler := handlers.NewTwiMLHandler(twilioService)

	// Setup router
	router := gin.Default()

	// Add CORS middleware
	router.Use(corsMiddleware())

	// Setup routes
	api.SetupRoutes(router, callHandler, twimlHandler)

	// Start server
	addr := ":" + config.AppConfig.Port
	log.Printf("Starting Q&I IVR API server on %s", addr)
	log.Printf("Twilio webhooks will be available at: %s/api/v1/twiml/welcome", config.AppConfig.ServerBaseURL)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
