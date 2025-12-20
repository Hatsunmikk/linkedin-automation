package messaging

import (
	"errors"
	"fmt"
	"time"

	"github.com/Hatsunmikk/linkedin-automation/internal/state"
)

// Message represents a follow-up message intent.
type Message struct {
	ProfileURL string
	Content    string
	SentAt     time.Time
}

// Manager handles messaging logic.
type Manager struct {
	state *state.State
}

// New creates a new messaging manager.
func New(st *state.State) *Manager {
	return &Manager{
		state: st,
	}
}

// IsConnectionAccepted simulates detection of an accepted connection.
//
// PoC behavior:
// If a connection request exists in state, we assume it was accepted.
func (m *Manager) IsConnectionAccepted(profileURL string) bool {
	_, sent := m.state.SentRequests[profileURL]
	return sent
}

// CanSendMessage checks if a message was already sent.
func (m *Manager) CanSendMessage(profileURL string) bool {
	_, sent := m.state.SentMessages[profileURL]
	return !sent
}

// SendFollowUp simulates sending a follow-up message.
func (m *Manager) SendFollowUp(profileURL, content string) (*Message, error) {
	if !m.IsConnectionAccepted(profileURL) {
		return nil, errors.New("connection not accepted")
	}

	if !m.CanSendMessage(profileURL) {
		return nil, errors.New("message already sent")
	}

	msg := &Message{
		ProfileURL: profileURL,
		Content:    content,
		SentAt:     time.Now(),
	}

	// Persist message state
	m.state.MarkMessageSent(profileURL)

	return msg, nil
}

// BuildTemplate builds a message from a template.
func BuildTemplate(name, company string) string {
	return fmt.Sprintf(
		"Hi %s, thanks for connecting! I’d love to learn more about your experience at %s.",
		name,
		company,
	)
}
