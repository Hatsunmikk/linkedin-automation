package browser

import (
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"

	"github.com/Hatsunmikk/linkedin-automation/internal/logger"
)

// Browser encapsulates the Rod browser instance.
// This abstraction centralizes browser lifecycle management
// and allows us to inject stealth configurations consistently.
type Browser struct {
	rodBrowser *rod.Browser
	log        *logger.Logger
}

// New creates and launches a new Chromium browser isntance
// Headless mode is configurable to support both development
// (visible browser) and automated execution.
func New(headless bool, log *logger.Logger) (*Browser, error) {
	launch := launcher.New().
		// These flags improve stability on Windows environments
		// and reduce common Chromium startup failures.
		Set("disable-gpu", "true").
		Set("no-sandbox", "true").
		Set("disable-dev-shm-usage", "true")

	if headless {
		launch = launch.Headless(true)
	} else {
		launch = launch.Headless(false)
	}

	// Launch Chromium and obtain the control URL.
	url, err := launch.Launch()
	if err != nil {
		log.Error("Chromium launch failed: " + err.Error())
		return nil, err
	}

	// Connect Rod to the launched browser instance.
	rb := rod.New().ControlURL(url)
	if err := rb.Connect(); err != nil {
		log.Error("Failed to connect to Chromium: " + err.Error())
		return nil, err
	}

	log.Info("Chromium browser launched successfully")

	return &Browser{
		rodBrowser: rb,
		log:        log,
	}, nil
}

// Close shuts down the borwser instance gracefully
func (b *Browser) Close() {
	b.log.Info("Closing browser instance")
	_ = b.rodBrowser.Close()
}
