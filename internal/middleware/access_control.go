package middleware

import (
	"net/http"
	"strings"

	"github.com/zchelalo/sa_api_gateway/pkg/bootstrap"
)

func (mdw *Middleware) AccessControl(h http.Handler) http.Handler {
	config := bootstrap.GetConfig()
	origins := make(map[string]bool)
	for _, origin := range strings.Split(config.AllowedOrigins, ",") {
		origin = strings.TrimSpace(origin)
		origins[origin] = true
	}
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if origin := req.Header.Get("Origin"); origins[origin] {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		} else {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, DELETE, OPTIONS, HEAD")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Cache-Control, Content-Type, DNT, If-Modified-Since, Keep-Alive, Origin, User-Agent, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if req.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		h.ServeHTTP(w, req)
	})
}
