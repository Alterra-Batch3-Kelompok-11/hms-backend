package dto

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type DoctorScheduleReq struct {
	DoctorId  uint   `json:"doctor_id" validate:"required"`
	DayInt    int    `json:"day_int" form:"day_int"`
	StartTime string `json:"start_time" form:"start_time" validate:"required"`
	EndTime   string `json:"end_time" form:"end_time" validate:"required"`
}

func (ctrl *DoctorScheduleReq) Validate() error {
	validate := validator.New()
	err := validate.Struct(ctrl)
	if err != nil {
		return err
	}

	return nil
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

type DoctorProfileScheduleRes struct {
	ID        uint      `json:"id" form:"id"`
	DoctorId  uint      `json:"doctor_id"`
	Date      time.Time `json:"date" form:"date"`
	DateIndo  string    `json:"date_indo" form:"date_indo"`
	DayInt    int       `json:"day_int" form:"day_int"`
	DayString string    `json:"day_string" form:"day_string"`
	StartTime string    `json:"start_time" form:"start_time"`
	EndTime   string    `json:"end_time" form:"end_time"`
}
