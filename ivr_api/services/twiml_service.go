package services

import (
	"fmt"
	"strings"

	"github.com/prabhatkumar/ivrcalling/models"
)

// TwiMLGenerator generates TwiML responses for IVR
type TwiMLGenerator struct {
	language string
	strings  LanguageStrings
}

func NewTwiMLGenerator(language string) *TwiMLGenerator {
	return &TwiMLGenerator{
		language: language,
		strings:  GetLanguageStrings(language),
	}
}

// GenerateDynamicWelcome generates the welcome message TwiML with campaign intro and actions
func (g *TwiMLGenerator) GenerateDynamicWelcome(customerName string, campaign *models.Campaign) string {
	greeting := fmt.Sprintf(g.strings.Welcome, customerName)
	if customerName == "" {
		greeting = strings.Replace(g.strings.Welcome, "%s, ", "", 1)
	}

	// Build intro text
	introText := campaign.IntroText
	if introText == "" {
		introText = g.strings.MainMenu
	}

	// Build menu from actions
	menuText := g.buildMenuFromActions(campaign.Actions)

	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice" language="%s">%s</Say>
    <Say voice="alice" language="%s">%s</Say>
    <Gather action="/api/webhook/gather" method="POST" numDigits="1" timeout="5">
        <Say voice="alice" language="%s">%s</Say>
    </Gather>
    <Say voice="alice" language="%s">%s</Say>
    <Redirect>/api/webhook/voice</Redirect>
</Response>`,
		g.getVoiceLanguage(), greeting,
		g.getVoiceLanguage(), introText,
		g.getVoiceLanguage(), menuText,
		g.getVoiceLanguage(), g.strings.InvalidInput)
}

// GenerateDynamicResponse generates TwiML based on action configuration
func (g *TwiMLGenerator) GenerateDynamicResponse(action *models.IVRAction, campaign *models.Campaign) string {
	if action.ActionType == "forward" {
		return g.GenerateForward(action.ForwardPhone, action.Message)
	}

	// Information type - check if message is URL or text
	message := action.Message
	if message == "" {
		message = g.strings.InvalidInput
	}

	// Check if message is a URL (starts with http:// or https://)
	if strings.HasPrefix(message, "http://") || strings.HasPrefix(message, "https://") {
		return g.GeneratePlayAudio(message, campaign)
	}

	// Otherwise, use text-to-speech
	return g.GenerateTextToSpeech(message, campaign)
}

// GeneratePlayAudio generates TwiML to play audio file
func (g *TwiMLGenerator) GeneratePlayAudio(audioURL string, campaign *models.Campaign) string {
	menuText := g.buildMenuFromActions(campaign.Actions)

	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Play>%s</Play>
    <Gather action="/api/webhook/gather" method="POST" numDigits="1" timeout="5">
        <Say voice="alice" language="%s">%s</Say>
    </Gather>
    <Redirect>/api/webhook/voice</Redirect>
</Response>`,
		audioURL,
		g.getVoiceLanguage(), menuText)
}

// GenerateTextToSpeech generates TwiML for text-to-speech
func (g *TwiMLGenerator) GenerateTextToSpeech(message string, campaign *models.Campaign) string {
	menuText := g.buildMenuFromActions(campaign.Actions)

	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice" language="%s">%s</Say>
    <Gather action="/api/webhook/gather" method="POST" numDigits="1" timeout="5">
        <Say voice="alice" language="%s">%s</Say>
    </Gather>
    <Redirect>/api/webhook/voice</Redirect>
</Response>`,
		g.getVoiceLanguage(), message,
		g.getVoiceLanguage(), menuText)
}

// GenerateForward generates TwiML to forward call to another number
func (g *TwiMLGenerator) GenerateForward(phoneNumber string, message string) string {
	if message != "" {
		return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice" language="%s">%s</Say>
    <Dial>%s</Dial>
    <Say voice="alice" language="%s">%s</Say>
    <Hangup/>
</Response>`,
			g.getVoiceLanguage(), message,
			phoneNumber,
			g.getVoiceLanguage(), g.strings.Goodbye)
	}

	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice" language="%s">Forwarding your call. Please wait.</Say>
    <Dial>%s</Dial>
    <Say voice="alice" language="%s">%s</Say>
    <Hangup/>
</Response>`,
		g.getVoiceLanguage(),
		phoneNumber,
		g.getVoiceLanguage(), g.strings.Goodbye)
}

// buildMenuFromActions creates menu text from campaign actions
func (g *TwiMLGenerator) buildMenuFromActions(actions []models.IVRAction) string {
	if len(actions) == 0 {
		// Return a simple prompt when no actions are defined
		return "Press 0 to hear this message again"
	}

	var menuParts []string
	for _, action := range actions {
		var actionDesc string
		if action.ActionType == "forward" {
			actionDesc = fmt.Sprintf("Press %s to speak with an agent", action.ActionInput)
		} else {
			// Use first few words of message as description
			words := strings.Fields(action.Message)
			if len(words) > 0 {
				desc := strings.Join(words[:min(5, len(words))], " ")
				if len(words) > 5 {
					desc += "..."
				}
				actionDesc = fmt.Sprintf("Press %s for %s", action.ActionInput, desc)
			} else {
				actionDesc = fmt.Sprintf("Press %s for more information", action.ActionInput)
			}
		}
		menuParts = append(menuParts, actionDesc)
	}

	// Add option to return to main menu or repeat
	menuParts = append(menuParts, "Press 0 to repeat this menu")

	return strings.Join(menuParts, ". ")
}

// Helper function to get minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// GenerateWelcome generates the welcome message TwiML
func (g *TwiMLGenerator) GenerateWelcome(customerName string) string {
	greeting := fmt.Sprintf(g.strings.Welcome, customerName)
	if customerName == "" {
		greeting = strings.Replace(g.strings.Welcome, "%s, ", "", 1)
	}

	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice" language="%s">%s</Say>
    <Say voice="alice" language="%s">%s</Say>
    <Gather action="/api/webhook/gather" method="POST" numDigits="1" timeout="5">
        <Say voice="alice" language="%s">%s</Say>
    </Gather>
    <Say voice="alice" language="%s">%s</Say>
    <Redirect>/api/webhook/voice</Redirect>
</Response>`,
		g.getVoiceLanguage(), greeting,
		g.getVoiceLanguage(), g.strings.MainMenu,
		g.getVoiceLanguage(), g.strings.PressToRepeat,
		g.getVoiceLanguage(), g.strings.InvalidInput)
}

// GenerateMainMenu generates the main menu TwiML
func (g *TwiMLGenerator) GenerateMainMenu() string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Gather action="/api/webhook/gather" method="POST" numDigits="1" timeout="5">
        <Say voice="alice" language="%s">%s</Say>
    </Gather>
    <Say voice="alice" language="%s">%s</Say>
    <Redirect>/api/webhook/voice</Redirect>
</Response>`,
		g.getVoiceLanguage(), g.strings.MainMenu,
		g.getVoiceLanguage(), g.strings.InvalidInput)
}

// GenerateProductInfo generates product information TwiML
func (g *TwiMLGenerator) GenerateProductInfo() string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice" language="%s">%s</Say>
    <Gather action="/api/webhook/gather" method="POST" numDigits="1" timeout="5">
        <Say voice="alice" language="%s">%s</Say>
    </Gather>
    <Redirect>/api/webhook/voice</Redirect>
</Response>`,
		g.getVoiceLanguage(), g.strings.ProductInfo,
		g.getVoiceLanguage(), g.strings.PressForInfo)
}

// GenerateOfferDetails generates offer details TwiML
func (g *TwiMLGenerator) GenerateOfferDetails() string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice" language="%s">%s</Say>
    <Gather action="/api/webhook/gather" method="POST" numDigits="1" timeout="5">
        <Say voice="alice" language="%s">%s</Say>
    </Gather>
    <Redirect>/api/webhook/voice</Redirect>
</Response>`,
		g.getVoiceLanguage(), g.strings.OfferDetails,
		g.getVoiceLanguage(), g.strings.PressForInfo)
}

// GenerateOptOut generates opt-out confirmation TwiML
func (g *TwiMLGenerator) GenerateOptOut() string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Gather action="/api/webhook/gather" method="POST" numDigits="1" timeout="5">
        <Say voice="alice" language="%s">%s</Say>
    </Gather>
    <Redirect>/api/webhook/voice</Redirect>
</Response>`,
		g.getVoiceLanguage(), g.strings.PressToOptOut)
}

// GenerateOptOutConfirm generates opt-out final confirmation TwiML
func (g *TwiMLGenerator) GenerateOptOutConfirm() string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice" language="%s">%s</Say>
    <Say voice="alice" language="%s">%s</Say>
    <Hangup/>
</Response>`,
		g.getVoiceLanguage(), g.strings.OptOutConfirm,
		g.getVoiceLanguage(), g.strings.Goodbye)
}

// GenerateGoodbye generates goodbye message TwiML
func (g *TwiMLGenerator) GenerateGoodbye() string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice" language="%s">%s</Say>
    <Say voice="alice" language="%s">%s</Say>
    <Hangup/>
</Response>`,
		g.getVoiceLanguage(), g.strings.ThankYou,
		g.getVoiceLanguage(), g.strings.Goodbye)
}

// GenerateInvalidInput generates invalid input message TwiML
func (g *TwiMLGenerator) GenerateInvalidInput() string {
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<Response>
    <Say voice="alice" language="%s">%s</Say>
    <Redirect>/api/webhook/voice</Redirect>
</Response>`,
		g.getVoiceLanguage(), g.strings.InvalidInput)
}

// getVoiceLanguage maps language codes to Twilio voice language codes
func (g *TwiMLGenerator) getVoiceLanguage() string {
	voiceMap := map[string]string{
		"en": "en-US",
		"es": "es-ES",
		"fr": "fr-FR",
		"de": "de-DE",
		"hi": "hi-IN",
	}

	if voice, ok := voiceMap[g.language]; ok {
		return voice
	}
	return "en-US" // default
}
