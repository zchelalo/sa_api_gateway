package userInfrastructure

import (
	"context"
	"net/http"

	userApplication "github.com/zchelalo/sa_api_gateway/internal/modules/user/application"
	userProto "github.com/zchelalo/sa_api_gateway/pkg/proto/user"
	"google.golang.org/grpc"
)

type UserRouter struct {
	ctx         context.Context
	router      *http.ServeMux
	userHandler *UserHandler
}

var userGRPCClient userProto.UserServiceClient

func NewUserRouter(ctx context.Context, cc grpc.ClientConnInterface, router *http.ServeMux) *UserRouter {
	userGRPCClient = userProto.NewUserServiceClient(cc)
	userRepository := NewGRPCRepository(ctx, userGRPCClient)
	userUseCases := userApplication.NewUserUseCases(ctx, userRepository)
	UserHandler := NewUserHandler(ctx, userUseCases)

	return &UserRouter{
		ctx:         ctx,
		router:      router,
		userHandler: UserHandler,
	}
}

func (r *UserRouter) SetRoutes() {
	r.router.Handle("GET /users/{id}", http.HandlerFunc(r.userHandler.Get))
}
