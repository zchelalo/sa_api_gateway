package userDomain

import (
	"regexp"

	"github.com/google/uuid"
	userError "github.com/zchelalo/sa_api_gateway/internal/modules/user/error"
)

const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func IsIdValid(id string) error {
	if id == "" {
		return userError.ErrIdRequired
	}
	if _, err := uuid.Parse(id); err != nil {
		return userError.ErrIdInvalid
	}
	return nil
}

func IsNameValid(name string) error {
	if name == "" {
		return userError.ErrNameRequired
	}
	if len(name) < 3 {
		return userError.ErrNameInvalid
	}
	return nil
}

func IsEmailValid(email string) error {
	if email == "" {
		return userError.ErrEmailRequired
	}
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return userError.ErrEmailInvalid
	}
	return nil
}

func IsPasswordValid(password string) error {
	if password == "" {
		return userError.ErrPasswordRequired
	}
	if len(password) < 8 {
		return userError.ErrPasswordInvalid
	}
	return nil
}
