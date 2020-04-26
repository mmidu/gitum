package models

// User is the strict that holds the user data
type User struct {
	Identifier  string      `json:"identifier"`
	Credentials Credentials `json:"credentials"`
}
