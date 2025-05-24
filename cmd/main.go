package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joaobaronii/to-do-list-go/configs"
	"github.com/joaobaronii/to-do-list-go/internal/database"
	model "github.com/joaobaronii/to-do-list-go/internal/models"
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
		fmt.Println("--------------------------------")
		fmt.Println("Welcome to the To-Do List App!")
		fmt.Println("1. Add task")
		fmt.Println("2. List tasks")
		fmt.Println("3. Mark task as done")
		fmt.Println("4. Delete task")
		fmt.Println("5. Exit")
		fmt.Print("Select an option: ")
		fmt.Scan(&option)

		switch option {
			case 1:
				err = model.AddTask(db)	
				if err != nil {
					panic(err)
				}
				fmt.Println("Task added successfully!")
			case 2:
				err = menuListTasks(db)
				if err != nil {
					panic(err)
				}
			case 3:
				err = taskDone(db)
				if err != nil {
					panic(err)
				}
			case 4:
				err = menuDeleteTask(db)
				if err != nil {
					panic(err)
				}
			case 5: 
				fmt.Println("Exiting the application...")
				return
			default:
				fmt.Println("Invalid option. Please try again.")
		}
			
	}
}

func menuListTasks(db *sql.DB) error{
	var op int

	for op != 4 {
		fmt.Println("--------------------------------")
		fmt.Println("1. List all tasks")
		fmt.Println("2. List done tasks")
		fmt.Println("3. List not done tasks")
		fmt.Println("4. Back to main menu")
		fmt.Print("Select an option: ")

		fmt.Scan(&op)

		switch op {
			case 1:
				tasks, err := database.SelectAllTasks(db)
				if err != nil {
					return err
				}
				model.ListTasks(db, tasks)
				return nil
			case 2:
				tasks, err := database.SelectTasksByStatus(db, true)
				if err != nil {
					return err
				}
				model.ListTasks(db, tasks)
				return nil
			case 3: 
				tasks, err := database.SelectTasksByStatus(db, false)
				if err != nil {
					return err
				}
				model.ListTasks(db, tasks)
				return nil
			case 4:
				fmt.Println("Back to main menu")
				return nil
			default:
				fmt.Println("Invalid option. Please try again.")
		}	
	}
	fmt.Println("--------------------------------")
	return nil
}

func taskDone(db *sql.DB) error {
	var taskName string
	fmt.Print("Enter the task to mark as done: ")
	
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	taskName = strings.TrimSpace(input)
	if taskName == "" {
		fmt.Println("Task name cannot be empty.")
		return nil
	}

	err = database.MarkTaskAsDone(db, taskName)
	if err != nil {
		return err
	}
	fmt.Println("Task marked as done successfully!")
	return nil
}

func menuDeleteTask(db *sql.DB) error {
	var op int

	for op != 4 {
		fmt.Println("--------------------------------")
		fmt.Println("1. Delete all tasks")
		fmt.Println("2. Delete done tasks")
		fmt.Println("3. Delete task by name")
		fmt.Println("4. Back to main menu")
		fmt.Print("Select an option: ")

		fmt.Scan(&op)

		switch op {
			case 1:
				err := database.DeleteAllTasks(db)
				if err != nil {
					return err
				}
				fmt.Println("All tasks deleted successfully!")
				return nil
			case 2:
				err := database.DeleteAllDoneTasks(db)
				if err != nil {
					return err
				}
				fmt.Println("All done tasks deleted successfully!")
				return nil
			case 3: 
				err := model.DeleteTaskByName(db)
				if err != nil {
					return err
				}
				fmt.Println("Task removed successfully!")
				return nil
			case 4:
				fmt.Println("Back to main menu")
				return nil
			default:
				fmt.Println("Invalid option. Please try again.")
		}	
	}
	fmt.Println("--------------------------------")
	return nil
}