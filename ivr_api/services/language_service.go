package services

// LanguageStrings contains all IVR messages for different languages
type LanguageStrings struct {
	Welcome         string
	MainMenu        string
	PressForInfo    string
	PressToOptOut   string
	PressToRepeat   string
	ThankYou        string
	Goodbye         string
	InvalidInput    string
	ProductInfo     string
	OfferDetails    string
	OptOutConfirm   string
	TransferMessage string
}

var Languages = map[string]LanguageStrings{
	"en": { // English
		Welcome:         "Hello %s, welcome to our marketing campaign.",
		MainMenu:        "Press 1 for product information. Press 2 for special offers. Press 3 to opt out. Press 9 to repeat this menu.",
		PressForInfo:    "Press 1 to hear more details or press 0 to return to the main menu.",
		PressToOptOut:   "Press 1 to confirm opt out or press 0 to return to the main menu.",
		PressToRepeat:   "Press 9 to repeat the menu.",
		ThankYou:        "Thank you for your time.",
		Goodbye:         "Goodbye!",
		InvalidInput:    "Sorry, that was not a valid option. Please try again.",
		ProductInfo:     "Our new product offers cutting-edge features designed to improve your daily life. It includes advanced technology and user-friendly design.",
		OfferDetails:    "Special offer! Get 20 percent off on your first purchase. Use promo code WELCOME20 at checkout. Offer valid for 30 days.",
		OptOutConfirm:   "You have been successfully removed from our calling list. We apologize for any inconvenience.",
		TransferMessage: "Please hold while we transfer you to our customer service representative.",
	},
	"es": { // Spanish
		Welcome:         "Hola %s, bienvenido a nuestra campaña de marketing.",
		MainMenu:        "Presione 1 para información del producto. Presione 2 para ofertas especiales. Presione 3 para excluirse. Presione 9 para repetir este menú.",
		PressForInfo:    "Presione 1 para escuchar más detalles o presione 0 para volver al menú principal.",
		PressToOptOut:   "Presione 1 para confirmar la exclusión o presione 0 para volver al menú principal.",
		PressToRepeat:   "Presione 9 para repetir el menú.",
		ThankYou:        "Gracias por su tiempo.",
		Goodbye:         "¡Adiós!",
		InvalidInput:    "Lo sentimos, esa no fue una opción válida. Por favor, inténtelo de nuevo.",
		ProductInfo:     "Nuestro nuevo producto ofrece características de vanguardia diseñadas para mejorar su vida diaria. Incluye tecnología avanzada y diseño fácil de usar.",
		OfferDetails:    "¡Oferta especial! Obtenga un 20 por ciento de descuento en su primera compra. Use el código promocional WELCOME20 al pagar. Oferta válida por 30 días.",
		OptOutConfirm:   "Ha sido eliminado exitosamente de nuestra lista de llamadas. Nos disculpamos por cualquier inconveniente.",
		TransferMessage: "Por favor espere mientras lo transferimos a nuestro representante de servicio al cliente.",
	},
	"fr": { // French
		Welcome:         "Bonjour %s, bienvenue dans notre campagne marketing.",
		MainMenu:        "Appuyez sur 1 pour les informations sur le produit. Appuyez sur 2 pour les offres spéciales. Appuyez sur 3 pour vous désinscrire. Appuyez sur 9 pour répéter ce menu.",
		PressForInfo:    "Appuyez sur 1 pour entendre plus de détails ou appuyez sur 0 pour revenir au menu principal.",
		PressToOptOut:   "Appuyez sur 1 pour confirmer la désinscription ou appuyez sur 0 pour revenir au menu principal.",
		PressToRepeat:   "Appuyez sur 9 pour répéter le menu.",
		ThankYou:        "Merci pour votre temps.",
		Goodbye:         "Au revoir!",
		InvalidInput:    "Désolé, ce n'était pas une option valide. Veuillez réessayer.",
		ProductInfo:     "Notre nouveau produit offre des fonctionnalités de pointe conçues pour améliorer votre vie quotidienne. Il comprend une technologie avancée et une conception conviviale.",
		OfferDetails:    "Offre spéciale! Obtenez 20 pour cent de réduction sur votre premier achat. Utilisez le code promo WELCOME20 lors du paiement. Offre valable pendant 30 jours.",
		OptOutConfirm:   "Vous avez été supprimé avec succès de notre liste d'appels. Nous nous excusons pour tout inconvénient.",
		TransferMessage: "Veuillez patienter pendant que nous vous transférons à notre représentant du service client.",
	},
	"de": { // German
		Welcome:         "Hallo %s, willkommen zu unserer Marketingkampagne.",
		MainMenu:        "Drücken Sie 1 für Produktinformationen. Drücken Sie 2 für Sonderangebote. Drücken Sie 3, um sich abzumelden. Drücken Sie 9, um dieses Menü zu wiederholen.",
		PressForInfo:    "Drücken Sie 1, um weitere Details zu hören, oder drücken Sie 0, um zum Hauptmenü zurückzukehren.",
		PressToOptOut:   "Drücken Sie 1, um die Abmeldung zu bestätigen, oder drücken Sie 0, um zum Hauptmenü zurückzukehren.",
		PressToRepeat:   "Drücken Sie 9, um das Menü zu wiederholen.",
		ThankYou:        "Vielen Dank für Ihre Zeit.",
		Goodbye:         "Auf Wiedersehen!",
		InvalidInput:    "Entschuldigung, das war keine gültige Option. Bitte versuchen Sie es erneut.",
		ProductInfo:     "Unser neues Produkt bietet modernste Funktionen, die entwickelt wurden, um Ihr tägliches Leben zu verbessern. Es umfasst fortschrittliche Technologie und benutzerfreundliches Design.",
		OfferDetails:    "Sonderangebot! Erhalten Sie 20 Prozent Rabatt auf Ihren ersten Einkauf. Verwenden Sie den Promocode WELCOME20 beim Bezahlen. Angebot gültig für 30 Tage.",
		OptOutConfirm:   "Sie wurden erfolgreich von unserer Anrufliste entfernt. Wir entschuldigen uns für etwaige Unannehmlichkeiten.",
		TransferMessage: "Bitte warten Sie, während wir Sie zu unserem Kundendienstmitarbeiter verbinden.",
	},
	"hi": { // Hindi
		Welcome:         "नमस्ते %s, हमारे मार्केटिंग अभियान में आपका स्वागत है।",
		MainMenu:        "उत्पाद जानकारी के लिए 1 दबाएं। विशेष ऑफर के लिए 2 दबाएं। ऑप्ट आउट करने के लिए 3 दबाएं। इस मेनू को दोहराने के लिए 9 दबाएं।",
		PressForInfo:    "अधिक विवरण सुनने के लिए 1 दबाएं या मुख्य मेनू पर वापस जाने के लिए 0 दबाएं।",
		PressToOptOut:   "ऑप्ट आउट की पुष्टि करने के लिए 1 दबाएं या मुख्य मेनू पर वापस जाने के लिए 0 दबाएं।",
		PressToRepeat:   "मेनू को दोहराने के लिए 9 दबाएं।",
		ThankYou:        "आपके समय के लिए धन्यवाद।",
		Goodbye:         "अलविदा!",
		InvalidInput:    "क्षमा करें, यह एक वैध विकल्प नहीं था। कृपया पुनः प्रयास करें।",
		ProductInfo:     "हमारा नया उत्पाद आपके दैनिक जीवन को बेहतर बनाने के लिए डिज़ाइन की गई अत्याधुनिक सुविधाएं प्रदान करता है। इसमें उन्नत तकनीक और उपयोगकर्ता के अनुकूल डिज़ाइन शामिल है।",
		OfferDetails:    "विशेष ऑफर! अपनी पहली खरीदारी पर 20 प्रतिशत की छूट पाएं। चेकआउट पर प्रोमो कोड WELCOME20 का उपयोग करें। ऑफर 30 दिनों के लिए वैध है।",
		OptOutConfirm:   "आपको हमारी कॉलिंग सूची से सफलतापूर्वक हटा दिया गया है। किसी भी असुविधा के लिए हम क्षमा चाहते हैं।",
		TransferMessage: "कृपया प्रतीक्षा करें जब तक हम आपको हमारे ग्राहक सेवा प्रतिनिधि से स्थानांतरित करते हैं।",
	},
}

// GetLanguageStrings returns the language strings for the given language code
func GetLanguageStrings(lang string) LanguageStrings {
	if strings, ok := Languages[lang]; ok {
		return strings
	}
	// Default to English if language not found
	return Languages["en"]
}

// GetSupportedLanguages returns a list of all supported language codes
func GetSupportedLanguages() []string {
	langs := make([]string, 0, len(Languages))
	for lang := range Languages {
		langs = append(langs, lang)
	}
	return langs
}
