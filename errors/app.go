package errors

import "github.com/pkg/errors"

const (
	errLoadConfigMsg        = "load config"
	errProvideHTTPServerMsg = "provide http server"
	errOpenDatabaseMsg      = "open database"
	errProvideRouterMsg     = "provide router"
	errStartHTTPServerMsg   = "start http Server"
)

func ErrLoadConfig(w error) error {
	return errors.Wrap(w, errLoadConfigMsg)
}

func ErrOpenDatabase(w error) error {
	return errors.Wrap(w, errOpenDatabaseMsg)
}

func ErrProvideRouter(w error) error {
	return errors.Wrap(w, errProvideRouterMsg)
}

func ErrProvideHTTPServer(w error) error {
	return errors.Wrap(w, errProvideHTTPServerMsg)
}

func ErrStartHTTPServer(w error) error {
	return errors.Wrap(w, errStartHTTPServerMsg)
}
