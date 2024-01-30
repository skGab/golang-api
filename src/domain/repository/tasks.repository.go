package repository

type TasksRepository interface {
	FindAll(id string) *Status
	Create(text string, userID string) *Status
	Delete(id string) *Status
	DeleteAllTodo(id string) *Status
	DeleteAllDone(id string) *Status
	UpdateState(id string) *Status
	UpdateTask(id string) *Status
}
