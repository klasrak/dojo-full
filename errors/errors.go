package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type Type string

const (
	BadRequest Type = "BAD_REQUEST"
	Internal   Type = "INTERNAL_SERVER_ERROR"
	NotFound   Type = "NOT_FOUND"
)

type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Status() int {
	switch e.Type {
	case BadRequest:
		return http.StatusBadRequest
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}

func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}

	return http.StatusInternalServerError
}

/*
* Error "Factories"
 */

// NewBadRequest to create 400 errors
func NewBadRequest(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad request. Reason: %v", reason),
	}
}

// NewInternal for 500 errors
func NewInternal() *Error {
	return &Error{
		Type:    Internal,
		Message: "Internal server error.",
	}
}

// NewNotFound to create an error for 404
func NewNotFound(name string, value string) *Error {
	var message string

	if name != "" && value == "" {
		message = fmt.Sprintf("resource: %v not found", name)
	} else {
		message = fmt.Sprintf("resource: %v with id: %v not found", name, value)
	}

	return &Error{
		Type:    NotFound,
		Message: message,
	}
}
