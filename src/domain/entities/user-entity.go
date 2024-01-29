package entitie

import "github.com/go-api/src/domain/valueObjects"

type UserEntity struct {
	ID       string `gorm:"primaryKey"`
	Email    string
	Password string
	User     string
	Tasks    []valueObjects.TasksVo `gorm:"foreignKey:UserID"`
}
