package services

import (
	"time"

	"github.com/samirprakash/go-bookstore/users/domain/users"
	"github.com/samirprakash/go-bookstore/users/utils/errors"
)

// CreateUser creates a new user
func CreateUser(user users.User) (*users.User, *errors.REST) {
	user.Created = time.Now().Format("2006-01-02 15:04:05")
	return &user, nil
}
