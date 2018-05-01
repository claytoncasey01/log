// Package level implements a level filter handler.
package level

import "github.com/claytoncasey01/log"

// Handler implementation.
type Handler struct {
	Level   log.Level
	Handler log.Handler
}

// New handler.
func New(h log.Handler) *Handler {
	return &Handler{
		Level:   log.InfoLevel,
		Handler: h,
	}
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
	if e.Level < h.Level {
		return nil
	}

	return h.Handler.HandleLog(e)
}
