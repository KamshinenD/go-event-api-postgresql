package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"events.com/rest-api/utils"
)


func Authenticate(context *gin.Context){
	token:=context.Request.Header.Get("Authorization")

	if token==""{
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Not Authorized"})
		return
	}

	userId, err:=utils.VerifyToken(token)

	if err !=nil{
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	context.Set("userId", userId)
	context.Next()
}

