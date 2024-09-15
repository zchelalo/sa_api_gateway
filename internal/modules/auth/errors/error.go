package authErrors

import (
	"errors"
	"fmt"

	"github.com/zchelalo/sa_api_gateway/pkg/constants"
)

var ErrUnauthorized = errors.New("unauthorized")
var ErrDataInvalid = errors.New("data invalid")
var ErrSignUpFailed = errors.New("sign up failed")

type ErrTokenRequired struct {
	Name constants.TokenConstants
}

func (e *ErrTokenRequired) Error() string {
	return fmt.Sprintf("%s token required", e.Name)
}

type ErrTokenInvalid struct {
	Name constants.TokenConstants
}

func (e *ErrTokenInvalid) Error() string {
	return fmt.Sprintf("%s token invalid", e.Name)
}

type ErrTokenExpired struct {
	Name constants.TokenConstants
}

func (e *ErrTokenExpired) Error() string {
	return fmt.Sprintf("%s token is expired", e.Name)
}
