package models

// User is the strict that holds the user data
type User struct {
	Identifier  string      `json:"identifier"`
	Credentials Credentials `json:"credentials"`
}

// NewUser makes a new user
func NewUser(identifier string, credentials Credentials) *User {
	user := User{
		Identifier:  identifier,
		Credentials: credentials,
	}
	return &user
}
