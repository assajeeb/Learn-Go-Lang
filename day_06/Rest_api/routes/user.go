package routes

import (
	"net/http"

	"example.com/rest-api/models"
	"github.com/gin-gonic/gin"
)

func SignUpUser(context *gin.Context) {

	var NewUser models.User

	err := context.ShouldBindJSON(&NewUser)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to bind Account", "error": err})
		return
	}
	err = NewUser.Save()

	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message": "Unable to Create Account", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "data": NewUser})

}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Unable to bind User"})
		return
	}
	err = user.ValidateUser()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized user login"})
		return
	}

	token, err := user.GenerateToken()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Session creation fail", "error": err})
		return
	}

	context.JSON(http.StatusAccepted, gin.H{"message": "User Login Successful!", "token": token})

}
