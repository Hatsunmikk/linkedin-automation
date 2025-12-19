package stealth

import "time"

// IsWithinBusinessHours checks whether the current
// time falls within allowed automation hours.
func IsWithinBusinessHours(startHour, endHour int) bool {
	now := time.Now()
	hour := now.Hour()

	return hour >= startHour && hour < endHour
}
