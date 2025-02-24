package classManagementApplication

import (
	"context"

	classManagementDomain "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/domain"
	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
)

type GetClassCodeRequest struct {
	UserID  string `json:"user_id"`
	ClassID string `json:"class_id"`
}

func (useCases *UseCases) GetClassCode(ctx context.Context, getClassCodeRequest *GetClassCodeRequest) (string, error) {
	if err := userDomain.IsIdValid(getClassCodeRequest.UserID); err != nil {
		return "", err
	}

	if err := classManagementDomain.IsIdValid(getClassCodeRequest.ClassID); err != nil {
		return "", err
	}

	return useCases.classManagementRepository.GetClassCode(ctx, getClassCodeRequest.UserID, getClassCodeRequest.ClassID)
}
