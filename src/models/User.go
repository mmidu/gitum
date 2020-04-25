package models

// Users is users list struct
type Users struct {
	Users []User `json:"users"`
}

// Contains checks whether the Users list contains the passed identifier
func (u Users) Contains(identifier string) bool {
	for _, user := range u.Users {
		if user.Identifier == identifier {
			return true
		}
	}
	return false
}

// Get retrieves a user by its identifier
func (u Users) Get(identifier string) User {
	var user User
	for _, user := range u.Users {
		if user.Identifier == identifier {
			return user
		}
	}
	return user
}

// List lists the available identifiers
func (u Users) List() []string {
	var list []string
	for _, user := range u.Users {
		list = append(list, user.Identifier)
	}
	return list
}

// User is the strict that holds the user data
type User struct {
	Identifier  string      `json:"identifier"`
	Credentials Credentials `json:"credentials"`
}

// Credentials is the struct that holds the user credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Domain   string `json:"domain"`
}
