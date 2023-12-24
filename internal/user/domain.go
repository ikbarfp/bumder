package user

// User . . .
type User struct {
	ID                 string `json:"-"`
	Name               string `json:"name"`
	IsPremium          bool   `json:"is_premium"`
	DailyActionCounter int8   `json:"-"`
}

// ActionUser . . .
type ActionUser map[string][]string
