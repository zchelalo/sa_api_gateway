package classManagementGRPCRepo

import (
	classManagementDomain "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/domain"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
)

type GRPCRepository struct {
	classClient  proto.ClassServiceClient
	memberClient proto.MemberServiceClient
}

func New(classClient proto.ClassServiceClient, memberClient proto.MemberServiceClient) classManagementDomain.ClassManagementRepository {
	return &GRPCRepository{
		classClient:  classClient,
		memberClient: memberClient,
	}
}
