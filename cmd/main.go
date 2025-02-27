package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/zchelalo/sa_api_gateway/internal/middleware"
	authApplication "github.com/zchelalo/sa_api_gateway/internal/modules/auth/application"
	authREST "github.com/zchelalo/sa_api_gateway/internal/modules/auth/infrastructure/adapters/rest"
	authGRPCRepo "github.com/zchelalo/sa_api_gateway/internal/modules/auth/infrastructure/repositories/grpc"
	classManagementApplication "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/application"
	classManagementREST "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/infrastructure/adapters/rest"
	classManagementGRPCRepo "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/infrastructure/repositories/grpc"
	userApplication "github.com/zchelalo/sa_api_gateway/internal/modules/user/application"
	userREST "github.com/zchelalo/sa_api_gateway/internal/modules/user/infrastructure/adapters/rest"
	userGRPCRepo "github.com/zchelalo/sa_api_gateway/internal/modules/user/infrastructure/repositories/grpc"
	"github.com/zchelalo/sa_api_gateway/internal/server"
	"github.com/zchelalo/sa_api_gateway/pkg/bootstrap"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
	"google.golang.org/grpc"
)

func main() {
	logger := bootstrap.GetLogger()

	config, err := bootstrap.LoadConfig(".")
	if err != nil {
		logger.Fatal("cannot load config:", err)
	}

	address := fmt.Sprintf("0.0.0.0:%d", config.Port)

	services := map[constants.GRPCConstants]string{
		constants.UserMicroserviceDomain:            config.UserMicroserviceDomain,
		constants.AuthMicroserviceDomain:            config.AuthMicroserviceDomain,
		constants.ClassManagementMicroserviceDomain: config.ClassManagementMicroserviceDomain,
		constants.MemberMicroserviceDomain:          config.ClassManagementMicroserviceDomain,
	}
	for name, addr := range services {
		if err := bootstrap.InitGRPCClient(addr, name); err != nil {
			logger.Fatalf("cannot init grpc client for %s: %v", name, err)
		}
	}

	clients := make(map[constants.GRPCConstants]*grpc.ClientConn)
	for name := range services {
		client, err := bootstrap.GetGRPCClient(name)
		if err != nil {
			logger.Fatalf("cannot get grpc client for %s: %v", name, err)
		}
		clients[name] = client
	}

	userRepository := userGRPCRepo.New(proto.NewUserServiceClient(clients[constants.UserMicroserviceDomain]))
	authRepository := authGRPCRepo.New(proto.NewAuthServiceClient(clients[constants.AuthMicroserviceDomain]))
	classRepository := classManagementGRPCRepo.New(
		proto.NewClassServiceClient(clients[constants.ClassManagementMicroserviceDomain]),
		proto.NewMemberServiceClient(clients[constants.MemberMicroserviceDomain]),
	)

	userUseCases := userApplication.New(userRepository)
	authUseCases := authApplication.New(authRepository)
	classUseCases := classManagementApplication.New(classRepository)

	mdw := middleware.New(authUseCases)

	s, err := server.New(address, mdw,
		func(router *http.ServeMux) { userREST.New(router, userUseCases, mdw).SetRoutes() },
		func(router *http.ServeMux) { authREST.New(router, authUseCases, mdw).SetRoutes() },
		func(router *http.ServeMux) { classManagementREST.New(router, classUseCases, mdw).SetRoutes() },
	)
	if err != nil {
		logger.Fatal("cannot create server:", err)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigs
		logger.Println("shutting down gracefully...")
		bootstrap.CloseGRPCClients()
		os.Exit(0)
	}()

	s.Start()
}
