package controllers

import (
	"hms-backend/constants"
	"hms-backend/databases"
	"hms-backend/models"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func PatientsCreate(c echo.Context) error {

	tm, _ := time.Parse(constants.TimeLayout, "2021-01-01 00:00:00")

	var Patient struct {
		NIK           string    `json:"nik"`
		Name          string    `json:"name"`
		BirthDate     time.Time `json:"birth_date"`
		Gender        bool
		Address       string
		MaritalStatus bool
		ReligionID    uint
		StatusID      uint
	}

	err := c.Bind(&Patient)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	post := models.Patient{Nik: Patient.NIK, Name: Patient.Name, BirthDate: tm, Gender: true, Address: Patient.Address, MaritalStatus: true, ReligionID: Patient.ReligionID, StatusID: Patient.StatusID}

	result := databases.DB.Create(&post)

	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, result.Error)
	}

	return c.JSON(http.StatusOK, "Adding patient succes")
}

func PatientsIndex(c echo.Context) error {

	var posts []models.Patient
	databases.DB.Find(&posts)

	return c.JSON(http.StatusOK, posts)

}

func PatientShow(c echo.Context) error {

	id := c.Param("id")

	var post models.Patient
	databases.DB.First(&post, id)

	return c.JSON(http.StatusOK, post)
}

func PatientsUpdate(c echo.Context) error {

	id := c.Param("id")

	var Patient struct {
		NIK           string    `json:"nik"`
		Name          string    `json:"name"`
		BirthDate     time.Time `json:"birth_date"`
		Gender        bool
		Address       string
		MaritalStatus bool
		ReligionID    uint
		StatusID      uint
	}

	err := c.Bind(&Patient)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	var post models.Patient
	databases.DB.First(&post, id)

	databases.DB.Model(&post).Updates(models.Patient{
		Nik:        Patient.NIK,
		Name:       Patient.Name,
		Address:    Patient.Address,
		ReligionID: Patient.ReligionID,
		StatusID:   Patient.StatusID,
	})

	return c.JSON(http.StatusOK, "Updated Patient Data")
}

func PatientsDelete(c echo.Context) error {

	id := c.Param("id")

	databases.DB.Delete(&models.Patient{}, id)

	return c.JSON(http.StatusOK, "Deleted Patient Data")
}
