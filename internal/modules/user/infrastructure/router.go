package userInfrastructure

import (
	"context"
	"net/http"

	"github.com/zchelalo/sa_api_gateway/internal/middleware"
	userApplication "github.com/zchelalo/sa_api_gateway/internal/modules/user/application"
	"github.com/zchelalo/sa_api_gateway/pkg/bootstrap"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	userProto "github.com/zchelalo/sa_api_gateway/pkg/proto/user"
)

type UserRouter struct {
	router      *http.ServeMux
	userHandler *UserHandler
	middleware  *middleware.Middleware
}

var userGRPCClient userProto.UserServiceClient

func NewUserRouter(router *http.ServeMux) *UserRouter {
	ctx := context.Background()

	clientConn := bootstrap.GetGRPCClient(constants.UserMicroserviceDomain)
	userGRPCClient = userProto.NewUserServiceClient(clientConn)
	userRepository := NewGRPCRepository(ctx, userGRPCClient)
	userUseCases := userApplication.NewUserUseCases(ctx, userRepository)
	userHandler := NewUserHandler(ctx, userUseCases)

	middleware := middleware.NewMiddleware(ctx)

	return &UserRouter{
		router:      router,
		userHandler: userHandler,
		middleware:  middleware,
	}
}

func (r *UserRouter) SetRoutes() {
	r.router.Handle("GET /users/{id}", middleware.ApplyMiddlewares(http.HandlerFunc(r.userHandler.Get), r.middleware.Logger))
}
