package models

import 
(
	"fmt"
	"errors"
	"events.com/rest-api/db"
	"events.com/rest-api/utils"
)

type User struct{
	ID int64
	FirstName string
	LastName string
	Email string `binding:"required"`
	Password string `binding:"required"`
}


func (u User) Save() error{
	query:="INSERT INTO users(firstname, lastname, email, password) VALUES (?, ?, ?, ?)"
	stmt, err :=db.DB.Prepare(query)

	if err !=nil{
		return err
	}

	defer stmt.Close()

	hashedPassword, err:= utils.HashPassword(u.Password)

	if err !=nil{
		return err
	}

	result, err :=stmt.Exec(u.FirstName, u.LastName, u.Email, hashedPassword)

	if err !=nil{
		return err
	}

	userId, err :=result.LastInsertId()

	u.ID=userId

	return err
	
}

// func(u User) ValidateCredentials() error{
// 	query:= "SELECT password, email FROM users WHERE email=?"
// 	row, err :=db.DB.Query(query, u.Email)
	
// 	var retrievedPassword string
// 	err = row.Scan(&retrievedPassword)

// 	if err !=nil{
// 		return errors.New("Credentials Invalid")
// 	}

// 	passwordIsValid:= utils.CheckPasswordHash(u.Password, retrievedPassword)

// 	if !passwordIsValid{
// 		return errors.New("Credentials Invalid")
// 	}

// 	return nil
// }

func (u *User) ValidateCredentials() error {
    query := "SELECT id, password FROM users WHERE email=?"
    
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