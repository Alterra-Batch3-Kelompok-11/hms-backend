package models

import "gorm.io/gorm"

type DoctorNurse struct {
	gorm.Model
	DoctorId uint `json:"doctor_id" form:"doctor_id"`
	NurseId  uint `json:"nurse_id" form:"nurse_id"`
	Doctor   Doctor
	Nurse    Nurse
}
