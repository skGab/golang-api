package server

import (
	gin "github.com/gin-gonic/gin"
	"github.com/go-api/src/domain/entities"
	"github.com/go-api/src/infrastructure/adapter"
	"github.com/go-api/src/infrastructure/factory"
	"github.com/go-api/src/infrastructure/routes"
)

func Up() {
	// REMOVE GIN LOGS FROM TERMINAL
	gin.SetMode(gin.ReleaseMode)

	// SET UP AN GIN ROUTER
	router := gin.Default()

	// CONNECT WITH DATABASE
	db := adapter.Connect()

	// MIGRATE MODELS
	db.AutoMigrate(&entities.UserEntity{}, &entities.TasksEntity{})

	// CREATING AND INJECTING INSTANCES
	userController, tasksController := factory.InjectInstances(db)

	// REGISTER THE ROUTES AND HANDLES
	routes.Register(router, userController, tasksController)

	// RUN THE SERVER
	router.Run()
}
