package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func initDB() {
	var err error

	if DB, err = sql.Open("sqlite", "dev.db"); err != nil {
		log.Fatal(err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Database connected!")

	createTables()
	fmt.Println("Tables created!")
}

func createTables() {
	query := `
		CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT NOT NULL,
		description TEXT,
		degree_of_difficulty INTEGER,
		deadline DATETIME,
		repeatable BOOLEAN,
		active BOOLEAN,
		parent_id INTEGER,

		FOREIGN KEY(parent_id) REFERENCES tasks(id)
		);

		CREATE TABLE IF NOT EXISTS task_events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			task_id INTEGER,
			event_type TEXT NOT NULL,
			timestamp DATETIME NOT NULL,

			FOREIGN KEY(task_id) REFERENCES tasks(id)
		);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}
