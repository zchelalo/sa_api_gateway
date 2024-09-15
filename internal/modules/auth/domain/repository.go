package authDomain

type AuthRepository interface {
	SignIn(email, password string) (*AuthEntity, error)
	SignUp(name, email, password string) (*AuthEntity, error)
	SignOut(refreshToken string) error
	IsAuthorized(accessToken, refreshToken string) (*AuthorizeEntity, error)
}
