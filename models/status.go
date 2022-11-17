package models

import "gorm.io/gorm"

type Status struct {
	gorm.Model
	Name string `json:"name" form:"name"`
}
