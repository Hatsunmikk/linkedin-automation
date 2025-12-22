package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// ScrollHumanLike scrolls the page in small, irregular steps
// with pauses to simulate natural reading behavior.
func ScrollHumanLike(page *rod.Page, totalScroll int) {
	scrollRemaining := totalScroll

	for scrollRemaining > 0 {
		// Scroll in small chunks
		step := rand.Intn(200) + 100 // 100–300px
		if step > scrollRemaining {
			step = scrollRemaining
		}

		page.Mouse.Scroll(0, float64(step), 1)
		scrollRemaining -= step

		// Pause as if reading
		time.Sleep(time.Duration(rand.Intn(600)+300) * time.Millisecond)

		// Occasional small scroll back
		if rand.Float64() < 0.15 {
			page.Mouse.Scroll(0, float64(-rand.Intn(80)), 1)
			time.Sleep(time.Duration(rand.Intn(300)+150) * time.Millisecond)
		}
	}
}
