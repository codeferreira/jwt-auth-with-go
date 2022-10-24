package main

import (
	"net/http"
	"os"

	"github.com/codeferreira/jwt-auth-with-go/controllers"
	"github.com/codeferreira/jwt-auth-with-go/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	router := gin.Default()

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)

	router.Run(os.Getenv("PORT"))
}
