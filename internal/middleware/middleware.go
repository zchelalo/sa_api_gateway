package middleware

import (
	"net/http"

	authApplication "github.com/zchelalo/sa_api_gateway/internal/modules/auth/application"
)

type Middleware struct {
	authUseCases *authApplication.UseCases
}

func NewMiddleware(authUseCases *authApplication.UseCases) *Middleware {
	return &Middleware{
		authUseCases: authUseCases,
	}
}

func ApplyMiddlewares(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
