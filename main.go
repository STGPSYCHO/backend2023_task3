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

//	@title			CRUD Application
//	@version		1.2
//	@description	API Server for CRUD Application

//	@host		localhost:8080
//	@BasePath	/

func main() {

	route := gin.Default()

	docs.SwaggerInfo.BasePath = "/"

	route.LoadHTMLGlob("templates/*")

	models.ConnectDB()

	store := gormsessions.NewStore(models.DB, true, []byte("secret"))
	route.Use(sessions.Sessions("mysession", store))

	// CRUD блогов для UI
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

	// CRUD блогов для UI
	swag := route.Group("/swag", controllers.UserValidation)
	{
		swag.GET("/blogs/:id", controllers.GetBlog)
		swag.GET("/blogs", controllers.GetBlogs)
		swag.POST("/create-blog", controllers.CreateBlog)
		swag.POST("/remove-blog/:id", controllers.DeleteBlog)
		swag.PATCH("/blogs/:id", controllers.UpdateBlog)

		swag.POST("/create-category", controllers.CreateCategory)
		swag.POST("/remove-category/:id", controllers.DeleteCategory)
		swag.PATCH("/category/:id", controllers.UpdateCategory)

		swag.POST("/create-comment", controllers.CreateComment)

		swag.POST("/tag", controllers.AssosiateTag)
		swag.POST("/create-tag", controllers.CreateTag)
		swag.POST("/remove-tag", controllers.RemoveTag)
	}

	// Авторизация, регистрация Swag
	route.POST("/login-verify", controllers.LoginSwag)
	route.POST("/register", controllers.RegisterSwag)

	// Авторизация, регистрация
	route.GET("/login", controllers.LoginPage)
	route.POST("/login-verification", controllers.Login)

	// Базовый эндпоинт
	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Hello World!"})
	})

	// Swagger
	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	route.Run()
}

// swag init --parseDependency --parseInternal
