package classManagementApplication

import classManagementDomain "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/domain"

type UseCases struct {
	classManagementRepository classManagementDomain.ClassManagementRepository
}

func New(classManagementRepository classManagementDomain.ClassManagementRepository) *UseCases {
	return &UseCases{
		classManagementRepository: classManagementRepository,
	}
}
