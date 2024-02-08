package usecases

import (
	"fmt"

	"github.com/go-api/src/app/dto"
	"github.com/go-api/src/domain/entities"
	"github.com/go-api/src/domain/repository"
)

// CREATE USER
func CreateUser(userRepository repository.UserRepository, user *entities.UserEntity) (*dto.NewUserDTO, error) {
	// CREATE THE USER
	err, newUser := userRepository.Create(user)

	// RETURN ERROR IF STATUS IS FALSE
	if err != nil {
		return nil, error(response.Error)
	}

	userData := response.Data.(entities.UserEntity)

	// MAP ENTITY TO DTO
	return dto.NewUser(userData), nil
}

// GET ALL USERS
func FindAllUsers(userRepository repository.UserRepository) ([]dto.UsersDTO, error) {
	// GET USERS FROM DB
	err, userRepository := userRepository.FindAll()

	// RETURN ERROR RESPONSE IF ERROR
	if err != nil {
		fmt.Errorf(err.Error())
		return nil, error(err)
	}

	usersEntities := response.Data.(*[]entities.UserEntity)

	// MAP ENTITY TO DTO
	return dto.NewUsersDtos(usersEntities), nil
}
