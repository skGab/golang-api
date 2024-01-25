package server

import (
	gin "github.com/gin-gonic/gin"
	controller "github.com/go-api/src/app/controllers"
	entitie "github.com/go-api/src/domain/entities"
	"github.com/go-api/src/domain/repository"
	"github.com/go-api/src/domain/valueObjects"
	"github.com/go-api/src/infrastructure/adapter"
	"github.com/go-api/src/infrastructure/database"
)

func Up() {
	// REMOVE GIN LOGS FROM TERMINAL
	gin.SetMode(gin.ReleaseMode)

	// SET UP AN GIN ROUTER
	router := gin.Default()

	// CONNECT WITH DATABASE
	db := adapter.Connect()

	// MIGRATE MODELS
	db.AutoMigrate(&entitie.UserEntity{}, &valueObjects.TasksVo{})

	// INSTANCE OF USER HANDLER and Repository
	var _ repository.UserRepository = &database.UserDatabase{}
	userRepository := &database.UserDatabase{}

	userController := controller.UserController{
		UserRepository: userRepository,
	}

	// REGISTER THE ROUTES AND HANDLES
	router.GET("/users", userController.FindAll)
	router.POST("/users/create", userController.Create)

	// RUN THE SERVER
	router.Run()
}
