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
	// CHECK IF USER EXISTS
	userStatus := db.Adapter.Raw("SELECT id FROM user_entities WHERE email = ?", userEntity.Email).Scan(&userEntity)

	if userStatus.Error != nil {
		return nil, userStatus.Error
	}

	if userStatus.RowsAffected != 0 {
		return nil, errors.New("usuário já cadastrado")
	}

	// CREATE USERS
	response := db.Adapter.Create(&userEntity)

	if response.Error != nil {
		return nil, response.Error
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
		return nil, errors.New("nenhum usuario encontrado")
	}

	return users, nil
}
