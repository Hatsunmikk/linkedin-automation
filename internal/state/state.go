package state

import (
	"encoding/json"
	"os"
	"sync"
	"time"
)

// State represents persisted automation state.
// This allows resuming automation after restarts
// and enforcing limits across sessions.
type State struct {
	SentRequests        map[string]time.Time `json:"sent_requests"`
	AcceptedConnections map[string]time.Time `json:"accepted_connections"`
	SentMessages        map[string]time.Time `json:"sent_messages"`

	mu sync.Mutex `json:"-"`
}

// Load loads state from disk or initializes a new one.
func Load(path string) (*State, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return newEmptyState(), nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var s State
	if err := json.Unmarshal(data, &s); err != nil {
		return nil, err
	}

	if s.SentRequests == nil {
		s.SentRequests = make(map[string]time.Time)
	}
	if s.AcceptedConnections == nil {
		s.AcceptedConnections = make(map[string]time.Time)
	}
	if s.SentMessages == nil {
		s.SentMessages = make(map[string]time.Time)
	}

	return &s, nil
}

// Save persists state to disk.
func (s *State) Save(path string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// MarkRequestSent records a connection request.
func (s *State) MarkRequestSent(profileURL string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.SentRequests[profileURL] = time.Now()
}

// MarkConnectionAccepted records an accepted connection.
func (s *State) MarkConnectionAccepted(profileURL string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.AcceptedConnections[profileURL] = time.Now()
}

// MarkMessageSent records a sent message.
func (s *State) MarkMessageSent(profileURL string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.SentMessages[profileURL] = time.Now()
}

func newEmptyState() *State {
	return &State{
		SentRequests:        make(map[string]time.Time),
		AcceptedConnections: make(map[string]time.Time),
		SentMessages:        make(map[string]time.Time),
	}
}
