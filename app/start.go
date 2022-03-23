package app

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/goava/di"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"blog/api/router"
	"blog/api/server"
	"blog/config"
	"blog/datastore"
	api "blog/errors"
	"blog/service"
)

func startCommand(ctx context.Context, cfg *config.Config, dic *di.Container, app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "start",
		Usage: "Starts " + app.Usage,
		Before: func(context *cli.Context) error {
			// load data from config file.
			if err := cfg.Load(); err != nil {
				return api.ErrLoadConfig(err)
			}

			return provideServices(dic)
		},
		Action: func(ctx *cli.Context) error {
			return invokeServices(dic)
		},
		After: func(c *cli.Context) error {
			<-c.Done()
			var wg *sync.WaitGroup
			if err := dic.Resolve(&wg); err != nil {
				return errors.Wrap(err, "resolve application wait group filed")
			}

			ctx.Done()
			wg.Wait()
			log.Println("Application shutdown complete.")

			return nil
		},
	})
}

// invokeServices tries to invoke required
// services from application DI container.
func invokeServices(dic *di.Container) error {
	// invoke database connection starter
	if err := dic.Invoke(datastore.Start); err != nil {
		if !errors.Is(err, http.ErrServerClosed) {
			return api.ErrOpenDatabase(err)
		}
	}

	// invoke models service starter.
	if err := dic.Invoke(service.Note.Start); err != nil {
		return api.ErrStartNoteService(err)
	}

	// invoke grpc server starter
	if err := dic.Invoke(server.GRPC.Start); err != nil {
		if !errors.As(err, &http.ErrServerClosed) {
			return api.ErrStartHTTPServer(err)
		}
	}

	// invoke api router starter
	if err := dic.Invoke(router.StartAPI); err != nil {
		if !errors.As(err, &http.ErrServerClosed) {
			return api.ErrStartGRPCPServer(err)
		}
	}

	return nil
}

// provideServices provides cli command specific
// services from application DI container.
func provideServices(dic *di.Container) error {
	// provide chi router interface
	var ri chi.Router
	if err := dic.Provide(chi.NewRouter, di.As(&ri)); err != nil {
		return api.ErrProvideRouter(err)
	}

	// provide HTTP server interface.
	if err := dic.Provide(server.NewHTTP); err != nil {
		return api.ErrProvideHTTPServer(err)
	}

	// provide GRPC server interface.
	if err := dic.Provide(server.NewGRPC); err != nil {
		return api.ErrProvideGRPCServer(err)
	}

	// provide User Service.
	if err := dic.Provide(service.NewNoteService); err != nil {
		return api.ErrProvideUserService(err)
	}

	return nil
}
