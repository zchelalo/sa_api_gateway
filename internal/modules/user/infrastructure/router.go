package userInfrastructure

import (
	"context"
	"net/http"

	"github.com/zchelalo/sa_api_gateway/internal/middleware"
	userApplication "github.com/zchelalo/sa_api_gateway/internal/modules/user/application"
	userProto "github.com/zchelalo/sa_api_gateway/pkg/proto/user"
	"google.golang.org/grpc"
)

type UserRouter struct {
	router      *http.ServeMux
	userHandler *UserHandler
	middleware  *middleware.Middleware
}

var userGRPCClient userProto.UserServiceClient

func NewUserRouter(cc grpc.ClientConnInterface, router *http.ServeMux) *UserRouter {
	ctx := context.Background()

	userGRPCClient := userProto.NewUserServiceClient(cc)
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
