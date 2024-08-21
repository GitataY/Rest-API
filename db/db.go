package db

import (
	"database/sql"
	"log"
	

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") // Assign to the package-level DB

	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	// Ping the database to ensure the connection is established
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Could not establish a connection: %v", err)
	}

	createTables()
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL UNIQUE,
		Password TEXT NOT NULL
	)`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatalf("Could not create users table: %v", err)
	}

	createEventsTable := `CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		Name TEXT NOT NULL,
		Location TEXT NOT NULL,
		Description TEXT NOT NULL,
		DateTime DATETIME NOT NULL,
		UserID INTEGER,
		FOREIGN KEY(UserID) REFERENCES users(id)
	)`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Fatalf("Could not create events table: %v", err)
	}
}