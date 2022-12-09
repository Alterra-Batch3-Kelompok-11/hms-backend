package dto

import "time"

type InsertPatientCondition struct {
	OutpatientSessionId uint   `json:"outpatient_session_id" form:"outpatient_session_id"`
	Description         string `json:"description" form:"description"`
	Medicine            string `json:"medicine" form:"medicine"`
	Allergy             string `json:"allergy" form:"allergy"`
}

type InsertPatientConditionRes struct {
	OutpatientSessionId uint      `json:"outpatient_session_id" form:"outpatient_session_id"`
	Description         string    `json:"description" form:"description"`
	Medicine            string    `json:"medicine" form:"medicine"`
	Allergy             string    `json:"allergy" form:"allergy"`
	IsFinish            bool      `json:"is_finish" form:"is_finish"`
	FinishedAt          time.Time `json:"finished_at" form:"finished_at"`
	FinishedAtIndo      string    `json:"finished_at_indo" form:"finished_at_indo"`
}

type PatientConditionRes struct {
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
	Schedule         time.Time `json:"schedule" form:"schedule"`
	ScheduleDate     string    `json:"schedule_date" form:"schedule_date"`
	ScheduleDateIndo string    `json:"schedule_date_indo" form:"schedule_date_indo"`
	ScheduleTime     string    `json:"schedule_time" form:"schedule_time"`
	Complaint        string    `json:"complaint" form:"complaint"`
	IsApproved       int       `json:"is_approved" form:"is_approved"`
	Description      string    `json:"description" form:"description"`
	Medicine         string    `json:"medicine" form:"medicine"`
	Allergy          string    `json:"allergy" form:"allergy"`
	IsFinish         bool      `json:"is_finish" form:"is_finish"`
	FinishedAt       time.Time `json:"finished_at" form:"finished_at"`
	FinishedAtIndo   string    `json:"finished_at_indo" form:"finished_at_indo"`
}
