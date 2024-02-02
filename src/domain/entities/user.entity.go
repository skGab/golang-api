package entities

type UserEntity struct {
	ID       string `gorm:"primaryKey"`
	User     string
	Email    string
	Password string
	Tasks    []TaskEntity `gorm:"foreignKey:UserID"`
}
