package users

import (
	"context"
	"fmt"

	"github.com/samirprakash/go-bookstore/users/datasources/psql/users"
	"github.com/samirprakash/go-bookstore/users/utils/errors"
)

// Save saves a new user to the database
func (user *User) Save() *errors.REST {
	insertUserQuery := `INSERT INTO users (first_name, last_name, email, created, status, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`
	r := users.DB.QueryRow(context.Background(), insertUserQuery, user.FirstName, user.LastName, user.Email, user.Created, user.Status, user.Password)
	err := r.Scan(&user.ID)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error inserting user: %s", err.Error()))
	}
	return nil
}

// Get returns a user from the database
func (user *User) Get() *errors.REST {
	getUserQuery := `SELECT id, first_name, last_name, email, created, status, password FROM users WHERE id = $1;`
	r := users.DB.QueryRow(context.Background(), getUserQuery, user.ID)
	err := r.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Created, &user.Status, &user.Password)
	if err != nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.ID))
	}
	return nil
}

// Update updates an existing user in the database
func (user *User) Update() *errors.REST {
	updateUserQuery := `UPDATE users SET first_name = $1, last_name = $2, email = $3 WHERE id = $4 RETURNING id;`
	_, err := users.DB.Exec(context.Background(), updateUserQuery, user.FirstName, user.LastName, user.Email, user.ID)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error updating user: %s", err.Error()))
	}
	return nil
}

// Delete deletes an existing user from the database
func (user *User) Delete() *errors.REST {
	deleteUserQuery := `DELETE FROM users WHERE id = $1;`
	_, err := users.DB.Exec(context.Background(), deleteUserQuery, user.ID)
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error deleting user: %s", err.Error()))
	}
	return nil
}

// FindByStatus returns a list of users by their status
func (user *User) FindByStatus(status string) ([]User, *errors.REST) {
	var uu []User
	getUsersQuery := `SELECT id, first_name, last_name, email, created, status, password FROM users WHERE status = $1;`
	rows, err := users.DB.Query(context.Background(), getUsersQuery, status)
	if err != nil {
		return nil, errors.NewInternalServerError(fmt.Sprintf("error getting users: %s", err.Error()))
	}
	defer rows.Close()

	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.FirstName, &u.LastName, &u.Email, &u.Created, &u.Status, &u.Password)
		if err != nil {
			return nil, errors.NewInternalServerError(fmt.Sprintf("error getting users: %s", err.Error()))
		}
		uu = append(uu, u)
	}
	return uu, nil
}
