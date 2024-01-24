package database

import (
	"log"

	"github.com/go-api/src/domain/valueObjects"
)

type UserDatabase struct{}

// CREATE USER
func (*UserDatabase) Create() valueObjects.CreateUserVO {
	log.Print("Rota chamada")
	return valueObjects.CreateUserVO{}
}

// GET USERS
func (*UserDatabase) FindAll() ([]valueObjects.UserVo, error) {
	users := []valueObjects.UserVo{userVo, userVo, userVo, userVo}

	return users, nil
}
