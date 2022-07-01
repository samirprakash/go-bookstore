package main

import (
	"github.com/samirprakash/go-bookstore/oauth/app"
	"github.com/samirprakash/go-bookstore/oauth/clients/cassandra"
)

func main() {
	// initialize the cassandra cluster
	cassandra.New()

	// start the application
	app.Start()
}
