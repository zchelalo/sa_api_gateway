package authREST

import (
	"encoding/json"
	"net/http"

	authApplication "github.com/zchelalo/sa_api_gateway/internal/modules/auth/application"
	userError "github.com/zchelalo/sa_api_gateway/internal/modules/user/error"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/response"
	"github.com/zchelalo/sa_api_gateway/pkg/util"
)

func (handler *Handler) SignUp(w http.ResponseWriter, req *http.Request) {
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

	auth, err := handler.useCases.SignUp(req.Context(), request)
	if err != nil {
		badRequestErrors := []error{
			userError.ErrNameRequired,
			userError.ErrNameInvalid,
			userError.ErrEmailRequired,
			userError.ErrEmailInvalid,
			userError.ErrPasswordRequired,
			userError.ErrPasswordInvalid,
		}
		for _, badRequestError := range badRequestErrors {
			if err == badRequestError {
				resp := response.BadRequest("", err.Error())
				response.WriteErrorResponse(w, resp)
				return
			}
		}

		if err == userError.ErrEmailAlreadyExists {
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
