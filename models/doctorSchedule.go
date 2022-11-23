package models

import "gorm.io/gorm"

type DoctorSchedule struct {
	gorm.Model
	DoctorId  uint   `json:"doctor_id" form:"doctor_id"`
	Day       int    `json:"day" form:"day"`
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
	Doctor    Doctor
}
