package auth

import "context"

type (
	Service interface {
		Login(ctx context.Context, body RequestLogin) (userAuth *Auth, err error)
	}

	Repository interface {
		FindByMobilePhone(ctx context.Context, mobilePhone string) (userAuth *Auth, err error)
		CreateAuth(ctx context.Context, auth *Auth) error
		GenerateToken(ctx context.Context, userID, pin string) error
		ValidateToken(ctx context.Context, token string) bool
		ValidatePin(ctx context.Context, userID, pin string) error
	}
)
