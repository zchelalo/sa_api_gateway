package classManagementApplication

import (
	"context"

	classManagementDomain "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/domain"
	memberDomain "github.com/zchelalo/sa_api_gateway/internal/modules/member/domain"
	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
	"github.com/zchelalo/sa_api_gateway/pkg/meta"
)

type ListMembersRequest struct {
	UserID  string `json:"user_id"`
	ClassID string `json:"class_id"`
	Page    int32  `json:"page"`
	Limit   int32  `json:"limit"`
}

func (useCases *UseCases) ListMembers(ctx context.Context, listMembersRequest *ListMembersRequest) ([]*memberDomain.MemberEntity, *meta.Meta, error) {
	if err := userDomain.IsIdValid(listMembersRequest.UserID); err != nil {
		return nil, nil, err
	}

	if err := classManagementDomain.IsIdValid(listMembersRequest.ClassID); err != nil {
		return nil, nil, err
	}

	if err := classManagementDomain.IsPageValid(listMembersRequest.Page); err != nil {
		return nil, nil, err
	}

	if err := classManagementDomain.IsLimitValid(listMembersRequest.Limit); err != nil {
		return nil, nil, err
	}

	return useCases.classManagementRepository.ListMembers(ctx, listMembersRequest.UserID, listMembersRequest.ClassID, listMembersRequest.Page, listMembersRequest.Limit)
}
