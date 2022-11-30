package religionUseCase

import (
	"hms-backend/dto"
	"hms-backend/repositories/religionRepository"
)

type ReligionUseCase interface {
	GetAll() ([]dto.Religion, error)
	GetById(id uint) (dto.Religion, error)
}

type religionUseCase struct {
	religionRepository religionRepository.ReligionRepository
}

func New(rlgRep religionRepository.ReligionRepository) *religionUseCase {
	return &religionUseCase{rlgRep}
}
func (uc *religionUseCase) GetAll() ([]dto.Religion, error) {
	var res []dto.Religion
	religions, err := uc.religionRepository.GetAll()
	if err != nil {
		return res, err
	}

	for _, religion := range religions {
		res = append(res, dto.Religion{
			ID:        religion.ID,
			CreatedAt: religion.CreatedAt,
			UpdatedAt: religion.UpdatedAt,
			DeletedAt: religion.DeletedAt,
			Name:      religion.Name,
		})
	}

	return res, nil
}
func (uc *religionUseCase) GetById(id uint) (dto.Religion, error) {
	var res dto.Religion
	religion, err := uc.religionRepository.GetById(id)
	if err != nil {
		return res, err
	}

	res.ID = religion.ID
	res.CreatedAt = religion.CreatedAt
	res.UpdatedAt = religion.UpdatedAt
	res.DeletedAt = religion.DeletedAt
	res.Name = religion.Name

	return res, nil
}
