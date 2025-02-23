package authGRPCRepo

import (
	"context"
	"errors"

	authDomain "github.com/zchelalo/sa_api_gateway/internal/modules/auth/domain"
	authErrors "github.com/zchelalo/sa_api_gateway/internal/modules/auth/errors"
	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
	userErrors "github.com/zchelalo/sa_api_gateway/internal/modules/user/errors"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (r *GRPCRepository) SignUp(ctx context.Context, name, email, password string) (*authDomain.AuthEntity, error) {
	auth, err := r.client.SignUp(ctx, &proto.SignUpRequest{
		Name:     name,
		Email:    email,
		Password: password,
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
		if int32(codes.AlreadyExists) == errorCode {
			return nil, userErrors.ErrEmailAlreadyExists
		}
		if int32(codes.Internal) == errorCode {
			return nil, errors.New(errorMessage)
		}

		return nil, errors.New(errorMessage)
	}

	authObtained := auth.GetAuth()
	if authObtained == nil {
		return nil, authErrors.ErrSignUpFailed
	}

	user := userDomain.UserEntity{
		ID:       authObtained.User.GetId(),
		Name:     authObtained.User.GetName(),
		Email:    authObtained.User.GetEmail(),
		Verified: authObtained.User.GetVerified(),
	}

	return &authDomain.AuthEntity{
		User:         user,
		AccessToken:  authObtained.GetAccessToken(),
		RefreshToken: authObtained.GetRefreshToken(),
		ExpiresAt:    authObtained.GetExpiresAt(),
	}, nil
}
