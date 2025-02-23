package server

import (
	"log"
	"net/http"
	"time"

	authREST "github.com/zchelalo/sa_api_gateway/internal/modules/auth/infrastructure/adapters/rest"
	userREST "github.com/zchelalo/sa_api_gateway/internal/modules/user/infrastructure/adapters/rest"
	"github.com/zchelalo/sa_api_gateway/pkg/bootstrap"
)

type Server struct {
	router  *http.ServeMux
	address string
}

func NewServer(address string) *Server {
	return &Server{
		router:  http.NewServeMux(),
		address: address,
	}
}

func (s *Server) Start() {
	userRouter := userREST.New(s.router)
	userRouter.SetRoutes()

	authRouter := authREST.NewAuthRouter(s.router)
	authRouter.SetRoutes()

	server := &http.Server{
		Addr:         s.address,
		Handler:      accessControl(s.router),
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

func accessControl(h http.Handler) http.Handler {
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
