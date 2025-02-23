package authREST

import (
	authApplication "github.com/zchelalo/sa_api_gateway/internal/modules/auth/application"
)

type Handler struct {
	useCases *authApplication.UseCases
}

func NewHandler(useCases *authApplication.UseCases) *Handler {
	return &Handler{
		useCases: useCases,
	}
}
