package database

import (
	"log"

	entitie "github.com/go-api/src/domain/entities"
	"github.com/go-api/src/domain/repository"
	"gorm.io/gorm"
)

type UserDatabase struct {
	Adapter *gorm.DB
}

// CREATE USER
func (db *UserDatabase) Create(user *entitie.UserEntity) *repository.Status {
	// CHECK IF HAS DATA

	// CHECK IF USER EXISTS

	// CREATE USERS
	log.Print("Usuario", user)
	db.Adapter.Create(user)

	response := &repository.Status{Status: true, Message: "Usuario criado"}

	return response
}

// GET USERS
func (*UserDatabase) FindAll() (*[]entitie.UserEntity, *repository.Status) {
	// FIND ALL USERS
	var users *[]entitie.UserEntity

	return users, nil
}
