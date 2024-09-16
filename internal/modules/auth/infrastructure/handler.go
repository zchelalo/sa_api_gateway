package authInfrastructure

import (
	"context"
	"encoding/json"
	"net/http"

	authApplication "github.com/zchelalo/sa_api_gateway/internal/modules/auth/application"
	userErrors "github.com/zchelalo/sa_api_gateway/internal/modules/user/errors"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/response"
	"github.com/zchelalo/sa_api_gateway/pkg/util"
)

type AuthHandler struct {
	ctx          context.Context
	authUseCases *authApplication.AuthUseCases
}

func NewAuthHandler(ctx context.Context, authUseCases *authApplication.AuthUseCases) *AuthHandler {
	return &AuthHandler{
		ctx:          ctx,
		authUseCases: authUseCases,
	}
}

func (authHandler *AuthHandler) SignIn(w http.ResponseWriter, req *http.Request) {
	request := &authApplication.SignInRequest{}

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		resp := response.BadRequest("", err.Error())
		response.WriteErrorResponse(w, resp)
		return
	}

	auth, err := authHandler.authUseCases.SignIn(request)
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
	request := &authApplication.SignUpRequest{}

	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		resp := response.BadRequest("", err.Error())
		response.WriteErrorResponse(w, resp)
		return
	}

	auth, err := authHandler.authUseCases.SignUp(request)
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
