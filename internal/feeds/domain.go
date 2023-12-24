package feeds

// UnseenUser . . .
type UnseenUser struct {
	UserID    string `json:"user_id"`
	Name      string `json:"name"`
	IsPremium bool   `json:"is_premium"`
}
