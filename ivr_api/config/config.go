package config

import (
	"os"
)

type Config struct {
	Port              string
	Environment       string
	TwilioAccountSID  string
	TwilioAuthToken   string
	TwilioPhoneNumber string
	MongoDBURI        string
	MongoDBDatabase   string
	DefaultLanguage   string
	WebhookBaseURL    string
}

func LoadConfig() *Config {
	return &Config{
		Port:              getEnv("PORT", "8080"),
		Environment:       getEnv("ENV", "development"),
		TwilioAccountSID:  getEnv("TWILIO_ACCOUNT_SID", ""),
		TwilioAuthToken:   getEnv("TWILIO_AUTH_TOKEN", ""),
		TwilioPhoneNumber: getEnv("TWILIO_PHONE_NUMBER", ""),
		MongoDBURI:        getEnv("MONGODB_URI", "mongodb://localhost:27017"),
		MongoDBDatabase:   getEnv("MONGODB_DATABASE", "ivr_calling_system"),
		DefaultLanguage:   getEnv("DEFAULT_LANGUAGE", "en"),
		WebhookBaseURL:    getEnv("WEBHOOK_BASE_URL", "http://localhost:8080"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
