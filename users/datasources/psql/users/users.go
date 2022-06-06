package users

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://root:secret@localhost:5432/users?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err := DB.Ping(); err != nil {
		panic(err)
	}
	log.Println("Successfully connected to go-bookstore-user database")
}
