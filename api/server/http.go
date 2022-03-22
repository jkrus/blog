package server

import (
	"context"
	"log"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/go-chi/chi/v5"

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
		wg       *sync.WaitGroup
		listener net.Listener
		server   *http.Server
	}
)

func NewHTTP(ctx context.Context, cfg *config.Config, r chi.Router) HTTP {
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
		wg:       &sync.WaitGroup{},
		listener: listener,
		server:   server,
	}
}

// Start implements HTTP interface.
func (h *httpService) Start() error {
	log.Println("Start HTTP server...")
	h.wg.Add(1)
	go h.createContextHandler()

	log.Println("HTTP server listen on:", h.server.Addr, "and serve...")

	return h.server.Serve(h.listener)
}

// Stop implements HTTP interface.
func (h *httpService) Stop() error {
	defer h.wg.Done()
	log.Println("Stop HTTP Server...")
	if err := h.server.Shutdown(context.Background()); err != nil {
		return err
	}

	log.Println("HTTP Server stop complete.")

	if h.listener != nil {
		return h.listener.Close()
	}

	return nil
}

// createContextHandler creates a context handler goroutine.
func (h *httpService) createContextHandler() {
	for {
		<-h.ctx.Done()
		_ = h.Stop()
	}
}
