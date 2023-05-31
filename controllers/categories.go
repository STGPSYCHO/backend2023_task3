package controllers

import (
	"net/http"

	"github.com/STGPSYCHO/backend2023_task3/models"
	"github.com/gin-gonic/gin"
)

// POST /create-category
// Создаем категорию
func CreateCategory(c *gin.Context) {

	var category models.Category

	category_name, ok := c.GetPostForm("category_name")

	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не передали Id"})
		return
	}

	category.Category_name = category_name

	result := models.DB.Create(&category)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/api/blogs")
}

// POST /remove-category/:id
// Удаляем категорию
func DeleteCategory(c *gin.Context) {

	var category models.Category
	result := models.DB.Where("id = ?", c.Param("id")).First(&category)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	models.DB.Delete(&category)

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

// PATCH /category/:id
// Обновляем существующую категорию
func UpdateCategory(c *gin.Context) {

	// Get the blog ID from the request URL parameters
	categotyID := c.Param("id")

	// Retrieve the blog from the database
	var category models.Category
	result := models.DB.First(&category, categotyID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	// Bind the updated blog data from the request body
	var updatedCategory models.Category
	if err := c.ShouldBindJSON(&updatedCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the fields of the blog
	category.Category_name = updatedCategory.Category_name

	// Save the updated blog to the database
	result = models.DB.Save(&category)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}

	c.JSON(http.StatusOK, category)
}
