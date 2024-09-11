package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Status  uint32      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details"`
}

const (
	defaultErrorMessage string = "error"
)

func errorResponse(status uint32, message string, details interface{}) *Error {
	if message == "" {
		message = defaultErrorMessage
	}

	return &Error{
		Status:  status,
		Message: message,
		Details: details,
	}
}

func WriteErrorResponse(w http.ResponseWriter, resp *Error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(resp.Status))
	json.NewEncoder(w).Encode(resp)
}

func BadRequest(message string, details interface{}) *Error {
	return errorResponse(http.StatusBadRequest, message, details)
}

func Unauthorized(message string, details interface{}) *Error {
	return errorResponse(http.StatusUnauthorized, message, details)
}

func Forbidden(message string, details interface{}) *Error {
	return errorResponse(http.StatusForbidden, message, details)
}

func NotFound(message string, details interface{}) *Error {
	return errorResponse(http.StatusNotFound, message, details)
}

func MethodNotAllowed(message string, details interface{}) *Error {
	return errorResponse(http.StatusMethodNotAllowed, message, details)
}

func Conflict(message string, details interface{}) *Error {
	return errorResponse(http.StatusConflict, message, details)
}

func InternalServerError(message string, details interface{}) *Error {
	return errorResponse(http.StatusInternalServerError, message, details)
}
