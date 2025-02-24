package classManagementREST

import (
	"encoding/json"
	"net/http"

	classManagementApplication "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/application"
	classManagementError "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/error"
	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
	userError "github.com/zchelalo/sa_api_gateway/internal/modules/user/error"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/response"
	"github.com/zchelalo/sa_api_gateway/pkg/util"
)

func (handler *Handler) Create(w http.ResponseWriter, req *http.Request) {
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

	request := &classManagementApplication.CreateRequest{}
	if err := json.NewDecoder(req.Body).Decode(&request); err != nil {
		resp := response.BadRequest("", err.Error())
		response.WriteErrorResponse(w, resp)
		return
	}
	request.UserID = id

	class, err := handler.useCases.Create(req.Context(), request)
	if err != nil {
		badRequestErrors := []error{
			userError.ErrIdInvalid,
			userError.ErrIdRequired,

			classManagementError.ErrNameRequired,
			classManagementError.ErrNameTooShort,
			classManagementError.ErrGradeRequired,
			classManagementError.ErrGradeTooShort,
			classManagementError.ErrSubjectRequired,
			classManagementError.ErrSubjectTooShort,
		}
		if util.IsErrorType(err, badRequestErrors) {
			resp := response.BadRequest("", err.Error())
			response.WriteErrorResponse(w, resp)
			return
		}

		if err == userError.ErrUserNotFound {
			resp := response.NotFound("", err.Error())
			response.WriteErrorResponse(w, resp)
			return
		}

		resp := response.InternalServerError("", err.Error())
		response.WriteErrorResponse(w, resp)
		return
	}

	resp := response.OK("", class)
	response.WriteSuccessResponse(w, resp)
}
