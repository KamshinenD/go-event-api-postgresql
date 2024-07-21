package db

import (
	"database/sql"
	"log"

	 _ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10) //number of connections that can be opened simultaneausly
	DB.SetMaxIdleConns(5)  //no. of connections to be allowed when idle

	createTables()
}

func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER
	)
	`

	_, err := DB.Exec(createEventsTable)
	if err != nil {
		// panic("Could not create events table")
		log.Fatal("Could not create events table:", err)
	}
}
