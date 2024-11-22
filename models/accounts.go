package models

// User account
type Account struct {
	ID        string  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Status    string `json:"status"` // "active" or "deleted"
	Balance   int64  `json:"balance"`
}
