package user

// RequestRegister . . .
type RequestRegister struct {
	FullName        string `json:"full_name" validate:"required"`
	MobilePhone     string `json:"mobile_phone" validate:"required,e164"`
	Pin             string `json:"pin" validate:"required,len=6"`
	ConfirmationPin string `json:"confirmation_pin" validate:"required,len=6"`
}
