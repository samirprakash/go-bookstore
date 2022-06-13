package main

import (
	"log"

	"github.com/samirprakash/go-bookstore/oauth/app"
	"github.com/samirprakash/go-bookstore/oauth/clients/cassandra"
)

func main() {
	// check connectivity to the cassandra cluster
	// close the session after the check
	session, err := cassandra.GetSession()
	if err != nil {
		log.Fatal(err)
	}
	session.Close()

	// start the application
	app.Start()
}
