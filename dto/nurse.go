package dto

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type NurseRes struct {
	ID            uint           `json:"id" form:"id"`
	UserID        uint           `json:"user_id" form:"user_id" validate:"required"`
	LicenseNumber string         `json:"license_number" form:"license_number" validate:"required"`
	CreatedAt     time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
}

func (ctrl *NurseRes) Validate() error {
	validate := validator.New()
	err := validate.Struct(ctrl)
	if err != nil {
		return err
	}

	return nil
}
