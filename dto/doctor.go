package dto

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type DoctorReq struct {
	Name          string `json:"name" form:"name" validate:"required"`
	LicenseNumber string `json:"license_number" form:"license_number" validate:"required"`
	Password      string `json:"password" form:"password" validate:"required"`
	SpecialityID  uint   `json:"speciality_id" form:"speciality_id" validate:"required"`
	ProfilePic    string `json:"profile_pic" form:"profile_pic"`
	BirthDate     string `json:"birth_date" form:"birth_date"`
	Phone         string `json:"phone" form:"phone" validate:"required"`
	MaritalStatus bool   `json:"marital_status" form:"address"`
	Email         string `json:"email" form:"email"`
}

type DoctorRes struct {
	ID                  uint                       `json:"id" form:"id" validate:"required"`
	CreatedAt           time.Time                  `json:"created_at" form:"created_at" validate:"required"`
	UpdatedAt           time.Time                  `json:"updated_at" form:"updated_at" validate:"required"`
	DeletedAt           gorm.DeletedAt             `gorm:"index" json:"deleted_at" form:"deleted_at"`
	Name                string                     `json:"name" validate:"required"`
	SpecialityId        uint                       `json:"speciality_id" form:"speciality_id" validate:"required"`
	LicenseNumber       string                     `json:"license_number" form:"license_number" validate:"required"`
	SpecialityName      string                     `json:"speciality_name" form:"speciality_name"`
	ProfilePic          string                     `json:"profile_pic" form:"profile_pic"`
	BirthDate           time.Time                  `json:"birth_date" form:"birth_date"`
	BirthDateString     string                     `json:"birth_date_string" form:"birth_date_string"`
	BirthDateStringIndo string                     `json:"birth_date_string_indo" form:"birth_date_string_indo"`
	Phone               string                     `json:"phone" form:"phone"`
	MaritalStatus       bool                       `json:"marital_status" form:"address"`
	Email               string                     `json:"email" form:"email"`
	DoctorSchedules     []DoctorProfileScheduleRes `json:"doctor_schedules" form:"doctor_schedules"`
}

func (ctrl *DoctorReq) Validate() error {
	validate := validator.New()
	err := validate.Struct(ctrl)
	if err != nil {
		return err
	}

	return nil
}

func (ctrl *DoctorRes) Validate() error {
	validate := validator.New()
	err := validate.Struct(ctrl)
	if err != nil {
		return err
	}

	return nil
}

type TodayDoctorRes struct {
	Name           string `json:"name"`
	LicenseNumber  string `json:"license_number" form:"license_number"`
	SpecialityName string `json:"speciality_name" form:"speciality_name"`
	ProfilePic     string `json:"profile_pic" form:"profile_pic"`
	DayInt         int    `json:"day_int" form:"day_int"`
	DayString      string `json:"day_string" form:"day_string"`
	StartTime      string `json:"start_time" form:"start_time"`
	EndTime        string `json:"end_time" form:"end_time"`
}
