package server

import (
	"context"
	"log"
	"net/http"
	"time"

	userInfrastructure "github.com/zchelalo/sa_api_gateway/internal/modules/user/infrastructure"
	"github.com/zchelalo/sa_api_gateway/pkg/bootstrap"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
)

type Server struct {
	ctx     context.Context
	logger  *log.Logger
	router  *http.ServeMux
	address string
}

func NewServer(ctx context.Context, logger *log.Logger, address string) *Server {
	return &Server{
		ctx:     ctx,
		logger:  logger,
		router:  http.NewServeMux(),
		address: address,
	}
}

func (s *Server) Start() {
	userGRPCClient := bootstrap.GetGRPCClient(constants.UserMicroserviceDomain)
	userRouter := userInfrastructure.NewUserRouter(s.ctx, s.logger, userGRPCClient, s.router)
	userRouter.SetRoutes()

	server := &http.Server{
		Addr:         s.address,
		Handler:      accessControl(s.router),
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	errCh := make(chan error)
	go func() {
		s.logger.Printf("Server is listening on %s", s.address)
		errCh <- server.ListenAndServe()
	}()

	err := <-errCh
	if err != nil {
		log.Fatal(err)
	}
}

func accessControl(h http.Handler) http.Handler {
	allowOrigins := map[string]bool{
		"http://localhost:5173": true,
	}
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if origin := req.Header.Get("Origin"); allowOrigins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "null")
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS, HEAD")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Cache-Control, Content-Type, DNT, If-Modified-Since, Keep-Alive, Origin, User-Agent, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if req.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, req)
	})
}
