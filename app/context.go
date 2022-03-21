package app

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// NewContext construct context of the application.
func NewContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	go handleOsSig(cancel)

	return ctx
}

func handleOsSig(cancel context.CancelFunc) {
	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGQUIT, syscall.SIGABRT, syscall.SIGTERM)
	for {
		sig := <-sigCh
		println()
		log.Printf("Received %s os signal, shutting down...", sig)
		cancel()
	}
}
