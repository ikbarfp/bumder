package service

import (
	"context"
	"github.com/ikbarfp/bumder/internal/auth"
	"github.com/ikbarfp/bumder/internal/user"
)

type userService struct {
	userRepository user.Repository
	authRepository auth.Repository
}

// New . . .
func New(userRepo user.Repository, authRepo auth.Repository) user.Service {
	return &userService{userRepository: userRepo, authRepository: authRepo}
}

// Register . . .
func (s userService) Register(ctx context.Context, body user.RequestRegister) error {

	newUser := &user.User{
		Name:               body.FullName,
		IsPremium:          false,
		DailyActionCounter: 0,
	}

	userID, err := s.userRepository.CreateUser(ctx, newUser)
	if err != nil {
		return err
	}

	newAuth := &auth.Auth{
		UserID:      userID,
		MobilePhone: body.MobilePhone,
		Pin:         body.Pin,
	}

	if err = s.authRepository.CreateAuth(ctx, newAuth); err != nil {
		return err
	}

	return nil
}

// GetProfile . . .
func (s userService) GetProfile(ctx context.Context, userID string) (*user.User, error) {

	usr, err := s.userRepository.FindByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	return usr, nil
}
