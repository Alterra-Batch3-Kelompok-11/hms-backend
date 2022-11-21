package models

import (
	"gorm.io/gorm"
)

type Treatment struct {
	gorm.Model
	SessionId   uint
	Diagnose    string
	Description string
	Medicines   string
}
