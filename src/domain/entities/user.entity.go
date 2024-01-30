package entities

type UserEntity struct {
	ID       string `gorm:"primaryKey"`
	Email    string
	Password string
	User     string
	Tasks    []TasksEntity `gorm:"foreignKey:UserID"`
}
