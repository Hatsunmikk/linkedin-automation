package main

import (
	"strconv"

	"github.com/joho/godotenv"

	"github.com/Hatsunmikk/linkedin-automation/internal/config"
	"github.com/Hatsunmikk/linkedin-automation/internal/logger"
)

// main bootstraps the application by loading configuration,
// initializing core dependencies, and starting the automation workflow.
func main() {
	// Load environment variables from .env file if present.
	// This is useful for local development and testing.
	_ = godotenv.Load()

	log := logger.New(true)

	cfg, err := config.Load()
	if err != nil {
		log.Error("Failed to load configuration")
		return
	}

	// Log configuration values to ensure they are actively
	// used and validated at startup.
	log.Info("Configuration loaded successfully")
	log.Debug("Daily connection limit: " + strconv.Itoa(cfg.DailyConnectionLimit))
}
