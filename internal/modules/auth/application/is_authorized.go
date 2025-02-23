package authApplication

import (
	"context"

	authDomain "github.com/zchelalo/sa_api_gateway/internal/modules/auth/domain"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
)

func (useCases *UseCases) IsAuthorized(ctx context.Context, accessToken, refreshToken string) (*authDomain.AuthorizeEntity, error) {
	err := authDomain.IsTokenValid(accessToken, constants.AccessToken)
	if err != nil {
		return nil, err
	}

	err = authDomain.IsTokenValid(refreshToken, constants.RefreshToken)
	if err != nil {
		return nil, err
	}

	return useCases.authRepository.IsAuthorized(ctx, accessToken, refreshToken)
}
