package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/goava/di"

	"blog/api/handlers/http"
	"blog/api/server"
)

func StartAPI(r chi.Router, httpServer server.HTTP, dic *di.Container) error {

	r.Route("/api", func(route chi.Router) {
		route.Route("/v1", func(route chi.Router) {
			http.Register(route, dic)
		})
	})

	return httpServer.Start()
}
