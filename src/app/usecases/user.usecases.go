package usecases

import (
	"github.com/go-api/src/app/dto"
	"github.com/go-api/src/domain/entities"
	"github.com/go-api/src/domain/repository"
)

func FindAllUsers(userRepository repository.UserRepository) (*[]dto.UsersDTO, error) {
	// GET USERS FROM DB
	response := userRepository.FindAll()

	// RETURN ERROR RESPONSE IF ERROR
	if !response.Status {
		return nil, error(response.Error)
	}

	usersEntities := response.Data.(*[]entities.UserEntity)

	// MAP ENTITY TO DTO
	return dto.NewUsersDtos(usersEntities), nil
}

func CreateUser(userRepository repository.UserRepository, user *entities.UserEntity) (*dto.NewUserDTO, error) {
	// CREATE THE USER
	response := userRepository.Create(user)

	// RETURN ERROR IF STATUS IS FALSE
	if !response.Status {
		return nil, error(response.Error)
	}

	userData := response.Data.(*entities.UserEntity)

	// MAP ENTITY TO DTO
	return dto.NewUser(userData), nil
}
