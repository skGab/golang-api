package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-api/src/app/usecases"
)

type TasksController struct {
	TasksUsecases *usecases.TasksUsecases
}

// GET THE TASKS BY USER ID
func (tc *TasksController) FindAllTasks(gin *gin.Context) {
	// GET THE USER ID
	userID := gin.Param("id")

	// FETCH TASKS
	response, err := tc.TasksUsecases.FindAll(userID)

	if err != nil {
		gin.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	gin.JSON(http.StatusOK, response)
}

// CREATE USER TASKS
func (tc *TasksController) CreateTasks(gin *gin.Context) {
	// GET THE USER ID
	userID := gin.Param("id")

	if userID == "" {
		gin.JSON(http.StatusBadRequest, "id não encontrado no corpo da requisição")
		return
	}

	// GET THE TEXT FROM REQUEST BODY
	var task struct {
		Text string `json:"text"`
	}

	if err := gin.BindJSON(&task); err != nil {
		gin.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// CREATE THE TASK
	taskDTO, err := tc.TasksUsecases.Create(task.Text, userID)

	if err != nil {
		gin.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	gin.JSON(http.StatusOK, taskDTO)
}

// DELETE USER TASK
func (tc *TasksController) DeleteTask(gin *gin.Context) {
	// GET THE USER ID
	userID := gin.Param("id")

	if userID == "" {
		gin.JSON(http.StatusBadRequest, "ID não encontrado no corpo da requisição")
		return
	}

	// DELETE THE TASK
	response, err := tc.TasksUsecases.Delete(string(userID))

	if err != nil {
		gin.JSON(http.StatusInternalServerError, "Erro interno no servidor")
		return
	}

	// RETURN RESPONSE
	gin.JSON(http.StatusOK, response)
}

// DELETE ALL TODO TASKS
func (tc *TasksController) DeleteAllTodoTasks(gin *gin.Context) {
	// GET THE USER ID
	userID := gin.Param("id")

	// DELETE THE TASK
	response, err := tc.TasksUsecases.DeleteAllTodo(userID)

	if err != nil {
		gin.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// RETURN RESPONSE
	gin.JSON(http.StatusOK, response)
}

// DELETE ALL DONE TASKS
func (tc *TasksController) DeleteAllDoneTasks(gin *gin.Context) {
	// GET THE USER ID
	userID := gin.Param("id")

	// DELETE THE TASK
	response, err := tc.TasksUsecases.DeleteAllDone(userID)

	if err != nil {
		gin.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// RETURN RESPONSE
	gin.JSON(http.StatusOK, response)
}

// UPDATE THE TASK STATE (COMPLETE/UNCOMPLETE)
func (tc *TasksController) UpdateStateTask(gin *gin.Context) {
	// GET THE USER ID
	ID := gin.Param("id")

	// GET TASKS STATE
	var state struct {
		Complete bool `json:"complete"`
	}

	if err := gin.ShouldBindJSON(&state); err != nil {
		gin.JSON(http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(state, state.Complete)

	// UPDATE THE TASK STATE
	response, err := tc.TasksUsecases.UpdateState(state.Complete, ID)

	if err != nil {
		gin.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// RETURN RESPONSE
	gin.JSON(http.StatusOK, response)
}

// UPDATE THE TASK'S TEXT
func (tc *TasksController) UpdateTask(gin *gin.Context) {
	// GET THE USER ID
	userID := gin.Param("id")

	// GET THE TASK TO BE UPDATED
	var body struct {
		Text string `json:"text"`
	}

	if err := gin.ShouldBindJSON(&body); err != nil {
		gin.JSON(http.StatusBadRequest, "texto não encontrado no corpo da requisição")
		return
	}

	// UPDATE THE TASK
	task, err := tc.TasksUsecases.UpdateTask(body.Text, userID)

	if err != nil {
		gin.JSON(http.StatusInternalServerError, "erro interno no servidor")
		return
	}

	// RETURN RESPONSE
	gin.JSON(http.StatusOK, task)
}
