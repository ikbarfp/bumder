package user

import "time"

// RequestRegister . . .
type RequestRegister struct {
	FullName        string    `json:"full_name"`
	DateOfBirth     time.Time `json:"date_of_birth"`
	MobilePhone     string    `json:"mobile_phone"`
	Pin             string    `json:"pin"`
	ConfirmationPin string    `json:"confirmation_pin"`
}
