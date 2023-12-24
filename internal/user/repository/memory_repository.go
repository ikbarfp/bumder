package repository

import (
	"context"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/ikbarfp/bumder/internal/user"
	"github.com/ikbarfp/bumder/pkg/response"
)

var (
	numOfRandomUser = 30

	// From here its just in-memory datastore for simulation
	userData       []*user.User
	likedUserData  []user.ActionUser
	passedUserData []user.ActionUser
)

type userRepository struct {
}

// New . . .
func New() user.Repository {
	populateUser(numOfRandomUser)

	return &userRepository{}
}

func populateUser(qty int) {
	for i := 0; i < qty; i++ {
		usr := &user.User{
			ID:        uuid.NewString(),
			Name:      faker.Name(),
			IsPremium: false,
		}
		userData = append(userData, usr)

		populateInitialLikes(usr.ID)
		populateInitialPasses(usr.ID)
	}
}

func populateInitialLikes(userID string) {
	initialLikes := make(user.ActionUser)
	initialLikes[userID] = []string{}
	likedUserData = append(likedUserData, initialLikes)
}

func populateInitialPasses(userID string) {
	initialPasses := make(user.ActionUser)
	initialPasses[userID] = []string{}
	passedUserData = append(passedUserData, initialPasses)
}

// FindByID . . .
func (r userRepository) FindByID(ctx context.Context, id string) (*user.User, error) {

	for _, usr := range userData {
		if id == usr.ID {
			return usr, nil
		}
	}

	return nil, response.ErrUserNotFound
}

// FindLikedUserByID . . .
func (r userRepository) FindLikedUserByID(ctx context.Context, userID string) ([]*user.User, error) {

	var (
		res          []*user.User
		likedUserIDs []string
	)

	for _, likedUser := range likedUserData {
		values, exists := likedUser[userID]
		if exists {
			likedUserIDs = values
			break
		}
	}

	if len(likedUserIDs) == 0 {
		return nil, response.ErrDataNotFound
	}

	for _, likedUserID := range likedUserIDs {
		for _, usrData := range userData {
			if usrData.ID == likedUserID {
				res = append(res, usrData)
			}
		}
	}

	return res, nil
}

// FindPassedUserByID . . .
func (r userRepository) FindPassedUserByID(ctx context.Context, userID string) ([]*user.User, error) {

	var (
		res           []*user.User
		passedUserIDs []string
	)

	for _, passedUser := range passedUserData {
		values, exists := passedUser[userID]
		if exists {
			passedUserIDs = values
			break
		}
	}

	if len(passedUserIDs) == 0 {
		return nil, response.ErrDataNotFound
	}

	for _, passedUserID := range passedUserIDs {
		for _, usrData := range userData {
			if usrData.ID == passedUserID {
				res = append(res, usrData)
			}
		}
	}

	return res, nil
}

// FindActionCounterByID . . .
func (r userRepository) FindActionCounterByID(ctx context.Context, userID string) (int8, error) {

	for _, usr := range userData {
		if usr.ID == userID {
			return usr.DailyActionCounter, nil
		}
	}

	return 0, response.ErrUserNotFound
}

// FindAllUser . . .
func (r userRepository) FindAllUser(ctx context.Context) ([]*user.User, error) {

	return userData, nil
}

// CreateUser . . .
func (r userRepository) CreateUser(ctx context.Context, user *user.User) (string, error) {
	userID := uuid.NewString()

	user.ID = userID
	userData = append(userData, user)

	populateInitialPasses(userID)
	populateInitialLikes(userID)

	return userID, nil
}
