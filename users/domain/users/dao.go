package users

import (
	"fmt"

	"github.com/samirprakash/go-bookstore/users/utils/errors"
)

var db = make(map[int64]*User)

// Get returns a user from the database
func (user *User) Get() *errors.REST {
	res := db[user.ID]
	if res == nil {
		return errors.NewNotFound(fmt.Sprintf("user %d not found", user.ID))
	}

	user.ID = res.ID
	user.FirstName = res.FirstName
	user.LastName = res.LastName
	user.Email = res.Email
	user.Created = res.Created

	return nil
}

// Save saves a new user to the database
func (user *User) Save() *errors.REST {
	cu := db[user.ID]
	if cu != nil {
		if cu.Email == user.Email {
			return errors.NewBadRequest(fmt.Sprintf("user email %s is already registered", user.Email))
		}
		return errors.NewBadRequest(fmt.Sprintf("user %d already exists", user.ID))
	}
	db[user.ID] = user

	return nil
}
