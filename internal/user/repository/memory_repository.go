package repository

import (
	"context"
	"github.com/go-faker/faker/v4"
	"github.com/google/uuid"
	"github.com/ikbarfp/bumder/internal/user"
	"github.com/ikbarfp/bumder/pkg/response"
	"time"
)

var (
	numOfRandomUser = 15
	userData        []*user.User
)

type userRepository struct {
}

// New . . .
func New() user.Repository {
	populateUser(numOfRandomUser)

	return &userRepository{}
}

func populateUser(qty int) {
	for i := 1; i <= qty; i++ {
		loc, _ := time.LoadLocation("Asia/Jakarta")
		userData = append(userData, &user.User{
			ID:          uuid.NewString(),
			Name:        faker.Name(),
			IsPremium:   false,
			DateOfBirth: time.Date(1990+i, time.January, 1+i, 0, 0, 0, 0, loc),
		})
	}
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

// CreateUser . . .
func (r userRepository) CreateUser(ctx context.Context, user *user.User) (string, error) {
	user.ID = uuid.NewString()
	userData = append(userData, user)

	return user.ID, nil
}
