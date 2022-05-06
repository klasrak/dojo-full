package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run() error {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	URLMapping(r)

	if err := http.ListenAndServe(":3000", r); err != nil {
		return err
	}

	return nil
}