package authApplication

import (
	"context"

	authDomain "github.com/zchelalo/sa_api_gateway/internal/modules/auth/domain"
	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
)

type SignUpRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (useCases *UseCases) SignUp(ctx context.Context, signUpRequest *SignUpRequest) (*authDomain.AuthEntity, error) {
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

	return useCases.authRepository.SignUp(ctx, signUpRequest.Name, signUpRequest.Email, signUpRequest.Password)
}
