package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-api/src/app/usecases"
	"github.com/go-api/src/domain/repository"
)

type UserController struct {
	UserRepository repository.UserRepository
}

// GET ALL USERS
func (uc *UserController) FindAll(gin *gin.Context) {

	usersDTOs := usecases.FindAllUsers(uc.UserRepository, gin)

	// RETURN THE USERS WITH SUCCESS STATUS CODE
	gin.JSON(http.StatusOK, usersDTOs)
	return
}

// CREATE USER
func (uc *UserController) Create(gin *gin.Context) {

	newUserDTO := usecases.CreateUser(uc.UserRepository, gin)

	if newUserDTO != nil {
		// RETURN STATUS CODE 200 AND MESSAGE
		gin.JSON(http.StatusOK, newUserDTO)
		return
	}

	return
}
