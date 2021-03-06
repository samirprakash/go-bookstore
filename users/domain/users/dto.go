package users

import (
	"strings"

	"github.com/samirprakash/go-bookstore/users/utils/errors"
)

const (
	StatusActive = "active"
)

// User is a struct that represents a user
type User struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Created   string `json:"created"`
	Status    string `json:"status"`
	Password  string `json:"password"`
}

type Users []User

// Validate validates the user
func (user *User) Validate() *errors.REST {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("invalid email address")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewBadRequestError("invalid password")
	}
	return nil
}
