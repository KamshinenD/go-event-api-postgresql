package routes

import (
	"github.com/gin-gonic/gin"
	"events.com/rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine){
	server.GET("/events", getEvents)    //GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent)
	server.POST("/events", middlewares.Authenticate, createEvent) //GET, POST, PUT, PATCH, DELETE
	server.PUT("/events/:id", middlewares.Authenticate, updateEvent)
	server.DELETE("/events/:id", middlewares.Authenticate, deleteEvent)
	server.POST("/signup", signUp)
	server.GET("/users", getAllUsers)
	server.POST("/login", login)
	server.POST("/events/register/:id", middlewares.Authenticate, EventRegistration)
	server.DELETE("/events/register/:id", middlewares.Authenticate, cancelRegistration)
	server.GET("/events/registrations", middlewares.Authenticate, getAllRegistrations)
}