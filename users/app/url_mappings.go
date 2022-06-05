package app

import "github.com/samirprakash/go-bookstore/users/controllers"

func mapURLs() {
	router.GET("/ping", controllers.Ping)
}
