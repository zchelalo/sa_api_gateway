package authDomain

import userDomain "github.com/zchelalo/sa_api_gateway/internal/modules/user/domain"

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

type AuthEntity struct {
	AccessToken  string                `json:"access_token"`
	RefreshToken string                `json:"refresh_token"`
	ExpiresAt    int64                 `json:"expires_at"`
	User         userDomain.UserEntity `json:"user"`
}

type AuthorizeEntity struct {
	Tokens       Tokens `json:"tokens"`
	IsAuthorized bool   `json:"is_authorized"`
}
