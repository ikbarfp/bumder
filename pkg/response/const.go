package response

import "errors"

var (
	ErrDataNotFound       = errors.New("data not found")
	ErrUserNotFound       = errors.New("user not found")
	ErrUnauthorized       = errors.New("unauthorized request")
	ErrInvalidRequest     = errors.New("invalid request")
	ErrBadRequest         = errors.New("something bad happen ")
	ErrUnknown            = errors.New("unknown error")
	ErrInvalidPin         = errors.New("invalid pin")
	ErrInsufficientCredit = errors.New("insufficient credit")
)
