package authREST

import (
	"net/http"

	"github.com/zchelalo/sa_api_gateway/internal/middleware"
	authApplication "github.com/zchelalo/sa_api_gateway/internal/modules/auth/application"
)

type AuthRouter struct {
	router     *http.ServeMux
	handler    *Handler
	middleware *middleware.Middleware
}

func New(router *http.ServeMux, authUseCases *authApplication.UseCases, middleware *middleware.Middleware) *AuthRouter {
	handler := NewHandler(authUseCases)

	return &AuthRouter{
		router:     router,
		handler:    handler,
		middleware: middleware,
	}
}

func (r *AuthRouter) SetRoutes() {
	r.router.Handle("POST /auth/sign-in", middleware.ApplyMiddlewares(http.HandlerFunc(r.handler.SignIn), r.middleware.Logger))
	r.router.Handle("POST /auth/sign-up", middleware.ApplyMiddlewares(http.HandlerFunc(r.handler.SignUp), r.middleware.Logger))
	r.router.Handle("POST /auth/sign-out", middleware.ApplyMiddlewares(http.HandlerFunc(r.handler.SignOut), r.middleware.Auth, r.middleware.Logger))
}
