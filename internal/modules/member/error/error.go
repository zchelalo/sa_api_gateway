package memberError

import "errors"

var (
	ErrMemberNotFound  = errors.New("member not found")
	ErrMembersNotFound = errors.New("members not found")

	ErrNoMeta = errors.New("no meta")
)
