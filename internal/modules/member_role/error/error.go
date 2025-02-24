package memberRoleError

import "errors"

var ErrIdRequired = errors.New("id is required")
var ErrIdInvalid = errors.New("id is invalid")
