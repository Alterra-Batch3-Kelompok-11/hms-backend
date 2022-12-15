package dto

import (
	"time"

	"github.com/go-playground/validator/v10"
)

type Patient struct {
	NIK           string    `json:"nik" validate:"required"`
	Name          string    `json:"name" validate:"required"`
	BirthDate     time.Time `json:"birth_date"`
	Gender        int       `json:"gender" validate:"required"`
	Phone         string    `json:"phone" validate:"required"`
	Address       string    `json:"address" validate:"required"`
	MaritalStatus bool      `json:"marital_status"`
	ReligionID    uint      `json:"religion_id"`
}

func (ctrl *Patient) Validate() error {
	validate := validator.New()
	err := validate.Struct(ctrl)
	if err != nil {
		return err
	}

	return nil
}

type PatientToday struct {
	Name             string    `json:"name"`
	ScheduleDate     string    `json:"schedule_date" form:"schedule_date"`
	ScheduleDateIndo string    `json:"schedule_date_indo" form:"schedule_date_indo"`
	ScheduleTime     string    `json:"schedule_time" form:"schedule_time"`
	Schedule         time.Time `json:"schedule" form:"schedule"`
	Complaint        string    `json:"complaint" form:"complaint"`
}
