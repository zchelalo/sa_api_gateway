package response

import (
	"encoding/json"
	"net/http"

	"github.com/zchelalo/sa_api_gateway/pkg/meta"
)

type Success struct {
	Status  uint32      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    *meta.Meta  `json:"meta"`
}

const (
	defaultSuccessMessage string = "success"
)

func successResponse(status uint32, message string, data interface{}, meta *meta.Meta) *Success {
	if message == "" {
		message = defaultSuccessMessage
	}

	return &Success{
		Status:  status,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}

func WriteSuccessResponse(w http.ResponseWriter, resp *Success) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(resp.Status))
	json.NewEncoder(w).Encode(resp)
}

func OK(message string, data interface{}, meta *meta.Meta) *Success {
	return successResponse(http.StatusOK, message, data, meta)
}

func Created(message string, data interface{}, meta *meta.Meta) *Success {
	return successResponse(http.StatusCreated, message, data, meta)
}

func Accepted(message string, data interface{}, meta *meta.Meta) *Success {
	return successResponse(http.StatusAccepted, message, data, meta)
}

func NonAuthorativeInfo(message string, data interface{}, meta *meta.Meta) *Success {
	return successResponse(http.StatusNonAuthoritativeInfo, message, data, meta)
}

func NoContent(message string, data interface{}, meta *meta.Meta) *Success {
	return successResponse(http.StatusNoContent, message, data, meta)
}

func ResetContent(message string, data interface{}, meta *meta.Meta) *Success {
	return successResponse(http.StatusResetContent, message, data, meta)
}

func PartialContent(message string, data interface{}, meta *meta.Meta) *Success {
	return successResponse(http.StatusPartialContent, message, data, meta)
}
