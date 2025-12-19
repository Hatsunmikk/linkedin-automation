package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// ApplyFingerprintMask applies browser fingerprint masking
// before any site JavaScript executes.
func ApplyFingerprintMask(page *rod.Page) error {
	rand.Seed(time.Now().UnixNano())

	// Randomize viewport size to mimic real user devices
	width := rand.Intn(400) + 1200 // 1200–1600
	height := rand.Intn(300) + 700 // 700–1000

	page.MustSetViewport(width, height, 1, false)

	// Inject stealth script on every new document.
	// This is the correct Rod v0.116 way to mask webdriver.
	page.MustEvalOnNewDocument(`
		Object.defineProperty(Navigator.prototype, 'webdriver', {
			get: () => undefined,
			configurable: true
		});
	`)

	return nil
}
