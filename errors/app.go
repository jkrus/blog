package errors

import "github.com/pkg/errors"

const (
	errLoadConfigMsg        = "load config"
	errProvideHTTPServerMsg = "provide http server"
	errProvideRouterMsg     = "provide router"
	errStartHTTPServerMsg   = "start http Server"
)

func ErrLoadConfig(w error) error {
	return errors.Wrap(w, errLoadConfigMsg)
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
