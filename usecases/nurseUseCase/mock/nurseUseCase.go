package mock

import (
	"errors"
	"hms-backend/dto"

	"github.com/stretchr/testify/mock"
)

type NurseUseCaseMock struct {
	mock.Mock
}

func New() *NurseUseCaseMock {
	return &NurseUseCaseMock{}
}

func (uc *NurseUseCaseMock) GetAll() ([]dto.NurseRes, error) {
	args := uc.Called()

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).([]dto.NurseRes)

	if isSuccess {
		return data, nil
	} else {
		return []dto.NurseRes{}, errors.New("testing error")
	}
}

func (uc *NurseUseCaseMock) GetById(id uint) (dto.NurseRes, error) {
	args := uc.Called(id)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.NurseRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.NurseRes{}, errors.New("testing error")
	}
}

func (uc *NurseUseCaseMock) GetByLicenseNumber(licenseNumber string) (dto.NurseRes, error) {
	args := uc.Called(licenseNumber)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.NurseRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.NurseRes{}, errors.New("testing error")
	}
}

func (uc *NurseUseCaseMock) Create(payload dto.NurseReq) (dto.NurseRes, error) {
	args := uc.Called(payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.NurseRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.NurseRes{}, errors.New("testing error")
	}
}

func (uc *NurseUseCaseMock) Update(id uint, payload dto.NurseReq) (dto.NurseRes, error) {
	args := uc.Called(id, payload)

	isSuccess := args.Get(0).(bool)
	data := args.Get(1).(dto.NurseRes)

	if isSuccess {
		return data, nil
	} else {
		return dto.NurseRes{}, errors.New("testing error")
	}
}

func (uc *NurseUseCaseMock) Delete(id uint) error {
	args := uc.Called(id)

	isSuccess := args.Get(0).(bool)

	if isSuccess {
		return nil
	} else {
		return errors.New("testing error")
	}
}
