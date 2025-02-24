package classManagementGRPCRepo

import (
	"context"
	"errors"

	classManagementDomain "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/domain"
	classManagementError "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/error"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (r *GRPCRepository) Join(ctx context.Context, userID, code string) (*classManagementDomain.ClassEntity, error) {
	class, err := r.classClient.JoinClass(ctx, &proto.JoinClassRequest{
		UserId: userID,
		Code:   code,
	})
	if err != nil {
		return nil, err
	}

	errorObtained := class.GetError()
	if errorObtained != nil {
		errorCode := errorObtained.GetCode()
		errorMessage := errorObtained.GetMessage()

		if int32(codes.InvalidArgument) == errorCode {
			return nil, classManagementError.ErrDataInvalid
		}
		if int32(codes.NotFound) == errorCode {
			return nil, classManagementError.ErrClassNotFound
		}
		if int32(codes.Internal) == errorCode {
			return nil, errors.New(errorMessage)
		}

		return nil, errors.New(errorMessage)
	}

	classObtained := class.GetClass()
	if classObtained == nil {
		return nil, classManagementError.ErrClassNotFound
	}

	return &classManagementDomain.ClassEntity{
		ID:      classObtained.GetId(),
		Name:    classObtained.GetName(),
		Subject: classObtained.GetSubject(),
		Grade:   classObtained.GetGrade(),
		Code:    classObtained.GetCode(),
	}, nil
}
