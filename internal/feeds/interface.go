package feeds

import "context"

type (
	Service interface {
		Unseen(ctx context.Context) ([]*UnseenUser, error)
		SwipeRight(ctx context.Context, userID int, likedUserID int) error
		SwipeLeft(ctx context.Context, userID int, dislikedUserID int) error
	}

	Repository interface {
		IncrementLikeByUserID(ctx context.Context, userID int, likedUserID int) error
		IncrementPassByUserID(ctx context.Context, userID int, dislikedUserID int) error
	}
)
