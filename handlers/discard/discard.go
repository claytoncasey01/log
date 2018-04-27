// Package discard implements a no-op handler useful for benchmarks and tests.
package discard

import (
	"github.com/claytoncasey01/log"
)

// Default handler.
var Default = New()

// Handler implementation.
type Handler struct{}

// New handler.
func New() *Handler {
	return &Handler{}
}

// GetLevel gets the log level for the handler.
func (h *Handler) GetLevel() log.Level {
	return nil
}

// SetLevel sets the handler log level.
func (h *Handler) SetLevel(l log.Level) {}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
	return nil
}
