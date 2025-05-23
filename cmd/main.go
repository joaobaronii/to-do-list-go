package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joaobaronii/to-do-list-go/configs"
	"github.com/joaobaronii/to-do-list-go/internal/database"
)

func main() {
	config := configs.LoadConfig("cmd")	
	fmt.Printf("Config completa: %+v\n", *config)

	url := config.GetPostgresURL()

	db, err := sql.Open("pgx", url)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	database.CreateTable(db)
}