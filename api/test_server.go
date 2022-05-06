package api

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/go-chi/chi/v5"
)

const (
	ErrorAPICall = "Error API didn't response the call"
)

type Response struct {
	StatusCode int
	Headers    http.Header
	Body       []byte
}

func (response *Response) StringBody() string {
	return string(response.Body)
}

func GetTestRouter() *chi.Mux {
	router := chi.NewRouter()

	URLMapping(router)

	return router
}

func DoRequest(method string, url string, headers http.Header, body string) *Response {
	var bodyReader io.Reader

	if body != "" {
		bodyReader = strings.NewReader(body)
	}

	request := httptest.NewRequest(method, url, bodyReader)

	response := httptest.NewRecorder()

	request.Header = headers

	router := GetTestRouter()

	router.ServeHTTP(response, request)

	bytes := []byte(ErrorAPICall)

	if response.Body != nil {
		bytes = response.Body.Bytes()
	}

	return &Response{
		StatusCode: response.Code,
		Headers:    response.Header(),
		Body:       bytes,
	}
}
