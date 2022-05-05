package main

import (
	"github.com/go-chi/chi/v5"
)

func URLMapping(router *chi.Mux) {
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/spaceships", GetSpaceshipHandler)
	})
}
