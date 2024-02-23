package factory

import (
	controller "github.com/go-api/src/app/controllers"
	"github.com/go-api/src/app/usecases"
	"github.com/go-api/src/infrastructure/database"
	"gorm.io/gorm"
)

func InjectDepedencies(db *gorm.DB) (*controller.UserController, *controller.TasksController) {
	// CREATE DATABASE INSTANCES
	userRepository := &database.UserDatabase{Adapter: db}
	tasksRepository := &database.TasksDatabase{Adapter: db}

	// CREATE USECASE INSTANCES
	tasksUsecases := &usecases.TasksUsecases{TasksRepository: tasksRepository}
	userUsecases := &usecases.UserUsecases{UserRepository: userRepository}

	// INJECT DEPEDENCIES ON USERCONTROLLER STRUCT
	userController := &controller.UserController{
		UserUsecases: *userUsecases,
	}

	// INJECT DEPEDENCIES ON TASKSCONTROLER STRUCT
	tasksController := &controller.TasksController{
		TasksUsecases: tasksUsecases,
	}

	return userController, tasksController
}
