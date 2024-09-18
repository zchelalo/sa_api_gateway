package util

import (
	"net/http"
	"time"

	"github.com/zchelalo/sa_api_gateway/pkg/constants"
)

func secondsToTime(seconds int64) time.Time {
	return time.Now().Add(time.Duration(seconds) * time.Second)
}

func CreateCookie(name constants.CookieConstants, value string, expiresAt int64) *http.Cookie {
	expires := secondsToTime(expiresAt)
	return &http.Cookie{
		Name:     string(name),
		Value:    value,
		HttpOnly: true,
		SameSite: http.SameSiteNoneMode,
		Secure:   false,
		Expires:  expires,
		Path:     "/",
	}
}
