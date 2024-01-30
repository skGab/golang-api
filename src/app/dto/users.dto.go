package dto

type UsersDTO struct {
	ID       string
	User     string
	Email    string
	Password string
	Tasks    []TasksDTO
}
