package classManagementREST

import (
	"net/http"
	"strconv"

	classManagementApplication "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/application"
	classManagementError "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/error"
	userError "github.com/zchelalo/sa_api_gateway/internal/modules/user/error"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/response"
	"github.com/zchelalo/sa_api_gateway/pkg/util"
)

func (handler *Handler) List(w http.ResponseWriter, req *http.Request) {
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

	queries := req.URL.Query()

	page, err := strconv.Atoi(queries.Get("page"))
	if err != nil {
		resp := response.BadRequest("", "page is invalid")
		response.WriteErrorResponse(w, resp)
		return
	}

	limit, err := strconv.Atoi(queries.Get("limit"))
	if err != nil {
		resp := response.BadRequest("", "limit is invalid")
		response.WriteErrorResponse(w, resp)
		return
	}

	request := &classManagementApplication.ListRequest{
		UserID: id,
		Page:   int32(page),
		Limit:  int32(limit),
	}

	classes, meta, err := handler.useCases.List(req.Context(), request)
	if err != nil {
		badRequestErrors := []error{
			userError.ErrIdInvalid,
			userError.ErrIdRequired,

			classManagementError.ErrPageInvalid,
			classManagementError.ErrLimitInvalid,
		}
		if util.IsErrorType(err, badRequestErrors) {
			resp := response.BadRequest("", err.Error())
			response.WriteErrorResponse(w, resp)
			return
		}

		if err == userError.ErrUserNotFound || err == classManagementError.ErrClassesNotFound {
			resp := response.NotFound("", err.Error())
			response.WriteErrorResponse(w, resp)
			return
		}

		resp := response.InternalServerError("", err.Error())
		response.WriteErrorResponse(w, resp)
		return
	}

	resp := response.OK("", classes, meta)
	response.WriteSuccessResponse(w, resp)
}
