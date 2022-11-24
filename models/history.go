package models

import (
	"time"

	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	DoctorId    uint      `json:"doctor_id" form:"doctor_id"`
	PatientId   uint      `json:"patient_id" form:"patient_id"`
	Diagnose    string    `json:"diagnose" form:"diagnose"`
	Description string    `json:"description" form:"description"`
	Medicine    string    `json:"medicine" form:"medicine"`
	Allergy     string    `json:"allergy" form:"allergy"`
	FinishedAt  time.Time `json:"finished_at" form:"finished_at"`
	Doctor      Doctor
	Patient     Patient
}
