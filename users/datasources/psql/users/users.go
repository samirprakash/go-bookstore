package users

import (
	"context"

	"github.com/jackc/pgx/v4"
	"github.com/samirprakash/go-bookstore/users/logger"
)

var (
	DB *pgx.Conn
)

func init() {
	var err error
	DB, err = pgx.Connect(context.Background(), "postgres://root:secret@localhost:5432/bookstore_users?sslmode=disable")
	if err != nil {
		logger.Error("error when connecting to database: %s", err)
	}
}
