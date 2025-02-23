package middleware

import (
	"context"
	"net/http"

	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/response"
	"github.com/zchelalo/sa_api_gateway/pkg/util"
)

func (mdw *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		accessToken, err := req.Cookie(string(constants.CookieAccessToken))
		if err != nil {
			errorResponse := response.Unauthorized("", "access token is required")
			response.WriteErrorResponse(w, errorResponse)
			return
		}

		refreshToken, err := req.Cookie(string(constants.CookieRefreshToken))
		if err != nil {
			errorResponse := response.Unauthorized("", "refresh token is required")
			response.WriteErrorResponse(w, errorResponse)
			return
		}

		auth, err := mdw.authenticator.IsAuthorized(req.Context(), accessToken.Value, refreshToken.Value)
		if err != nil {
			errorResponse := response.Unauthorized("", err.Error())
			response.WriteErrorResponse(w, errorResponse)
			return
		}

		if !auth.IsAuthorized {
			errorResponse := response.Unauthorized("", "unauthorized")
			response.WriteErrorResponse(w, errorResponse)
			return
		}

		if auth.Tokens.AccessToken != accessToken.Value {
			http.SetCookie(w, util.CreateCookie(constants.CookieAccessToken, auth.Tokens.AccessToken, auth.Tokens.ExpiresAt))
		}

		if auth.Tokens.RefreshToken != refreshToken.Value {
			http.SetCookie(w, util.CreateCookie(constants.CookieRefreshToken, auth.Tokens.RefreshToken, auth.Tokens.ExpiresAt))
		}

		ctx := context.WithValue(req.Context(), constants.ContextUserID, auth.UserID)

		next.ServeHTTP(w, req.WithContext(ctx))
	})
}
