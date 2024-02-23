package entities

import "github.com/google/uuid"

type UserEntity struct {
	ID       string `gorm:"primaryKey"`
	User     string
	Email    string
	Password string
	Tasks    []TaskEntity `gorm:"foreignKey:UserID"`
}

func NewUser(user *UserEntity) *UserEntity {
	var userID string

	if user.ID == "" {
		userID = uuid.New().String()
	} else {
		userID = user.ID
	}

	return &UserEntity{
		ID:       userID,
		User:     user.User,
		Email:    user.Email,
		Password: user.Password,
		Tasks:    user.Tasks,
	}
}
