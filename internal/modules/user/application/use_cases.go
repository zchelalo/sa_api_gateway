package userApplication

import (
	"context"

	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
)

type UseCases struct {
	userRepository userDomain.UserRepository
}

func New(userRepository userDomain.UserRepository) *UseCases {
	return &UseCases{
		userRepository: userRepository,
	}
}

func (useCases *UseCases) Get(ctx context.Context, id string) (*userDomain.UserEntity, error) {
	err := userDomain.IsIdValid(id)
	if err != nil {
		return nil, err
	}

	return useCases.userRepository.Get(ctx, id)
}
