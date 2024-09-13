package middleware

import (
	"context"
	"net/http"
)

type Middleware struct {
	ctx context.Context
}

func NewMiddleware(ctx context.Context) *Middleware {
	return &Middleware{
		ctx: ctx,
	}
}

func ApplyMiddlewares(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
