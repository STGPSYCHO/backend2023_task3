package controllers

import (
	"net/http"
	"strconv"

	"github.com/STGPSYCHO/backend2023_task3/models"
	"github.com/gin-gonic/gin"
)

// POST /create-comment
// Создаем Коммент
func CreateComment(c *gin.Context) {
	var comment models.Comment

	userID := GetUserId(c)
	comment.UserID = userID

	blog_id, ok := c.GetPostForm("id")
	v, _ := strconv.ParseUint(blog_id, 10, 32)
	comment.BlogID = uint(v)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не передали Id"})
		return
	}

	text, ok := c.GetPostForm("comment")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не передали текст"})
		return
	}
	comment.Text = text

	result := models.DB.Create(&comment)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/api/blogs/"+blog_id)
}
