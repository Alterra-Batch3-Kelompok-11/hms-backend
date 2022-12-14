package dto

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Speciality struct {
	ID        uint           `json:"id" form:"id" validate:"required"`
	CreatedAt time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	Name      string         `json:"name" form:"name" validate:"required"`
}

func (ctrl *Speciality) Validate() error {
	validate := validator.New()
	err := validate.Struct(ctrl)
	if err != nil {
		return err
	}

	return nil
}
