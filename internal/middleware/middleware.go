package middleware

import (
	"context"
	"log"
	"net/http"
)

type Middleware struct {
	logger *log.Logger
}

func NewMiddleware(ctx context.Context, logger *log.Logger) *Middleware {
	return &Middleware{
		logger: logger,
	}
}

func ApplyMiddlewares(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
