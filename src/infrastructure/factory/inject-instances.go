package factory

import (
	controller "github.com/go-api/src/app/controllers"
	"github.com/go-api/src/infrastructure/database"
	"gorm.io/gorm"
)

func InjectInstances(db *gorm.DB) (*controller.UserController, *controller.TasksController) {
	// USER
	// CREATE DATABASE INSTANCE
	userRepository := &database.UserDatabase{Adapter: db}

	// INJECT INSTANCE ON USERCONTROLLER STRUCT
	userController := &controller.UserController{
		UserRepository: userRepository,
	}

	// TASKS
	// CREATE DATABASE INSTANCE
	tasksRepository := &database.TasksDatabase{Adapter: db}

	// INJECT INSTANCE ON TASKSCONTROLER STRUCT
	tasksController := &controller.TasksController{
		TasksRepository: tasksRepository,
	}

	return userController, tasksController
}
