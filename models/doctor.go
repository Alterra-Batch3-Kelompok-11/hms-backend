package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	UserId        uint   `json:"user_id" form:"user_id"`
	SpecialityId  uint   `json:"speciality_id" form:"speciality_id"`
	LicenseNumber string `json:"license_number" form:"license_number"`
	User          User
	Speciality    Speciality
}
