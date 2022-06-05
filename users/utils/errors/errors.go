package errors

import "net/http"

// REST represents an error response
type REST struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// NewBadRequestError returns a new bad request error
func NewBadRequestError(message string) *REST {
	return &REST{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}
