package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (app *Config) routes() http.Handler {
	router := chi.NewRouter()

	router.HandleFunc("/", app.handlerIndex)
	app.publicRoutes(router)

	return router
}

func (app *Config) publicRoutes(router *chi.Mux) {
	router.Group(func(r chi.Router) {
		r.Get("/health", app.handlerHealth)
		r.Get("/modules", app.handlerGetModules)
		r.Post("/modules", app.handlerModules)
	})
}
