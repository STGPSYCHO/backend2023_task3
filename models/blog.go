package models

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title      string `json:"blog_title"`
	UserID     uint   `json:"user_id" binding:"required"`
	Content    string `json:"blog_content"`
	CategoryID uint   `json:"category_id"`
	Tags       []*Tag `gorm:"many2many:blog_tags;"`
}
