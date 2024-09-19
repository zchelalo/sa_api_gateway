package userInfrastructure

import (
	"net/http"

	"github.com/zchelalo/sa_api_gateway/internal/middleware"
	authApplication "github.com/zchelalo/sa_api_gateway/internal/modules/auth/application"
	authInfrastructure "github.com/zchelalo/sa_api_gateway/internal/modules/auth/infrastructure"
	userApplication "github.com/zchelalo/sa_api_gateway/internal/modules/user/application"
	"github.com/zchelalo/sa_api_gateway/pkg/bootstrap"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	authProto "github.com/zchelalo/sa_api_gateway/pkg/proto/auth"
	userProto "github.com/zchelalo/sa_api_gateway/pkg/proto/user"
)

type UserRouter struct {
	router      *http.ServeMux
	userHandler *UserHandler
	middleware  *middleware.Middleware
}

var userGRPCClient userProto.UserServiceClient
var authGRPCClient authProto.AuthServiceClient

func NewUserRouter(router *http.ServeMux) *UserRouter {
	userClientConn := bootstrap.GetGRPCClient(constants.UserMicroserviceDomain)
	userGRPCClient = userProto.NewUserServiceClient(userClientConn)
	userRepository := NewGRPCRepository(userGRPCClient)
	userUseCases := userApplication.NewUserUseCases(userRepository)
	userHandler := NewUserHandler(userUseCases)

	authClientConn := bootstrap.GetGRPCClient(constants.AuthMicroserviceDomain)
	authGRPCClient = authProto.NewAuthServiceClient(authClientConn)
	authRepository := authInfrastructure.NewGRPCRepository(authGRPCClient)
	authUseCases := authApplication.NewAuthUseCases(authRepository)

	middleware := middleware.NewMiddleware(authUseCases)

	return &UserRouter{
		router:      router,
		userHandler: userHandler,
		middleware:  middleware,
	}
}

func (r *UserRouter) SetRoutes() {
	r.router.Handle("GET /profile", middleware.ApplyMiddlewares(http.HandlerFunc(r.userHandler.Get), r.middleware.Auth, r.middleware.Logger))
}
