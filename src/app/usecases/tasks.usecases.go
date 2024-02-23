package usecases

import (
	"fmt"

	"github.com/go-api/src/app/dto"
	"github.com/go-api/src/domain/repository"
)

type TasksUsecases struct {
	TasksRepository repository.TasksRepository
}

// FIND ALL TASKS
func (tu *TasksUsecases) FindAll(userID string) ([]dto.TasksDTO, error) {
	// FETCH THE TASKS
	tasks, err := tu.TasksRepository.FindAll(userID)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// MAP TO ENTITIES TO DTO
	tasksDTOs := make([]dto.TasksDTO, 0, len(tasks))

	for _, task := range tasks {
		taskDTO := dto.TasksDTO{
			ID:        task.ID,
			Complete:  task.Complete,
			Text:      task.Text,
			CreatedAt: task.CreatedAt,
			UserID:    task.UserID,
		}

		tasksDTOs = append(tasksDTOs, taskDTO)
	}

	return tasksDTOs, nil
}

// CREATE TASK
func (tu *TasksUsecases) Create(text string, userEmail string) (*dto.TasksDTO, error) {
	// CREATE AND GET THE TASK ENTITY
	taskEntity, err := tu.TasksRepository.Create(text, userEmail)

	// CHECK FOR ERROS
	if err != nil {
		fmt.Print(err.Error())
		return nil, err
	}

	// MAP THE ENTITIE TO DTO
	taskDTO := &dto.TasksDTO{
		ID:        taskEntity.ID,
		Complete:  taskEntity.Complete,
		Text:      taskEntity.Text,
		CreatedAt: taskEntity.CreatedAt,
		UserID:    taskEntity.UserID,
	}

	return taskDTO, nil
}

// DELETE TASK
func (tu *TasksUsecases) Delete(userID string) (string, error) {
	// DELETE THE TASK
	response, err := tu.TasksRepository.Delete(userID)

	// CHECK FOR ERROS
	if err != nil {
		fmt.Print(err.Error())
		return "", err
	}

	return response, nil
}

// DELETE ALL TODO TASKS
func (tu *TasksUsecases) DeleteAllTodo(userID string) (string, error) {
	// EXECUTE THE CASE
	response, err := tu.TasksRepository.DeleteAllTodo(userID)

	// CHECK FOR ERROS
	if err != nil {
		fmt.Print(err.Error())
		return "", err
	}

	return response, nil
}

// DELETE ALL DONE TASKS
func (tu *TasksUsecases) DeleteAllDone(userID string) (string, error) {
	// EXECUTE THE CASE
	response, err := tu.TasksRepository.DeleteAllDone(userID)

	// CHECK FOR ERROS
	if err != nil {
		fmt.Print(err.Error())
		return "", err
	}

	return response, err
}

// UPDATE TASK STATE
func (tu *TasksUsecases) UpdateState(complete bool, id string) (string, error) {
	response, err := tu.TasksRepository.UpdateState(id, complete)

	// CHECK
	if err != nil {
		fmt.Print(err.Error())
		return "nil", err
	}

	return response, nil
}

// UPDATE TASK
func (tu *TasksUsecases) UpdateTask(text string, id string) (string, error) {
	// UPDATE THE TASK
	updatedTask, err := tu.TasksRepository.UpdateTask(id, text)

	// CHECK FOR ERROS
	if err != nil {
		fmt.Print(err.Error())
		return "", err
	}

	// RETURN RESPONSE
	return updatedTask, nil
}
