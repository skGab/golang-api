package database

import (
	"github.com/go-api/src/domain/entities"
	"github.com/go-api/src/domain/repository"
	"gorm.io/gorm"
)

type TasksDatabase struct {
	Adapter *gorm.DB
}

// FIND ALL TASKS
func (db *TasksDatabase) FindAll(userID string) *repository.Status {
	var tasks []entities.TaskEntity

	response := db.Adapter.Where("userId = ?", userID).Find(&tasks)

	if response.Error != nil {
		return &repository.Status{Status: false, Message: "Erro interno", Error: response.Error}

	}

	if response.RowsAffected == 0 {
		return &repository.Status{Status: false, Message: "Nenhuma tarefa encontrada"}
	}

	// RETURN RESPONSE
	return &repository.Status{Status: true, Data: &tasks}
}

// // CREATE TASK
func (db *TasksDatabase) Create(text string, userID string) *repository.Status {
	task := entities.NewTask(text, userID)

	response := db.Adapter.Create(&task)

	if response.Error != nil {
		return &repository.Status{Status: false, Message: "Erro interno", Error: response.Error}
	}

	if task == nil {
		return &repository.Status{Status: false, Message: "Dados não encontrados para criação da tarefa"}
	}

	return &repository.Status{Status: true, Data: &task}
}

// // DELETE TASK
func (db *TasksDatabase) Delete(id string) *repository.Status {
	response := db.Adapter.Where("id = ?", id).Delete(&entities.TaskEntity{})

	if response.Error != nil {
		return &repository.Status{Status: false, Message: "Erro interno", Error: response.Error}
	}

	if response.RowsAffected == 0 {
		return &repository.Status{Status: false, Message: "Tarefa não encontrada para exclusão"}
	}

	return &repository.Status{Status: true, Message: "Tarefa deletada"}
}

// // DELETE ALL TODO TASKS
func (db *TasksDatabase) DeleteAllTodo(userID string) *repository.Status {
	response := db.Adapter.Where("userID = ? AND complete = ?", userID, false).Delete(&entities.TaskEntity{})

	if response.Error != nil {
		return &repository.Status{Status: false, Message: "Erro interno", Error: response.Error}
	}

	if response.RowsAffected == 0 {
		return &repository.Status{Status: false, Message: "Erro ao deletar tarefas"}
	}

	return &repository.Status{Status: true, Message: "Tarefas deletadas"}
}

// // DELETE ALL DONE TASKS
func (db *TasksDatabase) DeleteAllDone(id string) *repository.Status {
	response := db.Adapter.Where("id = ? AND complete = ?", id, true).Delete(&entities.TaskEntity{})

	if response.Error != nil {
		return &repository.Status{Status: false, Message: "Erro interno", Error: response.Error}
	}

	if response.RowsAffected == 0 {
		return &repository.Status{Status: false, Message: "Erro ao deletar tarefas"}
	}

	return &repository.Status{Status: true, Message: "Tarefas deletadas"}
}

// // UPDATE TASK STATE
func (db *TasksDatabase) UpdateState(id string, complete bool) *repository.Status {
	var updatedTask *entities.TaskEntity

	// Or if you want to update multiple tasks based on some condition
	response := db.Adapter.Model(&entities.TaskEntity{}).Where("id = ?", id).Update("complete", !complete)

	if response.Error != nil {
		return &repository.Status{Status: false, Message: "Erro interno", Error: response.Error}
	}

	if response.RowsAffected == 0 {
		return &repository.Status{Status: false, Message: "O estado da task não foi atualizado"}
	}

	return &repository.Status{Status: true, Data: &updatedTask}
}

// // UPDATE TASK
func (db *TasksDatabase) UpdateTask(id string, text bool) *repository.Status {
	var updatedTask *entities.TaskEntity

	// Or if you want to update multiple tasks based on some condition
	response := db.Adapter.Model(&entities.TaskEntity{}).Where("id = ?", id).Update("text", text)

	if response.Error != nil {
		return &repository.Status{Status: false, Message: "Erro interno", Error: response.Error}
	}

	if response.RowsAffected == 0 {
		return &repository.Status{Status: false, Message: "A task não foi atualizada"}
	}

	return &repository.Status{Status: true, Data: &updatedTask}
}

