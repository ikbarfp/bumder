package user

// RequestRegister . . .
type RequestRegister struct {
	FullName        string `json:"full_name"`
	MobilePhone     string `json:"mobile_phone"`
	Pin             string `json:"pin"`
	ConfirmationPin string `json:"confirmation_pin"`
}
