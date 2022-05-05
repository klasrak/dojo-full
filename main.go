package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewMux()
	r.Use(middleware.Logger)

	URLMapping(r)

	if err := http.ListenAndServe(":3001", r); err != nil {
		panic(err)
	}
}
