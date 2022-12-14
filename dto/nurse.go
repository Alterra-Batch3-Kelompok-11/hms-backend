package dto

import (
	"time"

	"gorm.io/gorm"
)

type NurseRes struct {
	ID            uint           `json:"id" form:"id"`
	UserID        uint           `json:"user_id" form:"user_id"`
	LicenseNumber string         `json:"license_number" form:"license_number"`
	BirthDate     string         `json:"birth_date" form:"birth_date"`
	Phone         string         `json:"phone" form:"phone"`
	MaritalStatus bool           `json:"marital_status" form:"address"`
	Email         string         `json:"email" form:"email"`
	Nira          string         `json:"nira" form:"nira"`
	SIP           string         `json:"sip" form:"sip"`
	CreatedAt     time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
}
