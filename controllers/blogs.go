package controllers

import (
	"fmt"
	"net/http"

	"github.com/STGPSYCHO/backend2023_task3/models"
	"github.com/gin-gonic/gin"
)

type BlogsInfo struct {
	ID            uint
	Title         string
	Content       string
	First_name    string
	Category_name string
}

type TagsInfo struct {
	Name string
	ID   uint
}

// Blogs example
// @Summary GetBlog
// @Security ApiKeyAuth
// @Tags api
// @Description get blog by id
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /api/blogs [get]
func GetBlog(c *gin.Context) {

	var comms []models.Comment
	var raws BlogsInfo
	var msg string
	var tags []TagsInfo
	var tagsAdd []TagsInfo

	query_blogs := "select b.id, b.title, b.content, u.first_name, c.category_name from blogs b join users u on u.ID = b.user_id left join categories c on b.category_id = c.ID where b.ID = ?"
	query_comms := "select c.text, u.first_name from comments c left join users u on u.ID = c.user_id where c.blog_id = ?"
	query_tags := "select t.name, t.id from blogs b join blog_tags bt on b.id = bt.blog_id join tags t on t.id = bt.tag_id where b.id = ?"
	query_tagsAdd := "select name, id from tags"

	result := models.DB.Raw(query_blogs, c.Param("id")).Scan(&raws)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": result.Error})
		return
	}

	result_2 := models.DB.Raw(query_comms, c.Param("id")).Scan(&comms)
	if result_2.RowsAffected == 0 {
		msg = "еще не создавали комментариев"
	} else if result_2.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": result_2.Error})
		return
	}

	result_3 := models.DB.Raw(query_tags, c.Param("id")).Scan(&tags)
	if result_2.RowsAffected == 0 {
		msg = "еще не создавали тегов"
	} else if result_2.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": result_3.Error})
		return
	}

	result_4 := models.DB.Raw(query_tagsAdd).Scan(&tagsAdd)
	if result_2.RowsAffected == 0 {
		msg = "еще не создавали тегов"
	} else if result_2.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": result_4.Error})
		return
	}

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

	var raws []BlogsInfo
	query := "select b.id, b.title, b.content, u.first_name, c.category_name from blogs b join users u on u.ID = b.user_id left join categories c on b.category_id = c.ID where b.deleted_at is null"

	result := models.DB.Raw(query).Scan(&raws)
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": result.Error})
		return
	} else if result.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": result.Error})
		return
	}

	c.HTML(
		http.StatusOK,
		"blogs.html",
		gin.H{
			"Blogs": raws,
		},
	)
}

// POST /create-blog
// Создаем блог
func CreateBlog(c *gin.Context) {

	userID := GetUserId(c)

	var blog models.Blog
	blog.UserID = userID

	fmt.Print(blog.UserID)
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := models.DB.Create(&blog)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create blog post"})
		return
	}

	c.Redirect(301, "/api/blogs")
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
