package models

import (
	"time"
	"events.com/rest-api/db"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int
}

// var events = []Event{}
var events = []Event{
	{
		ID:          1,
		Name:        "Initial Event",
		Description: "This is the initial event",
		Location:    "New York",
		DateTime:    time.Date(2024, time.July, 13, 10, 0, 0, 0, time.UTC),
		UserId:      1,
	},
}

func (e Event) Save() error {
	//later add it to database
	query:= `INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES(?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err !=nil{
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err !=nil{
		return err
	}

	id, err :=result.LastInsertId()
	e.ID= id
	return err
	// events = append(events, e)
}

func GetAllEvents() ([]Event, error) {
	query:= "SELECT * FROM events"
	rows, err :=db.DB.Query(query)
	if err !=nil{
		return nil, err
	}
	defer rows.Close()
	var events []Event
	
	for rows.Next(){
		var event Event 
		err:= rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId )
		if err !=nil{
			return nil, err
		}
		events= append(events, event)
	}
	return events, nil
	return events, nil
}

func GetEventByID(id int64)(*Event, error){
	query:= "SELECT * FROM events WHERE id = ?" //we add ? instead of inputting id to avoid sql injection
	row :=db.DB.QueryRow(query, id)

	var event Event
	err:= row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId )
	if err !=nil{
		return nil, err
	}
	return &event, nil
	//pls note that we had to use pointer for event so that it can take a nil value when there is an error
}
