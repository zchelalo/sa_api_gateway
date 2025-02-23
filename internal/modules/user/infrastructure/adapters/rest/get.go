package userREST

import (
	"net/http"

	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
	userErrors "github.com/zchelalo/sa_api_gateway/internal/modules/user/errors"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/response"
)

func (handler *Handler) Get(w http.ResponseWriter, req *http.Request) {
	idContext := req.Context().Value(constants.ContextUserID)

	if idContext == nil {
		resp := response.Unauthorized("", "unauthorized")
		response.WriteErrorResponse(w, resp)
		return
	}

	id, ok := idContext.(string)
	if !ok {
		resp := response.Unauthorized("", "unauthorized")
		response.WriteErrorResponse(w, resp)
	}

	if err := userDomain.IsIdValid(id); err != nil {
		resp := response.BadRequest("", err.Error())
		response.WriteErrorResponse(w, resp)
		return
	}

	user, err := handler.useCases.Get(req.Context(), id)
	if err != nil {
		if err == userErrors.ErrIdInvalid || err == userErrors.ErrIdRequired {
			resp := response.BadRequest("", err.Error())
			response.WriteErrorResponse(w, resp)
			return
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

	resp := response.OK("", user)
	response.WriteSuccessResponse(w, resp)
}
