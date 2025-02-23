package authApplication

import (
	authDomain "github.com/zchelalo/sa_api_gateway/internal/modules/auth/domain"
)

type UseCases struct {
	authRepository authDomain.AuthRepository
}

func New(authRepository authDomain.AuthRepository) *UseCases {
	return &UseCases{
		authRepository: authRepository,
	}
}
