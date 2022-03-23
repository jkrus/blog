package errors

import "github.com/pkg/errors"

const (
	errLoadConfigMsg         = "load config"
	errProvideHTTPServerMsg  = "provide http server"
	errProvideGRPCServerMsg  = "provide grpc server"
	errProvideNoteServiceMsg = "provide models service"
	errOpenDatabaseMsg       = "open database"
	errProvideRouterMsg      = "provide router"
	errStartHTTPServerMsg    = "start http Server"
	errStartGRPCServerMsg    = "start grpc Server"
	errStartNoteServiceMsg   = "start models service"
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

func ErrProvideGRPCServer(w error) error {
	return errors.Wrap(w, errProvideGRPCServerMsg)
}

func ErrStartHTTPServer(w error) error {
	return errors.Wrap(w, errStartHTTPServerMsg)
}

func ErrStartGRPCPServer(w error) error {
	return errors.Wrap(w, errStartGRPCServerMsg)
}

func ErrStartNoteService(w error) error {
	return errors.Wrap(w, errStartNoteServiceMsg)
}

func ErrProvideUserService(w error) error {
	return errors.Wrap(w, errProvideNoteServiceMsg)
}
