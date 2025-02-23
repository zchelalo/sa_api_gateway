package userGRPCRepo

import (
	"context"
	"errors"

	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
	userErrors "github.com/zchelalo/sa_api_gateway/internal/modules/user/errors"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (r *GRPCRepository) Get(ctx context.Context, id string) (*userDomain.UserEntity, error) {
	user, err := r.client.GetUser(ctx, &proto.GetUserRequest{
		Id: id,
	})
	if err != nil {
		return nil, err
	}

	errorObtained := user.GetError()
	if errorObtained != nil {
		errorCode := errorObtained.GetCode()
		errorMessage := errorObtained.GetMessage()

		if int32(codes.InvalidArgument) == errorCode {
			return nil, userErrors.ErrIdInvalid
		}
		if int32(codes.NotFound) == errorCode {
			return nil, userErrors.ErrUserNotFound
		}
		if int32(codes.Internal) == errorCode {
			return nil, errors.New(errorMessage)
		}

		return nil, errors.New(errorMessage)
	}

	userObtained := user.GetUser()
	if userObtained == nil {
		return nil, userErrors.ErrUserNotFound
	}

	return &userDomain.UserEntity{
		ID:       userObtained.GetId(),
		Name:     userObtained.GetName(),
		Email:    userObtained.GetEmail(),
		Verified: userObtained.GetVerified(),
	}, nil
}
