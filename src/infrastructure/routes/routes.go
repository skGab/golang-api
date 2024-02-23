package routes

import (
	gin "github.com/gin-gonic/gin"
	controller "github.com/go-api/src/app/controllers"
)

func Register(router *gin.Engine, userController *controller.UserController, tasksController *controller.TasksController) {
	// USERS ROUTE
	router.GET("/users", userController.FindAll)
	router.POST("/users/create", userController.Create)

	// TASKS ROUTE
	router.GET("/tasks/:id", tasksController.FindAllTasks)
	router.POST("/tasks/create/:id", tasksController.CreateTasks)
	router.DELETE("/tasks/delete/:id", tasksController.DeleteTask)
	router.DELETE("/tasks/deleteAllTodo/:id", tasksController.DeleteAllTodoTasks)
	router.DELETE("/tasks/deleteAllDone/:id", tasksController.DeleteAllDoneTasks)
	router.PUT("/tasks/state/:id", tasksController.UpdateStateTask)
	router.PUT("/tasks/update/:id", tasksController.UpdateTask)
}
