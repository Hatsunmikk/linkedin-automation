package main

import (
	"strconv"

	"github.com/joho/godotenv"

	"github.com/Hatsunmikk/linkedin-automation/internal/browser"
	"github.com/Hatsunmikk/linkedin-automation/internal/config"
	"github.com/Hatsunmikk/linkedin-automation/internal/logger"
	"github.com/Hatsunmikk/linkedin-automation/internal/stealth"
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

	//Initialize the browser with headless configuration
	br, err := browser.New(cfg.Headless, log)
	if err != nil {
		log.Error("Failed to launch browser")
		return
	}
	defer br.Close()

	log.Info("Browser ready for automation")

	page, err := br.NewPage()
	if err != nil {
		log.Error("Failed to create browser page")
		return
	}

	if err := stealth.ApplyFingerprintMask(page); err != nil {
		log.Error("Failed to apply fingerprint masking")
		return
	}

	log.Info("Stealth fingerprint masking applied")

	// Simulate human pause before further actions
	stealth.Think(800, 1800)
	log.Debug("Human-like think time applied")

	// Demonstrate human-like mouse movement
	stealth.MoveMouseHumanLike(page, 100, 100, 600, 400)
	log.Debug("Human-like mouse movement executed")

	// Demonstrate natural scrolling behavior
	stealth.ScrollHumanLike(page, 1200)
	log.Debug("Human-like scrolling executed")

	// Demonstrate human-like typing (no real submission)
	page.MustNavigate("https://example.com")
	stealth.Think(1000, 2000)

	body := page.MustElement("body")
	stealth.TypeHumanLike(body, "Human-like typing simulation test")
	log.Debug("Human-like typing simulation executed")

}
