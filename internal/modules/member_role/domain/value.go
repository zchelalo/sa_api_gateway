package memberRoleDomain

import (
	"github.com/google/uuid"
	memberRoleError "github.com/zchelalo/sa_api_gateway/internal/modules/member_role/error"
)

func IsIdValid(id string) error {
	if id == "" {
		return memberRoleError.ErrIdRequired
	}
	if _, err := uuid.Parse(id); err != nil {
		return memberRoleError.ErrIdInvalid
	}
	return nil
}
