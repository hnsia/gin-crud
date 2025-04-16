package main

import (
	"github.com/hnsia/gin-crud/initializers"
	"github.com/hnsia/gin-crud/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
