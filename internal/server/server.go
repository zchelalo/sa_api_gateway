package server

import (
	"log"
	"net/http"
	"time"

	"github.com/zchelalo/sa_api_gateway/internal/middleware"
	"github.com/zchelalo/sa_api_gateway/pkg/bootstrap"
)

type Server struct {
	router  http.Handler
	address string
}

func New(address string, mdw *middleware.Middleware, routerRegistrations ...func(*http.ServeMux)) (*Server, error) {
	router := http.NewServeMux()

	for _, register := range routerRegistrations {
		register(router)
	}

	finalRouter := mdw.Logger(mdw.AccessControl(router))

	return &Server{
		router:  finalRouter,
		address: address,
	}, nil
}

func (s *Server) Start() {
	server := &http.Server{
		Addr:         s.address,
		Handler:      s.router,
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
