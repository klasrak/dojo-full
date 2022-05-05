package main

import "net/http"

func GetStarshipHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Hello, Starships"))
}
