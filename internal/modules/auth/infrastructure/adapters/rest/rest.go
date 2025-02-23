package authREST

import (
	"net/http"

	"github.com/zchelalo/sa_api_gateway/internal/middleware"
	authApplication "github.com/zchelalo/sa_api_gateway/internal/modules/auth/application"
	authGRPCRepo "github.com/zchelalo/sa_api_gateway/internal/modules/auth/infrastructure/repositories/grpc"
	"github.com/zchelalo/sa_api_gateway/pkg/bootstrap"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
)

type AuthRouter struct {
	router     *http.ServeMux
	handler    *Handler
	middleware *middleware.Middleware
}

var authGRPCClient proto.AuthServiceClient

func NewAuthRouter(router *http.ServeMux) *AuthRouter {
	authClientConn := bootstrap.GetGRPCClient(constants.AuthMicroserviceDomain)
	authGRPCClient = proto.NewAuthServiceClient(authClientConn)
	authRepository := authGRPCRepo.New(authGRPCClient)
	authUseCases := authApplication.New(authRepository)
	handler := NewHandler(authUseCases)

	middleware := middleware.NewMiddleware(authUseCases)

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
