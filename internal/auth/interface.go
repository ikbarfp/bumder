package auth

import "context"

type (
	// Usecase . . .
	Usecase interface {
		Login(ctx context.Context, userID int) error
	}

	// Repository . . .
	Repository interface {
	}
)
