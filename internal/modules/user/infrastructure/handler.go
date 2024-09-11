package userInfrastructure

import (
	"context"
	"net/http"

	userApplication "github.com/zchelalo/sa_api_gateway/internal/modules/user/application"
	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
	userErrors "github.com/zchelalo/sa_api_gateway/internal/modules/user/errors"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/response"
)

type UserHandler struct {
	ctx          context.Context
	userUseCases *userApplication.UserUseCases
}

func NewUserHandler(ctx context.Context, userUseCases *userApplication.UserUseCases) *UserHandler {
	return &UserHandler{
		ctx:          ctx,
		userUseCases: userUseCases,
	}
}

func (userHandler *UserHandler) Get(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue(string(constants.ID))

	if err := userDomain.IsIdValid(id); err != nil {
		resp := response.BadRequest("", err.Error())
		response.WriteErrorResponse(w, resp)
		return
	}

	user, err := userHandler.userUseCases.Get(id)
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
