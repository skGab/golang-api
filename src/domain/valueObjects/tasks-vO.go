package valueObjects

import "time"

type TasksVo struct {
	ID        string
	Complete  bool
	Text      string
	CreatedAt time.Time
	UserID    string
}
