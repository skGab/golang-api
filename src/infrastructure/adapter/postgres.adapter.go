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

	dns := databaseURL

	// TRY TO CONNECT TO DBe
	db, err := gorm.Open(postgres.Open(dns))

	// RETURN ERROR IF CONNECTION FAILS
	if err != nil {
		fmt.Println(err)
		panic("Falha ao conectar com o banco")
	}

	return db
}
