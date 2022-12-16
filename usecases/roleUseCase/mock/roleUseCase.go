package mock

import (
	"errors"
	"hms-backend/dto"

	"github.com/stretchr/testify/mock"
)

type RoleUseCaseMock struct {
	mock.Mock
}

func New() *RoleUseCaseMock {
	return &RoleUseCaseMock{}
}

func (uc *RoleUseCaseMock) GetAll() ([]dto.Role, error) {
	args := uc.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]dto.Role)

	if isSuccess {
		return data, nil
	} else {
		return []dto.Role{}, errors.New("testing error")
	}
}

func (uc *RoleUseCaseMock) GetById(id uint) (dto.Role, error) {
	args := uc.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.Role)

	if isSuccess {
		return data, nil
	} else {
		return dto.Role{}, errors.New("testing error")
	}
}
