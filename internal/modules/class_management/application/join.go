package classManagementApplication

import (
	"context"

	classManagementDomain "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/domain"
	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
)

type JoinRequest struct {
	UserID string `json:"user_id"`
	Code   string `json:"code"`
}

func (useCases *UseCases) Join(ctx context.Context, joinRequest *JoinRequest) (*classManagementDomain.ClassEntity, error) {
	if err := userDomain.IsIdValid(joinRequest.UserID); err != nil {
		return nil, err
	}

	if err := classManagementDomain.IsCodeValid(joinRequest.Code); err != nil {
		return nil, err
	}

	return useCases.classManagementRepository.Join(ctx, joinRequest.UserID, joinRequest.Code)
}
