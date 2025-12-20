package auth

import (
	"errors"
	"os"

	"github.com/go-rod/rod"
)

// AuthResult represents the outcome of an authentication attempt
type AuthResult struct {
	Success       bool
	CheckpointHit bool
	FailureReason string
}

// Login perfroms a simulated LinkedIn login flow
//
// IMPORTANT:
// This is a proof-of-concept implementation intended to demonstrate
// automation structure, error handling, and checkpoint detection.
// It must NOT be used against real LinkedIn accounts.
func Login(page *rod.Page) (*AuthResult, error) {
	email := os.Getenv("LINKEDIN_EMAIL")
	password := os.Getenv("LINKEDIN_PASSWORD")

	if email == "" || password == "" {
		return nil, errors.New("missing LINKEDIN_EMAIL or LINKEDIN_PASSWORD env variables")

	}

	//Navigate to a placeholder login page (safe demo)
	page.MustNavigate("https://example.com/login")

	//--- Simulated login steps ---
	// In a real implementation, this would:
	// 1. Locate email/password fields
	// 2. Type credentials using human-like typing
	// 3. Click submit
	// 4. Observe post-login state

	// --- Checkpoint detection (simulated) ---
	html, _ := page.HTML()

	if containsCheckpoint(html) {
		return &AuthResult{
			Success:       false,
			CheckpointHit: true,
			FailureReason: "security checkpoint detected (captcha / 2FA)",
		}, nil
	}

	// --- Login failure simulation ---
	if email == "fail@example.com" {
		return &AuthResult{
			Success:       false,
			CheckpointHit: false,
			FailureReason: "invalid credentials",
		}, nil
	}

	return &AuthResult{
		Success: true,
	}, nil
}

// containsCheckpoint simulates detection of security challenges
func containsCheckpoint(html string) bool {
	keywords := []string{
		"captcha",
		"verify",
		"security check",
		"unusual activity",
	}

	for _, k := range keywords {
		if containsIgnoreCase(html, k) {
			return true
		}
	}
	return false
}

// containsIgnoreCase checks substring match safely
func containsIgnoreCase(text, sub string) bool {
	return len(text) > 0 && len(sub) > 0 &&
		(stringContainsFold(text, sub))

}
