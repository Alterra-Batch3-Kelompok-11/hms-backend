package mock

import (
	"errors"
	"hms-backend/dto"

	"github.com/stretchr/testify/mock"
)

type SpecialityUseCaseMock struct {
	mock.Mock
}

func New() *SpecialityUseCaseMock {
	return &SpecialityUseCaseMock{}
}

func (uc *SpecialityUseCaseMock) GetAll() ([]dto.Speciality, error) {
	args := uc.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]dto.Speciality)

	if isSuccess {
		return data, nil
	} else {
		return []dto.Speciality{}, errors.New("testing error")
	}
}

func (uc *SpecialityUseCaseMock) GetById(id uint) (dto.Speciality, error) {
	args := uc.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.Speciality)

	if isSuccess {
		return data, nil
	} else {
		return dto.Speciality{}, errors.New("testing error")
	}
}

func (uc *SpecialityUseCaseMock) Create(payload dto.Speciality) (dto.Speciality, error) {
	args := uc.Called(payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.Speciality)

	if isSuccess {
		return data, nil
	} else {
		return dto.Speciality{}, errors.New("testing error")
	}
}

func (uc *SpecialityUseCaseMock) Update(id uint, payload dto.Speciality) (dto.Speciality, error) {
	args := uc.Called(id, payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.Speciality)

	if isSuccess {
		return data, nil
	} else {
		return dto.Speciality{}, errors.New("testing error")
	}
}

func (uc *SpecialityUseCaseMock) Delete(id uint) error {
	args := uc.Called(id)

	isSuccess := args.Get(0).(bool)

	if isSuccess {
		return nil
	} else {
		return errors.New("testing error")
	}
}
