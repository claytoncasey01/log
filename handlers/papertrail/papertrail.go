// Package papertrail implements a papertrail logfmt format handler.
package papertrail

import (
	"bytes"
	"fmt"
	"log/syslog"
	"net"
	"os"
	"sync"
	"time"

	"github.com/claytoncasey01/log"
	"github.com/go-logfmt/logfmt"
)

// TODO: syslog portion is ad-hoc for my serverless use-case,
// I don't really need hostnames etc, but this should be improved

// Config for Papertrail.
type Config struct {
	// Papertrail settings.
	Host string // Host subdomain such as "logs4"
	Port int    // Port number

	// Application settings
	Hostname string // Hostname value
	Tag      string // Tag value
}

// Handler implementation.
type Handler struct {
	*Config

	mu    sync.Mutex
	conn  net.Conn
	Level log.Level
}

// New handler.
func New(config *Config) *Handler {
	conn, err := net.Dial("udp", fmt.Sprintf("%s.papertrailapp.com:%d", config.Host, config.Port))
	if err != nil {
		panic(err)
	}

	return &Handler{
		Config: config,
		conn:   conn,
		Level:  log.InfoLevel,
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
	ts := time.Now().Format(time.Stamp)

	var buf bytes.Buffer

	enc := logfmt.NewEncoder(&buf)
	enc.EncodeKeyval("level", e.Level.String())
	enc.EncodeKeyval("message", e.Message)

	for k, v := range e.Fields {
		enc.EncodeKeyval(k, v)
	}

	enc.EndRecord()

	msg := []byte(fmt.Sprintf("<%d>%s %s %s[%d]: %s\n", syslog.LOG_KERN, ts, h.Hostname, h.Tag, os.Getpid(), buf.String()))

	h.mu.Lock()
	_, err := h.conn.Write(msg)
	h.mu.Unlock()

	return err
}
