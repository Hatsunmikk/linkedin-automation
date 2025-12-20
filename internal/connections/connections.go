package connections

import (
	"errors"
	"fmt"
	"time"

	"github.com/Hatsunmikk/linkedin-automation/internal/state"
)

// Request represents a connection request intent.
type Request struct {
	ProfileURL string
	Note       string
	SentAt     time.Time
}

// Manager handles connection request logic.
type Manager struct {
	dailyLimit int
	state      *state.State
}

// New creates a new connection request manager.
func New(dailyLimit int, st *state.State) *Manager {
	return &Manager{
		dailyLimit: dailyLimit,
		state:      st,
	}
}

// CanSend checks whether a connection request can be sent.
func (m *Manager) CanSend(profileURL string) bool {
	_, alreadySent := m.state.SentRequests[profileURL]
	return !alreadySent
}

// Send simulates sending a connection request.
//
// IMPORTANT:
// This is a proof-of-concept implementation. It does not
// interact with LinkedIn and only records intent.
func (m *Manager) Send(profileURL, note string) (*Request, error) {
	if !m.CanSend(profileURL) {
		return nil, errors.New("connection request already sent")
	}

	if len(m.state.SentRequests) >= m.dailyLimit {
		return nil, errors.New("daily connection request limit reached")
	}

	req := &Request{
		ProfileURL: profileURL,
		Note:       note,
		SentAt:     time.Now(),
	}

	// Persist state
	m.state.MarkRequestSent(profileURL)

	return req, nil
}

// BuildPersonalizedNote generates a simple personalized note.
func BuildPersonalizedNote(name, company string) string {
	return fmt.Sprintf(
		"Hi %s, I came across your profile and would love to connect and learn more about your work at %s.",
		name,
		company,
	)
}
