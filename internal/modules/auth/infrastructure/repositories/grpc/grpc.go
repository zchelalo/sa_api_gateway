package authGRPCRepo

import (
	authDomain "github.com/zchelalo/sa_api_gateway/internal/modules/auth/domain"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
)

type GRPCRepository struct {
	client proto.AuthServiceClient
}

func New(client proto.AuthServiceClient) authDomain.AuthRepository {
	return &GRPCRepository{
		client: client,
	}
}
