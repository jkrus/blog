package app

import (
	"context"
	"log"

	"github.com/goava/di"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"

	"blog/config"
)

func startCommand(ctx context.Context, cfg *config.Config, dic *di.Container, app *cli.App) {
	app.Commands = append(app.Commands, &cli.Command{
		Name:  "start",
		Usage: "Starts " + app.Usage,
		Before: func(context *cli.Context) error {
			// load data from config file.
			if err := cfg.Load(); err != nil {
				return errors.Wrap(err, "load config")
			}

			return provideServices(dic)
		},
		Action: func(ctx *cli.Context) error {
			return invokeServices(dic)
		},
		After: func(c *cli.Context) error {
			c.Done()

			ctx.Done()

			log.Println("Application shutdown complete.")

			return nil
		},
	})
}

// invokeServices tries to invoke required
// services from application DI container.
func invokeServices(dic *di.Container) error {
	return nil
}

// provideServices provides cli command specific
// services from application DI container.
func provideServices(dic *di.Container) error {
	return nil
}
