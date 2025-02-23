package authInfrastructure

import (
	"net/http"

	"github.com/zchelalo/sa_api_gateway/internal/middleware"
	authApplication "github.com/zchelalo/sa_api_gateway/internal/modules/auth/application"
	"github.com/zchelalo/sa_api_gateway/pkg/bootstrap"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
)

type AuthRouter struct {
	router      *http.ServeMux
	authHandler *AuthHandler
	middleware  *middleware.Middleware
}

var authGRPCClient proto.AuthServiceClient

func NewAuthRouter(router *http.ServeMux) *AuthRouter {
	authClientConn := bootstrap.GetGRPCClient(constants.AuthMicroserviceDomain)
	authGRPCClient = proto.NewAuthServiceClient(authClientConn)
	authRepository := NewGRPCRepository(authGRPCClient)
	authUseCases := authApplication.NewAuthUseCases(authRepository)
	authHandler := NewAuthHandler(authUseCases)

	middleware := middleware.NewMiddleware(authUseCases)

	return &AuthRouter{
		router:      router,
		authHandler: authHandler,
		middleware:  middleware,
	}
}

func (r *AuthRouter) SetRoutes() {
	r.router.Handle("POST /auth/sign-in", middleware.ApplyMiddlewares(http.HandlerFunc(r.authHandler.SignIn), r.middleware.Logger))
	r.router.Handle("POST /auth/sign-up", middleware.ApplyMiddlewares(http.HandlerFunc(r.authHandler.SignUp), r.middleware.Logger))
	r.router.Handle("POST /auth/sign-out", middleware.ApplyMiddlewares(http.HandlerFunc(r.authHandler.SignOut), r.middleware.Auth, r.middleware.Logger))
}
