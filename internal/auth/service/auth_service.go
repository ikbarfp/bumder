package service

import (
	"context"
	"github.com/ikbarfp/bumder/internal/auth"
)

type authService struct {
	authRepository auth.Repository
}

// New . . .
func New(authRepo auth.Repository) auth.Service {
	return &authService{authRepository: authRepo}
}

// Login . . .
func (s authService) Login(ctx context.Context, body auth.RequestLogin) (*auth.Auth, error) {

	userAuth, err := s.authRepository.FindByMobilePhone(ctx, body.MobilePhone)
	if err != nil {
		return nil, err
	}

	if err = s.authRepository.ValidatePin(ctx, userAuth.UserID, body.Pin); err != nil {
		return nil, err
	}

	if err = s.authRepository.GenerateToken(ctx, userAuth.UserID, body.Pin); err != nil {
		return nil, err
	}

	return userAuth, nil
}
