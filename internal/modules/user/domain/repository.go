package userDomain

type UserRepository interface {
	Get(id string) (*UserEntity, error)
}
