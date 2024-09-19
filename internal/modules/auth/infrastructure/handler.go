package authInfrastructure

import (
	"encoding/json"
	"errors"
	"net/http"

	authApplication "github.com/zchelalo/sa_api_gateway/internal/modules/auth/application"
	authErrors "github.com/zchelalo/sa_api_gateway/internal/modules/auth/errors"
	userErrors "github.com/zchelalo/sa_api_gateway/internal/modules/user/errors"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/response"
	"github.com/zchelalo/sa_api_gateway/pkg/util"
)

type AuthHandler struct {
	authUseCases *authApplication.AuthUseCases
}

func NewAuthHandler(authUseCases *authApplication.AuthUseCases) *AuthHandler {
	return &AuthHandler{
		authUseCases: authUseCases,
	}
}

func (authHandler *AuthHandler) SignIn(w http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie(string(constants.CookieRefreshToken))
	if err == nil {
		resp := response.Unauthorized("", "you are already signed in")
		response.WriteErrorResponse(w, resp)
		return
	}

	request := &authApplication.SignInRequest{}

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		resp := response.BadRequest("", err.Error())
		response.WriteErrorResponse(w, resp)
		return
	}

	auth, err := authHandler.authUseCases.SignIn(req.Context(), request)
	if err != nil {
		badRequestErrors := []error{
			userErrors.ErrEmailRequired,
			userErrors.ErrEmailInvalid,
			userErrors.ErrPasswordRequired,
			userErrors.ErrPasswordInvalid,
		}
		for _, badRequestError := range badRequestErrors {
			if err == badRequestError {
				resp := response.BadRequest("", err.Error())
				response.WriteErrorResponse(w, resp)
				return
			}
		}

		if err == userErrors.ErrUserNotFound {
			resp := response.NotFound("", err.Error())
			response.WriteErrorResponse(w, resp)
			return
		}

		resp := response.InternalServerError("", err.Error())
		response.WriteErrorResponse(w, resp)
		return
	}

	http.SetCookie(w, util.CreateCookie(constants.CookieAccessToken, auth.AccessToken, auth.ExpiresAt))
	http.SetCookie(w, util.CreateCookie(constants.CookieRefreshToken, auth.RefreshToken, auth.ExpiresAt))

	resp := response.OK("", auth.User)
	response.WriteSuccessResponse(w, resp)
}

func (authHandler *AuthHandler) SignUp(w http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie(string(constants.CookieRefreshToken))
	if err == nil {
		resp := response.Unauthorized("", "you are already signed in")
		response.WriteErrorResponse(w, resp)
		return
	}

	request := &authApplication.SignUpRequest{}

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		resp := response.BadRequest("", err.Error())
		response.WriteErrorResponse(w, resp)
		return
	}

	auth, err := authHandler.authUseCases.SignUp(req.Context(), request)
	if err != nil {
		badRequestErrors := []error{
			userErrors.ErrNameRequired,
			userErrors.ErrNameInvalid,
			userErrors.ErrEmailRequired,
			userErrors.ErrEmailInvalid,
			userErrors.ErrPasswordRequired,
			userErrors.ErrPasswordInvalid,
		}
		for _, badRequestError := range badRequestErrors {
			if err == badRequestError {
				resp := response.BadRequest("", err.Error())
				response.WriteErrorResponse(w, resp)
				return
			}
		}

		if err == userErrors.ErrEmailAlreadyExists {
			resp := response.Conflict("", err.Error())
			response.WriteErrorResponse(w, resp)
			return
		}

		resp := response.InternalServerError("", err.Error())
		response.WriteErrorResponse(w, resp)
		return
	}

	http.SetCookie(w, util.CreateCookie(constants.CookieAccessToken, auth.AccessToken, auth.ExpiresAt))
	http.SetCookie(w, util.CreateCookie(constants.CookieRefreshToken, auth.RefreshToken, auth.ExpiresAt))

	resp := response.OK("", auth.User)
	response.WriteSuccessResponse(w, resp)
}

func (authHandler *AuthHandler) SignOut(w http.ResponseWriter, req *http.Request) {
	refreshToken, err := req.Cookie(string(constants.CookieRefreshToken))
	if err != nil {
		resp := response.Unauthorized("", err.Error())
		response.WriteErrorResponse(w, resp)
		return
	}

	err = authHandler.authUseCases.SignOut(req.Context(), refreshToken.Value)
	if err != nil {
		if errors.As(err, &authErrors.ErrTokenInvalid{}) || errors.As(err, &authErrors.ErrTokenExpired{}) {
			resp := response.Unauthorized("", err.Error())
			response.WriteErrorResponse(w, resp)
			return
		}

		if err == authErrors.ErrSignOutFailed {
			resp := response.InternalServerError("", err.Error())
			response.WriteErrorResponse(w, resp)
			return
		}

		resp := response.InternalServerError("", err.Error())
		response.WriteErrorResponse(w, resp)
		return
	}

	http.SetCookie(w, util.CreateCookie(constants.CookieAccessToken, "", 0))
	http.SetCookie(w, util.CreateCookie(constants.CookieRefreshToken, "", 0))

	resp := response.OK("", nil)
	response.WriteSuccessResponse(w, resp)
}
