package models

import (
	"gorm.io/gorm"
)

type Treatment struct {
	gorm.Model
	SessionId   uint   `json:"session_id" form:"session_id"`
	Diagnose    string `json:"diagnose" form:"diagnose"`
	Description string `json:"description" form:"description"`
	Medicine    string `json:"medicine" form:"medicine"`
	Allergy     string `json:"allergy" form:"allergy"`
}
