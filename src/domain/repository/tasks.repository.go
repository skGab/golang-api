package repository

import "github.com/go-api/src/domain/entities"

type TasksRepository interface {
	FindAll(userID string) (interface{}, error)
	Create(text string, userID string) (*entities.TaskEntity, error)
	Delete(id string) (string, error)
	DeleteAllTodo(id string) (string, error)
	DeleteAllDone(id string) (string, error)
	UpdateState(id string, complete bool) (string, error)
	UpdateTask(id string, text string) (string, error)
}
