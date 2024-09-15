package authApplication

import (
	"context"

	authDomain "github.com/zchelalo/sa_api_gateway/internal/modules/auth/domain"
	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
)

type AuthUseCases struct {
	ctx            context.Context
	authRepository authDomain.AuthRepository
}

func NewAuthUseCases(ctx context.Context, authRepository authDomain.AuthRepository) *AuthUseCases {
	return &AuthUseCases{
		ctx:            ctx,
		authRepository: authRepository,
	}
}

func (authUseCases *AuthUseCases) SignIn(email, password string) (*authDomain.AuthEntity, error) {
	err := userDomain.IsEmailValid(email)
	if err != nil {
		return nil, err
	}

	err = userDomain.IsPasswordValid(password)
	if err != nil {
		return nil, err
	}

	return authUseCases.authRepository.SignIn(email, password)
}

func (authUseCases *AuthUseCases) SignUp(name, email, password string) (*authDomain.AuthEntity, error) {
	err := userDomain.IsNameValid(name)
	if err != nil {
		return nil, err
	}

	err = userDomain.IsEmailValid(email)
	if err != nil {
		return nil, err
	}

	err = userDomain.IsPasswordValid(password)
	if err != nil {
		return nil, err
	}

	return authUseCases.authRepository.SignUp(name, email, password)
}

func (authUseCases *AuthUseCases) SignOut(refreshToken string) error {
	err := authDomain.IsTokenValid(refreshToken, constants.RefreshToken)
	if err != nil {
		return err
	}

	return authUseCases.authRepository.SignOut(refreshToken)
}

func (authUseCases *AuthUseCases) IsAuthorized(accessToken, refreshToken string) (*authDomain.AuthorizeEntity, error) {
	err := authDomain.IsTokenValid(accessToken, constants.AccessToken)
	if err != nil {
		return nil, err
	}

	err = authDomain.IsTokenValid(refreshToken, constants.RefreshToken)
	if err != nil {
		return nil, err
	}

	return authUseCases.authRepository.IsAuthorized(accessToken, refreshToken)
}
