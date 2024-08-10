package models

import (
	"time"
	"events.com/rest-api/db"
)

type Event struct {
	ID          string    `json:"id"`          // UUID represented as a string
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserId      string    `json:"userId"`      // UUID represented as a string
}


// var events = []Event{}
var events = []Event{
	// {
	// 	ID:          1,
	// 	Name:        "Initial Event",
	// 	Description: "This is the initial event",
	// 	Location:    "New York",
	// 	DateTime:    time.Date(2024, time.July, 13, 10, 0, 0, 0, time.UTC),
	// 	UserId:      1,
	// },
}

// func (e *Event) Save() error {
// 	//later add it to database
// 	query := `INSERT INTO events(name, description, location, dateTime, user_id) 
// 	VALUES($1, $2, $3, $4, $5) RETURNING id`
// 	stmt, err := db.DB.Prepare(query)
// 	if err !=nil{
// 		return err
// 	}
// 	defer stmt.Close()
// 	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)
// 	if err !=nil{
// 		return err
// 	}

// 	id, err :=result.LastInsertId()
// 	e.ID= id
// 	return err
// 	// events = append(events, e)
// }

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES($1, $2, $3, $4, $5) RETURNING id`

	// Use QueryRow to execute the query and retrieve the generated ID
	err := db.DB.QueryRow(query, e.Name, e.Description, e.Location, e.DateTime, e.UserId).Scan(&e.ID)
	if err != nil {
		return err
	}
	return nil
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
	// return events, nil
}

func GetEventByID(id string)(*Event, error){
	query := "SELECT * FROM events WHERE id = $1"//we add ? instead of inputting id to avoid sql injection
	row :=db.DB.QueryRow(query, id)

	var event Event
	err:= row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId )
	if err !=nil{
		return nil, err
	}
	return &event, nil
	//pls note that we had to use pointer for event so that it can take a nil value when there is an error
}


func (event Event) Update() error{
	query := `
		UPDATE events
		SET name=$1, description=$2, location=$3, dateTime=$4
		WHERE id=$5
	`

	stmt, err := db.DB.Prepare(query)

	if err !=nil{
		return err
	}

	defer stmt.Close()

	_, err=stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}


func (e Event) Delete() error {
   query := "DELETE FROM events WHERE id = $1"
    
    stmt, err := db.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()

    // result, err := stmt.Exec(e.ID)
    _, err = stmt.Exec(e.ID)
 
	if err != nil {
        return err
    }

    // rowsAffected, err := result.RowsAffected()
    // if err != nil {
    //     return fmt.Errorf("failed to get affected rows: %w", err)
    // }

    // if rowsAffected == 0 {
    //     return fmt.Errorf("no event found with id %d", e.ID)
    // }

    return nil
}