package app

import "github.com/samirprakash/go-bookstore/users/controllers/ping"

func mapRoutes() {
	router.GET("/ping", ping.Ping)
}
