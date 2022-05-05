package main

import "net/http"

func GetSpaceshipHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusOK)
	rw.Write([]byte("Hello, Spaceships"))
}
