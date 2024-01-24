package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-api/src/domain/repository"
	"github.com/go-api/src/domain/valueObjects"
)

type UserController struct {
	UserRepository repository.UserRepository
}

// GET ALL USERS
func (uc *UserController) FindAll(gin *gin.Context) {
	// GET USERS FROM DB
	users, err := uc.UserRepository.FindAll()

	// RETURN ERROR RESPONSE IF ERROR
	if err != nil {
		gin.JSON(http.StatusInternalServerError, err)
		return
	}

	// RETURN THE USERS WITH SUCCESS STATUS CODE
	gin.JSON(http.StatusOK, users)
}

// CREATE USER
func (uc *UserController) Create(gin *gin.Context) {
	var user valueObjects.CreateUserVO

	if err := gin.ShouldBindJSON(&user); err != nil {
		gin.JSON(http.StatusBadRequest, err)
		return
	}

	log.Print(user)
}
