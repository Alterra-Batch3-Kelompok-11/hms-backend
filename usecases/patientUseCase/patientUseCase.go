package patientUseCase

import (
	"gorm.io/gorm"
	"hms-backend/dto"
	"hms-backend/models"
	"hms-backend/repositories/patientRepository"
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
		res = append(res, dto.PatientRes{
			ID:            patient.ID,
			CreatedAt:     patient.CreatedAt,
			UpdatedAt:     patient.UpdatedAt,
			DeletedAt:     patient.DeletedAt,
			Nik:           patient.Nik,
			Name:          patient.Name,
			BirthDate:     patient.BirthDate,
			Gender:        patient.Gender,
			Address:       patient.Address,
			Phone:         patient.Phone,
			MaritalStatus: patient.MaritalStatus,
			ReligionID:    patient.ReligionID,
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

	res = dto.PatientRes{
		ID:            patient.ID,
		CreatedAt:     patient.CreatedAt,
		UpdatedAt:     patient.UpdatedAt,
		DeletedAt:     patient.DeletedAt,
		Nik:           patient.Nik,
		Name:          patient.Name,
		BirthDate:     patient.BirthDate,
		Gender:        patient.Gender,
		Address:       patient.Address,
		Phone:         patient.Phone,
		MaritalStatus: patient.MaritalStatus,
		ReligionID:    patient.ReligionID,
	}

	return res, nil
}

func (uc *patientUseCase) Create(payload dto.Patient) (dto.PatientRes, error) {
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
		return dto.PatientRes{}, err
	}

	res := dto.PatientRes{
		ID:            patUc.ID,
		CreatedAt:     patUc.CreatedAt,
		UpdatedAt:     patUc.UpdatedAt,
		DeletedAt:     patUc.DeletedAt,
		Nik:           patUc.Nik,
		Name:          patUc.Name,
		BirthDate:     patUc.BirthDate,
		Gender:        patUc.Gender,
		Address:       patUc.Address,
		Phone:         patUc.Phone,
		MaritalStatus: patUc.MaritalStatus,
		ReligionID:    patUc.ReligionID,
	}

	return res, nil
}

func (uc *patientUseCase) Update(id uint, payload dto.Patient) (dto.PatientRes, error) {
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
		return dto.PatientRes{}, err
	}

	res := dto.PatientRes{
		ID:            patUc.ID,
		CreatedAt:     patUc.CreatedAt,
		UpdatedAt:     patUc.UpdatedAt,
		DeletedAt:     patUc.DeletedAt,
		Nik:           patUc.Nik,
		Name:          patUc.Name,
		BirthDate:     patUc.BirthDate,
		Gender:        patUc.Gender,
		Address:       patUc.Address,
		Phone:         patUc.Phone,
		MaritalStatus: patUc.MaritalStatus,
		ReligionID:    patUc.ReligionID,
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
