package initializers

import "github.com/codeferreira/jwt-auth-with-go/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
