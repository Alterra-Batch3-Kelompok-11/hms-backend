package dto

import "time"

type Patient struct {
	ID            uint      `json:"id"`
	NIK           string    `json:"nik"`
	Name          string    `json:"name"`
	BirthDate     time.Time `json:"birth_date"`
	Gender        int       `json:"gender"`
	Phone         string    `json:"phone"`
	Address       string    `json:"address"`
	MaritalStatus bool      `json:"marital_status"`
	ReligionID    uint      `json:"religion_id"`
}
