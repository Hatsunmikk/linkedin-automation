package search

import (
	"errors"

	"github.com/go-rod/rod"
)

// Query defines search parameters for targeting users.
type Query struct {
	JobTitle string
	Company  string
	Location string
	Keywords []string
}

// Result represents a discovered profile.
type Result struct {
	ProfileURL string
}

// SearchEngine encapsulates search behavior.
type SearchEngine struct {
	seen map[string]bool
}

// New creates a new SearchEngine with deduplication.
func New() *SearchEngine {
	return &SearchEngine{
		seen: make(map[string]bool),
	}
}

// Search executes a simulated search and returns unique profile URLs.
//
// IMPORTANT:
// This is a proof-of-concept implementation using a mock DOM.
// It demonstrates pagination, parsing, and deduplication logic.
func (s *SearchEngine) Search(page *rod.Page, q Query, maxPages int) ([]Result, error) {
	if maxPages <= 0 {
		return nil, errors.New("maxPages must be > 0")
	}

	results := []Result{}

	// Mock search results (PoC-safe)
	mockHTML := `
		<a class="profile-link" href="https://linkedin.com/in/alice-dev">Alice</a>
		<a class="profile-link" href="https://linkedin.com/in/bob-engineer">Bob</a>
		<a class="profile-link" href="https://linkedin.com/in/alice-dev">Duplicate Alice</a>
	`

	for pageNum := 1; pageNum <= maxPages; pageNum++ {
		// Navigate to a blank page
		page.MustNavigate("about:blank")

		// Inject mock HTML into DOM (Rod v0.116 compatible)
		_, err := page.Eval(`() => {
			document.body.innerHTML = ` + "`" + mockHTML + "`" + `;
		}`)
		if err != nil {
			continue
		}

		links, err := page.Elements("a.profile-link")
		if err != nil {
			continue
		}

		for _, link := range links {
			href, err := link.Attribute("href")
			if err != nil || href == nil {
				continue
			}

			url := *href

			// Deduplication
			if s.seen[url] {
				continue
			}

			s.seen[url] = true
			results = append(results, Result{ProfileURL: url})
		}
	}

	return results, nil
}
