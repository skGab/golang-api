package adapter

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dns := ""

	// TRY TO CONNECT TO DB
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	// RETURN ERROR IF CONNECTION FAILS
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
