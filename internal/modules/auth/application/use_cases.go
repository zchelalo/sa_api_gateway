package authApplication

import (
	"context"

	authDomain "github.com/zchelalo/sa_api_gateway/internal/modules/auth/domain"
	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
)

type AuthUseCases struct {
	authRepository authDomain.AuthRepository
}

func NewAuthUseCases(authRepository authDomain.AuthRepository) *AuthUseCases {
	return &AuthUseCases{
		authRepository: authRepository,
	}
}

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (authUseCases *AuthUseCases) SignIn(ctx context.Context, signInRequest *SignInRequest) (*authDomain.AuthEntity, error) {
	err := userDomain.IsEmailValid(signInRequest.Email)
	if err != nil {
		return nil, err
	}

	err = userDomain.IsPasswordValid(signInRequest.Password)
	if err != nil {
		return nil, err
	}

	return authUseCases.authRepository.SignIn(ctx, signInRequest.Email, signInRequest.Password)
}

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (authUseCases *AuthUseCases) SignUp(ctx context.Context, signUpRequest *SignUpRequest) (*authDomain.AuthEntity, error) {
	err := userDomain.IsNameValid(signUpRequest.Name)
	if err != nil {
		return nil, err
	}

	err = userDomain.IsEmailValid(signUpRequest.Email)
	if err != nil {
		return nil, err
	}

	err = userDomain.IsPasswordValid(signUpRequest.Password)
	if err != nil {
		return nil, err
	}

	return authUseCases.authRepository.SignUp(ctx, signUpRequest.Name, signUpRequest.Email, signUpRequest.Password)
}

func (authUseCases *AuthUseCases) SignOut(ctx context.Context, refreshToken string) error {
	err := authDomain.IsTokenValid(refreshToken, constants.RefreshToken)
	if err != nil {
		return err
	}

	return authUseCases.authRepository.SignOut(ctx, refreshToken)
}

func (authUseCases *AuthUseCases) IsAuthorized(ctx context.Context, accessToken, refreshToken string) (*authDomain.AuthorizeEntity, error) {
	err := authDomain.IsTokenValid(accessToken, constants.AccessToken)
	if err != nil {
		return nil, err
	}

	err = authDomain.IsTokenValid(refreshToken, constants.RefreshToken)
	if err != nil {
		return nil, err
	}

	return authUseCases.authRepository.IsAuthorized(ctx, accessToken, refreshToken)
}
