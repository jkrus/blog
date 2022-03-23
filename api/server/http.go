package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/goava/di"
	"github.com/pkg/errors"

	"blog/config"
	"blog/service"
)

type (
	// HTTP implements HTTP server interface.
	HTTP interface {
		// Service use embedded common interface.
		service.Service
	}

	// httpService implemented HTTP interface.
	httpService struct {
		ctx      context.Context
		dic      *di.Container
		wg       *sync.WaitGroup
		listener net.Listener
		server   *http.Server
	}
)

func NewHTTP(ctx context.Context, cfg *config.Config, r chi.Router, dic *di.Container) HTTP {
	if cfg.HTTP.Port <= 0 || cfg.Host == "" {
		log.Fatal("Can't create HTTP Server: config is not specified")
	}

	addr := cfg.Host + ":" + strconv.Itoa(cfg.HTTP.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Can't create HTTP Server: can't create listener")
	}

	server := &http.Server{
		Addr:    addr,
		Handler: r,
	}

	return &httpService{
		ctx:      ctx,
		dic:      dic,
		wg:       &sync.WaitGroup{},
		listener: listener,
		server:   server,
	}
}

// Start implements HTTP interface.
func (h *httpService) Start() error {
	log.Println("Start HTTP server...")
	if err := h.dic.Resolve(&h.wg); err != nil {
		return errors.Wrap(err, "resolve application wait group filed")
	}
	h.createContextHandler()

	log.Println("HTTP server listen on:", h.server.Addr, "and serve...")

	return h.server.Serve(h.listener)
}

// Stop implements HTTP interface.
func (h *httpService) Stop() error {
	log.Println("Stop HTTP Server...")
	if err := h.server.Shutdown(context.Background()); err != nil {
		return err
	}

	log.Println("HTTP Server stopped.")

	if h.listener != nil {
		return h.listener.Close()
	}

	return nil
}

// createContextHandler creates a context handler goroutine.
func (h *httpService) createContextHandler() {
	h.wg.Add(1)
	go func() {
		for {
			<-h.ctx.Done()
			_ = h.Stop()
			h.wg.Done()
			return
		}
	}()

}
