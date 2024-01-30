package entities

import "time"

type TasksEntity struct {
	ID        string `gorm:"primaryKey"`
	Complete  bool
	Text      string
	CreatedAt time.Time
	UserID    string
}
