// Package implements a Graylog-backed handler.
package graylog

import (
	"github.com/aphistic/golf"
	"github.com/claytoncasey01/log"
)

// Handler implementation.
type Handler struct {
	logger *golf.Logger
	client *golf.Client
	Level  log.Level
}

// New handler.
// Connection string should be in format "udp://<ip_address>:<port>".
// Server should have GELF input enabled on that port.
func New(url string, l log.Level) (*Handler, error) {
	c, err := golf.NewClient()
	if err != nil {
		return nil, err
	}

	err = c.Dial(url)
	if err != nil {
		return nil, err
	}

	l, err := c.NewLogger()
	if err != nil {
		return nil, err
	}

	return &Handler{
		logger: l,
		client: c,
		Level:  l,
	}, nil
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
	switch e.Level {
	case log.DebugLevel:
		return h.logger.Dbgm(e.Fields, e.Message)
	case log.InfoLevel:
		return h.logger.Infom(e.Fields, e.Message)
	case log.WarnLevel:
		return h.logger.Warnm(e.Fields, e.Message)
	case log.ErrorLevel:
		return h.logger.Errm(e.Fields, e.Message)
	case log.FatalLevel:
		return h.logger.Critm(e.Fields, e.Message)
	}

	return nil
}

// Closes connection to server, flushing message queue.
func (h *Handler) Close() error {
	return h.client.Close()
}
