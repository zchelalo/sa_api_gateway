package userDomain

import "context"

type UserRepository interface {
	Get(ctx context.Context, id string) (*UserEntity, error)
}
