package classManagementREST

import classManagementApplication "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/application"

type Handler struct {
	useCases *classManagementApplication.UseCases
}

func NewHandler(useCases *classManagementApplication.UseCases) *Handler {
	return &Handler{
		useCases: useCases,
	}
}
