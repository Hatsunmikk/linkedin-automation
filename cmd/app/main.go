package main

import (
	"strconv"

	"github.com/joho/godotenv"

	"time"

	"github.com/Hatsunmikk/linkedin-automation/internal/auth"
	"github.com/Hatsunmikk/linkedin-automation/internal/browser"
	"github.com/Hatsunmikk/linkedin-automation/internal/config"
	"github.com/Hatsunmikk/linkedin-automation/internal/connections"
	"github.com/Hatsunmikk/linkedin-automation/internal/logger"
	"github.com/Hatsunmikk/linkedin-automation/internal/messaging"
	"github.com/Hatsunmikk/linkedin-automation/internal/search"
	"github.com/Hatsunmikk/linkedin-automation/internal/state"
	"github.com/Hatsunmikk/linkedin-automation/internal/stealth"
)

// main bootstraps the application by loading configuration,
// initializing core dependencies, and starting the automation workflow.
func main() {
	// Load environment variables from .env file if present.
	// This is useful for local development and testing.
	_ = godotenv.Load()

	log := logger.New(true)

	statePath := "state.json"

	appState, err := state.Load(statePath)
	if err != nil {
		log.Error("Failed to load state: " + err.Error())
		return
	}

	defer func() {
		if err := appState.Save(statePath); err != nil {
			log.Error("Failed to save state: " + err.Error())
		}
	}()

	cfg, err := config.Load()
	if err != nil {
		log.Error("Failed to load configuration")
		return
	}

	// Log configuration values to ensure they are actively
	// used and validated at startup.
	log.Info("Configuration loaded successfully")
	log.Debug("Daily connection limit: " + strconv.Itoa(cfg.DailyConnectionLimit))

	// Initialize rate limiter (human-like pacing)
	limiter := stealth.NewRateLimiter(3, 10*time.Second)

	// Enforce activity scheduling (business hours only)
	if !stealth.IsWithinBusinessHours(9, 18) {
		log.Warn("Outside business hours. Automation will not run.")
		return
	}

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

	authResult, err := auth.Login(page)
	if err != nil {
		log.Error("Authentication error: " + err.Error())
		return
	}

	if authResult.CheckpointHit {
		log.Warn("Security checkpoint detected. Aborting automation.")
		return
	}

	if !authResult.Success {
		log.Warn("Login failed: " + authResult.FailureReason)
		return
	}

	log.Info("Authentication successful")

	// ---- Search & Targeting (PoC-safe) ----
	engine := search.New()

	query := search.Query{
		JobTitle: "Software Engineer",
		Company:  "SubSpace",
		Location: "Bangalore",
		Keywords: []string{"Go", "Backend"},
	}

	results, err := engine.Search(page, query, 2)
	if err != nil {
		log.Error("Search failed: " + err.Error())
		return
	}

	log.Info("Search completed. Profiles found: " + strconv.Itoa(len(results)))

	for _, r := range results {
		log.Debug("Discovered profile: " + r.ProfileURL)
	}

	// ---- Connection Requests (PoC-safe) ----
	connManager := connections.New(cfg.DailyConnectionLimit, appState)

	for _, r := range results {
		if !connManager.CanSend(r.ProfileURL) {
			log.Debug("Skipping already-contacted profile: " + r.ProfileURL)
			continue
		}

		note := connections.BuildPersonalizedNote("there", "SubSpace")

		req, err := connManager.Send(r.ProfileURL, note)
		if err != nil {
			log.Warn("Could not send request: " + err.Error())
			continue
		}

		log.Info("Simulated connection request sent to: " + req.ProfileURL)
		log.Debug("Note content: " + req.Note)
	}

	// ---- Messaging System (PoC-safe) ----
	msgManager := messaging.New(appState)

	for _, r := range results {
		if !msgManager.IsConnectionAccepted(r.ProfileURL) {
			log.Debug("Connection not accepted yet: " + r.ProfileURL)
			continue
		}

		if !msgManager.CanSendMessage(r.ProfileURL) {
			log.Debug("Message already sent to: " + r.ProfileURL)
			continue
		}

		content := messaging.BuildTemplate("there", "SubSpace")

		msg, err := msgManager.SendFollowUp(r.ProfileURL, content)
		if err != nil {
			log.Warn("Could not send follow-up: " + err.Error())
			continue
		}

		log.Info("Simulated follow-up message sent to: " + msg.ProfileURL)
		log.Debug("Message content: " + msg.Content)
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
	limiter.Allow()
	stealth.MoveMouseHumanLike(page, 100, 100, 600, 400)
	log.Debug("Human-like mouse movement executed")

	// Demonstrate natural scrolling behavior
	limiter.Allow()
	stealth.ScrollHumanLike(page, 1200)
	log.Debug("Human-like scrolling executed")

	// Demonstrate human-like typing (no real submission)
	page.MustNavigate("https://example.com")
	stealth.Think(1000, 2000)

	body := page.MustElement("body")
	limiter.Allow()
	stealth.TypeHumanLike(body, "Human-like typing simulation test")
	log.Debug("Human-like typing simulation executed")

	// Demonstrate human-like hover behavior
	el := page.MustElement("h1")
	stealth.HoverHumanLike(page, el)
	log.Debug("Human-like hover behavior executed")

	// ---- State persistence demo (PoC-safe) ----
	// This simulates tracking of automation actions
	// without interacting with real LinkedIn profiles.

	testProfile := "https://linkedin.com/in/example-profile"

	if _, exists := appState.SentRequests[testProfile]; !exists {
		log.Info("Recording sent connection request")
		appState.MarkRequestSent(testProfile)
	}

	if _, exists := appState.SentMessages[testProfile]; !exists {
		log.Info("Recording sent message")
		appState.MarkMessageSent(testProfile)
	}

}
