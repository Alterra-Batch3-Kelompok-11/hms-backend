package roleUseCase

import (
	"hms-backend/dto"
	"hms-backend/repositories/roleRepository"
)

type RoleUseCase interface {
	GetAll() ([]dto.Role, error)
	GetById(id uint) (dto.Role, error)
}

type roleUseCase struct {
	roleRepository roleRepository.RoleRepository
}

func New(roleRep roleRepository.RoleRepository) *roleUseCase {
	return &roleUseCase{roleRep}
}
func (uc *roleUseCase) GetAll() ([]dto.Role, error) {
	var res []dto.Role
	roles, err := uc.roleRepository.GetAll()
	if err != nil {
		return res, err
	}

	for _, role := range roles {
		res = append(res, dto.Role{
			ID:        role.ID,
			CreatedAt: role.CreatedAt,
			UpdatedAt: role.UpdatedAt,
			DeletedAt: role.DeletedAt,
			Name:      role.Name,
		})
	}

	return res, nil
}
func (uc *roleUseCase) GetById(id uint) (dto.Role, error) {
	var res dto.Role
	role, err := uc.roleRepository.GetById(id)
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