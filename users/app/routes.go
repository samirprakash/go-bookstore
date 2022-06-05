package app

import (
	"github.com/samirprakash/go-bookstore/users/controllers/ping"
	"github.com/samirprakash/go-bookstore/users/controllers/users"
)

func mapRoutes() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
}
