package auth

// Auth . . .
type Auth struct {
	UserID      string `json:"-"`
	Token       string `json:"-"`
	MobilePhone string `json:"mobile_phone"`
	Pin         string `json:"pin"`
}
