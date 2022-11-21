package models

import (
	"time"

	"gorm.io/gorm"
)

type Patient struct {
	gorm.Model
	ReligionId    uint
	Nik           string    `json:"nik" form:"nik"`
	Name          string    `json:"name" form:"name"`
	BirthDate     time.Time `json:"birth_date" form:"birth_date"`
	Gender        bool      `json:"gender" form:"gender"`
	Address       string    `json:"address" form:"address"`
	MaritalStatus bool      `json:"marital_status" form:"address"`
	Religion      Religion
	Status        Status
}
