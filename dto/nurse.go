package dto

import (
	"github.com/go-playground/validator/v10"
	"time"

	"gorm.io/gorm"
)

type NurseRes struct {
	ID                  uint           `json:"id" form:"id"`
	UserID              uint           `json:"user_id" form:"user_id" validate:"required"`
	Name                string         `json:"name" validate:"required"`
	SpecialityId        uint           `json:"speciality_id" form:"speciality_id" validate:"required"`
	LicenseNumber       string         `json:"license_number" form:"license_number" validate:"required"`
	SpecialityName      string         `json:"speciality_name" form:"speciality_name"`
	BirthDate           time.Time      `json:"birth_date" form:"birth_date"`
	BirthDateString     string         `json:"birth_date_string" form:"birth_date_string"`
	BirthDateStringIndo string         `json:"birth_date_string_indo" form:"birth_date_string_indo"`
	ProfilePic          string         `json:"profile_pic" form:"profile_pic"`
	Phone               string         `json:"phone" form:"phone"`
	MaritalStatus       bool           `json:"marital_status" form:"address"`
	Email               string         `json:"email" form:"email"`
	Nira                string         `json:"nira" form:"nira"`
	SIP                 string         `json:"sip" form:"sip"`
	CreatedAt           time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
}

func (ctrl *NurseRes) Validate() error {
	validate := validator.New()
	err := validate.Struct(ctrl)
	if err != nil {
		return err
	}

	return nil
}
