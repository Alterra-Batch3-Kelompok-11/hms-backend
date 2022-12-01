package dto

import (
	"gorm.io/gorm"
	"time"
)

type DoctorRes struct {
	ID              uint                `json:"id" form:"id"`
	CreatedAt       time.Time           `json:"created_at" form:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at" form:"updated_at"`
	DeletedAt       gorm.DeletedAt      `gorm:"index" json:"deleted_at" form:"deleted_at"`
	Name            string              `json:"name"`
	SpecialityId    uint                `json:"speciality_id" form:"speciality_id"`
	LicenseNumber   string              `json:"license_number" form:"license_number"`
	SpecialityName  string              `json:"speciality_name" form:"speciality_name"`
	DoctorSchedules []DoctorScheduleRes `json:"doctor_schedules" form:"doctor_schedules"`
}
