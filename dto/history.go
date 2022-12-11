package dto

import "time"

type History struct {
	PatientName      string    `json:"patient_name"`
	Schedule         time.Time `json:"schedule"`
	ScheduleDate     string    `json:"schedule_date"`
	ScheduleDateIndo string    `json:"schedule_date_indo"`
	ScheduleTime     string    `json:"schedule_time"`
	Status           string    `json:"status"`
}
