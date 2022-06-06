package users

import (
	"context"
	"fmt"

	"github.com/samirprakash/go-bookstore/users/datasources/psql/users"
	"github.com/samirprakash/go-bookstore/users/utils/date"
	"github.com/samirprakash/go-bookstore/users/utils/errors"
)

// Get returns a user from the database
func (user *User) Get() *errors.REST {
	getUserQuery := `SELECT id, first_name, last_name, email, created FROM users WHERE id = $1;`
	r := users.DB.QueryRow(context.Background(), getUserQuery, user.ID)
	err := r.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Created)
	if err != nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	return nil
}

// Save saves a new user to the database
func (user *User) Save() *errors.REST {
	insertUserQuery := `INSERT INTO users (first_name, last_name, email, created) VALUES ($1, $2, $3, $4) RETURNING id;`
	user.Created = date.GetCurrentAsString()
	r := users.DB.QueryRow(context.Background(), insertUserQuery, user.FirstName, user.LastName, user.Email, user.Created)
	err := r.Scan(&user.ID)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error inserting user: %s", err.Error()))
	}
	return nil
}
