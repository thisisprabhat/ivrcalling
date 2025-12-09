package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qandi/ivr-calling-api/internal/service"
)

type TwiMLHandler struct {
	twilioService *service.TwilioService
}

func NewTwiMLHandler(twilioService *service.TwilioService) *TwiMLHandler {
	return &TwiMLHandler{
		twilioService: twilioService,
	}
}

// WelcomeMessage godoc
// @Summary TwiML Welcome Message
// @Description Generates TwiML for the welcome message and menu options
// @Tags twiml
// @Produce xml
// @Success 200 {string} string "TwiML XML response"
// @Router /api/v1/twiml/welcome [get]
// @Router /api/v1/twiml/welcome [post]
func (h *TwiMLHandler) WelcomeMessage(c *gin.Context) {
	twiml := h.twilioService.GenerateWelcomeTwiML()
	c.Header("Content-Type", "application/xml")
	c.String(http.StatusOK, twiml)
}

// HandleInput godoc
// @Summary TwiML Handle User Input
// @Description Processes user's digit input and generates appropriate TwiML response
// @Tags twiml
// @Accept application/x-www-form-urlencoded
// @Produce xml
// @Param Digits formData string false "Digit pressed by user"
// @Success 200 {string} string "TwiML XML response"
// @Router /api/v1/twiml/handle-input [post]
func (h *TwiMLHandler) HandleInput(c *gin.Context) {
	digit := c.PostForm("Digits")

	twiml := h.twilioService.GenerateHandleInputTwiML(digit)
	c.Header("Content-Type", "application/xml")
	c.String(http.StatusOK, twiml)
}
