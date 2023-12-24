package service

import (
	"context"
	"github.com/ikbarfp/bumder/internal/feeds"
	"github.com/ikbarfp/bumder/internal/user"
	"github.com/ikbarfp/bumder/pkg/constant"
	"github.com/ikbarfp/bumder/pkg/response"
)

type feedsService struct {
	userRepository user.Repository
}

// New . . .
func New(userRepo user.Repository) feeds.Service {
	return &feedsService{userRepository: userRepo}
}

// Unseen . . .
func (s feedsService) Unseen(ctx context.Context, userID string) ([]*feeds.UnseenUser, error) {

	var res []*feeds.UnseenUser
	users, err := s.userRepository.FindAllUser(ctx)
	if err != nil {
		return nil, err
	}

	for _, usr := range users {
		res = append(res, &feeds.UnseenUser{
			UserID:    usr.ID,
			Name:      usr.Name,
			IsPremium: usr.IsPremium,
		})
	}

	return res, nil
}

// SwipeRight . . .
func (s feedsService) SwipeRight(ctx context.Context, userID, likedUserID string) error {

	isSufficient, err := s.isSufficientCredit(ctx, userID)
	if err != nil {
		return err
	}

	if !isSufficient {
		return response.ErrInsufficientCredit
	}

	return nil
}

// SwipeLeft . . .
func (s feedsService) SwipeLeft(ctx context.Context, userID, dislikedUserID string) error {

	isSufficient, err := s.isSufficientCredit(ctx, userID)
	if err != nil {
		return err
	}

	if !isSufficient {
		return response.ErrInsufficientCredit
	}

	return nil
}

func (s feedsService) isSufficientCredit(ctx context.Context, userID string) (bool, error) {

	counter, err := s.userRepository.FindActionCounterByID(ctx, userID)
	if err != nil {
		return false, err
	}

	return counter <= constant.DailyCreditLimit, nil
}
