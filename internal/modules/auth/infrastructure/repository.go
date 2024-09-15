package authInfrastructure

import (
	"context"
	"errors"

	authDomain "github.com/zchelalo/sa_api_gateway/internal/modules/auth/domain"
	authErrors "github.com/zchelalo/sa_api_gateway/internal/modules/auth/errors"
	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
	userErrors "github.com/zchelalo/sa_api_gateway/internal/modules/user/errors"
	authProto "github.com/zchelalo/sa_api_gateway/pkg/proto/auth"
	"google.golang.org/grpc/codes"
)

type GRPCRepository struct {
	ctx    context.Context
	client authProto.AuthServiceClient
}

func NewGRPCRepository(ctx context.Context, client authProto.AuthServiceClient) authDomain.AuthRepository {
	return &GRPCRepository{
		ctx:    ctx,
		client: client,
	}
}

func (r *GRPCRepository) SignIn(email, password string) (*authDomain.AuthEntity, error) {
	auth, err := r.client.SignIn(r.ctx, &authProto.SignInRequest{
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
		if int32(codes.NotFound) == errorCode {
			return nil, userErrors.ErrUserNotFound
		}
		if int32(codes.Unauthenticated) == errorCode || int32(codes.PermissionDenied) == errorCode {
			return nil, authErrors.ErrUnauthorized
		}
		if int32(codes.Internal) == errorCode {
			return nil, errors.New(errorMessage)
		}

		return nil, errors.New(errorMessage)
	}

	authObtained := auth.GetAuth()
	if authObtained == nil {
		return nil, userErrors.ErrUserNotFound
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
	}, nil
}

func (r *GRPCRepository) SignUp(name, email, password string) (*authDomain.AuthEntity, error) {
	auth, err := r.client.SignUp(r.ctx, &authProto.SignUpRequest{
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
	}, nil
}

func (r *GRPCRepository) SignOut(refreshToken string) error {
	_, err := r.client.SignOut(r.ctx, &authProto.SignOutRequest{
		RefreshToken: refreshToken,
	})
	if err != nil {
		return err
	}

	return nil
}

func (r *GRPCRepository) IsAuthorized(accessToken, refreshToken string) (*authDomain.AuthorizeEntity, error) {
	auth, err := r.client.IsAuthorized(r.ctx, &authProto.IsAuthorizedRequest{
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
		Tokens: authDomain.Tokens{
			AccessToken:  authObtained.Tokens.GetAccessToken(),
			RefreshToken: authObtained.Tokens.GetRefreshToken(),
		},
	}, nil
}
