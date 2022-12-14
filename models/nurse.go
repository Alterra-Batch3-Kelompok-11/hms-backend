package models

import "gorm.io/gorm"

type Nurse struct {
	gorm.Model
	UserId        uint
	LicenseNumber string `json:"license_number" form:"license_number"`
	ProfilePic    string `json:"profile_pic" form:"profile_pic"`
	BirthDate     string `json:"birth_date" form:"birth_date"`
	Phone         string `json:"phone" form:"phone"`
	MaritalStatus bool   `json:"marital_status" form:"address"`
	Email         string `json:"email" form:"email"`
	Nira          string `json:"nira" form:"nira"`
	SIP           string `json:"sip" form:"sip"`
	User          User
}
