package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"` // Remember to never store plain-text passwords
	// ... any other fields you might need
}
