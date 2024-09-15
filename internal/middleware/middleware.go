package middleware

import (
	"context"
	"net/http"

	authApplication "github.com/zchelalo/sa_api_gateway/internal/modules/auth/application"
)

type Middleware struct {
	ctx          context.Context
	authUseCases *authApplication.AuthUseCases
}

func NewMiddleware(ctx context.Context, authUseCases *authApplication.AuthUseCases) *Middleware {
	return &Middleware{
		ctx:          ctx,
		authUseCases: authUseCases,
	}
}

func ApplyMiddlewares(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
