package models

import (
	"events.com/rest-api/db"
)


type Registration struct {
	ID      int64  `db:"id"`
	EventId int64  `db:"event_id"`
	UserId  int64  `db:"user_id"`
}


func (e Event) RegisterEvent(userId int64) error{
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err !=nil{
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func GetAllEventsRegistration() ([]Registration, error) {
	query:= "SELECT * FROM registrations"
	rows, err :=db.DB.Query(query)
	if err !=nil{
		return nil, err
	}
	defer rows.Close()

	var registrations []Registration
	
	for rows.Next(){
		var u Registration 
		err:= rows.Scan(&u.ID, &u.EventId, &u.UserId)
		if err !=nil{
			return nil, err
		}
		registrations= append(registrations, u)
	}
	return registrations, nil
}


func GetRegistrationByID(id int64)(*Registration, error){
	query:= "SELECT * FROM registrations WHERE id = ?" //we add ? instead of inputting id to avoid sql injection
	row :=db.DB.QueryRow(query, id)

	var reg Registration
	err:= row.Scan(&reg.ID, &reg.EventId, &reg.UserId )
	if err !=nil{
		return nil, err
	}
	return &reg, nil
	//pls note that we had to use pointer for event so that it can take a nil value when there is an error
}


func (r Registration) DeleteReg() error {
    query := "DELETE FROM registrations WHERE id = ?"
    
    stmt, err := db.DB.Prepare(query)
    if err != nil {
        return err
    }
    defer stmt.Close()
    _, err = stmt.Exec(r.ID)
 
	if err != nil {
        return err
    }
    return nil
}


func (r Registration) CancelReg(userId int64) error{
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"

	stmt, err := db.DB.Prepare(query)

	if err !=nil{
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.ID, userId)

	return err
}