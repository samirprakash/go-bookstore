package app

import (
	"github.com/samirprakash/go-bookstore/users/controllers/ping"
	"github.com/samirprakash/go-bookstore/users/controllers/users"
)

func mapRoutes() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:id", users.Get)
	router.PUT("/users/:id", users.Update)
	router.PATCH("/users/:id", users.Update)
	router.DELETE("/users/:id", users.Delete)
	router.GET("/internal/users/search", users.Search)
}
