package entitie

import "github.com/go-api/src/domain/valueObjects"

type UserEntity struct {
	ID       string
	Email    int
	Password string
	User     string
	Tasks    []valueObjects.TasksVo
}
