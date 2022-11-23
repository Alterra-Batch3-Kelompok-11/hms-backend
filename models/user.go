package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	RoleId   uint   `json:"role_id" form:"role_id"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Name     string `json:"name" form:"name"`
	Role     Role
}
