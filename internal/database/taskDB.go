package database

import (
	"database/sql"

	_ "github.com/jackc/pgx/v5"
)



func CreateTable(db *sql.DB) error{
	query := `CREATE TABLE IF NOT EXISTS tasks(
				id TEXT PRIMARY KEY,
				name TEXT NOT NULL,
				status BOOL DEFAULT false);`
	
	_, err := db.Exec(query)
	return err
}
