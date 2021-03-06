package app

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/samirprakash/go-bookstore/users/datasources/psql/users"
)

var (
	router = gin.Default()
)

func Start() {
	// Initialize the database
	if err := users.DB.Ping(context.Background()); err != nil {
		panic(err)
	}

	// Initialize the routes
	mapRoutes()

	// Start the server
	router.Run(":8080")
}
