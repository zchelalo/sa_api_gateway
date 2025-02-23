package authApplication

import (
	"context"

	authDomain "github.com/zchelalo/sa_api_gateway/internal/modules/auth/domain"
	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
)

type SignInRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (useCases *UseCases) SignIn(ctx context.Context, signInRequest *SignInRequest) (*authDomain.AuthEntity, error) {
	err := userDomain.IsEmailValid(signInRequest.Email)
	if err != nil {
		return nil, err
	}

	err = userDomain.IsPasswordValid(signInRequest.Password)
	if err != nil {
		return nil, err
	}

	return useCases.authRepository.SignIn(ctx, signInRequest.Email, signInRequest.Password)
}
