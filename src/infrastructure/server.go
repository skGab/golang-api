package server

import (
	"net/http"

	controller "github.com/go-api/src/app/controllers"
	entitie "github.com/go-api/src/domain/entities"
)

func Up() {
	// Create a new request multiplexer
	// Take incoming requests and dispatch them to the matching handlers
	mux := http.NewServeMux()

	// Register the routes and handlers
	mux.Handle("/", &controller.UsersHandler{UserDomain: &entitie.UsersHandler{}})

	// Run the server
	http.ListenAndServe(":8080", mux)
}
