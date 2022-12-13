package dto

import "time"

type Patient struct {
	NIK           string    `json:"nik"`
	Name          string    `json:"name"`
	BirthDate     time.Time `json:"birth_date"`
	Gender        int       `json:"gender"`
	Phone         string    `json:"phone"`
	Address       string    `json:"address"`
	MaritalStatus bool      `json:"marital_status"`
	ReligionID    uint      `json:"religion_id"`
}

type PatientToday struct {
	Name             string    `json:"name"`
	ScheduleDate     string    `json:"schedule_date" form:"schedule_date"`
	ScheduleDateIndo string    `json:"schedule_date_indo" form:"schedule_date_indo"`
	ScheduleTime     string    `json:"schedule_time" form:"schedule_time"`
	Schedule         time.Time `json:"schedule" form:"schedule"`
	Complaint        string    `json:"complaint" form:"complaint"`
}
