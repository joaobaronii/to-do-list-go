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

	url := config.GetPostgresURL()

	db, err := sql.Open("pgx", url)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = database.CreateTable(db)
	if err != nil {
		panic(err)
	}

	var option int

	for option != 5 {
		fmt.Println("Welcome to the To-Do List App!")
		fmt.Println("--------------------------------")
		fmt.Println("1. Add task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Mark task as done")
		fmt.Println("4. Delete task")
		fmt.Println("5. Exit")
		fmt.Print("Select an option: ")

		fmt.Scan(&option)

		switch option {
			case 1:
				
		}
	}
}