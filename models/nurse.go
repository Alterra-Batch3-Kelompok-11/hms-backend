package models

import "gorm.io/gorm"

type Nurse struct {
	gorm.Model
	UserId        uint
	LicenseNumber string `json:"license_number" form:"license_number"`
	User          User
}
