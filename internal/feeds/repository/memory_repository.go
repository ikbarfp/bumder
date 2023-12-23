package repository

import (
	"context"
	"github.com/ikbarfp/bumder/internal/feeds"
)

type feedsRepository struct {
}

func New() feeds.Repository {
	return &feedsRepository{}
}

func (r feedsRepository) IncrementLikeByUserID(ctx context.Context, userID int, likedUserID int) error {

	return nil
}

func (r feedsRepository) IncrementPassByUserID(ctx context.Context, userID int, dislikedUserID int) error {

	return nil
}
