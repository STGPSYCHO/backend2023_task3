package controllers

import (
	"net/http"

	"github.com/STGPSYCHO/backend2023_task3/models"
	"github.com/gin-gonic/gin"
)

// POST /create-tag
// Создаем блог
func CreateTag(c *gin.Context) {

	var tag models.Tag

	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := models.DB.Create(&tag)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create tag"})
		return
	}

	c.JSON(http.StatusOK, tag)
}

// POST /remove-tag/:id
// Удаляем блог
func DeleteTag(c *gin.Context) {
	var tag models.Tag

	result := models.DB.Where("id = ?", c.Param("id")).First(&tag)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	models.DB.Delete(&tag)

	c.JSON(http.StatusOK, gin.H{"message": "Tag deleted successfully"})
}

// POST /tag
// Коннект блога к тегу
func AssosiateTag(c *gin.Context) {

	query := "insert into blog_tags values (?,?)"

	blog_id, ok := c.GetPostForm("blog")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "form blog not found"})
	}

	tag_id, ok := c.GetPostForm("tag")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "form tag not found"})
	}

	result := models.DB.Exec(query, tag_id, blog_id)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Инсерт не произошел"})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": result.Error})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/api/blogs/"+blog_id)
}
