package middleware

import (
	"net/http"
	"time"
)

func (mdw *Middleware) Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		mdw.logger.Printf("Method: %s, URI: %s, Duration: %v", r.Method, r.RequestURI, duration)
	})
}
