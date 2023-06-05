package controllers

import (
	"fmt"
	"net/http"

	"github.com/STGPSYCHO/backend2023_task3/models"
	"github.com/gin-contrib/sessions"

	// "golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

type Signin struct {
	Login    string
	Password string
}

// Register User
//
//		@Summary		Register user
//		@Description	Register my user
//		@Accept json
//		@Produce json
//	 	@Param 			login body models.User true "Account ID"
//	 	@Param			password body string true "Password"
//		@Success		200		{string}	string
//		@Failure		400,404	{object}	errorResponse
//		@Failure		500		{object}	errorResponse
//		@Failure		default	{object}	errorResponse
//		@Router			/register [post]
func RegisterSwag(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := models.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {

	var user models.User
	username, ok := c.GetPostForm("username")
	if !ok {
		c.HTML(http.StatusForbidden, "login.html", nil)
		c.Abort()
		return
	}
	password, ok := c.GetPostForm("password")
	if !ok {
		c.HTML(http.StatusForbidden, "login.html", nil)
		c.Abort()
		return
	}

	result := models.DB.Where("username = ?", username).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	// bcrypt.CompareHashAndPassword([]byte(loginData.Password) ,[]byte(user.Password))
	if user.Password != password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}
	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	if err := session.Save(); err != nil {
		c.HTML(http.StatusForbidden, "login.html", nil)
		c.Abort()
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/api/blogs")
}

// Login logic
//
//		@Summary		Login user
//		@Description	Login my user
//		@Accept json
//		@Produce json
//	 	@Param 			login body models.User true "Account ID"
//		@Success		200		{string}	string
//		@Failure		400,404	{object}	errorResponse
//		@Failure		500		{object}	errorResponse
//		@Failure		default	{object}	errorResponse
//		@Router			/login-verify [post]
func LoginSwag(c *gin.Context) {

	var user models.User

	fmt.Print(c.Request.Body)

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := models.DB.Where("username = ? and password = ?", user.Username, user.Password).First(&user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found or invalid password"})
		return
	}
	// bcrypt.CompareHashAndPassword([]byte(loginData.Password) ,[]byte(user.Password))

	session := sessions.Default(c)
	session.Set("user_id", user.ID)
	if err := session.Save(); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "Cannot save session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in"})
}

// Get /login
// Отрисовка формы авторизации
func LoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{},
	)
}
