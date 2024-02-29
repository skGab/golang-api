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
func (db *UserDatabase) Create(userEntity *entitie.UserEntity) (*entitie.UserEntity, error) {
	// GET ALL USERS
	var existingUsers []entitie.UserEntity
	response := db.Adapter.Find(&existingUsers)

	if response.Error != nil {
		return nil, response.Error
	}

	// LOOP TROUGHT USERS TO CHECK IF ALREADY EXISTS
	for _, user := range existingUsers {
		if user.User == userEntity.User {
			return nil, errors.New("usuario já cadastrado")
		}
	}

	// CREATE USERS
	userStatus := db.Adapter.Create(&userEntity)

	if userStatus.Error != nil {
		return nil, userStatus.Error
	}

	return userEntity, nil
}

// GET USERS
func (db *UserDatabase) FindAll() ([]entitie.UserEntity, error) {
	// FIND ALL USERS
	var users []entitie.UserEntity

	// Find all users. The result will be stored in the 'users' slice.
	result := db.Adapter.Find(&users)

	// Check for errors in finding users
	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		println("nenhum usuario encontrado")
		return make([]entitie.UserEntity, 0), nil
	}

	return users, nil
}
