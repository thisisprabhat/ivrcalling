package routes

import (
"github.com/gin-gonic/gin"
"github.com/prabhatkumar/ivrcalling/config"
"github.com/prabhatkumar/ivrcalling/database"
"github.com/prabhatkumar/ivrcalling/handlers"
"github.com/prabhatkumar/ivrcalling/services"
)

func SetupRoutes(router *gin.Engine, db *database.MongoDB, cfg *config.Config) {
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
}
