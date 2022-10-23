package controllers

import (
	"net/http"

	"github.com/codeferreira/jwt-auth-with-go/initializers"
	"github.com/codeferreira/jwt-auth-with-go/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(context *gin.Context) {
	var body struct {
		Name     string
		Email    string
		Password string
	}

	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to hash the password",
		})

		return
	}

	user := models.User{Name: body.Name, Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create the user",
		})

		return
	}

	context.JSON(http.StatusCreated, user)
}
