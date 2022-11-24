package specialityUseCase

import (
	"gorm.io/gorm"
	"hms-backend/dto"
	"hms-backend/models"
	"hms-backend/repositories/specialityRepository"
)

type SpecialityUseCase interface {
	GetAll() ([]dto.Speciality, error)
	GetById(id uint) (dto.Speciality, error)
	Create(payload dto.Speciality) (dto.Speciality, error)
	Update(id uint, payload dto.Speciality) (dto.Speciality, error)
	Delete(id uint) error
}

type specialityUseCase struct {
	specialityRep specialityRepository.SpecialityRepository
}

func New(spcRep specialityRepository.SpecialityRepository) *specialityUseCase {
	return &specialityUseCase{spcRep}
}
func (uc *specialityUseCase) GetAll() ([]dto.Speciality, error) {
	var res []dto.Speciality
	roles, err := uc.specialityRep.GetAll()
	if err != nil {
		return res, err
	}

	for _, role := range roles {
		res = append(res, dto.Speciality{
			ID:        role.ID,
			CreatedAt: role.CreatedAt,
			UpdatedAt: role.UpdatedAt,
			DeletedAt: role.DeletedAt,
			Name:      role.Name,
		})
	}

	return res, nil
}
func (uc *specialityUseCase) GetById(id uint) (dto.Speciality, error) {
	var res dto.Speciality
	role, err := uc.specialityRep.GetById(id)
	if err != nil {
		return res, err
	}

	res.ID = role.ID
	res.CreatedAt = role.CreatedAt
	res.UpdatedAt = role.UpdatedAt
	res.DeletedAt = role.DeletedAt
	res.Name = role.Name

	return res, nil
}
func (uc *specialityUseCase) Create(payload dto.Speciality) (dto.Speciality, error) {
	speciality := models.Speciality{
		Model: gorm.Model{},
		Name:  payload.Name,
	}

	resUc, err := uc.specialityRep.Create(speciality)
	if err != nil {
		return payload, err
	}

	payload.ID = resUc.ID
	payload.CreatedAt = resUc.CreatedAt
	payload.DeletedAt = resUc.DeletedAt
	payload.UpdatedAt = resUc.UpdatedAt
	payload.Name = resUc.Name

	return payload, nil
}
func (uc *specialityUseCase) Update(id uint, payload dto.Speciality) (dto.Speciality, error) {
	speciality := models.Speciality{
		Model: gorm.Model{},
		Name:  payload.Name,
	}

	resUc, err := uc.specialityRep.Update(id, speciality)
	if err != nil {
		return payload, err
	}

	payload.ID = id
	payload.CreatedAt = resUc.CreatedAt
	payload.DeletedAt = resUc.DeletedAt
	payload.UpdatedAt = resUc.UpdatedAt
	payload.Name = resUc.Name

	return payload, nil
}
func (uc *specialityUseCase) Delete(id uint) error {

	err := uc.specialityRep.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
