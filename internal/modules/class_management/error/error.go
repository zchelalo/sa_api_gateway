package classManagementError

import "errors"

var (
	ErrNameRequired = errors.New("name is required")
	ErrNameTooShort = errors.New("name must be at least 3 characters")

	ErrSubjectRequired = errors.New("subject is required")
	ErrSubjectTooShort = errors.New("subject must be at least 2 characters")

	ErrGradeRequired = errors.New("grade is required")
	ErrGradeTooShort = errors.New("grade must be at least 1 character")

	ErrClassNotFound   = errors.New("class not found")
	ErrClassesNotFound = errors.New("classes not found")

	ErrPageInvalid  = errors.New("page is invalid")
	ErrLimitInvalid = errors.New("limit is invalid")

	ErrCodeRequired = errors.New("code is required")
	ErrCodeInvalid  = errors.New("code is invalid")

	ErrIdRequired = errors.New("id is required")
	ErrIdInvalid  = errors.New("id is invalid")

	ErrUnauthorized = errors.New("unauthorized access")

	ErrAlreadyJoined = errors.New("already joined")

	ErrDataInvalid = errors.New("data invalid")

	ErrNoMeta = errors.New("no meta")
)
