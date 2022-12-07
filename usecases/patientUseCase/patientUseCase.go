package patientUseCase

import (
	"hms-backend/dto"
	"hms-backend/models"
	"hms-backend/repositories/patientRepository"

	"gorm.io/gorm"
)

type PatientUseCase interface {
	GetAll() ([]dto.Patient, error)
	GetById(id uint) (dto.Patient, error)
	Create(payload dto.Patient) (dto.Patient, error)
	Update(id uint, payload dto.Patient) (dto.Patient, error)
	Delete(id uint) error
}

type patientUseCase struct {
	patientRep patientRepository.PatientRepository
}

func New(patRep patientRepository.PatientRepository) *patientUseCase {
	return &patientUseCase{patRep}
}

func (uc *patientUseCase) GetAll() ([]dto.Patient, error) {
	var res []dto.Patient
	roles, err := uc.patientRep.GetAll()
	if err != nil {
		return res, err
	}

	for _, role := range roles {
		res = append(res, dto.Patient{
			ID:            role.ID,
			NIK:           role.Nik,
			Name:          role.Name,
			BirthDate:     role.BirthDate,
			Gender:        role.Gender,
			Phone:         role.Phone,
			Address:       role.Address,
			MaritalStatus: role.MaritalStatus,
			ReligionID:    role.ReligionID,
		})
	}

	return res, nil
}

func (uc *patientUseCase) GetById(id uint) (dto.Patient, error) {
	var res dto.Patient
	role, err := uc.patientRep.GetById(id)
	if err != nil {
		return res, err
	}

	res.ID = role.ID
	res.NIK = role.Nik
	res.Name = role.Name
	res.BirthDate = role.BirthDate
	res.Gender = role.Gender
	res.Address = role.Address
	res.Phone = role.Phone
	res.MaritalStatus = role.MaritalStatus
	res.ReligionID = role.ReligionID

	return res, nil
}

func (uc *patientUseCase) Create(payload dto.Patient) (dto.Patient, error) {
	patient := models.Patient{
		Model:         gorm.Model{},
		Nik:           payload.NIK,
		Name:          payload.Name,
		BirthDate:     payload.BirthDate,
		Gender:        payload.Gender,
		Address:       payload.Address,
		Phone:         payload.Phone,
		MaritalStatus: payload.MaritalStatus,
		ReligionID:    payload.ReligionID,
	}

	patUc, err := uc.patientRep.Create(patient)
	if err != nil {
		return payload, err
	}

	payload.NIK = patUc.Nik
	payload.Name = patUc.Name
	payload.BirthDate = patUc.BirthDate
	payload.Gender = patUc.Gender
	payload.Address = patUc.Address
	payload.Phone = patUc.Phone
	payload.MaritalStatus = patUc.MaritalStatus
	payload.ReligionID = patUc.ReligionID

	return payload, nil
}

func (uc *patientUseCase) Update(id uint, payload dto.Patient) (dto.Patient, error) {
	patient := models.Patient{
		Model:         gorm.Model{},
		Nik:           payload.NIK,
		Name:          payload.Name,
		BirthDate:     payload.BirthDate,
		Gender:        payload.Gender,
		Address:       payload.Address,
		Phone:         payload.Phone,
		MaritalStatus: payload.MaritalStatus,
		ReligionID:    payload.ReligionID,
	}

	patUc, err := uc.patientRep.Update(id, patient)
	if err != nil {
		return payload, err
	}

	payload.NIK = patUc.Nik
	payload.Name = patUc.Name
	payload.BirthDate = patUc.BirthDate
	payload.Gender = patUc.Gender
	payload.Address = patUc.Address
	payload.Phone = patUc.Phone
	payload.MaritalStatus = patUc.MaritalStatus
	payload.ReligionID = patUc.ReligionID

	return payload, nil
}

func (uc *patientUseCase) Delete(id uint) error {

	err := uc.patientRep.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
