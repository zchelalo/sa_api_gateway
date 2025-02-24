package classManagementGRPCRepo

import (
	"context"
	"errors"

	classManagementDomain "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/domain"
	classManagementError "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/error"
	"github.com/zchelalo/sa_api_gateway/pkg/meta"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (r *GRPCRepository) List(ctx context.Context, userID string, page, limit int32) ([]*classManagementDomain.ClassEntity, *meta.Meta, error) {
	classes, err := r.classClient.ListClasses(ctx, &proto.ListClassesRequest{
		UserId: userID,
		Page:   page,
		Limit:  limit,
	})
	if err != nil {
		return nil, nil, err
	}

	errorObtained := classes.GetError()
	if errorObtained != nil {
		errorCode := errorObtained.GetCode()
		errorMessage := errorObtained.GetMessage()

		if int32(codes.InvalidArgument) == errorCode {
			return nil, nil, classManagementError.ErrDataInvalid
		}
		if int32(codes.NotFound) == errorCode {
			return nil, nil, classManagementError.ErrClassNotFound
		}
		if int32(codes.Internal) == errorCode {
			return nil, nil, errors.New(errorMessage)
		}

		return nil, nil, errors.New(errorMessage)
	}

	classesObtained := classes.GetData().GetClasses()
	if classesObtained == nil {
		return nil, nil, classManagementError.ErrClassNotFound
	}

	metaObtained := classes.GetData().GetMeta()
	if metaObtained == nil {
		return nil, nil, classManagementError.ErrNoMeta
	}

	var classesList []*classManagementDomain.ClassEntity
	for _, classObtained := range classesObtained {
		classesList = append(classesList, &classManagementDomain.ClassEntity{
			ID:      classObtained.GetId(),
			Name:    classObtained.GetName(),
			Subject: classObtained.GetSubject(),
			Grade:   classObtained.GetGrade(),
			Code:    classObtained.GetCode(),
		})
	}

	return classesList, &meta.Meta{
		Page:       metaObtained.GetPage(),
		PerPage:    metaObtained.GetPerPage(),
		PageCount:  metaObtained.GetCount(),
		TotalCount: metaObtained.GetTotalCount(),
	}, nil
}
