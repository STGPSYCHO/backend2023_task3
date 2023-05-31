package controllers

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func UserValidation(c *gin.Context) {

	session := sessions.Default(c)
	user_id := session.Get("user_id")
	if user_id == nil {
		c.HTML(http.StatusForbidden, "login.html", nil)
		c.Abort()
		return
	}
	c.Next()
}

func SessionValidation(c *gin.Context) {

	session := sessions.Default(c)
	user_id := session.Get("user_id")
	if user_id != nil {
		c.HTML(http.StatusForbidden, "login.html", gin.H{"message": "Please logout first"})
		c.Abort()
		return
	}
	c.Next()
}

func GetUserId(c *gin.Context) uint {

	session := sessions.Default(c)
	user_id := session.Get("user_id")
	if user_id == nil {
		c.HTML(http.StatusForbidden, "login.html", gin.H{"message": "No userID provided"})
		c.Abort()
		return 0
	}
	return user_id.(uint)
}
