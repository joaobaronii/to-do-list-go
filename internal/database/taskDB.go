package database

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5"
	"github.com/joaobaronii/to-do-list-go/internal/entity"
)

func CreateTable(db *sql.DB) error {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS tasks(
				id TEXT PRIMARY KEY,
				name TEXT NOT NULL,
				status BOOL DEFAULT false);`)
	return err
}

func InsertTask(db *sql.DB, task entity.Task) error {
	stmt, err := db.Prepare("INSERT INTO tasks (id, name, status) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(task.ID, task.Name, task.Status)
	if err != nil {
		return err
	}
	return nil
}

func SelectTasksByStatus(db *sql.DB, status bool) ([]entity.Task, error) {
	rows, err := db.Query("SELECT name, status FROM tasks WHERE status = $1", status)
	if err != nil {
		return nil, err
	}
	return rowsToSlice(rows)
}

func SelectAllTasks(db *sql.DB) ([]entity.Task, error) {
	rows, err := db.Query("SELECT name, status FROM tasks")
	if err != nil {
		return nil, err
	}
	return rowsToSlice(rows)
}

func rowsToSlice(rows *sql.Rows) ([]entity.Task, error) {
	defer rows.Close()
	var tasks []entity.Task
	for rows.Next() {
		var t entity.Task
		err := rows.Scan(&t.Name, &t.Status)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func MarkTaskAsDone(db *sql.DB, name string) error {
	stmt, err := db.Prepare("UPDATE tasks SET status = true WHERE name = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllCompletedTasks(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM tasks WHERE status = true")
	if err != nil {
		return err
	}
	return nil
}

func DeleteTask(db *sql.DB, name string) error {
	stmt, err := db.Prepare("DELETE FROM tasks WHERE name = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAllTasks(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM tasks")
	if err != nil {
		return err
	}
	return nil
}
