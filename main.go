package main

import (
	"net/http"

	docs "github.com/STGPSYCHO/backend2023_task3/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/STGPSYCHO/backend2023_task3/controllers"
	"github.com/STGPSYCHO/backend2023_task3/models"
	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/gin-gonic/gin"
)

// @title Todo App API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8000
// @BasePath /

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {

	route := gin.Default()

	docs.SwaggerInfo.BasePath = "/"
	route.LoadHTMLGlob("templates/*")

	models.ConnectDB()

	store := gormsessions.NewStore(models.DB, true, []byte("secret"))
	route.Use(sessions.Sessions("mysession", store))

	// CRUD блогов
	api := route.Group("/api", controllers.UserValidation)
	{
		api.GET("/blogs/:id", controllers.GetBlog)
		api.GET("/blogs", controllers.GetBlogs)
		api.POST("/create-blog", controllers.CreateBlog)
		api.POST("/remove-blog/:id", controllers.DeleteBlog)
		api.PATCH("/blogs/:id", controllers.UpdateBlog)

		api.POST("/create-category", controllers.CreateCategory)
		api.POST("/remove-category/:id", controllers.DeleteCategory)
		api.PATCH("/category/:id", controllers.UpdateCategory)

		api.POST("/create-comment", controllers.CreateComment)

		api.POST("/tag", controllers.AssosiateTag)
		api.POST("/create-tag", controllers.CreateTag)
		api.POST("/remove-tag", controllers.RemoveTag)
	}

	// Авторизация, регистрация
	route.GET("/login", controllers.LoginPage)
	route.POST("/register", controllers.Register)
	route.POST("/login-verification", controllers.Login)

	// Базовый эндпоинт
	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	// Swagger
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	route.Run()
}
