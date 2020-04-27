package models

// Credentials is the struct that holds the user credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
}

// NewCredentials makes a new credentials instance
func NewCredentials(username string, password string, domain string) *Credentials {
	credentials := Credentials{
		Username: username,
		Password: password,
		Domain:   domain,
	}
	return &credentials
}
