package adapter

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	dns := "postgres://go_api_db_user:Go6gpdR59ranA8t3grhSuFhI59m3BS6q@dpg-cms0cd21hbls73dr4j3g-a.oregon-postgres.render.com/go_api_db"

	// TRY TO CONNECT TO DB
	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})

	// RETURN ERROR IF CONNECTION FAILS
	if err != nil {
		fmt.Println(err)
		panic("Falha ao conectar com o banco")
	}

	return db
}
