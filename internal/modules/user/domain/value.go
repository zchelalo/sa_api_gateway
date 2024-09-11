package userDomain

import (
	"regexp"

	"github.com/google/uuid"
	userErrors "github.com/zchelalo/sa_api_gateway/internal/modules/user/errors"
)

const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

func IsIdValid(id string) error {
	if id == "" {
		return userErrors.ErrIdRequired
	}
	if _, err := uuid.Parse(id); err != nil {
		return userErrors.ErrIdInvalid
	}
	return nil
}

func IsNameValid(name string) error {
	if name == "" {
		return userErrors.ErrNameRequired
	}
	if len(name) < 3 {
		return userErrors.ErrNameInvalid
	}
	return nil
}

func IsEmailValid(email string) error {
	if email == "" {
		return userErrors.ErrEmailRequired
	}
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return userErrors.ErrEmailInvalid
	}
	return nil
}

func IsPasswordValid(password string) error {
	if password == "" {
		return userErrors.ErrPasswordRequired
	}
	if len(password) < 8 {
		return userErrors.ErrPasswordInvalid
	}
	return nil
}
