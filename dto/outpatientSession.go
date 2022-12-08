package dto

import (
	"gorm.io/gorm"
	"time"
)

type OutpatientSessionReq struct {
	PatientId    uint   `json:"patient_id" form:"patient_id"`
	DoctorId     uint   `json:"doctor_id" form:"doctor_id"`
	Complaint    string `json:"complaint" form:"complaint"`
	ScheduleDate string `json:"schedule_date" form:"schedule_date"`
	ScheduleTime string `json:"schedule_time" form:"schedule_time"`
}

type OutpatientSessionRes struct {
	ID               uint           `json:"id" form:"id"`
	CreatedAt        time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt        time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt        gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
	DoctorId         uint           `json:"doctor_id" form:"doctor_id"`
	PatientId        uint           `json:"patient_id" form:"patient_id"`
	Schedule         time.Time      `json:"schedule" form:"schedule"`
	Complaint        string         `json:"complaint" form:"complaint"`
	IsApproved       int            `json:"is_approved" form:"is_approved"`
	IsFinish         bool           `json:"is_finish" form:"is_finish"`
	FinishedAt       time.Time      `json:"finished_at" form:"finished_at"`
	ScheduleDate     string         `json:"schedule_date" form:"schedule_date"`
	ScheduleDateIndo string         `json:"schedule_date_indo" form:"schedule_date_indo"`
	ScheduleTime     string         `json:"schedule_time" form:"schedule_time"`
	Patient          Patient        `json:"patient" form:"patient"`
	Doctor           DoctorRes      `json:"doctor" form:"doctor"`
}

type ApprovalReq struct {
	IsApproved int `json:"is_approved" form:"is_approved"`
}

type OutpatientSessionDashboardRes struct {
	Patient struct {
		NIK           string    `json:"nik"`
		Name          string    `json:"name"`
		BirthDate     time.Time `json:"birth_date"`
		Gender        int       `json:"gender"`
		Age           int       `json:"age"`
		Phone         string    `json:"phone"`
		Address       string    `json:"address"`
		MaritalStatus bool      `json:"marital_status"`
		ReligionName  string    `json:"religion_name"`
	} `json:"patient"`
	Doctor struct {
		Name           string `json:"name"`
		LicenseNumber  string `json:"license_number" form:"license_number"`
		SpecialityName string `json:"speciality_name" form:"speciality_name"`
	} `json:"doctor"`
	Schedule     time.Time `json:"schedule" form:"schedule"`
	Complaint    string    `json:"complaint" form:"complaint"`
	IsApproved   int       `json:"is_approved" form:"is_approved"`
	IsFinish     bool      `json:"is_finish" form:"is_finish"`
	FinishedAt   time.Time `json:"finished_at" form:"finished_at"`
	ScheduleDate string    `json:"schedule_date" form:"schedule_date"`
	ScheduleTime string    `json:"schedule_time" form:"schedule_time"`
}
