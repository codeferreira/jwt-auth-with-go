package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/codeferreira/jwt-auth-with-go/initializers"
	"github.com/codeferreira/jwt-auth-with-go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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


func Login(context *gin.Context) {
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

	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 3).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to authenticate",
		})

		return
	}

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie("Authorization", tokenString, 3600 * 24 * 3, "", "", false, true)

	context.JSON(http.StatusOK, gin.H{
	})
}