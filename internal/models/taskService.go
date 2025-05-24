package model

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"

	"github.com/joaobaronii/to-do-list-go/internal/database"
	"github.com/joaobaronii/to-do-list-go/internal/entity"
)

func AddTask(db *sql.DB) error {
	var taskName string

	fmt.Print("Enter the task: ")

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

	t := entity.NewTask(taskName)

	err = database.InsertTask(db, t)
	if err != nil {
		return err
	} 
	return nil
}

func ListTasks(db *sql.DB, tasks []entity.Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	fmt.Println("Tasks:")
	for _, task := range tasks {
		var done string

		if task.Status {
			done = "✅"
		} else {
			done = "❌"
		}
		fmt.Printf("%s - %s\n", task.Name, done)
	}
	return
}

func DeleteTaskByName(db *sql.DB) error {
	var taskName string

	fmt.Print("Enter the task: ")

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

	err = database.DeleteTask(db, taskName)
	if err != nil {
		return err
	} 
	return nil
}