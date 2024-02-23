package repository

import (
	entitie "github.com/go-api/src/domain/entities"
)

type UserRepository interface {
	Create(*entitie.UserEntity) (*entitie.UserEntity, error)
	FindAll() ([]entitie.UserEntity, error)
}
