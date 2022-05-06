package api

import "github.com/go-chi/chi/v5"

func URLMapping(router *chi.Mux) {
	router.Route("/api/v1", func(r chi.Router) {
		r.Get("/starships/{id}", GetStarshipHandler)
		r.Get("/starships", GetStarshipsHandler)
		r.Get("/people/{id}", GetPeopleHandler)
		r.Get("/people", GetPeopleListHandler)
	})
}
