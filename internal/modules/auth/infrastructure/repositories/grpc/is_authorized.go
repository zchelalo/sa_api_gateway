package authGRPCRepo

import (
	"context"
	"errors"

	authDomain "github.com/zchelalo/sa_api_gateway/internal/modules/auth/domain"
	authErrors "github.com/zchelalo/sa_api_gateway/internal/modules/auth/errors"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (r *GRPCRepository) IsAuthorized(ctx context.Context, accessToken, refreshToken string) (*authDomain.AuthorizeEntity, error) {
	auth, err := r.client.IsAuthorized(ctx, &proto.IsAuthorizedRequest{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	})
	if err != nil {
		return nil, err
	}

	errorObtained := auth.GetError()
	if errorObtained != nil {
		errorCode := errorObtained.GetCode()
		errorMessage := errorObtained.GetMessage()

		if int32(codes.InvalidArgument) == errorCode {
			return nil, authErrors.ErrDataInvalid
		}
		if int32(codes.Unauthenticated) == errorCode || int32(codes.PermissionDenied) == errorCode {
			return nil, authErrors.ErrUnauthorized
		}
		if int32(codes.Internal) == errorCode {
			return nil, errors.New(errorMessage)
		}

		return nil, errors.New(errorMessage)
	}

	authObtained := auth.GetData()
	if authObtained == nil {
		return nil, authErrors.ErrUnauthorized
	}

	return &authDomain.AuthorizeEntity{
		IsAuthorized: authObtained.GetIsAuthorized(),
		UserID:       authObtained.GetUserId(),
		Tokens: authDomain.Tokens{
			AccessToken:  authObtained.Tokens.GetAccessToken(),
			RefreshToken: authObtained.Tokens.GetRefreshToken(),
			ExpiresAt:    authObtained.Tokens.GetExpiresAt(),
		},
	}, nil
}
