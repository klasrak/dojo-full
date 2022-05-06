package api

import (
	"net/http"
	"strconv"
	"swapi/errors"
	"swapi/httphelpers"
	"swapi/services"

	"github.com/go-chi/chi/v5"
)

func GetStarshipHandler(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		httphelpers.BadRequest(rw, errors.NewBadRequest("invalid id"))
		return
	}

	result, err := services.GetStarshipService(id)

	if err != nil {
		if errors.Status(err) == http.StatusNotFound {
			httphelpers.NotFound(rw, err)
			return
		} else {
			httphelpers.InternalServerError(rw)
			return
		}
	}

	httphelpers.OK(rw, result)
}

func GetStarshipsHandler(rw http.ResponseWriter, r *http.Request) {
	result, err := services.GetStarshipsService()

	if err != nil {
		if errors.Status(err) == http.StatusNotFound {
			httphelpers.NotFound(rw, err)
			return
		} else {
			httphelpers.InternalServerError(rw)
			return
		}
	}

	httphelpers.OK(rw, result)
}

func GetPeopleHandler(rw http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))

	if err != nil {
		httphelpers.BadRequest(rw, errors.NewBadRequest("invalid id"))
		return
	}

	result, err := services.GetPeopleService(id)

	if err != nil {
		if errors.Status(err) == http.StatusNotFound {
			httphelpers.NotFound(rw, err)
			return
		} else {
			httphelpers.InternalServerError(rw)
			return
		}
	}

	httphelpers.OK(rw, result)
}

func GetPeopleListHandler(rw http.ResponseWriter, r *http.Request) {
	result, err := services.GetPeopleListService()

	if err != nil {
		if errors.Status(err) == http.StatusNotFound {
			httphelpers.NotFound(rw, err)
			return
		} else {
			httphelpers.InternalServerError(rw)
			return
		}
	}

	httphelpers.OK(rw, result)
}
