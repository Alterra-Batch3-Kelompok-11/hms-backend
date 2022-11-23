package models

import "gorm.io/gorm"

type Speciality struct {
	gorm.Model
	Name string `json:"name" form:"name"`
}
