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
	fmt.Println("Task added successfully!")
	return nil
}