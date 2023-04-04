package models

import (
	"time"
)

type Comment struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Title       string `json:"blog_title"`
	Text        string `json:"comment_text"`
	BlogID      uint   `json:"blog_id"`
	Blog        Blog
	UserID      uint `json:"user_id"`
	User        User
	Date_create time.Time `json:"d_create"`
}
