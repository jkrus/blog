package server

import (
	"context"
	"log"
	"net"
	"strconv"
	"sync"

	"github.com/goava/di"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	handlers "blog/api/handlers/grpc"
	"blog/config"
	"blog/service"
)

type (
	// GRPC implements GRPC server interface.
	GRPC interface {
		// Service use embedded common interface.
		service.Service
	}

	// grpcService implemented GRPC interface.
	grpcService struct {
		ctx      context.Context
		dic      *di.Container
		wg       *sync.WaitGroup
		listener net.Listener
		server   *grpc.Server
	}
)

func NewGRPC(ctx context.Context, cfg *config.Config, dic *di.Container) GRPC {
	if cfg.GRPC.Port <= 0 || cfg.Host == "" {
		log.Fatal("Can't create GRPC Server: config is not specified")
	}

	addr := cfg.Host + ":" + strconv.Itoa(cfg.GRPC.Port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Can't create GRPC Server: can't create listener")
	}

	return &grpcService{
		ctx:      ctx,
		dic:      dic,
		wg:       &sync.WaitGroup{},
		listener: listener,
		server:   grpc.NewServer(),
	}
}

// Start implements GRPC interface.
func (h *grpcService) Start() error {
	log.Println("Start GRPC server...")
	if err := h.dic.Resolve(&h.wg); err != nil {
		return errors.Wrap(err, "resolve application wait group filed")
	}
	h.createContextHandler()

	handlers.RegisterNoteHandlers(h.dic, h.server)

	go func() {
		_ = h.server.Serve(h.listener)
	}()

	log.Println("GRPC server listen on:", h.listener.Addr().String(), "and serve...")

	return nil
}

// Stop implements HTTP interface.
func (h *grpcService) Stop() error {
	log.Println("Stop GRPC Server...")
	h.server.GracefulStop()

	log.Println("GRPC Server stopped.")

	if h.listener != nil {
		return h.listener.Close()
	}

	return nil
}

// createContextHandler creates a context handler goroutine.
func (h *grpcService) createContextHandler() {
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
