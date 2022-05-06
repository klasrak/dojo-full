package api

import (
	"encoding/json"
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

	jsonResp, err := json.Marshal(result)

	if err != nil {
		httphelpers.InternalServerError(rw)
		return
	}

	httphelpers.OK(rw, jsonResp)
}

func GetStarshipsHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("GetStarships"))
}

func GetPeopleHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("GetPeople"))
}

func GetPeoplesHandler(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("GetPeoples"))
}
