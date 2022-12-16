package mock

import (
	"errors"
	"hms-backend/dto"

	"github.com/stretchr/testify/mock"
)

type PatientUseCaseMock struct {
	mock.Mock
}

func New() *PatientUseCaseMock {
	return &PatientUseCaseMock{}
}

func (uc *PatientUseCaseMock) GetAll() ([]dto.PatientRes, error) {
	args := uc.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]dto.PatientRes)

	if isSuccess {
		return data, nil
	} else {
		return []dto.PatientRes{}, errors.New("testing error")
	}
}

func (uc *PatientUseCaseMock) GetById(id uint) (dto.PatientRes, error) {
	args := uc.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.PatientRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.PatientRes{}, errors.New("testing error")
	}
}

func (uc *PatientUseCaseMock) Create(payload dto.Patient) (dto.PatientRes, error) {
	args := uc.Called(payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.PatientRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.PatientRes{}, errors.New("testing error")
	}
}

func (uc *PatientUseCaseMock) Update(id uint, payload dto.Patient) (dto.PatientRes, error) {
	args := uc.Called(id, payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.PatientRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.PatientRes{}, errors.New("testing error")
	}
}

func (uc *PatientUseCaseMock) Delete(id uint) error {
	args := uc.Called(id)

	isSuccess := args.Get(0).(bool)

	if isSuccess {
		return nil
	} else {
		return errors.New("testing error")
	}
}
