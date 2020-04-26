package collections

import (
	"../models"
)

// Users is users list struct
type Users struct {
	Users []models.User `json:"users"`
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
func (u Users) Get(identifier string) models.User {
	var user models.User
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
