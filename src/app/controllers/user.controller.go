package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-api/src/app/usecases"
	"github.com/go-api/src/domain/entities"
	"github.com/go-api/src/domain/repository"
)

type UserController struct {
	UserRepository repository.UserRepository
}

// GET ALL USERS
func (uc *UserController) FindAll(gin *gin.Context) {

	// CALL THE USECASE METHOD
	usersDTOs, err := usecases.FindAllUsers(uc.UserRepository)

	// RETURN ERROR IF ANY
	if err != nil {
		gin.JSON(http.StatusInternalServerError, "Erro interno no servidor")
		return
	}

	// RETURN THE USERS WITH SUCCESS STATUS CODE
	gin.JSON(http.StatusOK, usersDTOs)
}

// CREATE USER
func (uc *UserController) Create(gin *gin.Context) {
	var user *entities.UserEntity

	// GET THE USER ON THE REQUEST
	if err := gin.ShouldBindJSON(&user); err != nil {
		gin.JSON(http.StatusBadRequest, err)
		return
	}

	// CALL THE USECASE METHOD
	newUserDTO, err := usecases.CreateUser(uc.UserRepository, user)

	// CHECK IF HAS A VALID USER DTO
	if err != nil {
		gin.JSON(http.StatusInternalServerError, err)
		return
	}

	// RETURN STATUS CODE 200 AND MESSAGE
	gin.JSON(http.StatusOK, newUserDTO)
}
