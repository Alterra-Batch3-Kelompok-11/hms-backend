package dto

import (
	"gorm.io/gorm"
	"time"
)

type Speciality struct {
	ID        uint           `json:"id" form:"id"`
	CreatedAt time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	Name      string         `json:"name" form:"name"`
}
