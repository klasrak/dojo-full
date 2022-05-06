package api

import "github.com/go-chi/chi/v5"

func URLMapping(router *chi.Mux) {
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/starship/{id}", GetStarshipHandler)
		r.Get("/starships", GetStarshipsHandler)
		r.Get("/people/{id}", GetPeopleHandler)
		r.Get("/peoples", GetPeoplesHandler)
	})
}
