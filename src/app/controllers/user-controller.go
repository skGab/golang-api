package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	entitie "github.com/go-api/src/domain/entities"
	"github.com/go-api/src/domain/repository"
)

type UserController struct {
	UserRepository repository.UserRepository
}

// GET ALL USERS
func (uc *UserController) FindAll(gin *gin.Context) {
	// GET USERS FROM DB
	users, err := uc.UserRepository.FindAll()

	// RETURN ERROR RESPONSE IF ERROR
	if !err.Status {
		gin.JSON(http.StatusInternalServerError, err)
		return
	}

	// RETURN THE USERS WITH SUCCESS STATUS CODE
	gin.JSON(http.StatusOK, users)
}

// CREATE USER
func (uc *UserController) Create(gin *gin.Context) {
	var user *entitie.UserEntity

	// GET THE USER ON THE REQUEST
	if err := gin.ShouldBindJSON(&user); err != nil {
		gin.JSON(http.StatusBadRequest, err)
		return
	}

	// CREATE THE USER
	response := uc.UserRepository.Create(user)

	if !response.Status {
		gin.JSON(http.StatusInternalServerError, response.Message)
		return
	}

	gin.JSON(http.StatusOK, response.Message)
}
