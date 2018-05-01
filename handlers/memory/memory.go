// Package memory implements an in-memory handler useful for testing, as the
// entries can be accessed after writes.
package memory

import (
	"sync"

	"github.com/claytoncasey01/log"
)

// Handler implementation.
type Handler struct {
	mu      sync.Mutex
	Entries []*log.Entry
	Level   log.Level
}

// New handler.
func New() *Handler {
	return &Handler{Level: log.InfoLevel}
}

// GetLevel returns the log level for the given Handler
func (h *Handler) GetLevel() log.Level {
	return h.Level
}

// SetLevel sets the handler log level.
func (h *Handler) SetLevel(l log.Level) {
	h.Level = l
}

// SetLevelFromString sets the handler log level from a string.
func (h *Handler) SetLevelFromString(s string) {
	h.Level = log.MustParseLevel(s)
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.Entries = append(h.Entries, e)
	return nil
}
