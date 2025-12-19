package stealth

import (
	"time"
)

// RateLimiter enforces action limits and cooldowns
// to prevent bot-like burst behavior.
type RateLimiter struct {
	maxActions int
	window     time.Duration
	actions    int
	startTime  time.Time
}

// NewRateLimiter creates a new rate limiter with
// a maximum number of actions per time window.
func NewRateLimiter(maxActions int, window time.Duration) *RateLimiter {
	return &RateLimiter{
		maxActions: maxActions,
		window:     window,
		startTime:  time.Now(),
	}
}

// Allow checks whether an action is permitted.
// If the limit is reached, it enforces a cooldown.
func (r *RateLimiter) Allow() {
	r.actions++

	elapsed := time.Since(r.startTime)

	if r.actions >= r.maxActions && elapsed < r.window {
		// Cooldown to avoid burst behavior
		time.Sleep(r.window - elapsed)
		r.startTime = time.Now()
		r.actions = 0
	}

	if elapsed >= r.window {
		// Reset window naturally
		r.startTime = time.Now()
		r.actions = 0
	}
}
