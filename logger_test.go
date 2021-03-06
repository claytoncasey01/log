package log_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/claytoncasey01/log"
	"github.com/claytoncasey01/log/handlers/discard"
	"github.com/claytoncasey01/log/handlers/memory"
	"github.com/stretchr/testify/assert"
)

func TestLogger_printf(t *testing.T) {
	h := memory.New()

	l := &log.Logger{
		Handler: h,
	}

	l.Infof("logged in %s", "Tobi")

	e := h.Entries[0]
	assert.Equal(t, e.Message, "logged in Tobi")
	assert.Equal(t, e.Level, log.InfoLevel)
}

func TestLogger_levels(t *testing.T) {
	h := memory.New()

	l := &log.Logger{
		Handler: h,
	}

	l.Debug("uploading")
	l.Info("upload complete")

	assert.Equal(t, 1, len(h.Entries))

	e := h.Entries[0]
	assert.Equal(t, e.Message, "upload complete")
	assert.Equal(t, e.Level, log.InfoLevel)
}

func TestLogger_WithFields(t *testing.T) {
	h := memory.New()

	l := &log.Logger{
		Handler: h,
	}

	ctx := l.WithFields(log.Fields{"file": "sloth.png"})
	ctx.Debug("uploading")
	ctx.Info("upload complete")

	assert.Equal(t, 1, len(h.Entries))

	e := h.Entries[0]
	assert.Equal(t, e.Message, "upload complete")
	assert.Equal(t, e.Level, log.InfoLevel)
	assert.Equal(t, log.Fields{"file": "sloth.png"}, e.Fields)
}

func TestLogger_WithField(t *testing.T) {
	h := memory.New()

	l := &log.Logger{
		Handler: h,
	}

	ctx := l.WithField("file", "sloth.png").WithField("user", "Tobi")
	ctx.Debug("uploading")
	ctx.Info("upload complete")

	assert.Equal(t, 1, len(h.Entries))

	e := h.Entries[0]
	assert.Equal(t, e.Message, "upload complete")
	assert.Equal(t, e.Level, log.InfoLevel)
	assert.Equal(t, log.Fields{"file": "sloth.png", "user": "Tobi"}, e.Fields)
}

func TestLogger_Trace_info(t *testing.T) {
	h := memory.New()

	l := &log.Logger{
		Handler: h,
	}

	func() (err error) {
		defer l.WithField("file", "sloth.png").Trace("upload").Stop(&err)
		return nil
	}()

	assert.Equal(t, 2, len(h.Entries))

	{
		e := h.Entries[0]
		assert.Equal(t, e.Message, "upload")
		assert.Equal(t, e.Level, log.InfoLevel)
		assert.Equal(t, log.Fields{"file": "sloth.png"}, e.Fields)
	}

	{
		e := h.Entries[1]
		assert.Equal(t, e.Message, "upload")
		assert.Equal(t, e.Level, log.InfoLevel)
		assert.Equal(t, "sloth.png", e.Fields["file"])
		assert.IsType(t, time.Duration(0), e.Fields["duration"])
	}
}

func TestLogger_Trace_error(t *testing.T) {
	h := memory.New()

	l := &log.Logger{
		Handler: h,
	}

	func() (err error) {
		defer l.WithField("file", "sloth.png").Trace("upload").Stop(&err)
		return fmt.Errorf("boom")
	}()

	assert.Equal(t, 2, len(h.Entries))

	{
		e := h.Entries[0]
		assert.Equal(t, e.Message, "upload")
		assert.Equal(t, e.Level, log.InfoLevel)
		assert.Equal(t, "sloth.png", e.Fields["file"])
	}

	{
		e := h.Entries[1]
		assert.Equal(t, e.Message, "upload")
		assert.Equal(t, e.Level, log.ErrorLevel)
		assert.Equal(t, "sloth.png", e.Fields["file"])
		assert.Equal(t, "boom", e.Fields["error"])
		assert.IsType(t, time.Duration(0), e.Fields["duration"])
	}
}

func TestLogger_Trace_nil(t *testing.T) {
	h := memory.New()

	l := &log.Logger{
		Handler: h,
	}

	func() {
		defer l.WithField("file", "sloth.png").Trace("upload").Stop(nil)
	}()

	assert.Equal(t, 2, len(h.Entries))

	{
		e := h.Entries[0]
		assert.Equal(t, e.Message, "upload")
		assert.Equal(t, e.Level, log.InfoLevel)
		assert.Equal(t, log.Fields{"file": "sloth.png"}, e.Fields)
	}

	{
		e := h.Entries[1]
		assert.Equal(t, e.Message, "upload")
		assert.Equal(t, e.Level, log.InfoLevel)
		assert.Equal(t, "sloth.png", e.Fields["file"])
		assert.IsType(t, time.Duration(0), e.Fields["duration"])
	}
}

func BenchmarkLogger_small(b *testing.B) {
	l := &log.Logger{
		Handler: discard.New(),
	}

	for i := 0; i < b.N; i++ {
		l.Info("login")
	}
}

func BenchmarkLogger_medium(b *testing.B) {
	l := &log.Logger{
		Handler: discard.New(),
	}

	for i := 0; i < b.N; i++ {
		l.WithFields(log.Fields{
			"file": "sloth.png",
			"type": "image/png",
			"size": 1 << 20,
		}).Info("upload")
	}
}

func BenchmarkLogger_large(b *testing.B) {
	l := &log.Logger{
		Handler: discard.New(),
	}

	err := fmt.Errorf("boom")

	for i := 0; i < b.N; i++ {
		l.WithFields(log.Fields{
			"file": "sloth.png",
			"type": "image/png",
			"size": 1 << 20,
		}).
			WithFields(log.Fields{
				"some":     "more",
				"data":     "here",
				"whatever": "blah blah",
				"more":     "stuff",
				"context":  "such useful",
				"much":     "fun",
			}).
			WithError(err).Error("upload failed")
	}
}
