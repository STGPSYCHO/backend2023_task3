package models

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	Title  string `json:"blog_title"`
	Text   string `json:"comment_text"`
	BlogID uint   `json:"blog_id"`
	UserID uint   `json:"user_id"`
}
