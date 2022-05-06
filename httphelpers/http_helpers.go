package httphelpers

import (
	"net/http"
	"swapi/errors"
	"swapi/utils"
)

func BadRequest(rw http.ResponseWriter, err error) {
	rw.WriteHeader(http.StatusBadRequest)
	rw.Header().Add("Content-Type", "application/json")
	rw.Write(utils.ToJSON(err))
}

func InternalServerError(rw http.ResponseWriter) {
	rw.WriteHeader(http.StatusInternalServerError)
	rw.Header().Add("Content-Type", "application/json")
	rw.Write(utils.ToJSON(errors.NewInternal()))
}

func NotFound(rw http.ResponseWriter, err error) {
	rw.WriteHeader(http.StatusNotFound)
	rw.Header().Add("Content-Type", "application/json")
	rw.Write(utils.ToJSON(err))
}

func OK(rw http.ResponseWriter, data interface{}) {
	rw.WriteHeader(http.StatusOK)
	rw.Header().Add("Content-Type", "application/json")
	rw.Write(utils.ToJSON(data))
}
