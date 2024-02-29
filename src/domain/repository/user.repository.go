package repository

import (
	entitie "github.com/go-api/src/domain/entities"
)

type UserRepository interface {
	Create(*entitie.UserEntity) (interface{}, error)
	FindAll() ([]entitie.UserEntity, error)
}
