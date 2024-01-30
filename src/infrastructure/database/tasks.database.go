package database

import (
	"github.com/go-api/src/domain/repository"
	"gorm.io/gorm"
)

type TasksDatabase struct {
	Adapter *gorm.DB
}

// FIND ALL TASKS
func (this *TasksDatabase) FindAll(id string) *repository.Status {}

// CREATE TASK
func (this *TasksDatabase) Create(text string, userID string) *repository.Status {}

// DELETE TASK
func (this *TasksDatabase) Delete(id string) *repository.Status {}

// DELETE ALL TODO TASKS
func (this *TasksDatabase) DeleteAllTodo(id string) *repository.Status {}

// DELETE ALL DONE TASKS
func (this *TasksDatabase) DeleteAllDone(id string) *repository.Status {}

// UPDATE TASK STATE
func (this *TasksDatabase) UpdateState(id string) *repository.Status {}

// UPDATE TASK
func (this *TasksDatabase) UpdateTask(id string) *repository.Status {}
