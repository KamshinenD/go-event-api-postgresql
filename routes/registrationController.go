package routes


import (
	// "strconv"
	"net/http"

	"events.com/rest-api/models"

	"github.com/gin-gonic/gin"

)
func EventRegistration(context *gin.Context){
	// userId:= context.GetInt64("userId")
	userId, exists := context.Get("userId")
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "user ID not found in context"})
		return
	}

	// Type assertion for userId
	userIdStr, ok := userId.(string)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "user ID is not a string"})
		return
	}

	// Get eventId from URL parameters as a string
	eventId := context.Param("id")

	// Parse and validate registration details from the request body
	var registrationData models.Registration
	err := context.BindJSON(&registrationData); 
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request data"})
		return
	}

	if registrationData.Name == "" || registrationData.Age == "" || registrationData.Address == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Name, age, and address are required"})
		return
	}

	event, err := models.GetEventByID(eventId)

	
	if err !=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not find event"})
		return
	}

	err= event.RegisterEvent(userIdStr, registrationData.Name, registrationData.Age, registrationData.Address)

	if err !=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event ", "error":err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "registration successful"})

	}	


func getAllRegistrations(context *gin.Context) {
	data, err := models.GetAllEventsRegistration()
	if err !=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch registrations.", "error":err.Error()})
		return
	}
	// context.JSON(http.StatusOK, events)
	context.JSON(http.StatusOK, gin.H{"data": data})
}




func cancelRegistration(context *gin.Context){
	eventId:= context.Param("id")
	// eventId, err:= strconv.ParseInt(context.Param("id"), 10, 64)
	// userId:=context.GetInt64("userId") //userId from the token, was passed through the authorization
	userId, exists:= context.Get("userId") //userId from the token, was passed through the authorization
	if !exists {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "user ID not found in context"})
		return
	}

	// Type assertion for userId to ensure it's string
	userIdStr, ok := userId.(string)
	if !ok {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "user ID is not a string"})
		return
	}

	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
	// 	return
	// }

	reg, err := models.GetRegistrationByID(eventId)

	if err !=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fecth registration data"})
		return
	}

	if reg.UserId != userIdStr{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to delete this registration"})
		return
	}

	err = reg.CancelReg(userIdStr)
	// err=reg.DeleteReg()
	if err !=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete registration", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "registration data deleted successfully"})
}


func getSingleRegistration (context *gin.Context){
	eventId:= context.Param("id")
	reg, err := models.GetRegistrationByID(eventId)

	if err !=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fecth registration data"})
		return
	}
	
	context.JSON(http.StatusOK, gin.H{"data": reg})
}