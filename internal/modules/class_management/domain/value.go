package classManagementDomain

import (
	"github.com/google/uuid"
	classManagementError "github.com/zchelalo/sa_api_gateway/internal/modules/class_management/error"
)

const codeLength = 6

func IsIdValid(id string) error {
	if id == "" {
		return classManagementError.ErrIdInvalid
	}

	if _, err := uuid.Parse(id); err != nil {
		return classManagementError.ErrIdInvalid
	}

	return nil
}

func IsNameValid(name string) error {
	if name == "" {
		return classManagementError.ErrNameRequired
	}

	if len(name) < 3 {
		return classManagementError.ErrNameTooShort
	}

	return nil
}

func IsSubjectValid(subject string) error {
	if subject == "" {
		return classManagementError.ErrSubjectRequired
	}

	if len(subject) < 2 {
		return classManagementError.ErrSubjectTooShort
	}

	return nil
}

func IsGradeValid(grade string) error {
	if grade == "" {
		return classManagementError.ErrGradeRequired
	}

	if len(grade) < 1 {
		return classManagementError.ErrGradeTooShort
	}

	return nil
}

func IsPageValid(page int32) error {
	if page < 1 {
		return classManagementError.ErrPageInvalid
	}

	return nil
}

func IsLimitValid(limit int32) error {
	if limit < 1 {
		return classManagementError.ErrLimitInvalid
	}

	return nil
}

func IsCodeValid(code string) error {
	if code == "" {
		return classManagementError.ErrCodeRequired
	}

	if len(code) != codeLength {
		return classManagementError.ErrCodeInvalid
	}

	return nil
}
