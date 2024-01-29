package database

import (
	"fmt"
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
	// CHECK IF USER EXISTS
	userStatus := db.Adapter.Find(&user)

	if userStatus.RowsAffected != 0 {
		fmt.Print(&user)
		return &repository.Status{Status: true, Message: "Usuario já cadastrado"}
	}

	if userStatus.Error != nil {
		return &repository.Status{Status: false, Message: "Erro ao buscar usuario" + userStatus.Error.Error()}
	}

	// CREATE USERS
	log.Print("Usuario", user)

	response := db.Adapter.Create(&user)
	if response.Error != nil {
		return &repository.Status{Status: false, Message: "Erro ao criar usuarios" + response.Error.Error()}
	}

	return &repository.Status{Status: true, Message: "Usuario cadastrado"}
}

// GET USERS
func (db *UserDatabase) FindAll() (*[]entitie.UserEntity, *repository.Status) {
	// FIND ALL USERS
	var users []entitie.UserEntity

	// Find all users. The result will be stored in the 'users' slice.
	result := db.Adapter.Find(&users)

	// Check for errors in finding users
	if result.Error != nil {
		return nil, &repository.Status{Status: false, Message: "Erro ao buscar usuarios: " + result.Error.Error()}
	}

	if result.RowsAffected == 0 {
		return nil, &repository.Status{Status: false, Message: "Nenhum usuario encontrado"}
	}

	return &users, &repository.Status{Status: true, Message: "Usuarios encontrados"}
}
