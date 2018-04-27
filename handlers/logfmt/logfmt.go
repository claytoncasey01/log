// Package logfmt implements a "logfmt" format handler.
package logfmt

import (
	"io"
	"os"
	"sync"

	"github.com/claytoncasey01/log"
	"github.com/go-logfmt/logfmt"
)

// Default handler outputting to stderr.
var Default = New(os.Stderr)

// Handler implementation.
type Handler struct {
	mu    sync.Mutex
	enc   *logfmt.Encoder
	Level log.Level
}

// New handler.
func New(w io.Writer, l log.Level) *Handler {
	return &Handler{
		enc: logfmt.NewEncoder(w),
		Level: l
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
	names := e.Fields.Names()

	h.mu.Lock()
	defer h.mu.Unlock()

	h.enc.EncodeKeyval("timestamp", e.Timestamp)
	h.enc.EncodeKeyval("level", e.Level.String())
	h.enc.EncodeKeyval("message", e.Message)

	for _, name := range names {
		h.enc.EncodeKeyval(name, e.Fields.Get(name))
	}

	h.enc.EndRecord()

	return nil
}
