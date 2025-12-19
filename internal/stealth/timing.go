package stealth

import (
	"math/rand"
	"time"
)

// Think pauses execution for a random duration
// to simulate human decision-making delays.
func Think(minMs, maxMs int) {
	if maxMs <= minMs {
		time.Sleep(time.Duration(minMs) * time.Millisecond)
		return
	}

	delay := rand.Intn(maxMs-minMs) + minMs
	time.Sleep(time.Duration(delay) * time.Millisecond)
}
