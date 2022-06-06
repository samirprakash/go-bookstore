package users

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

var (
	DB *pgx.Conn
)

func init() {
	var err error
	DB, err = pgx.Connect(context.Background(), "postgres://root:secret@localhost:5432/bookstore_users?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to go-bookstore-user database")
}
