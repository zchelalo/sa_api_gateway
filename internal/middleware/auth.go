package middleware

import (
	"net/http"

	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/response"
)

func (mdw *Middleware) Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessToken, err := r.Cookie(string(constants.CookieAccessToken))
		if err != nil {
			errorResponse := response.Unauthorized("", "access token is required")
			response.WriteErrorResponse(w, errorResponse)
			return
		}

		refreshToken, err := r.Cookie(string(constants.CookieRefreshToken))
		if err != nil {
			errorResponse := response.Unauthorized("", "refresh token is required")
			response.WriteErrorResponse(w, errorResponse)
			return
		}

		auth, err := mdw.authUseCases.IsAuthorized(accessToken.Value, refreshToken.Value)
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
			http.SetCookie(w, &http.Cookie{
				Name:     string(constants.CookieAccessToken),
				Value:    auth.Tokens.AccessToken,
				HttpOnly: true,
				SameSite: http.SameSiteNoneMode,
				Secure:   false,
			})
		}

		if auth.Tokens.RefreshToken != refreshToken.Value {
			http.SetCookie(w, &http.Cookie{
				Name:     string(constants.CookieRefreshToken),
				Value:    auth.Tokens.RefreshToken,
				HttpOnly: true,
				SameSite: http.SameSiteNoneMode,
				Secure:   false,
			})
		}

		next.ServeHTTP(w, r)
	})
}
