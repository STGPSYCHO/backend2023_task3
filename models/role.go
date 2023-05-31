package models

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	Name  string  `json:"role_name"`
	Users []*User `gorm:"many2many:user_roles;"`
}
