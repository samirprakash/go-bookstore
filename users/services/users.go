package services

import (
	"github.com/samirprakash/go-bookstore/users/domain/users"
	"github.com/samirprakash/go-bookstore/users/utils/errors"
)

// CreateUser creates a new user
func CreateUser(user users.User) (*users.User, *errors.REST) {
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

func UpdateUser(isPatch bool, user users.User) (*users.User, *errors.REST) {
	current, err := GetUser(user.ID)
	if err != nil {
		return nil, err
	}

	// check incoming user if http.MethodPatch == true
	// set one or more fields from the current user
	if isPatch {
		if user.FirstName != "" {
			current.FirstName = user.FirstName
		}
		if user.LastName != "" {
			current.LastName = user.LastName
		}
		if user.Email != "" {
			current.Email = user.Email
		}
	} else {
		// set the incoming user to the current user
		current.FirstName = user.FirstName
		current.LastName = user.LastName
		current.Email = user.Email
	}

	if err := current.Update(); err != nil {
		return nil, err
	}

	return current, nil
}
