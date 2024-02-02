package dto

import (
	"github.com/go-api/src/domain/entities"
	"github.com/google/uuid"
)

type NewUserDTO struct {
	ID       string
	User     string
	Email    string
	Password string
}

// CONSTRUCTOR
func NewUser(userData *entities.UserEntity) *NewUserDTO {
	// GENERATE RANDOM STRING
	var newUserID string

	// POPULATE NEW USER ID WITH RANDOM ID IF NOT FOUND ON THE ARGUMENTS
	if userData.ID == "" {
		newUserID = uuid.New().String()
	} else {
		newUserID = userData.ID
	}

	// RETURN THE USER DTO
	return &NewUserDTO{
		ID:       newUserID,
		User:     userData.User,
		Email:    userData.Email,
		Password: userData.Password,
	}

}
