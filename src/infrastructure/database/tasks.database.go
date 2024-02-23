package database

import (
	"errors"

	"github.com/go-api/src/domain/entities"
	"gorm.io/gorm"
)

type TasksDatabase struct {
	Adapter *gorm.DB
}

// FIND ALL TASKS
func (db *TasksDatabase) FindAll(userID string) ([]entities.TaskEntity, error) {
	var tasks []entities.TaskEntity

	response := db.Adapter.Where("user_id = ?", userID).Find(&tasks)

	if response.Error != nil {
		return nil, response.Error
	}

	if response.RowsAffected == 0 {
		return nil, errors.New("nenhuma tarefa encontrada")
	}

	// RETURN RESPONSE
	return tasks, nil
}

// // CREATE TASK
func (db *TasksDatabase) Create(text string, userID string) (*entities.TaskEntity, error) {
	// CHECK IF USE.ID EXISTS BEFORE CREATING THE TASK
	userStatus := db.Adapter.Raw("SELECT id FROM user_entities WHERE id = ?", userID).Scan(&entities.UserEntity{})

	if userStatus.Error != nil {
		return nil, userStatus.Error
	}

	if userStatus.RowsAffected == 0 {
		return nil, errors.New("precisa existir um usuario para criar tarefas")
	}

	// MAP THE DATA TO TASK ENTITY
	newTaskEntity := entities.NewTask(text, userID)

	// CREATE THE TASK
	response := db.Adapter.Create(&newTaskEntity)

	if response.Error != nil {
		return nil, response.Error
	}

	if response.RowsAffected == 0 {
		return nil, errors.New("dados não encontrados para criação da tarefa")
	}

	// RETURN RESPONSE
	return newTaskEntity, nil
}

// // DELETE TASK
func (db *TasksDatabase) Delete(id string) (string, error) {
	response := db.Adapter.Where("id = ?", id).Delete(&entities.TaskEntity{})

	if response.Error != nil {
		return "", response.Error
	}

	if response.RowsAffected == 0 {
		return "Tarefa não encontrada para exclusão", nil
	}

	return "Tarefa deletada", nil
}

// // DELETE ALL TODO TASKS
func (db *TasksDatabase) DeleteAllTodo(userID string) (string, error) {
	response := db.Adapter.Where("user_id = ? AND complete = ?", userID, false).Delete(&entities.TaskEntity{})

	if response.Error != nil {
		return "", response.Error
	}

	if response.RowsAffected == 0 {
		return "Algo aconteceu ao deletar as tarefas", response.Error
	}

	return "Tarefas deletadas", nil
}

// // DELETE ALL DONE TASKS
func (db *TasksDatabase) DeleteAllDone(userID string) (string, error) {
	response := db.Adapter.Where("user_id = ? AND complete = ?", userID, true).Delete(&entities.TaskEntity{})

	if response.Error != nil {
		return "", response.Error
	}

	if response.RowsAffected == 0 {
		return "Algo aconteceu ao deletar tarefas", nil
	}

	return "Tarefas deletedas", nil
}

// // UPDATE TASK STATE
func (db *TasksDatabase) UpdateState(id string, complete bool) (string, error) {
	response := db.Adapter.Model(&entities.TaskEntity{}).Where("id = ?", id).Update("complete", !complete)

	if response.Error != nil {
		return "nil", response.Error
	}

	if response.RowsAffected == 0 {
		return "nil", errors.New("o estado da task não foi atualizado")
	}

	return "Estado alterado", nil
}

// // UPDATE TASK
func (db *TasksDatabase) UpdateTask(id string, text string) (string, error) {
	response := db.Adapter.Model(&entities.TaskEntity{}).Where("id = ?", id).Update("text", text)

	if response.Error != nil {
		return "nil", response.Error
	}

	if response.RowsAffected == 0 {
		return "nil", errors.New("a task não foi atualizada")
	}

	return "Tarefa atualizada", nil
}
