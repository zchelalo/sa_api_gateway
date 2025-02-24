package classManagementApplication

import (
	"context"

	classManagementDomain "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/domain"
	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
)

type CreateRequest struct {
	UserID  string `json:"user_id"`
	Name    string `json:"name"`
	Grade   string `json:"grade"`
	Subject string `json:"subject"`
}

func (useCases *UseCases) Create(ctx context.Context, createRequest *CreateRequest) (*classManagementDomain.ClassEntity, error) {
	if err := userDomain.IsIdValid(createRequest.UserID); err != nil {
		return nil, err
	}

	if err := classManagementDomain.IsNameValid(createRequest.Name); err != nil {
		return nil, err
	}

	if err := classManagementDomain.IsGradeValid(createRequest.Grade); err != nil {
		return nil, err
	}

	if err := classManagementDomain.IsSubjectValid(createRequest.Subject); err != nil {
		return nil, err
	}

	return useCases.classManagementRepository.Create(ctx, createRequest.UserID, createRequest.Name, createRequest.Grade, createRequest.Subject)
}
