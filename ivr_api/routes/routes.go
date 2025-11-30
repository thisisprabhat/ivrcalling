package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/prabhatkumar/ivrcalling/config"
	"github.com/prabhatkumar/ivrcalling/database"
	"github.com/prabhatkumar/ivrcalling/handlers"
	"github.com/prabhatkumar/ivrcalling/services"
)

func SetupRoutes(router *gin.Engine, db *database.MongoDB, cfg *config.Config) {
	// Enable CORS for frontend
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	twilioService := services.NewTwilioService(cfg)
	campaignHandler := handlers.NewCampaignHandler(db)
	callHandler := handlers.NewCallHandler(db, twilioService)
	webhookHandler := handlers.NewWebhookHandler(db)

	api := router.Group("/api")
	{
		campaigns := api.Group("/campaigns")
		{
			campaigns.POST("", campaignHandler.CreateCampaign)
			campaigns.GET("", campaignHandler.ListCampaigns)
			campaigns.GET("/:id", campaignHandler.GetCampaign)
			campaigns.PUT("/:id", campaignHandler.UpdateCampaign)
			campaigns.DELETE("/:id", campaignHandler.DeleteCampaign)
			campaigns.GET("/:id/calls", callHandler.GetCampaignCalls)
		}

		calls := api.Group("/calls")
		{
			calls.POST("/bulk", callHandler.InitiateBulkCalls)
			calls.GET("/:id", callHandler.GetCallStatus)
		}

		webhook := api.Group("/webhook")
		{
			webhook.POST("/voice", webhookHandler.HandleVoiceWebhook)
			webhook.POST("/gather", webhookHandler.HandleGatherWebhook)
			webhook.POST("/status", callHandler.HandleStatusWebhook)
			webhook.POST("/optout", webhookHandler.HandleOptOutConfirm)
		}

		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "healthy",
				"service": "IVR Calling System",
			})
		})

		api.GET("/languages", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"languages": services.GetSupportedLanguages(),
			})
		})
	}

	// API Documentation endpoints
	router.GET("/docs", handlers.DocsHandler())
	router.GET("/docs/swagger.yaml", func(c *gin.Context) {
		c.File("./docs/swagger.yaml")
	})
}
