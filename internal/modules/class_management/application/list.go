package classManagementApplication

import (
	"context"

	classManagementDomain "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/domain"
	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
	"github.com/zchelalo/sa_api_gateway/pkg/meta"
)

type ListRequest struct {
	UserID string `json:"user_id"`
	Page   int32  `json:"page"`
	Limit  int32  `json:"limit"`
}

func (useCases *UseCases) List(ctx context.Context, listRequest *ListRequest) ([]*classManagementDomain.ClassEntity, *meta.Meta, error) {
	if err := userDomain.IsIdValid(listRequest.UserID); err != nil {
		return nil, nil, err
	}

	if err := classManagementDomain.IsPageValid(listRequest.Page); err != nil {
		return nil, nil, err
	}

	if err := classManagementDomain.IsLimitValid(listRequest.Limit); err != nil {
		return nil, nil, err
	}

	return useCases.classManagementRepository.List(ctx, listRequest.UserID, listRequest.Page, listRequest.Limit)
}
