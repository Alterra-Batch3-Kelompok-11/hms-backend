package patientUseCase

import (
	"fmt"
	"gorm.io/gorm"
	"hms-backend/constants"
	"hms-backend/dto"
	"hms-backend/models"
	"hms-backend/repositories/patientRepository"
	"strconv"
	"strings"
	"time"
)

type PatientUseCase interface {
	GetAll() ([]dto.PatientRes, error)
	GetById(id uint) (dto.PatientRes, error)
	Create(payload dto.Patient) (dto.PatientRes, error)
	Update(id uint, payload dto.Patient) (dto.PatientRes, error)
	Delete(id uint) error
}

type patientUseCase struct {
	patientRep patientRepository.PatientRepository
}

func New(patRep patientRepository.PatientRepository) *patientUseCase {
	return &patientUseCase{patRep}
}

func (uc *patientUseCase) GetAll() ([]dto.PatientRes, error) {
	var res []dto.PatientRes
	patients, err := uc.patientRep.GetAll()
	if err != nil {
		return res, err
	}

	for _, patient := range patients {
		birthDateString := fmt.Sprintf("%02d", patient.BirthDate.Day()) + "-" +
			strconv.Itoa(int(patient.BirthDate.Month())) + "-" +
			strconv.Itoa(patient.BirthDate.Year())

		birthDateIndoString := fmt.Sprintf("%02d", patient.BirthDate.Day()) + " " +
			constants.Bulan[int(patient.BirthDate.Month())] + " " +
			strconv.Itoa(patient.BirthDate.Year())

		res = append(res, dto.PatientRes{
			ID:                  patient.ID,
			CreatedAt:           patient.CreatedAt,
			UpdatedAt:           patient.UpdatedAt,
			DeletedAt:           patient.DeletedAt,
			Nik:                 patient.Nik,
			Name:                patient.Name,
			BirthDate:           patient.BirthDate,
			BirthDateString:     birthDateString,
			BirthDateStringIndo: birthDateIndoString,
			Gender:              patient.Gender,
			Address:             patient.Address,
			Phone:               patient.Phone,
			MaritalStatus:       patient.MaritalStatus,
			ReligionID:          patient.ReligionID,
		})
	}

	return res, nil
}

func (uc *patientUseCase) GetById(id uint) (dto.PatientRes, error) {
	var res dto.PatientRes
	patient, err := uc.patientRep.GetById(id)
	if err != nil {
		return res, err
	}

	birthDateString := fmt.Sprintf("%02d", patient.BirthDate.Day()) + "-" +
		strconv.Itoa(int(patient.BirthDate.Month())) + "-" +
		strconv.Itoa(patient.BirthDate.Year())

	birthDateIndoString := fmt.Sprintf("%02d", patient.BirthDate.Day()) + " " +
		constants.Bulan[int(patient.BirthDate.Month())] + " " +
		strconv.Itoa(patient.BirthDate.Year())

	res = dto.PatientRes{
		ID:                  patient.ID,
		CreatedAt:           patient.CreatedAt,
		UpdatedAt:           patient.UpdatedAt,
		DeletedAt:           patient.DeletedAt,
		Nik:                 patient.Nik,
		Name:                patient.Name,
		BirthDate:           patient.BirthDate,
		BirthDateString:     birthDateString,
		BirthDateStringIndo: birthDateIndoString,
		Gender:              patient.Gender,
		Address:             patient.Address,
		Phone:               patient.Phone,
		MaritalStatus:       patient.MaritalStatus,
		ReligionID:          patient.ReligionID,
	}

	return res, nil
}

func (uc *patientUseCase) Create(payload dto.Patient) (dto.PatientRes, error) {
	splitedBirthDate := strings.Split(payload.BirthDate[0:10], "-")

	dateTimeString := splitedBirthDate[2] + "-" + splitedBirthDate[1] + "-" + splitedBirthDate[0] + "T00:00:00+07:00"
	birthDateTime, err := time.Parse(time.RFC3339, dateTimeString)
	if err != nil {
		return dto.PatientRes{}, err
	}

	patient := models.Patient{
		Model:         gorm.Model{},
		Nik:           payload.NIK,
		Name:          payload.Name,
		BirthDate:     birthDateTime,
		Gender:        payload.Gender,
		Address:       payload.Address,
		Phone:         payload.Phone,
		MaritalStatus: payload.MaritalStatus,
		ReligionID:    payload.ReligionID,
	}

	patUc, err := uc.patientRep.Create(patient)
	if err != nil {
		return dto.PatientRes{}, err
	}

	birthDateString := fmt.Sprintf("%02d", patient.BirthDate.Day()) + "-" +
		strconv.Itoa(int(patient.BirthDate.Month())) + "-" +
		strconv.Itoa(patient.BirthDate.Year())

	birthDateIndoString := fmt.Sprintf("%02d", patient.BirthDate.Day()) + " " +
		constants.Bulan[int(patient.BirthDate.Month())] + " " +
		strconv.Itoa(patient.BirthDate.Year())

	res := dto.PatientRes{
		ID:                  patUc.ID,
		CreatedAt:           patUc.CreatedAt,
		UpdatedAt:           patUc.UpdatedAt,
		DeletedAt:           patUc.DeletedAt,
		Nik:                 patUc.Nik,
		Name:                patUc.Name,
		BirthDate:           patUc.BirthDate,
		BirthDateString:     birthDateString,
		BirthDateStringIndo: birthDateIndoString,
		Gender:              patUc.Gender,
		Address:             patUc.Address,
		Phone:               patUc.Phone,
		MaritalStatus:       patUc.MaritalStatus,
		ReligionID:          patUc.ReligionID,
	}

	return res, nil
}

func (uc *patientUseCase) Update(id uint, payload dto.Patient) (dto.PatientRes, error) {
	splitedBirthDate := strings.Split(payload.BirthDate[0:10], "-")

	dateTimeString := splitedBirthDate[2] + "-" + splitedBirthDate[1] + "-" + splitedBirthDate[0] + "T00:00:00+07:00"
	birthDateTime, err := time.Parse(time.RFC3339, dateTimeString)
	if err != nil {
		return dto.PatientRes{}, err
	}

	patient := models.Patient{
		Model:         gorm.Model{},
		Nik:           payload.NIK,
		Name:          payload.Name,
		BirthDate:     birthDateTime,
		Gender:        payload.Gender,
		Address:       payload.Address,
		Phone:         payload.Phone,
		MaritalStatus: payload.MaritalStatus,
		ReligionID:    payload.ReligionID,
	}

	patUc, err := uc.patientRep.Update(id, patient)
	if err != nil {
		return dto.PatientRes{}, err
	}

	birthDateString := fmt.Sprintf("%02d", patient.BirthDate.Day()) + "-" +
		strconv.Itoa(int(patient.BirthDate.Month())) + "-" +
		strconv.Itoa(patient.BirthDate.Year())

	birthDateIndoString := fmt.Sprintf("%02d", patient.BirthDate.Day()) + " " +
		constants.Bulan[int(patient.BirthDate.Month())] + " " +
		strconv.Itoa(patient.BirthDate.Year())

	res := dto.PatientRes{
		ID:                  patUc.ID,
		CreatedAt:           patUc.CreatedAt,
		UpdatedAt:           patUc.UpdatedAt,
		DeletedAt:           patUc.DeletedAt,
		Nik:                 patUc.Nik,
		Name:                patUc.Name,
		BirthDate:           patUc.BirthDate,
		BirthDateString:     birthDateString,
		BirthDateStringIndo: birthDateIndoString,
		Gender:              patUc.Gender,
		Address:             patUc.Address,
		Phone:               patUc.Phone,
		MaritalStatus:       patUc.MaritalStatus,
		ReligionID:          patUc.ReligionID,
	}

	return res, nil
}

func (uc *patientUseCase) Delete(id uint) error {

	err := uc.patientRep.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
