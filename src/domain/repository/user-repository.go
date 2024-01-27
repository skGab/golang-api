package repository

import (
	entitie "github.com/go-api/src/domain/entities"
)

type Status struct {
	Status  bool
	Message string
}

type UserRepository interface {
	FindAll() (*[]entitie.UserEntity, *Status)
	Create(*entitie.UserEntity) *Status
}
