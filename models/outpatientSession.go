package models

import (
	"time"

	"gorm.io/gorm"
)

type OutpatientSession struct {
	gorm.Model
	DoctorId   uint      `json:"doctor_id" form:"doctor_id"`
	PatientId  uint      `json:"patient_id" form:"patient_id"`
	Schedule   time.Time `json:"schedule" form:"schedule"`
	Complaint  string    `json:"complaint" form:"complaint"`
	IsApproved int       `json:"is_approved" form:"is_approved"`
	IsFinish   bool      `json:"is_finish" form:"is_finish"`
	FinishedAt time.Time `json:"finished_at" form:"finished_at"`
	Doctor     Doctor
	Patient    Patient
}
