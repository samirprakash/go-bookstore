package services

import (
	"github.com/samirprakash/go-bookstore/users/domain/users"
	"github.com/samirprakash/go-bookstore/users/utils/crypto"
	"github.com/samirprakash/go-bookstore/users/utils/date"
	"github.com/samirprakash/go-bookstore/users/utils/errors"
)

// UsersService is the service that handles all user related operations
var UsersService usersServiceInterface = &usersService{}

type usersService struct{}

// usersServiceInterface is the interface that wraps the UsersService struct
type usersServiceInterface interface {
	CreateUser(users.User) (*users.User, *errors.REST)
	GetUser(int64) (*users.User, *errors.REST)
	UpdateUser(bool, users.User) (*users.User, *errors.REST)
	DeleteUser(int64) *errors.REST
	Search(string) (users.Users, *errors.REST)
	LoginUser(users.LoginRequest) (*users.User, *errors.REST)
}

// CreateUser creates a new user
func (s *usersService) CreateUser(user users.User) (*users.User, *errors.REST) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	user.Created = date.GetCurrentAsString()
	user.Status = users.StatusActive
	user.Password = crypto.GetMD5Hash(user.Password)

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUser returns a user by its id
func (s *usersService) GetUser(userID int64) (*users.User, *errors.REST) {
	user := users.User{
		ID: userID,
	}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates a user
func (s *usersService) UpdateUser(isPatch bool, user users.User) (*users.User, *errors.REST) {
	current, err := s.GetUser(user.ID)
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

// DeleteUser deletes a user
func (s *usersService) DeleteUser(userID int64) *errors.REST {
	user := users.User{
		ID: userID,
	}

	if err := user.Delete(); err != nil {
		return err
	}

	return nil
}

// Search finds a user by its status
func (s *usersService) Search(status string) (users.Users, *errors.REST) {
	user := users.User{}
	users, err := user.FindByStatus(status)
	if err != nil {
		return nil, err
	}

	return users, nil
}

// LoginUser authenticates a user
func (s *usersService) LoginUser(req users.LoginRequest) (*users.User, *errors.REST) {
	user := &users.User{
		Email:    req.Email,
		Password: crypto.GetMD5Hash(req.Password),
	}

	if err := user.FindByEmailAndPassword(); err != nil {
		return nil, err
	}

	return user, nil
}
