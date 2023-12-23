package user

import "context"

type (
	// Usecase . . .
	Usecase interface {
		Register(ctx context.Context, user *User) error
		GetProfile(ctx context.Context, id int)
	}

	// Repository . . .
	Repository interface {
		// FindByName . . .
		FindByName(ctx context.Context, name string) (*User, error)

		// FindByID . . .
		FindByID(ctx context.Context, id int) (*User, error)

		// CreateUser . . .
		CreateUser(ctx context.Context) error
	}
)
