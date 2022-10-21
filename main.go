package main

import (
	"fmt"

	"github.com/codeferreira/jwt-auth-with-go/initializers"
)

func init() {
	initializers.LoadEnvVariables()
}

func main() {
	fmt.Println("Hello World!")
}