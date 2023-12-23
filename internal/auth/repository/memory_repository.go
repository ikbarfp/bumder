package repository

import (
	"context"
	"encoding/base64"
	"github.com/ikbarfp/bumder/internal/auth"
	"github.com/ikbarfp/bumder/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

var (
	authData []*auth.Auth
)

type authRepository struct {
}

// New . . .
func New() auth.Repository {
	return &authRepository{}
}

// FindByMobilePhone . . .
func (r authRepository) FindByMobilePhone(ctx context.Context, mobilePhone string) (*auth.Auth, error) {

	for _, userAuth := range authData {
		if userAuth.MobilePhone == mobilePhone {
			return userAuth, nil
		}
	}

	return nil, response.ErrUserNotFound
}

// CreateAuth . . .
func (r authRepository) CreateAuth(ctx context.Context, newAuth *auth.Auth) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(newAuth.Pin), bcrypt.DefaultCost)
	if err != nil {
		return response.ErrUnknown
	}

	newAuth.Pin = string(pass)
	authData = append(authData, newAuth)

	return nil
}

// GenerateToken . . .
func (r authRepository) GenerateToken(ctx context.Context, userID, pin string) error {

	for _, userAuth := range authData {
		if userAuth.UserID == userID {
			src := []byte(userID + ":" + pin)
			userAuth.Token = base64.StdEncoding.EncodeToString(src)

			return nil
		}
	}

	return response.ErrUserNotFound
}

// ValidateToken . . .
func (r authRepository) ValidateToken(ctx context.Context, token string) bool {

	for _, userAuth := range authData {
		if userAuth.Token == token {
			return true
		}
	}

	return false
}

// ValidatePin . . .
func (r authRepository) ValidatePin(ctx context.Context, userID, pin string) error {

	for _, userAuth := range authData {
		if userAuth.UserID == userID {
			if err := bcrypt.CompareHashAndPassword([]byte(userAuth.Pin), []byte(pin)); err != nil {
				return response.ErrInvalidPin
			}

			return nil
		}
	}

	return response.ErrUserNotFound
}
