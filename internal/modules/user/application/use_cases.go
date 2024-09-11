package userApplication

import (
	"context"

	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
)

type UserUseCases struct {
	ctx            context.Context
	userRepository userDomain.UserRepository
}

func NewUserUseCases(ctx context.Context, userRepository userDomain.UserRepository) *UserUseCases {
	return &UserUseCases{
		ctx:            ctx,
		userRepository: userRepository,
	}
}

func (userUseCases *UserUseCases) Get(id string) (*userDomain.UserEntity, error) {
	err := userDomain.IsIdValid(id)
	if err != nil {
		return nil, err
	}

	return userUseCases.userRepository.Get(id)
}
