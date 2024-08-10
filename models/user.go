package models

import 
(
	"fmt"
	"errors"
	"events.com/rest-api/db"
	"events.com/rest-api/utils"
)

type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email" binding:"required"`
	Password  string `json:"password" binding:"required"`
}


// func (u User) Save() error{
// 	query:="INSERT INTO users(firstname, lastname, email, password) VALUES ($1, $2, $3, $4)"
// 	stmt, err :=db.DB.Prepare(query)

// 	if err !=nil{
// 		return err
// 	}

// 	defer stmt.Close()

// 	hashedPassword, err:= utils.HashPassword(u.Password)

// 	if err !=nil{
// 		return err
// 	}

// 	result, err :=stmt.Exec(u.FirstName, u.LastName, u.Email, hashedPassword)

// 	if err !=nil{
// 		return err
// 	}

// 	userId, err :=result.LastInsertId()

// 	u.ID=userId

// 	return err
	
// }

func (u *User) Save() error {
	query := `INSERT INTO users(firstname, lastname, email, password) 
	VALUES($1, $2, $3, $4) RETURNING id`

	// Prepare and execute the query
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	// Use QueryRow to execute the query and retrieve the generated ID
	err = stmt.QueryRow(u.FirstName, u.LastName, u.Email, hashedPassword).Scan(&u.ID)
	if err != nil {
		return err
	}

	return nil
}


func (u *User) ValidateCredentials() error {
    query := "SELECT id, password FROM users WHERE email=$1"
    
	var retrievedPassword string
	row :=db.DB.QueryRow(query, u.Email)
    // err := db.DB.QueryRow(query, u.Email).Scan(&u.ID, &retrievedPassword) //binding the password. Weare also binding the UserID so that we can acccess it to generate jwt token during login 
    err := row.Scan(&u.ID, &retrievedPassword) //binding the password. Weare also binding the UserID so that we can acccess it to generate jwt token during login 

    if err != nil {
			fmt.Println(err)
            return errors.New("credentials Invalid...")
    }

    passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)

    if !passwordIsValid {
        return errors.New("credentials Invalid")
    }

    return nil
}



func GetAllUsers() ([]User, error) {
	query:= "SELECT * FROM users"
	rows, err :=db.DB.Query(query)
	if err !=nil{
		return nil, err
	}
	defer rows.Close()
	var allUsers []User
	
	for rows.Next(){
		var user User 
		err:= rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.Password )
		if err !=nil{
			return nil, err
		}
		allUsers= append(allUsers, user)
	}
	return allUsers, nil
}