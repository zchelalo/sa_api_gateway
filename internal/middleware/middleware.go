package middleware

import (
	"net/http"

	authApplication "github.com/zchelalo/sa_api_gateway/internal/modules/auth/application"
)

type Middleware struct {
	authUseCases *authApplication.AuthUseCases
}

func NewMiddleware(authUseCases *authApplication.AuthUseCases) *Middleware {
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
