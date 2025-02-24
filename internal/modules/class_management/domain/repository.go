package classManagementDomain

import (
	"context"

	memberDomain "github.com/zchelalo/sa_api_gateway/internal/modules/member/domain"
	"github.com/zchelalo/sa_api_gateway/pkg/meta"
)

type ClassManagementRepository interface {
	Create(ctx context.Context, userID, name, grade, subject string) (*ClassEntity, error)
	Join(ctx context.Context, userID, code string) (*ClassEntity, error)
	List(ctx context.Context, userID string, page, limit int32) ([]*ClassEntity, *meta.Meta, error)
	ListMembers(ctx context.Context, userID, classID string, page, limit int32) ([]*memberDomain.MemberEntity, *meta.Meta, error)
	GetClassCode(ctx context.Context, userID, classID string) (string, error)
}
