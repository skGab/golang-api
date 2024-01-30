package repository

import (
	entitie "github.com/go-api/src/domain/entities"
)

type Status struct {
	Status  bool
	Message string
	Data    interface{}
	Error   error
}

type UserRepository interface {
	FindAll() *Status
	Create(*entitie.UserEntity) *Status
}
