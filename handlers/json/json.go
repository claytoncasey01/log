// Package json implements a JSON handler.
package json

import (
	j "encoding/json"
	"io"
	"os"
	"sync"

	"github.com/claytoncasey01/log"
)

// Default handler outputting to stderr.
var Default = New(os.Stderr, log.InfoLevel)

// Handler implementation.
type Handler struct {
	*j.Encoder
	mu    sync.Mutex
	Level log.Level
}

// New handler.
func New(w io.Writer, l log.Level) *Handler {
	return &Handler{
		Encoder: j.NewEncoder(w),
		Level:   l,
	}
}

// HandleLog implements log.Handler.
func (h *Handler) HandleLog(e *log.Entry) error {
	h.mu.Lock()
	defer h.mu.Unlock()

	if e.Level == h.Level {
		return h.Encoder.Encode(e)
	}

	return nil
}
