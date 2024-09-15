package authDomain

import (
	"regexp"

	authErrors "github.com/zchelalo/sa_api_gateway/internal/modules/auth/errors"
	"github.com/zchelalo/sa_api_gateway/pkg/constants"
)

const tokenRegex = `^[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+\.[A-Za-z0-9-_]+$`

func IsTokenValid(token string, tokenType constants.TokenConstants) error {
	if token == "" {
		return &authErrors.ErrTokenRequired{
			Name: tokenType,
		}
	}
	re := regexp.MustCompile(tokenRegex)
	if !re.MatchString(token) {
		return &authErrors.ErrTokenInvalid{
			Name: tokenType,
		}
	}
	return nil
}
