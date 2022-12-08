package models

import "gorm.io/gorm"

type Nurse struct {
	gorm.Model
	UserId        uint
	LicenseNumber string `json:"license_number" form:"license_number"`
	ProfilePic    string `json:"profile_pic" form:"profile_pic"`
	User          User
}
