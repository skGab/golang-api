package repository

import (
	entitie "github.com/go-api/src/domain/entities"
)

type UserRepository interface {
	Create(*entitie.UserEntity) (error, *entitie.UserEntity)
	FindAll() (error, *[]entitie.UserEntity)
}
