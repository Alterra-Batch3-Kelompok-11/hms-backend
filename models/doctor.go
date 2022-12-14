package models

import (
	"gorm.io/gorm"
	"time"
)

type Doctor struct {
	gorm.Model
	UserId        uint      `json:"user_id" form:"user_id"`
	SpecialityId  uint      `json:"speciality_id" form:"speciality_id"`
	LicenseNumber string    `json:"license_number" form:"license_number"`
	ProfilePic    string    `json:"profile_pic" form:"profile_pic"`
	BirthDate     time.Time `json:"birth_date" form:"birth_date"`
	Phone         string    `json:"phone" form:"phone"`
	MaritalStatus bool      `json:"marital_status" form:"address"`
	Email         string    `json:"email" form:"email"`
	User          User
	Speciality    Speciality
}
