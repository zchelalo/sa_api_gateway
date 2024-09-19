package middleware

import (
	"net/http"
	"time"

	"github.com/zchelalo/sa_api_gateway/pkg/bootstrap"
)

func (mdw *Middleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logger := bootstrap.GetLogger()

		start := time.Now()
		next.ServeHTTP(w, req)
		duration := time.Since(start)
		logger.Printf("Method: %s, URI: %s, Duration: %v", req.Method, req.RequestURI, duration)
	})
}
