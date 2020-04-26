package models

// Credentials is the struct that holds the user credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
}
