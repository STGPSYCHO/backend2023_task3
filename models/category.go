package models

import (
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Category_name string `json:"category_name"`
	Blog          []Blog
}
