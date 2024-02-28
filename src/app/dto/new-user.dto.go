package dto

import (
	"github.com/go-api/src/domain/entities"
)

type NewUserDTO struct {
	ID       string
	User     string
	Password string
}

// CONSTRUCTOR
func NewUser(userEntity *entities.UserEntity) *NewUserDTO {
	// RETURN THE USER DTO
	return &NewUserDTO{
		ID:       userEntity.ID,
		User:     userEntity.User,
		Password: userEntity.Password,
	}
}
