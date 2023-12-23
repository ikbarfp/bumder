package auth

// RequestLogin . . .
type RequestLogin struct {
	MobilePhone string `json:"mobile_phone"`
	Pin         string `json:"pin"`
}
