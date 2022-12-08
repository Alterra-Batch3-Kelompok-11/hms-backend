package models

import "gorm.io/gorm"

type Doctor struct {
	gorm.Model
	UserId        uint   `json:"user_id" form:"user_id"`
	SpecialityId  uint   `json:"speciality_id" form:"speciality_id"`
	LicenseNumber string `json:"license_number" form:"license_number"`
	ProfilePic    string `json:"profile_pic" form:"profile_pic"`
	User          User
	Speciality    Speciality
}
