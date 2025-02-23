package userREST

import (
	"net/http"

	"github.com/zchelalo/sa_api_gateway/internal/middleware"
	authApplication "github.com/zchelalo/sa_api_gateway/internal/modules/auth/application"
	authGRPCRepo "github.com/zchelalo/sa_api_gateway/internal/modules/auth/infrastructure/repositories/grpc"
	userApplication "github.com/zchelalo/sa_api_gateway/internal/modules/user/application"
	userGRPCRepo "github.com/zchelalo/sa_api_gateway/internal/modules/user/infrastructure/repositories/grpc"
	"github.com/zchelalo/sa_api_gateway/pkg/bootstrap"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
)

type Router struct {
	router     *http.ServeMux
	handler    *Handler
	middleware *middleware.Middleware
}

var userGRPCClient proto.UserServiceClient
var authGRPCClient proto.AuthServiceClient

func New(router *http.ServeMux) *Router {
	userClientConn := bootstrap.GetGRPCClient(constants.UserMicroserviceDomain)
	userGRPCClient = proto.NewUserServiceClient(userClientConn)
	userRepository := userGRPCRepo.New(userGRPCClient)
	userUseCases := userApplication.New(userRepository)
	handler := NewHandler(userUseCases)

	authClientConn := bootstrap.GetGRPCClient(constants.AuthMicroserviceDomain)
	authGRPCClient = proto.NewAuthServiceClient(authClientConn)
	authRepository := authGRPCRepo.New(authGRPCClient)
	authUseCases := authApplication.New(authRepository)

	middleware := middleware.NewMiddleware(authUseCases)

	return &Router{
		router:     router,
		handler:    handler,
		middleware: middleware,
	}
}

func (r *Router) SetRoutes() {
	r.router.Handle("GET /profile", middleware.ApplyMiddlewares(http.HandlerFunc(r.handler.Get), r.middleware.Auth, r.middleware.Logger))
}
