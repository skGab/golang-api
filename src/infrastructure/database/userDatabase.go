package database

import (
	entitie "github.com/go-api/src/domain/entities"
	"github.com/go-api/src/domain/repository"
)

type UserDatabase struct{}

// CREATE USER
func (*UserDatabase) Create(*entitie.UserEntity) *repository.StatusResponse {
	// CHECK IF HAS DATA

	// MAP THE DATA TO ORM MODEL

	// CHECK IF USER EXISTS

	// CREATE USERS
	status := &repository.StatusResponse{Status: "Usuario criado"}

	return status
}

// GET USERS
func (*UserDatabase) FindAll() (*[]entitie.UserEntity, error) {
	// FIND ALL USERS
	var users *[]entitie.UserEntity

	return users, nil
}
