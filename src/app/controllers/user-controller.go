package controller

import (
	"net/http"

	entitie "github.com/go-api/src/domain/entities"
)

type UsersHandler struct {
	UserDomain *entitie.UsersHandler
}

// ServeHTTP implements http.Handler.
func (*UsersHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("This is my home page"))
}
