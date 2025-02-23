package authApplication

import (
	"context"

	authDomain "github.com/zchelalo/sa_api_gateway/internal/modules/auth/domain"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
)

func (useCases *UseCases) SignOut(ctx context.Context, refreshToken string) error {
	err := authDomain.IsTokenValid(refreshToken, constants.RefreshToken)
	if err != nil {
		return err
	}

	return useCases.authRepository.SignOut(ctx, refreshToken)
}
