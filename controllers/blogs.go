package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/STGPSYCHO/backend2023_task3/models"
	"github.com/STGPSYCHO/backend2023_task3/repository"
	"github.com/gin-gonic/gin"
)

// Blogs example
//
//	@Summary		GetBlog
//	@Tags			api
//	@Description	get blog by id
//	@Success		200		{integer}	integer	1
//	@Failure		400,404	{object}	errorResponse
//	@Failure		500		{object}	errorResponse
//	@Failure		default	{object}	errorResponse
//	@Router			/api/blogs [get]
func GetBlog(c *gin.Context) {

	var comms []models.Comment
	var raws repository.BlogsInfo
	var msg string
	var tags []repository.TagsInfo
	var tagsAdd []repository.TagsInfo

	raws = repository.GetBlogsInfo(c)
	comms, msg = repository.GetComments(c)
	tags, msg = repository.GetTags(c)
	tagsAdd, msg = repository.GetTagsAdd(c)

	c.HTML(
		http.StatusOK,
		"blog.html",
		gin.H{
			"Blog":     raws,
			"Comments": comms,
			"Tags":     tags,
			"TagsAdd":  tagsAdd,
			"Message":  msg,
		},
	)
}

// GET /blogs
// Получаем все блоги
func GetBlogs(c *gin.Context) {

	var raws []repository.BlogsInfo
	var categoriesAdd []repository.CategoriesInfo
	var msg string

	categoriesAdd, msg = repository.GetCategories(c)
	raws, msg = repository.GetBlogs(c)

	c.HTML(
		http.StatusOK,
		"blogs.html",
		gin.H{
			"Blogs":         raws,
			"CategoriesAdd": categoriesAdd,
			"msg":           msg,
		},
	)
}

// POST /create-blog
// Создаем блог
func CreateBlog(c *gin.Context) {

	userID := GetUserId(c)

	var blog models.Blog
	blog.UserID = userID
	title, ok := c.GetPostForm("title")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не передали title"})
		return
	}
	blog.Title = title

	content, ok := c.GetPostForm("content")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не передали content"})
		return
	}
	blog.Content = content

	category, ok := c.GetPostForm("category")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "не передали categoryId"})
		return
	}
	v, _ := strconv.ParseUint(category, 10, 32)
	blog.CategoryID = uint(v)

	result := models.DB.Create(&blog)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog post"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, "/api/blogs")
}

// POST /remove-blog/:id
// Удаляем блог
func DeleteBlog(c *gin.Context) {

	var blog models.Blog
	userID := GetUserId(c)

	result := models.DB.Where("id = ?", c.Param("id")).First(&blog)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog post not found"})
		return
	}

	fmt.Print(blog.UserID)
	fmt.Print(userID)
	if blog.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Blog post can only be deleted by is's owner"})
		return
	}

	models.DB.Delete(&blog)

	// c.JSON(http.StatusOK, gin.H{"message": "Blog post deleted successfully"})
	c.Redirect(301, "/api/blogs")
}

// PATCH /blogs/:id
// Обновляем существующий блог
func UpdateBlog(c *gin.Context) {

	// Get the blog ID from the request URL parameters
	blogID := c.Param("id")
	userID := GetUserId(c)

	// Retrieve the blog from the database
	var blog models.Blog
	result := models.DB.First(&blog, blogID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	// Bind the updated blog data from the request body
	var updatedBlog models.Blog
	if err := c.ShouldBindJSON(&updatedBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update the fields of the blog
	blog.Title = updatedBlog.Title
	blog.Content = updatedBlog.Content

	if blog.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Blog post can only be deleted by is's owner"})
		return
	}

	// Save the updated blog to the database
	result = models.DB.Save(&blog)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update blog"})
		return
	}

	c.JSON(http.StatusOK, blog)
}
