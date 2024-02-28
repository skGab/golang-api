package dto

import (
	"github.com/go-api/src/domain/entities"
)

type UsersDTO struct {
	ID       string
	User     string
	Password string
	Tasks    []TasksDTO
}

// CONSTRUCTOR
func NewUsersDtos(usersEntities []entities.UserEntity) []UsersDTO {
	usersDTOs := make([]UsersDTO, 0, len(usersEntities))

	for _, userEntity := range usersEntities {
		userDTO := UsersDTO{
			ID:       userEntity.ID,
			User:     userEntity.User,
			Password: userEntity.Password,
			Tasks:    NewTasksDTO(userEntity),
		}

		usersDTOs = append(usersDTOs, userDTO)
	}

	return usersDTOs
}
