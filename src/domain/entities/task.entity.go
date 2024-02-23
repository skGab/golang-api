package entities

import (
	"time"

	"github.com/google/uuid"
)

// NEED TO TEST THE JSON VALIDATION
type TaskEntity struct {
	ID        string `gorm:"primaryKey"`
	Complete  bool
	Text      string
	CreatedAt time.Time
	UserID    string
}

func NewTask(text string, UserID string) *TaskEntity {
	return &TaskEntity{
		ID:        uuid.New().String(),
		Complete:  false,
		Text:      text,
		CreatedAt: time.Now(),
		UserID:    UserID,
	}
}
