package usecases

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-api/src/app/dto"
	"github.com/go-api/src/domain/entities"
	"github.com/go-api/src/domain/repository"
)

func CreateUser(userRepository repository.UserRepository, gin *gin.Context) *dto.NewUserDTO {
	var user *entities.UserEntity

	// GET THE USER ON THE REQUEST
	if err := gin.ShouldBindJSON(&user); err != nil {
		gin.JSON(http.StatusBadRequest, err)
		return nil
	}

	// CREATE THE USER
	response := userRepository.Create(user)

	// RETURN ERROR IF STATUS IS FALSE
	if !response.Status {
		gin.JSON(http.StatusInternalServerError, response.Message)
		return nil
	}

	userData := response.Data.(*entities.UserEntity)

	// MAP ENTITY TO DTO
	return &dto.NewUserDTO{ID: userData.ID, User: userData.User, Email: userData.Email, Password: userData.Password}
}

func FindAllUsers(userRepository repository.UserRepository, gin *gin.Context) *[]dto.UsersDTO {
	// GET USERS FROM DB
	response := userRepository.FindAll()

	// RETURN ERROR RESPONSE IF ERROR
	if !response.Status {
		gin.JSON(http.StatusInternalServerError, response.Error)
		return nil
	}

	usersEntities := response.Data.(*[]entities.UserEntity)

	usersDTOs := make([]dto.UsersDTO, 0, len(*usersEntities))

	for _, userEntity := range *usersEntities {
		userDTO := dto.UsersDTO{
			ID:       userEntity.ID,
			User:     userEntity.User,
			Email:    userEntity.Email,
			Password: userEntity.Password,
		}

		usersDTOs = append(usersDTOs, userDTO)
	}

	// MAP ENTITY TO DTO
	return &usersDTOs
}
