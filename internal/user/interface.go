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
		FindLikedUserByID(ctx context.Context, userID string) ([]*User, error)
		FindPassedUserByID(ctx context.Context, userID string) ([]*User, error)
		FindActionCounterByID(ctx context.Context, userID string) (int8, error)
		FindAllUser(ctx context.Context) ([]*User, error)
		CreateUser(ctx context.Context, user *User) (string, error)
	}
)
