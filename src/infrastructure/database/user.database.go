package database

import (
	"errors"

	entitie "github.com/go-api/src/domain/entities"
	"gorm.io/gorm"
)

type UserDatabase struct {
	Adapter *gorm.DB
}

// CREATE USER
func (db *UserDatabase) Create(user *entitie.UserEntity) (error, *entitie.UserEntity) {
	// CHECK IF USER EXISTS
	userStatus := db.Adapter.Find(&user)

	if userStatus.Error != nil {
		return userStatus.Error, nil
	}

	if userStatus.RowsAffected != 0 {
		return errors.New("Usuario já cadastrado"), nil
	}

	// CREATE USERS
	response := db.Adapter.Create(&user)

	if response.Error != nil {
		return response.Error, nil
	}

	return nil, user
}

// GET USERS
func (db *UserDatabase) FindAll() (error, *[]entitie.UserEntity) {
	// FIND ALL USERS
	var users []entitie.UserEntity

	// Find all users. The result will be stored in the 'users' slice.
	result := db.Adapter.Find(&users)

	// Check for errors in finding users
	if result.Error != nil {
		return result.Error, nil
	}

	if result.RowsAffected == 0 {
		return errors.New("Nenhum usuario encontrado"), nil
	}

	return nil, &users
}
