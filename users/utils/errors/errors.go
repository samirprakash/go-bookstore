package errors

import "net/http"

// REST represents an error response
type REST struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// NewBadRequest returns a new bad request error
func NewBadRequest(message string) *REST {
	return &REST{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

// NewNotFound returns a new not found error
func NewNotFound(message string) *REST {
	return &REST{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}
