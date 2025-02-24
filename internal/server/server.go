package server

import (
	"log"
	"net/http"
	"time"

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
	"github.com/zchelalo/sa_api_gateway/pkg/bootstrap"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
)

type Server struct {
	router  *http.ServeMux
	address string
}

func NewServer(address string) *Server {
	router := http.NewServeMux()

	authGRPCClient := bootstrap.GetGRPCClient(constants.AuthMicroserviceDomain)
	userGRPCClient := bootstrap.GetGRPCClient(constants.UserMicroserviceDomain)
	classManagementGRPCClient := bootstrap.GetGRPCClient(constants.ClassManagementMicroserviceDomain)
	memberGRPCClient := bootstrap.GetGRPCClient(constants.MemberMicroserviceDomain)

	authRepository := authGRPCRepo.New(proto.NewAuthServiceClient(authGRPCClient))
	userRepository := userGRPCRepo.New(proto.NewUserServiceClient(userGRPCClient))
	classManagementRepository := classManagementGRPCRepo.New(proto.NewClassServiceClient(classManagementGRPCClient), proto.NewMemberServiceClient(memberGRPCClient))

	authUseCases := authApplication.New(authRepository)
	userUseCases := userApplication.New(userRepository)
	classManagementUseCases := classManagementApplication.New(classManagementRepository)

	mdw := middleware.New(authUseCases)

	authREST.New(router, authUseCases, mdw).SetRoutes()
	userREST.New(router, userUseCases, mdw).SetRoutes()
	classManagementREST.New(router, classManagementUseCases, mdw).SetRoutes()

	return &Server{
		router:  router,
		address: address,
	}
}

func (s *Server) Start() {
	server := &http.Server{
		Addr:         s.address,
		Handler:      s.accessControl(s.router),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	logger := bootstrap.GetLogger()

	errCh := make(chan error)
	go func() {
		logger.Printf("Server is listening on %s", s.address)
		errCh <- server.ListenAndServe()
	}()

	err := <-errCh
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) accessControl(h http.Handler) http.Handler {
	allowOrigins := map[string]bool{
		"http://localhost:5173": true,
	}
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if origin := req.Header.Get("Origin"); allowOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS, HEAD")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Cache-Control, Content-Type, DNT, If-Modified-Since, Keep-Alive, Origin, User-Agent, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if req.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, req)
	})
}
