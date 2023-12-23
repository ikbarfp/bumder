package user

import (
	"context"
)

type (
	Service interface {
		Register(ctx context.Context, body RequestRegister) error
		GetProfile(ctx context.Context, id string) (*User, error)
	}

	Repository interface {
		FindByID(ctx context.Context, id string) (*User, error)
		CreateUser(ctx context.Context, user *User) (string, error)
	}
)
