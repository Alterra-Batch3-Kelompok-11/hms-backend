package dto

import (
	"gorm.io/gorm"
	"time"
)

type DoctorRes struct {
	ID              uint                       `json:"id" form:"id"`
	CreatedAt       time.Time                  `json:"created_at" form:"created_at"`
	UpdatedAt       time.Time                  `json:"updated_at" form:"updated_at"`
	DeletedAt       gorm.DeletedAt             `gorm:"index" json:"deleted_at" form:"deleted_at"`
	Name            string                     `json:"name"`
	SpecialityId    uint                       `json:"speciality_id" form:"speciality_id"`
	LicenseNumber   string                     `json:"license_number" form:"license_number"`
	SpecialityName  string                     `json:"speciality_name" form:"speciality_name"`
	ProfilePic      string                     `json:"profile_pic" form:"profile_pic"`
	BirthDate       time.Time                  `json:"birth_date" form:"birth_date"`
	Phone           string                     `json:"phone" form:"phone"`
	MaritalStatus   bool                       `json:"marital_status" form:"address"`
	Email           string                     `json:"email" form:"email"`
	DoctorSchedules []DoctorProfileScheduleRes `json:"doctor_schedules" form:"doctor_schedules"`
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
