package valueObjects

import "time"

type TasksVo struct {
	ID        string `gorm:"primaryKey"`
	Complete  bool
	Text      string
	CreatedAt time.Time
	UserID    string
}
