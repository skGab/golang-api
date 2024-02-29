package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-api/src/app/usecases"
	"github.com/go-api/src/domain/entities"
)

type UserController struct {
	UserUsecases usecases.UserUsecases
}

// GET ALL USERS
func (uc *UserController) FindAll(gin *gin.Context) {

	// CALL THE USECASE METHOD
	usersDTOs, err := uc.UserUsecases.FindAllUsers()

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
	var rawUser *entities.UserEntity

	// GET THE USER ON THE REQUEST
	if err := gin.ShouldBindJSON(&rawUser); err != nil {
		gin.JSON(http.StatusBadRequest, err)
		return
	}

	// CALL THE USECASE METHOD
	response, err := uc.UserUsecases.CreateUser(rawUser)

	// CHECK IF HAS A VALID USER DTO
	if err != nil {
		gin.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// RETURN STATUS CODE 200 AND MESSAGE
	gin.JSON(http.StatusOK, response)
}
