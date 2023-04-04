package main

import (
	"net/http"

	"github.com/STGPSYCHO/backend2023_task3/controllers"
	"github.com/STGPSYCHO/backend2023_task3/models"
	"github.com/gin-gonic/gin"
)

func main() {

	route := gin.Default()
	route.LoadHTMLGlob("templates/*")

	models.ConnectDB()

	route.GET("/posts/:id", controllers.GetPost)
	// Работа с блогами
	// route.GET("/users", controllers.GetAllUsers)
	// route.GET("/users-by-headers", controllers.GetAllUsersByHeaders)
	// route.POST("/users", controllers.CreateUser)

	// route.PATCH("/users/:id", controllers.UpdateUser)
	// route.DELETE("/users/:id", controllers.DeleteUser)

	// Базовый эндпоинт
	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	route.Run()
}
