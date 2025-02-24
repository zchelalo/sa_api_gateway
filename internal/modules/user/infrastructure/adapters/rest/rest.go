package userREST

import (
	"net/http"

	"github.com/zchelalo/sa_api_gateway/internal/middleware"
	userApplication "github.com/zchelalo/sa_api_gateway/internal/modules/user/application"
)

type Router struct {
	router     *http.ServeMux
	handler    *Handler
	middleware *middleware.Middleware
}

func New(router *http.ServeMux, userUseCases *userApplication.UseCases, middleware *middleware.Middleware) *Router {
	handler := NewHandler(userUseCases)

	return &Router{
		router:     router,
		handler:    handler,
		middleware: middleware,
	}
}

func (r *Router) SetRoutes() {
	r.router.Handle("GET /profile", middleware.ApplyMiddlewares(http.HandlerFunc(r.handler.Get), r.middleware.Auth))
}
