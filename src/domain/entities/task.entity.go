package entities

import (
	"time"

	"github.com/google/uuid"
)

// NEED TO TEST THE JSON VALIDATION
type TaskEntity struct {
	ID        string
	Complete  bool
	Text      string
	CreatedAt time.Time
	UserID    string `gorm:"primaryKey"`
}

func NewTask(text string, userID string, ids ...string) *TaskEntity {
	var taskID string

	if len(ids) > 0 && ids[0] != "" {
		taskID = ids[0]
	} else {
		taskID = uuid.New().String()
	}

	return &TaskEntity{
		ID:        taskID,
		Complete:  false,
		Text:      text,
		CreatedAt: time.Now(),
		UserID:    userID,
	}
}
