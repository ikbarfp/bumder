package user

import "time"

// User . . .
type User struct {
	ID          int       `json:"-"`
	Name        string    `json:"name"`
	MobilePhone string    `json:"mobile_phone"`
	IsPremium   bool      `json:"is_premium"`
	DateOfBirth time.Time `json:"date_of_birth"`
}
