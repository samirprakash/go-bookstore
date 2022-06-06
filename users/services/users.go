package services

import (
	"time"

	"github.com/samirprakash/go-bookstore/users/domain/users"
	"github.com/samirprakash/go-bookstore/users/utils/errors"
)

// CreateUser creates a new user
func CreateUser(user users.User) (*users.User, *errors.REST) {
	user.Created = time.Now().Format("2006-01-02 15:04:05")
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser returns a user by its id
func GetUser(userID int64) (*users.User, *errors.REST) {
	user := users.User{
		ID: userID,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return &user, nil
}
