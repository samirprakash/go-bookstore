package rest

import (
	"encoding/json"
	"time"

	"github.com/mercadolibre/golang-restclient/rest"
	"github.com/samirprakash/go-bookstore/oauth/domain/users"
	"github.com/samirprakash/go-bookstore/oauth/utils/errors"
)

var (
	rc = rest.RequestBuilder{
		BaseURL: "https://api.bookstore.com",
		Timeout: 100 * time.Millisecond,
	}
)

type UserRepository interface {
	Login(string, string) (*users.User, *errors.REST)
}

type userRepository struct{}

func NewUserRepository() UserRepository {
	return &userRepository{}
}

// Login is a method that returns a user given a username and a password from the users service
func (r *userRepository) Login(email string, password string) (*users.User, *errors.REST) {
	req := users.LoginRequest{
		Email:    email,
		Password: password,
	}

	res := rc.Post("/users/login", req)
	if res == nil || res.Response == nil {
		return nil, errors.NewInternalServerError("invalid response from the users service when trying to login user")
	}

	if res.StatusCode > 299 {
		var rErr errors.REST
		err := json.Unmarshal(res.Bytes(), &rErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user")
		}
		return nil, &rErr
	}

	var user users.User
	if err := json.Unmarshal(res.Bytes(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal response from user service when trying to login user")
	}

	return &user, nil
}
