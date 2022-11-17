package models

import "gorm.io/gorm"

type Religion struct {
	gorm.Model
	Name string `json:"name" form:"name"`
}
