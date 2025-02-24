package classManagementREST

import (
	"net/http"

	"github.com/zchelalo/sa_api_gateway/internal/middleware"
	classManagementApplication "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/application"
)

type Router struct {
	router     *http.ServeMux
	handler    *Handler
	middleware *middleware.Middleware
}

func New(router *http.ServeMux, classUseCases *classManagementApplication.UseCases, middleware *middleware.Middleware) *Router {
	handler := NewHandler(classUseCases)

	return &Router{
		router:     router,
		handler:    handler,
		middleware: middleware,
	}
}

func (r *Router) SetRoutes() {
	r.router.Handle("POST /classes", middleware.ApplyMiddlewares(http.HandlerFunc(r.handler.Create), r.middleware.Auth))
	r.router.Handle("POST /classes/join", middleware.ApplyMiddlewares(http.HandlerFunc(r.handler.Join), r.middleware.Auth))
	r.router.Handle("GET /classes", middleware.ApplyMiddlewares(http.HandlerFunc(r.handler.List), r.middleware.Auth))
	r.router.Handle("GET /classes/{classID}/members", middleware.ApplyMiddlewares(http.HandlerFunc(r.handler.ListMembers), r.middleware.Auth))
	r.router.Handle("GET /classes/{classID}/code", middleware.ApplyMiddlewares(http.HandlerFunc(r.handler.GetClassCode), r.middleware.Auth))
}
