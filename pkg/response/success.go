package response

import (
	"encoding/json"
	"net/http"
)

type Success struct {
	Status  uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

const (
	defaultSuccessMessage string = "success"
)

func successResponse(status uint32, message string, data interface{}) *Success {
	if message == "" {
		message = defaultSuccessMessage
	}

	return &Success{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func WriteSuccessResponse(w http.ResponseWriter, resp *Success) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(resp.Status))
	json.NewEncoder(w).Encode(resp)
}

func OK(message string, data interface{}) *Success {
	return successResponse(http.StatusOK, message, data)
}

func Created(message string, data interface{}) *Success {
	return successResponse(http.StatusCreated, message, data)
}

func Accepted(message string, data interface{}) *Success {
	return successResponse(http.StatusAccepted, message, data)
}

func NonAuthorativeInfo(message string, data interface{}) *Success {
	return successResponse(http.StatusNonAuthoritativeInfo, message, data)
}

func NoContent(message string, data interface{}) *Success {
	return successResponse(http.StatusNoContent, message, data)
}

func ResetContent(message string, data interface{}) *Success {
	return successResponse(http.StatusResetContent, message, data)
}

func PartialContent(message string, data interface{}) *Success {
	return successResponse(http.StatusPartialContent, message, data)
}
