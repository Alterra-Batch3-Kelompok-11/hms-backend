package models

import "gorm.io/gorm"

type OutpatientSession struct {
	gorm.Model
	DoctorId  uint   `json:"doctor_id" form:"doctor_id"`
	PatientId uint   `json:"patient_id" form:"patient_id"`
	Complaint string `json:"complaint" form:"complaint"`
	Diagnose  string `json:"diagnose" form:"diagnose"`
	Doctor    Doctor
	Patient   Patient
}
