package dto

import (
	"time"

	"github.com/go-api/src/domain/entities"
)

type TasksDTO struct {
	ID        string
	Complete  bool
	Text      string
	CreatedAt time.Time
	UserID    string
}

func NewTasksDTO(userEntity entities.UserEntity) []TasksDTO {

	tasksDTO := make([]TasksDTO, 0, len(userEntity.Tasks))

	for _, task := range userEntity.Tasks {
		tasks := TasksDTO{
			ID:        task.ID,
			Complete:  task.Complete,
			Text:      task.Text,
			CreatedAt: task.CreatedAt,
			UserID:    task.UserID,
		}

		tasksDTO = append(tasksDTO, tasks)
	}

	return tasksDTO

}
