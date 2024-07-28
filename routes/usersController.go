package routes

import (
	"net/http"
	"fmt"

	"events.com/rest-api/models"
	"events.com/rest-api/utils"

	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	
	err=user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("Could not save data: %v", err)})
		return
	}

		context.JSON(http.StatusCreated, gin.H{"message":"user created successfuly"})

}

func login(context *gin.Context){
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}
	
	err=user.ValidateCredentials()

	// if err !=nil{
	// 	context.JSON(http.StatusUnauthorized, gin.H{"message":err})
	// 	return
	// }
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID) // we are able to access user.ID cos it has been binded in ValidateCredentials

	if err != nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate User"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message":"Login Successful", "token": token})

}

func getAllUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err !=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	// context.JSON(http.StatusOK, events)
	context.JSON(http.StatusOK, gin.H{"data": users})
}