package service

import (
	"context"
	"github.com/ikbarfp/bumder/internal/feeds"
)

type feedsService struct {
	feedsRepository feeds.Repository
}

// New . . .
func New(feedsRepo feeds.Repository) feeds.Service {
	return &feedsService{feedsRepository: feedsRepo}
}

// Unseen . . .
func (s feedsService) Unseen(ctx context.Context) ([]*feeds.UnseenUser, error) {

	return nil, nil
}

// SwipeRight . . .
func (s feedsService) SwipeRight(ctx context.Context, userID int, likedUserID int) error {

	return nil
}

// SwipeLeft . . .
func (s feedsService) SwipeLeft(ctx context.Context, userID int, dislikedUserID int) error {

	return nil
}
