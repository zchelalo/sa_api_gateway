package classManagementREST

import (
	"net/http"

	classManagementApplication "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/application"
	classManagementError "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/error"
	memberError "github.com/zchelalo/sa_api_gateway/internal/modules/member/error"
	userError "github.com/zchelalo/sa_api_gateway/internal/modules/user/error"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/response"
	"github.com/zchelalo/sa_api_gateway/pkg/util"
)

func (handler *Handler) GetClassCode(w http.ResponseWriter, req *http.Request) {
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

	classID := req.PathValue("classID")

	request := &classManagementApplication.GetClassCodeRequest{
		UserID:  id,
		ClassID: classID,
	}

	classCode, err := handler.useCases.GetClassCode(req.Context(), request)
	if err != nil {
		badRequestErrors := []error{
			userError.ErrIdInvalid,
			userError.ErrIdRequired,

			classManagementError.ErrIdRequired,
			classManagementError.ErrIdInvalid,
		}
		if util.IsErrorType(err, badRequestErrors) {
			resp := response.BadRequest("", err.Error())
			response.WriteErrorResponse(w, resp)
			return
		}

		if err == userError.ErrUserNotFound || err == classManagementError.ErrClassNotFound || err == memberError.ErrMemberNotFound {
			resp := response.NotFound("", err.Error())
			response.WriteErrorResponse(w, resp)
			return
		}

		if err == classManagementError.ErrUnauthorized {
			resp := response.Unauthorized("", err.Error())
			response.WriteErrorResponse(w, resp)
			return
		}

		resp := response.InternalServerError("", err.Error())
		response.WriteErrorResponse(w, resp)
		return
	}

	resp := response.OK("", classCode, nil)
	response.WriteSuccessResponse(w, resp)
}
