package models

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name  string  `json:"tag_name"`
	Blogs []*Blog `gorm:"many2many:blog_tags;"`
}
