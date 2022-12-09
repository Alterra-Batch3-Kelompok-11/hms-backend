package dto

import (
	"gorm.io/gorm"
	"time"
)

type UserReq struct {
	Name          string `json:"name" form:"name"`
	LicenseNumber string `json:"license_number" form:"license_number"`
	Username      string `json:"username" form:"username"`
	Password      string `json:"password" form:"password"`
	RoleID        uint   `json:"role_id" form:"role_id"`
	SpecialityID  uint   `json:"speciality_id" form:"speciality_id"`
}

type UserRes struct {
	ID            uint           `json:"id" form:"id"`
	CreatedAt     time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	Name          string         `json:"name" form:"name"`
	Username      string         `json:"username" form:"username"`
	RoleID        uint           `json:"role_id" form:"role_id"`
	Role          string         `json:"role" form:"role"`
	Password      string         `json:"password" form:"password"`
	LicenseNumber string         `json:"license_number" form:"license_number"`
}

type LoginRes struct {
	ID            uint        `json:"user_id" form:"user_id"`
	Name          string      `json:"name" form:"name"`
	Username      string      `json:"username" form:"username"`
	RoleID        uint        `json:"role_id" form:"role_id"`
	Token         string      `json:"token" form:"token"`
	LicenseNumber string      `json:"license_number" form:"license_number"`
	DoctorID      interface{} `json:"doctor_id" form:"doctor_id"`
	NurseID       interface{} `json:"nurse_id" form:"nurse_id"`
}
