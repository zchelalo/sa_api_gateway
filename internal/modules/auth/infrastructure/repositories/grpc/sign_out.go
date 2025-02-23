package authGRPCRepo

import (
	"context"
	"errors"

	authError "github.com/zchelalo/sa_api_gateway/internal/modules/auth/errors"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (r *GRPCRepository) SignOut(ctx context.Context, refreshToken string) error {
	auth, err := r.client.SignOut(ctx, &proto.SignOutRequest{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return err
	}

	errorObtained := auth.GetError()
	if errorObtained != nil {
		errorCode := errorObtained.GetCode()
		errorMessage := errorObtained.GetMessage()

		if int32(codes.InvalidArgument) == errorCode {
			return authError.ErrTokenInvalid{Name: constants.RefreshToken}
		}
		if int32(codes.Internal) == errorCode {
			return errors.New(errorMessage)
		}

		return errors.New(errorMessage)
	}

	success := auth.GetSuccess()
	if !success {
		return authError.ErrSignOutFailed
	}

	return nil
}
