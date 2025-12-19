package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// TypeHumanLike types text into the given element
// character-by-character with randomized delays
// to simulate human typing behavior.
func TypeHumanLike(el *rod.Element, text string) {
	for _, char := range text {
		// Input a single character
		el.Input(string(char))

		// Random delay between keystrokes (human rhythm)
		delay := rand.Intn(120) + 60 // 60–180ms
		time.Sleep(time.Duration(delay) * time.Millisecond)
	}
}
