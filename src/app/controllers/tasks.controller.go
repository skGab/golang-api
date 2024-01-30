package controller

import (
	"github.com/go-api/src/domain/repository"
)

type TasksController struct {
	TasksRepository *repository.TasksRepository
}
