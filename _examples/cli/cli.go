package main

import (
	"errors"
	"time"

	"github.com/claytoncasey01/log"
	"github.com/claytoncasey01/log/handlers/cli"
)

func main() {
	handler := cli.Default
	handler.SetLevel(log.DebugLevel)

	log.SetHandler(handler)

	ctx := log.WithFields(log.Fields{
		"file": "something.png",
		"type": "image/png",
		"user": "tobi",
	})

	go func() {
		for range time.Tick(time.Second) {
			ctx.Debug("doing stuff")
		}
	}()

	go func() {
		for range time.Tick(100 * time.Millisecond) {
			ctx.Info("uploading")
			ctx.Info("upload complete")
		}
	}()

	go func() {
		for range time.Tick(time.Second) {
			ctx.Warn("upload slow")
		}
	}()

	go func() {
		for range time.Tick(2 * time.Second) {
			err := errors.New("boom")
			ctx.WithError(err).Error("upload failed")
		}
	}()

	select {}
}
