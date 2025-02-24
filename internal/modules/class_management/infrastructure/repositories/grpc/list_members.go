package classManagementGRPCRepo

import (
	"context"
	"errors"

	classManagementError "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/error"
	memberDomain "github.com/zchelalo/sa_api_gateway/internal/modules/member/domain"
	memberError "github.com/zchelalo/sa_api_gateway/internal/modules/member/error"
	memberRoleDomain "github.com/zchelalo/sa_api_gateway/internal/modules/member_role/domain"
	userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"
	"github.com/zchelalo/sa_api_gateway/pkg/meta"
	"github.com/zchelalo/sa_api_gateway/pkg/proto"
	"google.golang.org/grpc/codes"
)

func (r *GRPCRepository) ListMembers(ctx context.Context, userID, classID string, page, limit int32) ([]*memberDomain.MemberEntity, *meta.Meta, error) {
	members, err := r.memberClient.ListMembers(ctx, &proto.ListMembersRequest{
		UserId:  userID,
		ClassId: classID,
		Page:    page,
		Limit:   limit,
	})
	if err != nil {
		return nil, nil, err
	}

	errorObtained := members.GetError()
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

	membersObtained := members.GetData().GetMembers()
	if membersObtained == nil {
		return nil, nil, memberError.ErrMembersNotFound
	}

	metaObtained := members.GetData().GetMeta()
	if metaObtained == nil {
		return nil, nil, memberError.ErrNoMeta
	}

	var membersList []*memberDomain.MemberEntity
	for _, memberObtained := range membersObtained {
		membersList = append(membersList, &memberDomain.MemberEntity{
			ID: memberObtained.GetId(),
			User: userDomain.UserEntity{
				ID:       memberObtained.GetUser().GetId(),
				Name:     memberObtained.GetUser().GetName(),
				Email:    memberObtained.GetUser().GetEmail(),
				Verified: memberObtained.GetUser().GetVerified(),
			},
			Role: memberRoleDomain.MemberRoleEntity{
				ID:  memberObtained.GetRole().GetId(),
				Key: memberObtained.GetRole().GetKey(),
			},
		})
	}

	return membersList, &meta.Meta{
		Page:       metaObtained.GetPage(),
		PerPage:    metaObtained.GetPerPage(),
		PageCount:  metaObtained.GetCount(),
		TotalCount: metaObtained.GetTotalCount(),
	}, nil
}
