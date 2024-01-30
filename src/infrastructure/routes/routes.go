package routes

import (
	gin "github.com/gin-gonic/gin"
	controller "github.com/go-api/src/app/controllers"
)

func Register(router *gin.Engine, userController *controller.UserController, tasksController *controller.TasksController) {
	router.GET("/users", userController.FindAll)
	router.POST("/users/create", userController.Create)
}
