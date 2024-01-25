package repository

import (
	entitie "github.com/go-api/src/domain/entities"
)

type StatusResponse struct {
	Status string
}
type UserRepository interface {
	FindAll() (*[]entitie.UserEntity, error)
	Create(*entitie.UserEntity) *StatusResponse
}
