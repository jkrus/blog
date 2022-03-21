package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"

	"blog/api/server"
)

func StartAPI(r chi.Router, httpServer server.HTTP) error {

	r.Route("/api", func(route chi.Router) {
		route.Route("/v1", func(route chi.Router) {
			r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
				writer.Write([]byte("Hello from HTTP"))
			})
		})
	})

	return httpServer.Start()
}
