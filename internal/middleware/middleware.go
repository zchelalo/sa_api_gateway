package middleware

import (
	"context"
	"net/http"

	authDomain "github.com/zchelalo/sa_api_gateway/internal/modules/auth/domain"
)

type Authenticator interface {
	IsAuthorized(ctx context.Context, accessToken, refreshToken string) (*authDomain.AuthorizeEntity, error)
}

type Middleware struct {
	authenticator Authenticator
}

func New(authenticator Authenticator) *Middleware {
	return &Middleware{
		authenticator: authenticator,
	}
}

func ApplyMiddlewares(handler http.Handler, middlewares ...func(http.Handler) http.Handler) http.Handler {
	for _, middleware := range middlewares {
		handler = middleware(handler)
	}
	return handler
}
