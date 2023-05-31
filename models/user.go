package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	First_name string  `json:"first_name"`
	Last_name  string  `json:"last_name"`
	Username   string  `json:"username" gorm:"unique"`
	Password   string  `json:"password"`
	Blogs      []Blog  `json:"blogs" gorm:"foreignkey:UserID"`
	Roles      []*Role `gorm:"many2many:user_roles;"`
}
