package repository

import (
	valueObjects "github.com/go-api/src/domain/valueObjects"
)

type UserRepository interface {
	FindAll() ([]valueObjects.UserVo, error)
	Create() valueObjects.CreateUserVO
}
