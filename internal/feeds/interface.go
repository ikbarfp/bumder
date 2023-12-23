package feeds

import "context"

type (
	// Usecase . . .
	Usecase interface {
		SwipeRight(ctx context.Context, userID int, likedUserID int) error
		SwipeLeft(ctx context.Context, userID int, dislikedUserID int) error
	}

	// Repository . . .
	Repository interface {
		IncrementLikeByID(ctx context.Context, userID int, likedUserID int) error
		IncrementDislikeByID(ctx context.Context, userID int, dislikedUserID int) error
	}
)
