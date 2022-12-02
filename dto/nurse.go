package dto

import (
	"time"

	"gorm.io/gorm"
)

type NurseRes struct {
	UserID        uint           `json:"user_id" form:"user_id"`
	LicenseNumber string         `json:"license_number" form:"license_number"`
	CreatedAt     time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
}
