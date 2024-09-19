package authDomain

import "context"

type AuthRepository interface {
	SignIn(ctx context.Context, email, password string) (*AuthEntity, error)
	SignUp(ctx context.Context, name, email, password string) (*AuthEntity, error)
	SignOut(ctx context.Context, refreshToken string) error
	IsAuthorized(ctx context.Context, accessToken, refreshToken string) (*AuthorizeEntity, error)
}
