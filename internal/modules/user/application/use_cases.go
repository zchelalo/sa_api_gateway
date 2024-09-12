package userApplication

import (
	"context"
	"log"

	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
)

type UserUseCases struct {
	ctx            context.Context
	logger         *log.Logger
	userRepository userDomain.UserRepository
}

func NewUserUseCases(ctx context.Context, logger *log.Logger, userRepository userDomain.UserRepository) *UserUseCases {
	return &UserUseCases{
		ctx:            ctx,
		logger:         logger,
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
