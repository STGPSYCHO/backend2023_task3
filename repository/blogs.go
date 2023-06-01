package repository

import (
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
type CategoriesInfo struct {
	Category_name string
	ID            uint
}

func GetBlogsInfo(c *gin.Context) BlogsInfo {

	var raws BlogsInfo

	query_blogs := "select b.id, b.title, b.content, u.first_name, c.category_name from blogs b join users u on u.ID = b.user_id left join categories c on b.category_id = c.ID where b.ID = ?"

	result := models.DB.Raw(query_blogs, c.Param("id")).Scan(&raws)
	if result.RowsAffected == 0 {
		return BlogsInfo{}
	} else if result.Error != nil {
		return BlogsInfo{}
	}

	return raws
}

func GetComments(c *gin.Context) (comm []models.Comment, msg string) {

	var comms []models.Comment

	query_comms := "select c.text, u.first_name from comments c left join users u on u.ID = c.user_id where c.blog_id = ?"

	result_2 := models.DB.Raw(query_comms, c.Param("id")).Scan(&comms)
	if result_2.RowsAffected == 0 {
		return []models.Comment{}, "еще не создавали комментариев"
	} else if result_2.Error != nil {
		return []models.Comment{}, ""
	}
	return comms, ""

}

func GetTags(c *gin.Context) (tags_arr []TagsInfo, msg string) {

	var tags []TagsInfo

	query_tags := "select t.name, t.id from blogs b join blog_tags bt on b.id = bt.blog_id join tags t on t.id = bt.tag_id where b.id = ?"

	result_3 := models.DB.Raw(query_tags, c.Param("id")).Scan(&tags)
	if result_3.RowsAffected == 0 {
		return []TagsInfo{}, "еще не создавали тегов"
	} else if result_3.Error != nil {
		return []TagsInfo{}, ""
	}
	return tags, ""
}

func GetTagsAdd(c *gin.Context) (tags_add []TagsInfo, msg string) {

	var tagsAdd []TagsInfo

	query_tagsAdd := "select name, id from tags"

	result_4 := models.DB.Raw(query_tagsAdd).Scan(&tagsAdd)
	if result_4.RowsAffected == 0 {
		return []TagsInfo{}, "еще не создавали тегов"
	} else if result_4.Error != nil {
		return []TagsInfo{}, ""
	}
	return tagsAdd, ""
}

func GetBlogs(c *gin.Context) (blogs []BlogsInfo, msg string) {

	var raws []BlogsInfo

	query := "select b.id, b.title, b.content, u.first_name, c.category_name from blogs b join users u on u.ID = b.user_id left join categories c on b.category_id = c.ID where b.deleted_at is null"

	result := models.DB.Raw(query).Scan(&raws)
	if result.RowsAffected == 0 {
		return []BlogsInfo{}, "не нашли такие блоги"
	} else if result.Error != nil {
		return []BlogsInfo{}, ""
	}
	return raws, ""
}

func GetCategories(c *gin.Context) (categories []CategoriesInfo, msg string) {

	var categoriesAdd []CategoriesInfo

	query_categoriesAdd := "select category_name, id from categories"

	result_5 := models.DB.Raw(query_categoriesAdd).Scan(&categoriesAdd)
	if result_5.RowsAffected == 0 {
		return []CategoriesInfo{}, "еще не создавали категорий"
	} else if result_5.Error != nil {
		return []CategoriesInfo{}, ""
	}
	return categoriesAdd, ""
}
