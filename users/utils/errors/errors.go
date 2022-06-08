package errors

import (
	"errors"
	"net/http"
)

// REST represents an error response
type REST struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewError(message string, err error) error {
	return errors.New(message)
}

// NewBadRequestError returns a new bad request error
func NewBadRequestError(message string) *REST {
	return &REST{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFoundError returns a new not found error
func NewNotFoundError(message string) *REST {
	return &REST{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

// NewInternalServerError returns a new internal server error
func NewInternalServerError(message string) *REST {
	return &REST{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_error",
	}
}
