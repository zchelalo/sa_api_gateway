package classManagementGRPCRepo

import (
	"context"
	"errors"

	classManagementError "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/error"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (r *GRPCRepository) GetClassCode(ctx context.Context, userID, classID string) (string, error) {
	classCode, err := r.classClient.GetClassCode(ctx, &proto.GetClassCodeRequest{
		UserId:  userID,
		ClassId: classID,
	})
	if err != nil {
		return "", err
	}

	errorObtained := classCode.GetError()
	if errorObtained != nil {
		errorCode := errorObtained.GetCode()
		errorMessage := errorObtained.GetMessage()

		if int32(codes.InvalidArgument) == errorCode {
			return "", classManagementError.ErrDataInvalid
		}
		if int32(codes.NotFound) == errorCode {
			return "", classManagementError.ErrClassNotFound
		}
		if int32(codes.Unauthenticated) == errorCode {
			return "", classManagementError.ErrUnauthorized
		}
		if int32(codes.PermissionDenied) == errorCode {
			return "", classManagementError.ErrUnauthorized
		}
		if int32(codes.Internal) == errorCode {
			return "", errors.New(errorMessage)
		}

		return "", errors.New(errorMessage)
	}

	classCodeObtained := classCode.GetCode()
	if classCodeObtained == "" {
		return "", classManagementError.ErrClassNotFound
	}

	return classCodeObtained, nil
}
