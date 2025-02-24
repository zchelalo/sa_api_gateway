package authREST

import (
	"errors"
	"net/http"

	authError "github.com/zchelalo/sa_api_gateway/internal/modules/auth/error"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/response"
	"github.com/zchelalo/sa_api_gateway/pkg/util"
)

func (handler *Handler) SignOut(w http.ResponseWriter, req *http.Request) {
	refreshToken, err := req.Cookie(string(constants.CookieRefreshToken))
	if err != nil {
		resp := response.Unauthorized("", err.Error())
		response.WriteErrorResponse(w, resp)
		return
	}

	err = handler.useCases.SignOut(req.Context(), refreshToken.Value)
	if err != nil {
		if errors.As(err, &authError.ErrTokenInvalid{}) || errors.As(err, &authError.ErrTokenExpired{}) {
			resp := response.Unauthorized("", err.Error())
			response.WriteErrorResponse(w, resp)
			return
		}

		if err == authError.ErrSignOutFailed {
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

	resp := response.OK("", nil, nil)
	response.WriteSuccessResponse(w, resp)
}
