package users

import (
	"strings"

	"github.com/samirprakash/go-bookstore/users/utils/errors"
)

// User is a struct that represents a user
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Created   string `json:"created"`
}

// Validate validates the user
func (user *User) Validate() *errors.REST {
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequest("invalid email address")
	}
	return nil
}
