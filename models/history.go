package models

import (
	"time"

	"gorm.io/gorm"
)

type History struct {
	gorm.Model
	PatientId  uint      `json:"patient_id" form:"patient_id"`
	Diagnose   string    `json:"diagnose" form:"diagnose"`
	FinishedAt time.Time `json:"finished_at" form:"finished_at"`
	Patient    Patient
}
