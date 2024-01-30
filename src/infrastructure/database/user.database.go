package database

import (
	"fmt"

	entitie "github.com/go-api/src/domain/entities"
	"github.com/go-api/src/domain/repository"
	"gorm.io/gorm"
)

type UserDatabase struct {
	Adapter *gorm.DB
}

// CREATE USER
func (this *UserDatabase) Create(user *entitie.UserEntity) *repository.Status {
	// CHECK IF USER EXISTS
	userStatus := this.Adapter.Find(&user)

	if userStatus.RowsAffected != 0 {
		fmt.Print(&user)
		return &repository.Status{Status: true, Data: &user, Message: "Usuario já cadastrado"}
	}

	if userStatus.Error != nil {
		return &repository.Status{Status: false, Message: "Erro ao buscar usuario", Error: userStatus.Error}
	}

	// CREATE USERS
	response := this.Adapter.Create(&user)

	if response.Error != nil {
		return &repository.Status{Status: false, Message: "Erro ao criar usuarios", Error: response.Error}
	}

	return &repository.Status{Status: true, Data: &user, Message: "Novo usuario registrado"}
}

// GET USERS
func (this *UserDatabase) FindAll() *repository.Status {
	// FIND ALL USERS
	var users []entitie.UserEntity

	// Find all users. The result will be stored in the 'users' slice.
	result := this.Adapter.Find(&users)

	// Check for errors in finding users
	if result.Error != nil {
		return &repository.Status{Status: false, Message: "Erro ao buscar usuarios: ", Error: result.Error}
	}

	if result.RowsAffected == 0 {
		return &repository.Status{Status: false, Message: "Nenhum usuario encontrado"}
	}

	return &repository.Status{Status: true, Data: &users, Message: "Usuarios encontrados"}
}
