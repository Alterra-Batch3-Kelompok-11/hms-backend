package models

import (
	"gorm.io/gorm"
	"time"
)

type Patient struct {
	gorm.Model
	ReligionId    uint
	StatusId      uint
	Nik           string    `json:"nik" form:"nik"`
	Name          string    `json:"name" form:"name"`
	BirthDate     time.Time `json:"birth_date" form:"birth_date"`
	Gender        bool      `json:"gender" form:"gender"`
	Address       string    `json:"address" form:"address"`
	MaritalStatus bool      `json:"marital_status" form:"address"`
	Religion      Religion
	Status        Status
}
