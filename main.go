package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hnsia/gin-crud/controllers"
	"github.com/hnsia/gin-crud/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	router.POST("/posts", controllers.PostsCreate)
	router.PUT("/posts/:id", controllers.PostsUpdate)

	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)

	router.DELETE("/posts/:id", controllers.PostsDelete)

	router.Run() // listen and serve on 0.0.0.0:8080 or chosen PORT number
}
