package userApplication

import (
	"context"

	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
)

type UserUseCases struct {
	userRepository userDomain.UserRepository
}

func NewUserUseCases(userRepository userDomain.UserRepository) *UserUseCases {
	return &UserUseCases{
		userRepository: userRepository,
	}
}

func (userUseCases *UserUseCases) Get(ctx context.Context, id string) (*userDomain.UserEntity, error) {
	err := userDomain.IsIdValid(id)
	if err != nil {
		return nil, err
	}

	return userUseCases.userRepository.Get(ctx, id)
}
