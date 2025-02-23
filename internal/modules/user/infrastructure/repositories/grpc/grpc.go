package userGRPCRepo

import (
	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
)

type GRPCRepository struct {
	client proto.UserServiceClient
}

func New(client proto.UserServiceClient) userDomain.UserRepository {
	return &GRPCRepository{
		client: client,
	}
}
