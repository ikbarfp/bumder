package user

import "time"

// User . . .
type User struct {
	ID          string    `json:"-"`
	Name        string    `json:"name"`
	IsPremium   bool      `json:"is_premium"`
	DateOfBirth time.Time `json:"date_of_birth"`
}
