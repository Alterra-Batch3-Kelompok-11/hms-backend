package dto

import (
	"gorm.io/gorm"
	"time"
)

type DoctorScheduleReq struct {
	DoctorId  uint   `json:"doctor_id"`
	DayInt    int    `json:"day_int" form:"day_int"`
	StartTime string `json:"start_time" form:"start_time"`
	EndTime   string `json:"end_time" form:"end_time"`
}

type DoctorScheduleRes struct {
	ID        uint           `json:"id" form:"id"`
	CreatedAt time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	DoctorId  uint           `json:"doctor_id"`
	DayInt    int            `json:"day_int" form:"day_int"`
	DayString string         `json:"day_string" form:"day_string"`
	StartTime string         `json:"start_time" form:"start_time"`
	EndTime   string         `json:"end_time" form:"end_time"`
}
