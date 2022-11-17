package models

import (
	"gorm.io/gorm"
	"time"
)

type Treatment struct {
	gorm.Model
	SessionId     uint
	TreatmentDate time.Time
	Description   string
}
