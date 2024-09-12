package userInfrastructure

import (
	"context"
	"log"
	"net/http"

	"github.com/zchelalo/sa_api_gateway/internal/middleware"
	userApplication "github.com/zchelalo/sa_api_gateway/internal/modules/user/application"
	userProto "github.com/zchelalo/sa_api_gateway/pkg/proto/user"
	"google.golang.org/grpc"
)

type UserRouter struct {
	ctx         context.Context
	logger      *log.Logger
	router      *http.ServeMux
	userHandler *UserHandler
	middleware  *middleware.Middleware
}

var userGRPCClient userProto.UserServiceClient

func NewUserRouter(ctx context.Context, logger *log.Logger, cc grpc.ClientConnInterface, router *http.ServeMux) *UserRouter {
	userGRPCClient = userProto.NewUserServiceClient(cc)
	userRepository := NewGRPCRepository(ctx, logger, userGRPCClient)
	userUseCases := userApplication.NewUserUseCases(ctx, logger, userRepository)
	userHandler := NewUserHandler(ctx, logger, userUseCases)

	middleware := middleware.NewMiddleware(logger)

	return &UserRouter{
		ctx:         ctx,
		logger:      logger,
		router:      router,
		userHandler: userHandler,
		middleware:  middleware,
	}
}

func (r *UserRouter) SetRoutes() {
	r.router.Handle("GET /users/{id}", middleware.ApplyMiddlewares(http.HandlerFunc(r.userHandler.Get), r.middleware.Logger))
}
