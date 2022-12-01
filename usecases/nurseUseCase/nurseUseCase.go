package nurseUseCase

import (
	"hms-backend/dto"
	"hms-backend/models"
	"hms-backend/repositories/nurseRepository"

	"gorm.io/gorm"
)

type NurseUseCase interface {
	GetAll() ([]dto.NurseRes, error)
	GetById(id uint) (dto.NurseRes, error)
	GetByLicenseNumber(licenseNumber string) (dto.NurseRes, error)
	Create(payload dto.NurseRes) (dto.NurseRes, error)
	Update(id uint, payload dto.NurseRes) (dto.NurseRes, error)
	Delete(id uint) error
}

type nurseUseCase struct {
	nurseRep nurseRepository.NurseRepository
}

func New(nurseRep nurseRepository.NurseRepository) *nurseUseCase {
	return &nurseUseCase{nurseRep}
}

func (uc *nurseUseCase) GetAll() ([]dto.NurseRes, error) {
	var res []dto.NurseRes
	nurses, err := uc.nurseRep.GetAll()
	if err != nil {
		return res, err
	}

	for _, nurse := range nurses {
		res = append(res, dto.NurseRes{
			UserID:        nurse.UserId,
			LicenseNumber: nurse.LicenseNumber,
			CreatedAt:     nurse.CreatedAt,
			UpdatedAt:     nurse.UpdatedAt,
			DeletedAt:     nurse.DeletedAt,
		})
	}

	return res, nil
}

func (uc *nurseUseCase) GetById(id uint) (dto.NurseRes, error) {
	var res dto.NurseRes
	nurse, err := uc.nurseRep.GetById(id)
	if err != nil {
		return res, err
	}

	res = dto.NurseRes{
		UserID:        nurse.UserId,
		LicenseNumber: nurse.LicenseNumber,
		CreatedAt:     nurse.CreatedAt,
		UpdatedAt:     nurse.UpdatedAt,
		DeletedAt:     nurse.DeletedAt,
	}

	return res, nil
}

func (uc *nurseUseCase) GetByLicenseNumber(licenseNumber string) (dto.NurseRes, error) {
	var res dto.NurseRes
	nurse, err := uc.nurseRep.GetByLicenseNumber(licenseNumber)
	if err != nil {
		return res, err
	}

	res = dto.NurseRes{
		UserID:        nurse.UserId,
		LicenseNumber: nurse.LicenseNumber,
		CreatedAt:     nurse.CreatedAt,
		UpdatedAt:     nurse.UpdatedAt,
		DeletedAt:     nurse.DeletedAt,
	}

	return res, nil
}

func (uc *nurseUseCase) Create(payload dto.NurseRes) (dto.NurseRes, error) {
	nurse := models.Nurse{
		Model:         gorm.Model{},
		UserId:        payload.UserID,
		LicenseNumber: payload.LicenseNumber,
	}

	var res dto.NurseRes
	nurse, err := uc.nurseRep.Create(nurse)
	if err != nil {
		return res, err
	}

	res = dto.NurseRes{
		UserID:        nurse.UserId,
		LicenseNumber: nurse.LicenseNumber,
		CreatedAt:     nurse.CreatedAt,
		UpdatedAt:     nurse.UpdatedAt,
		DeletedAt:     nurse.DeletedAt,
	}

	return res, nil
}

func (uc *nurseUseCase) Update(id uint, payload dto.NurseRes) (dto.NurseRes, error) {
	nurse := models.Nurse{
		Model:         gorm.Model{},
		UserId:        payload.UserID,
		LicenseNumber: payload.LicenseNumber,
	}

	var res dto.NurseRes
	nurse, err := uc.nurseRep.Update(id, nurse)
	if err != nil {
		return res, err
	}

	res = dto.NurseRes{
		UserID:        nurse.UserId,
		LicenseNumber: nurse.LicenseNumber,
		CreatedAt:     nurse.CreatedAt,
		UpdatedAt:     nurse.UpdatedAt,
		DeletedAt:     nurse.DeletedAt,
	}

	return res, nil
}

func (uc *nurseUseCase) Delete(id uint) error {

	err := uc.nurseRep.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
