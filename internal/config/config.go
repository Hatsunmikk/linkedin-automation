package config

import (
	"os"
	"strconv"
)

//Config hold all runtime configuration for the application.
//Centralizing configuration makes the system easier to reason about,
//test, and modify without touching business logic.

type Config struct {
	LinkedInEmail    string
	LinkedInPassword string

	Debug    bool
	Headless bool

	DailyConnectionLimit int
}

// Load reads configuration from environment variables
// and constructs a validated Config struct.
func Load() (*Config, error) {
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	headless, _ := strconv.ParseBool(os.Getenv("HEADLESS"))

	dailyLimit, err := strconv.Atoi(os.Getenv("DAILY_CONNECTION_LIMIT"))
	if err != nil {
		dailyLimit = 10 //sensible default
	}

	return &Config{
		LinkedInEmail:        os.Getenv("LINKEDIN_EMAIL"),
		LinkedInPassword:     os.Getenv("LINKEDIN_PASSWORD"),
		Debug:                debug,
		Headless:             headless,
		DailyConnectionLimit: dailyLimit,
	}, nil
}
