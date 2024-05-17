package adapter

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// err := godotenv.Load()

	// if err != nil {
	// 	panic(err)
	// }

	databaseURL := os.Getenv("DATABASE_URL")

	// TRY TO CONNECT TO DB
	db, err := gorm.Open(postgres.Open(databaseURL))

	// RETURN ERROR IF CONNECTION FAILS
	if err != nil {
		fmt.Println(err)
	}

	return db
}
