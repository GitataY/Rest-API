package routes

import (
	"example/com/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User

	// Binding the JSON to the user object
	if err := context.ShouldBindJSON(&user); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request"})
		return
	}
	
	// Saving the user object
	if err := user.Save(); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not save user"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}
