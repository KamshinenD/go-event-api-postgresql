package models

import (
	"events.com/rest-api/db"
)


type Registration struct {
	ID      string  `json:"id"`
	Name string `json:"name"`
	Age     string `json:"age"`
	Address string `json:"address"`
	EventId string `json:"event_id"`
	EventName string `json:"event_name"`
	UserId  string `json:"user_id"`
}


func (e Event) RegisterEvent(userId, name, age, address string) error{
	query := "INSERT INTO registrations(event_id, event_name, user_id, name, age, address) VALUES ($1, $2, $3, $4, $5, $6)"

	stmt, err := db.DB.Prepare(query)

	if err !=nil{
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(e.ID, e.Name, userId, name, age, address)

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
		err:= rows.Scan(&u.ID, &u.Name, &u.Age, &u.Address, &u.EventId, &u.EventName, &u.UserId)
		if err !=nil{
			return nil, err
		}
		registrations= append(registrations, u)
	}
	return registrations, nil
}


func GetRegistrationByID(id string)(*Registration, error){
	query:= "SELECT * FROM registrations WHERE id = $1" //we add ? instead of inputting id to avoid sql injection
	row :=db.DB.QueryRow(query, id)

	var reg Registration
	err:= row.Scan(&reg.ID, &reg.Name, &reg.Age, &reg.Address, &reg.EventId, &reg.EventName, &reg.UserId)

	if err !=nil{
		return nil, err
	}
	return &reg, nil
	//pls note that we had to use pointer for event so that it can take a nil value when there is an error
}


func (r Registration) DeleteReg() error {
    query := "DELETE FROM registrations WHERE id = $1"
    
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


func (r Registration) CancelReg(userId string) error{
	query := "DELETE FROM registrations WHERE event_id = $1 AND user_id = $2"

	stmt, err := db.DB.Prepare(query)

	if err !=nil{
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(r.ID, userId)

	return err
}