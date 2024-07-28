package routes


import (
	"strconv"
	"net/http"

	"events.com/rest-api/models"

	"github.com/gin-gonic/gin"

)
func EventRegistration(context *gin.Context){
	userId:= context.GetInt64("userId")
	eventId, err:= strconv.ParseInt(context.Param("id"), 10, 64)
	if err !=nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse event id"})
		return
	}
	event, err := models.GetEventByID(eventId)
	
	if err !=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not find event"})
		return
	}

	err= event.RegisterEvent(userId)

	if err !=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register user for event ", "error":err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "registration successful"})

	}	


func getAllRegistrations(context *gin.Context) {
	data, err := models.GetAllEventsRegistration()
	if err !=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch registrations."})
		return
	}
	// context.JSON(http.StatusOK, events)
	context.JSON(http.StatusOK, gin.H{"data": data})
}




func cancelRegistration(context *gin.Context){
	// eventId:= context.Param("id")
	eventId, err:= strconv.ParseInt(context.Param("id"), 10, 64)
	userId:=context.GetInt64("userId") //userId from the token, was passed through the authorization


	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data"})
		return
	}

	reg, err := models.GetRegistrationByID(eventId)

	if err !=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fecth registration data"})
		return
	}

	if reg.UserId != userId{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "You are not authorized to delete this registration"})
		return
	}

	err = reg.CancelReg(userId)
	// err=reg.DeleteReg()
	if err !=nil{
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete registration", "error": err})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "registration data deleted successfully"})
}