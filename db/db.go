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
	createUsersTable:=`
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		firstname TEXT NOT NULL,
		lastname TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		password TEXT NOT NULL
	)
	`
	_, err:= DB.Exec(createUsersTable)
	if err != nil {
		// panic("Could not create users table")
		log.Fatal("Could not create users table:", err)
	}

	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INTEGER,
		FOREIGN KEY( user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createEventsTable)
	if err != nil {
		// panic("Could not create events table")
		log.Fatal("Could not create events table:", err)
	}

	createRegistrationsTable:=`
	CREATE TABLE IF NOT EXISTS registrations (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_Id INTEGER,
	user_Id INTEGER,
	FOREIGN KEY(event_id) REFERENCES events(id),
	FOREIGN KEY(user_id) REFERENCES users(id)
	)
	`
	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		// panic("Could not create registrations table")
		log.Fatal("Could not create registration table:", err)
	}
}
