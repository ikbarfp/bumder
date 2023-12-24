package feeds

import "context"

type (
	Service interface {
		Unseen(ctx context.Context, userID string) ([]*UnseenUser, error)
		SwipeRight(ctx context.Context, userID string, likedUserID string) error
		SwipeLeft(ctx context.Context, userID string, dislikedUserID string) error
	}
)
