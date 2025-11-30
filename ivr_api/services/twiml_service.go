package services

import (
	"fmt"
	"strings"
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
