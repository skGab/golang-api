package adapter

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	databaseURL := os.Getenv("DATABASE_URL")

	dns := databaseURL

	// TRY TO CONNECT TO DB
	db, err := gorm.Open(postgres.Open(dns))

	// RETURN ERROR IF CONNECTION FAILS
	if err != nil {
		fmt.Println(err)
		panic("Falha ao conectar com o banco")
	}

	return db
}
