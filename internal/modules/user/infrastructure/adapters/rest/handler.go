package userREST

import (
	userApplication "github.com/zchelalo/sa_api_gateway/internal/modules/user/application"
)

type Handler struct {
	useCases *userApplication.UseCases
}

func NewHandler(useCases *userApplication.UseCases) *Handler {
	return &Handler{
		useCases: useCases,
	}
}
